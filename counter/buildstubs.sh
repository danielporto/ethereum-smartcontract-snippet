#!/usr/bin/env bash

abigen --sol solidity/Counter.sol --pkg contracts --out contracts/counter.go
echo "Done."