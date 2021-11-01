// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.8.10;

/**
* @title Counter
* @dev Keeps a counter
* https://www.youtube.com/watch?v=ooN6kZ9vqNQ
* https://www.youtube.com/watch?v=ZDAuELguhQQ
*/

contract Counter {
    uint counter = 0;
    event PrintConfirmation(string, uint); //helps identifying the original request
    event PrintValue(uint);


    function getCounter() public returns(uint) {
        emit PrintValue(counter);
        return counter;
    }

    function increment(uint value, string calldata id) public {
        counter +=value;
        emit PrintConfirmation(id, counter);
    }

    function decrement(uint value, string calldata id) public {
        counter -=value;
        emit PrintConfirmation(id, counter);
    }

    function set(uint value, string calldata id) public {
        counter = value;
        emit PrintConfirmation(id, counter);
    }

}
