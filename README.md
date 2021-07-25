# ethcall

code for demonstrate gas estimation problem

## contract 

```solidity
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
```

## test steps

step1: transfer 1 wei to contract

step2: simulate `sendETHBack()` by set `gasLimit = 0`, and `noSend` = true to do gas estimation 

step3: execute `sendETHBack()` 

## expected result

step1,2,3 all success

## actual result

step1,3 success

step2 failed

output:

```text
test contract deployed:  0x52cCf226B0f1E923ccDD9665fFdc75dDffe76b31
sleep for 20 second to make sure test contract deployed
sendEth tx 0xaee1a8379c2eef1de42ab0fb8b2defa7c3f6b39c1398be2e236277bc407c97c4 balanceBefore 0, balanceAfter 1
sendEthBackSimulation error failed to estimate gas needed: execution reverted: insufficient balance
sendEthBack tx 0xcc7cd7a970c2e5f999d658dcc5dc157c2461c39c83edbee2fb8b475d854e59ed balanceBefore 1, balanceAfter 0
```

