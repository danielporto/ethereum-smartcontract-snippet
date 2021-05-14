# Overview
Check  [here](https://www.linode.com/docs/guides/using-cobra/) to learn how create commands and subcommands.


## Requirements
Local development is easier with [Ganache](https://www.trufflesuite.com/ganache). Install it and use the quickstart button
to create the local testing network. Do a right click on any account to get the key used in all commands below.

## Contract cost:
Pay attention to the computation cost. If the execution fails, one common source of problems is insufficient budget.
- deploy               510465 gas
- sort 10 elements     116244 gas
- sort 100 elements:   407497 gas
- sort 1000 elements: 3566945 gas
- sort 1500 elements: 6000003 gas, above 1500 elements block gas limit was exceeded in ganache (can be configured)
- sort 2000 elements: 7649148 gas (takes a few seconds)

## Important: Listening to events
You can connect to the endpoints with `http` and `ws` sockets. However, tracking events are *only* supported via `websockets`. 


## Deploy a contract
It returns a contract address.

### Connecting to Quorum using http: 
```
./quicksort deploy --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```
### Connecting to Quorum using ws (note that the port changes for Quorum) 
```
./quicksort deploy --url "ws://146.193.41.166:23000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```
### Connecting to Ganache:
Use the same port, only change it to ws to listen to events.
```
./quicksort deploy --url "ws://localhost:7585" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```

## Run workload
There are three types of workload: `deploy, sort, mixed`. You can choose the workload using `-o`.

The `deploy` workload will issue operations to install contract. Ensure the wallet has enough budget as deplying is one of most expensive operations.

For the `sort` workload installs one contract and run sort commands. You can tune the size of the array using `-s`  to increase the usage of the CPU/cost of computation. The default size is 100.

The `mixed` worload altearnates between sorting an array and installing a new contract.

Workloads can be configured to run a certain number of operations `-c` or duration `-d`. 

```
./quicksort workload -o sort -c 1000  --url "ws://146.193.41.166:23000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```
