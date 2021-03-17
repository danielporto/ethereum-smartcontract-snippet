# brew install solidity
# install abigen:
# go get -u github.com/ethereum/go-ethereum/...
# cd $(go env GOPATH)/src/github.com/ethereum/go-ethereum
# make
# make devtools
#export GO111MODULE=on
abigen --sol Storage.sol --pkg main --out storage.go