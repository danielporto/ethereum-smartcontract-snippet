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
./counter deploy --url "http://192.168.10.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
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

## Print the content of a counter
```
./counter  --url "http://192.168.10.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0" print --contract "0x3debb8a04d6311cfdd72b44ac2db2c4920e9db2d" 
```

## Run increment/decrement
It requires a contract address. 
```
./counter increment --contract "0x3debb8a04d6311cfdd72b44ac2db2c4920e9db2d"  --url "http://192.168.10.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```

## Run workload
It creates a contract and run operations on it
```
./counter workload -o increment -c 10000  --url "http://192.168.10.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```
