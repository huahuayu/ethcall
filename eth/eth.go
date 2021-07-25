package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huahuayu/ethcall/eth/ethcall"
	"log"
	"math/big"
)

var (
	EthClient     *ethclient.Client
	Ethcall       *ethcall.Ethcall
	TransactorKey *ecdsa.PrivateKey
	FromAddress   common.Address
	ChainId       *big.Int
)

func Init() {
	var err error
	if EthClient, err = ethclient.DialContext(context.Background(), "wss://ropsten.infura.io/ws/v3/e9d43fcc8b60466c9b8c6c5b8215475c"); err != nil {
		log.Fatal("ethClient init: ", err)
	}
	if TransactorKey, err = crypto.HexToECDSA("806d0af138598ec82da4c28da4de7853ac336b2138917c7ce7c000f95b58b333"); err != nil {
		log.Fatal(err)
	} else {
		publicKey, _ := TransactorKey.Public().(*ecdsa.PublicKey)
		FromAddress = crypto.PubkeyToAddress(*publicKey) // 0x7E59f246E1A8A939ac7DFDc4F40D2F03b11D88f6
	}
	if ChainId, err = EthClient.ChainID(context.Background()); err != nil {
		log.Fatal(err)
	}
	transactor, err := GetTransactor()
	if err != nil {
		log.Fatal("GetTransactor: ", err)
	}
	contract, _, _, err := ethcall.DeployEthcall(transactor, EthClient)
	if err != nil {
		return
	}
	fmt.Println("test contract deployed: ", contract.String())
	if Ethcall, err = ethcall.NewEthcall(contract, EthClient); err != nil {
		log.Fatal("contract init: ", err)
	}
}

func GetTransactor() (transactor *bind.TransactOpts, err error) {
	transactor, _ = bind.NewKeyedTransactorWithChainID(TransactorKey, ChainId)
	transactor.Context = context.Background()
	transactor.GasPrice = big.NewInt(10000000000)
	transactor.GasLimit = 1000000
	transactor.From = FromAddress
	nonce, _ := EthClient.PendingNonceAt(context.Background(), FromAddress)
	transactor.Nonce = big.NewInt(int64(nonce))
	return transactor, nil
}
