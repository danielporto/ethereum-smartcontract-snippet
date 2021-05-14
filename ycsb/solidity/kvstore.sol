// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

contract KVstore {

  mapping(string=>string) store;

  function get(string memory _key) public view returns(string memory) {
    return store[_key];
  }
  
  function set(string memory _key, string memory _value) public {
    store[_key] = _value;
  }
}