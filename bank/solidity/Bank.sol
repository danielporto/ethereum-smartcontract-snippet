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
    event OperationsExecuted(uint);
    event PrintConfirmation(string,uint); //trx_id, total_operations

    // emit the total operations in the log
    function logTransferOperations() public {
        emit OperationsExecuted(operations);
    }

    // send money to the contract
    function receiveMoney() public payable {
        balanceReceived += msg.value;
        emit BalanceReceived(msg.value);
    }

    // withdraw all money from the contract to the sender address
    function withdrawMoney() public {
        address payable to = payable(msg.sender);
        to.transfer(getBalance());
    }


    // pay order a payment from contract balance to another account
    function payMoneyTo(address payable _to) public payable {
        _to.transfer(msg.value);
        emit BalanceTransferred(msg.value);
    }


    // transfer money to another account via smartcontract
    function transferMoneyTo(address payable _to, string calldata id) public payable {
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
