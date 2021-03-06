package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/huahuayu/ethcall/eth"
	"math/big"
	"time"
)

func main() {
	eth.Init()
	fmt.Println("sleep for 20 seconds to make sure test contract deployed")
	time.Sleep(20 * time.Second)
	txHash := sendEth()
	transaction, isPending, _ := eth.EthClient.TransactionByHash(context.Background(), txHash)
	fmt.Printf("txHash: %s isPending %v\n", transaction.Hash().String(), isPending)
	go sendEthBackSimulation() // simulation failed
	go ethCall()               // eth call success
	sendEthBack()              // tx execution success
}

func sendEth() (txHash common.Hash) {
	transactor, _ := eth.GetTransactor()
	transactor.Value = big.NewInt(1)
	balanceBefore, _ := eth.Ethcall.CheckETHBalance(&bind.CallOpts{Pending: true})
	tx, err := eth.Ethcall.Receive(transactor)
	if err != nil {
		fmt.Println("sendEth: ", err)
		return
	}
	balanceAfter, _ := eth.Ethcall.CheckETHBalance(&bind.CallOpts{Pending: true})
	fmt.Printf("sendEth tx %s balanceBefore %d, balanceAfter %d\n", tx.Hash(), balanceBefore, balanceAfter)
	return tx.Hash()
}

func ethCall() {
	var out []interface{}
	err := eth.Ethcall.Contract.Call(&bind.CallOpts{Pending: true}, &out, "sendETHBack", big.NewInt(1))
	if err != nil {
		fmt.Println("eth call error: ", err)
		return
	}
	result := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	fmt.Println("eth_call result: balance after ", result)
}

func sendEthBackSimulation() {
	transactor, _ := eth.GetTransactor()
	transactor.GasLimit = 0
	_, err := eth.Ethcall.SendETHBack(transactor, big.NewInt(1))
	if err != nil {
		fmt.Printf("sendEthBackSimulation error %s\n", err)
	} else {
		fmt.Println("sendEthBackSimulation success!")
	}
}

func sendEthBack() {
	transactor, _ := eth.GetTransactor()
	balanceBefore, _ := eth.Ethcall.CheckETHBalance(&bind.CallOpts{Pending: true})
	tx, err := eth.Ethcall.SendETHBack(transactor, big.NewInt(1))
	if err != nil {
		fmt.Println("sendEthBack: ", err)
	}
	balanceAfter, _ := eth.Ethcall.CheckETHBalance(&bind.CallOpts{Pending: true})
	fmt.Printf("sendEthBack tx %s balanceBefore %d, balanceAfter %d\n", tx.Hash(), balanceBefore, balanceAfter)
}
