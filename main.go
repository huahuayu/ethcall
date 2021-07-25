package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/huahuayu/ethcall/eth"
	"math/big"
	"time"
)

func main() {
	eth.Init()
	fmt.Println("sleep for 20 seconds to make sure test contract deployed")
	time.Sleep(20 * time.Second)
	sendEth()
	sendEthBackSimulation() // simulation failed
	sendEthBack()           // but actually tx execution success
}

func sendEth() {
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
