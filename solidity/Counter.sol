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
    event Increment(uint value);
    event Decrement(uint value);

    function getCount() view public returns (uint) {
        return counter;
    }

    function increment() public {
        counter +=1;
        emit Increment(counter);
    }

    function decrement() public {
        counter -=1;
        emit Decrement(counter);
    }

}