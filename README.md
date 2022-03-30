# cs5594-iota-htlc
IOTA Hash Time Locked Contract


## Installation
### 1. Set up smart contract directory
```sh
$ mkdir sc
```

### 2. Build the wasp node
```sh
$ docker build -f Dockerfile -t wasp-node ./wasp
```

### 3. Run the container with ports and mounted directory
```sh
$ docker run -p 127.0.0.1:4000:4000/tcp -p 127.0.0.1:5550:5550/tcp -p 127.0.0.1:7000:7000/tcp -p 127.0.0.1:8545:8545/tcp -p 127.0.0.1:9090:9090/tcp -v $(pwd)/contracts:/wasp/smart-contracts -d --name wasp-node wasp-node
```