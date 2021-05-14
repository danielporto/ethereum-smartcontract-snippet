// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.8.5;

/**
* @title Counter
* @dev Keeps a counter
* https://www.youtube.com/watch?v=ooN6kZ9vqNQ
* https://www.youtube.com/watch?v=ZDAuELguhQQ
*/

contract Counter {
    uint counter = 0;
    event Increment(uint);
    event Decrement(uint);
    event GetValue(uint);

    function getCounter() public {
        emit GetValue(counter);
    }

    function increment(uint value) public {
        counter +=value;
        emit Increment(counter);
    }

    function decrement(uint value) public {
        counter -=value;
        emit Decrement(counter);
    }

    function set(uint value) public {
        counter = value;
    }



}
