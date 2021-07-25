//SPDX-License-Identifier: undefined
pragma solidity 0.8.4;

contract EthCallTest{
    receive() external payable{}
    
    function checkETHBalance() external view returns(uint) {
        return address(this).balance;
    }
    
    function sendETHBack(uint amount) external returns(uint) {
        require(address(this).balance >= amount,"insufficient balance");
        payable(address(msg.sender)).transfer(amount);
        return amount;
    }
}
