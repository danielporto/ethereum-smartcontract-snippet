#!/usr/bin/env bash

abigen --sol solidity/SimpleStorage.sol --pkg contracts --out contracts/simplestorage.go
echo "Done."