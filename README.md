# cs5594-iota-htlc
IOTA Hash Time Locked Contract


## Installation
### 1. Update git submodule
```sh
$ git submodule update --init --recursive
```

### 2. Set up smart contract directory
```sh
$ mkdir sc
```

### 3. Build the wasp node
```sh
$ docker build -f Dockerfile -t wasp-node ./wasp
```

### 4. Run the container with ports and mounted directory
```sh
$ docker run -p 127.0.0.1:4000:4000/tcp -p 127.0.0.1:5550:5550/tcp -p 127.0.0.1:7000:7000/tcp -p 127.0.0.1:8545:8545/tcp -p 127.0.0.1:8546:8546/tcp -p 127.0.0.1:9090:9090/tcp -v $(pwd)/contracts:/wasp/smart-contracts -d --name wasp-node wasp-node
$ docker run -p 127.0.0.1:4001:4000/tcp -p 127.0.0.1:5551:5550/tcp -p 127.0.0.1:7001:7000/tcp -p 127.0.0.1:8547:8545/tcp -p 127.0.0.1:8548:8546/tcp -p 127.0.0.1:9091:9090/tcp -v $(pwd)/contracts:/wasp/smart-contracts -d --name wasp-node2 wasp-node
```

### 5. Config the wasp-cli
First, enter the running docker container
```sh
$ docker exec -it wasp-node /bin/sh
```

Inside the container, execute the following commands
```sh
$ ./wasp-cli init
```

This will create the `wasp-cli.json` file. Use `vim` to paste the testnet URL into the file
```
    "goshimmer": {
      "api": "https://api.goshimmer.sc.iota.org",
      "faucetpowtarget": -1
    },
    "wasp": {
      "0": {
        "api": "127.0.0.1:9090",
        "nanomsg": "127.0.0.1:5550",
        "peering": "127.0.0.1:4000"
      }
    }
```

Then deploy the chain
```sh
$ ./wasp-cli request-funds
$ ./wasp-cli chain deploy --committee=0 --quorum=1 --chain=devchain --description="chain1"
$ ./wasp-cli chain deposit IOTA:10000
```

Once finished, repeat this again for `chain2`

### 6. Compile and deploy contract
Inside the running container, use the following command to compile the HTLC smart contract
```sh
$ tinygo build -o htlc.wasm -target wasm smart-contracts/go/main.go
```

Then deploy it to the chain
```sh
$ ./wasp-cli chain deploy-contract wasmtime htlc "HTLC" htlc.wasm --verbose --debug
```

After deployment, use the following command to work with the other party
```sh
$ ./wasp-cli chain post-request htlc funcSetSecret string secret string <your-secret>
$ ./wasp-cli chain post-request htlc funcSetValue string value int <transaction>
$ ./wasp-cli chain post-request htlc funcSetReceivder string receivder address <address>
$ ./wasp-cli chain post-request htlc funcSetTime string time int <time>
$ ./wasp-cli chain post-request htlc funcTransfer string secret string <your-secret> string key string <your-key>
```

If a refund is needed, use the `funcWithdraw` after the contract expired
```sh
$ ./wasp-cli chain post-request htlc funcWithdraw
```