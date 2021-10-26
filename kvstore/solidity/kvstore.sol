// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

contract KVstore {

  mapping(string=>string) store;
  uint insertCounter = 0;


  event PrintKVAck(uint, string, string);
  event PrintInserts(uint);
  event PrintConfirmation(string, uint); //helps identifying the original request


  function get(string memory _key) public view returns(string memory) {
    return store[_key];
  }

  function set(string memory _key, string memory _value, string calldata id) public {
    store[_key] = _value;
    insertCounter +=1;
    //emit PrintKVAck(insertCounter, _key, _value);
    emit PrintConfirmation(id, insertCounter);
  }

  function PrintTotalInserts() public {
    emit PrintInserts(insertCounter);
  }

}