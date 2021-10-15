# Overview
This is a simple smartcontract example.
The contract itself is designed in solidity.

A client application designed in golang deploy a contract in a ethereum compatible blockchain network (Ganache, test net or GoQuorum). An easy way to test is installing Ganache (mentioned below).

Next, use the local golang client to deploy the contract
and run the workload.

# Requirements
## Ethereum compatible blockchain network
Local development is easier with [Ganache](https://www.trufflesuite.com/ganache). Install it, then use the quickstart button
to create the local testing network. 

Once you have Ganache (or another ethereum blockchain) working, make sure to get 

1. the connection url (http or ws) 
2. private key.

Using Ganache, you can easliy get this information from interface by right clicking in one of the generated accounts.


## Golang
Make sure you have Go installed. Installing go is beyond the scope of this document. 
An easy way to install it is using [ASDF](https://asdf-vm.com/) with [Go plugin](https://github.com/kennyp/asdf-golang). Other plugins are available [here](https://github.com/asdf-vm/asdf-plugins).
We used go version `1.16.7`. 

You will also need go ethereum source (go-ethereum) and tools (abigen) required to build our client application.

```
go get -u github.com/ethereum/go-ethereum@v1.10.9
```

After downloading the code to gopath, you are required to build the tools:
find where the source is installed:
```
go env GOPATH
```
Access the directory to build the source. (the location below can be slightly different for you)
```
cd $(go env GOPATH)/pkg/mod/github.com/ethereum/go-ethereum@v1.10.9
make
make devtools
```

Note: if you find a compilation error, that's likelly permission denied, simply create a bin dir inside build dir.:
```
mkdir -p $(go env GOPATH)/pkg/mod/github.com/ethereum/go-ethereum@v1.10.9/build/bin
```
also try to set permissions to it
```
sudo chmod -R u+rw $(go env GOPATH)/pkg/mod/github.com/ethereum/go-ethereum@v1.10.9 
```
If it fails (happened in osx big sur), open the path with finder and set permissions using GUI.

After that, the abigen command must be available.


# Build the blockchain client
Assuming the steps of Requeriments section was complete, you can create the client by running the build script.

```
bash buildstubs.sh
```
## Deploy the smartcontract
You can connect to the endpoints with `http` and `ws` sockets. However, tracking events are *only* supported via `websockets` (for that, just replace http with ws and set the correct port). 

Since Ganache run locally, you can replace the IP for `127.0.0.1`.
For Quorum, replace the IP with the corresponding one for the actual server node. 

It returns a contract address (use it for printing the counter at the end)
```
./ycsb deploy --host localhost --port 7545 --key "7d6419430c8f84d5305cb060c301214218338a8aa327ce4a7dc428397a17e930"
```
should output:
```
INFO[0000] Connecting to ethereum network...
Nonce: 0
Gas price: 20000000000
Transaction address: 0x843ef8f573414ae4631cf63ff95a05be90bd10f57c3b300abd5adcd854fae077
Contract address 0x24C4fbcd6fa5360Cd1AEA2CeD9F1D7a0cdbc1897
Contract deployed
```

## Run workload
It creates a contract and run operations on it
```
./ycsb workload --operation ycsb --host 127.0.0.1 --port 8545 --key a3cc08be595b8f4ce8cc76bab3f26eb20acdf82c55397711adcc5977341b9b25 --count 100 --dbkeysize 100 --dbvaluesize 200 --threads 10```
