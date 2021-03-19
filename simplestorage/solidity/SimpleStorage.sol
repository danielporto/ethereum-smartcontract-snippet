// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.8.5;



contract SimpleStorage {
    uint public storedData;

    constructor(uint _storedData) public {
        storedData = _storedData;
    }

    function set(uint x) public {
        storedData = x;
    }
    function add(uint x) public {
        storedData = storedData + x;
    }
    function get() public view returns (uint retVal) {
        return storedData;
    }
}
