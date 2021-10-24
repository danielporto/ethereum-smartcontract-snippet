// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.8.10;

/**
* @title Bank Transfer
* @dev Daniel Porto
*/

contract Bank {
    uint operations = 0;
    uint public balanceReceived;

    event BalanceReceived(uint);
    event BalanceTransferred(uint);
    event OperationsExecuted(uint,uint);
    event PrintConfirmation(string,uint); //trx_id, total_operations

    // emit the total operations in the log
    function logTransferOperations() public {
        emit OperationsExecuted(operations,address(this).balance);
    }

    // send money to the contract
    function receiveMoney() public payable {
        balanceReceived += msg.value;
        emit BalanceReceived(msg.value);
    }

    // send money to the contract
    function deposit(string calldata id) public payable {
        balanceReceived += msg.value;
        emit PrintConfirmation(id, operations);
    }


    // withdraw all money from the contract to the sender address
    function withdrawMoney(string calldata id) public {
        address payable to = payable(msg.sender);
        to.transfer(getBalance());
        emit PrintConfirmation(id, operations);
    }


    // withdraw all money from the contract to a given address
    function withdrawMoneyTo(string calldata id, address payable _to) public payable {
        _to.transfer(msg.value);
        emit PrintConfirmation(id, operations);
    }


    // transfer money to another account via smartcontract
    function directTransferTo(address payable _to, string calldata id) public payable {
        _to.transfer(msg.value);
        operations +=1;
        //emit BalanceTransferred(msg.value); //print amount transferred
        emit PrintConfirmation(id, operations);
    }

    // not mined, and free
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

    // not mined, and free
    function getOperations() public view returns (uint) {
        return operations;
    }

}
