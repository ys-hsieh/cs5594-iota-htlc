# VT Spring 2022 CS5594 Course Project: IOTA HTLC Bribery Attack

This is the course project for VT Spring 2022 CS5594 Course Project.

In this project, we develop HTLC that can be further tested with potentical cross-chain bribery attack.

There are two types of HTLCs:
1. WasmVM smart contract
2. EVM smart contract

The following instruction will walk you through both of them.

## Demo Video
1. WasmVM smart contract: [CS5594 Golang HTLC PoC](https://youtu.be/uInMscmOfCU)
2. EVM smart contract

## Requirement
- [GoShimmer](https://wiki.iota.org/goshimmer/welcome) node: In this project, we use the [GoShimmer Testnet](https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/testnet).
- [Wasp Node](https://wiki.iota.org/smart-contracts/overview)
- Docker: The tutorial is set up with Docker, but feel free to use local set ups.

## Installation
First we need to the Wasp Node

### 1. Update git submodule
```sh
$ git submodule update --init --recursive
```

### 2. Update the `docker_config.json`
Replace the `nodeconn` with the testnet `goshimmer.sc.iota.org:5000`


### 3. Build the wasp node
```sh
$ docker build -f Dockerfile -t wasp-node ./wasp
```

### 4. Run the container with ports and mounted directory
The following commands will create two running containers: `wasp-node` and `wasp-node2`
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
$ ./wasp-cli chain deploy --committee=0 --quorum=1 --chain=devchain --description="devchain1"
$ ./wasp-cli chain deposit IOTA:10000
```

Once finished, repeat this again for `devchain2`

### 6. Install EVM contract (only for solidity HTLC)
This step is only for solidity smart contracts. You can skip this is you plan to use WasmVM.

Detail steps please follow the [official documents](https://wiki.iota.org/smart-contracts/guide/evm/create-chain) and [tutorial videos](https://www.youtube.com/watch?v=JbUGX-9BTSo)

## Deploy WasmVM smart contract
### 1. Compile contract
Inside the running container, use the following command to compile the HTLC smart contract
```sh
$ tinygo build -o htlc.wasm -target wasm smart-contracts/go/main.go
```

### 2. Deploy contract
Then deploy it to the chain
```sh
$ ./wasp-cli chain deploy-contract wasmtime htlc "HTLC" htlc.wasm --verbose --debug
```

### 3. Interact with contract
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

Transactions and address records can be found on the [Goshammer Explorer](https://goshimmer.sc.iota.org/explorer)

## Deploy EVM smart contract
### 1. Deploy the EVM Chain Contract
Deploy EVM contract on the chain according to EVM account
```sh
./wasp-cli chain evm deploy -a mychain --alloc accountname:1000000000000000000000000
```

## 2. Run JSON-RPC Interface
Run the JSON-RPC server on port 8545 for you with Chain ID 1074. You can now point MetaMask or Hardhat to that server's address on port 8545
```sh
./wasp-cli chain evm jsonrpc
```
## 3. Complie Hash Time lock Smart Contract on Remix
enter remix website https://remix.ethereum.org/, and upload HTLC.sol file to complie contract
