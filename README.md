

# Overview
This is a set of smartcontract examples.
The contract itself is designed in solidity.

A client application designed in golang deploy a contract in a ethereum compatible blockchain network (Ganache, test net or GoQuorum). An easy way to test is installing Ganache (mentioned below).

Next, use the local golang client to deploy the contract
and run the workload.

# Requirements
* Solidity compiler
* abigen tool (go-ethereum devtools)
* nodejs
* protoc (protobuf dev libs)
* golang (1.14 and 1.16 worked)
* ganache (test network)

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