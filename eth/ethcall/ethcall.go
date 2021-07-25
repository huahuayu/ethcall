// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcall

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthcallABI is the input ABI used to generate the binding from.
const EthcallABI = "[{\"inputs\":[],\"name\":\"checkETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendETHBack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// EthcallBin is the compiled bytecode used for deploying new contracts.
var EthcallBin = "0x608060405234801561001057600080fd5b50610279806100206000396000f3fe60806040526004361061002d5760003560e01c8063131b1522146100395780633b6216d51461006457610034565b3661003457005b600080fd5b34801561004557600080fd5b5061004e6100a1565b60405161005b91906101cd565b60405180910390f35b34801561007057600080fd5b5061008b60048036038101906100869190610152565b6100a9565b60405161009891906101cd565b60405180910390f35b600047905090565b6000814710156100ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e5906101ad565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610134573d6000803e3d6000fd5b50819050919050565b60008135905061014c8161022c565b92915050565b60006020828403121561016457600080fd5b60006101728482850161013d565b91505092915050565b60006101886014836101e8565b915061019382610203565b602082019050919050565b6101a7816101f9565b82525050565b600060208201905081810360008301526101c68161017b565b9050919050565b60006020820190506101e2600083018461019e565b92915050565b600082825260208201905092915050565b6000819050919050565b7f696e73756666696369656e742062616c616e6365000000000000000000000000600082015250565b610235816101f9565b811461024057600080fd5b5056fea26469706673582212200c80d3bde75c05ca52adc64077565285187d95907d0abb1e778aece08c6a256e64736f6c63430008040033"

// DeployEthcall deploys a new Ethereum contract, binding an instance of Ethcall to it.
func DeployEthcall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethcall, error) {
	parsed, err := abi.JSON(strings.NewReader(EthcallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthcallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethcall{EthcallCaller: EthcallCaller{contract: contract}, EthcallTransactor: EthcallTransactor{contract: contract}, EthcallFilterer: EthcallFilterer{contract: contract}}, nil
}

// Ethcall is an auto generated Go binding around an Ethereum contract.
type Ethcall struct {
	EthcallCaller     // Read-only binding to the contract
	EthcallTransactor // Write-only binding to the contract
	EthcallFilterer   // Log filterer for contract events
}

// EthcallCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthcallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthcallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthcallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthcallSession struct {
	Contract     *Ethcall          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthcallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthcallCallerSession struct {
	Contract *EthcallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EthcallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthcallTransactorSession struct {
	Contract     *EthcallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthcallRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthcallRaw struct {
	Contract *Ethcall // Generic contract binding to access the raw methods on
}

// EthcallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthcallCallerRaw struct {
	Contract *EthcallCaller // Generic read-only contract binding to access the raw methods on
}

// EthcallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthcallTransactorRaw struct {
	Contract *EthcallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthcall creates a new instance of Ethcall, bound to a specific deployed contract.
func NewEthcall(address common.Address, backend bind.ContractBackend) (*Ethcall, error) {
	contract, err := bindEthcall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethcall{EthcallCaller: EthcallCaller{contract: contract}, EthcallTransactor: EthcallTransactor{contract: contract}, EthcallFilterer: EthcallFilterer{contract: contract}}, nil
}

// NewEthcallCaller creates a new read-only instance of Ethcall, bound to a specific deployed contract.
func NewEthcallCaller(address common.Address, caller bind.ContractCaller) (*EthcallCaller, error) {
	contract, err := bindEthcall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthcallCaller{contract: contract}, nil
}

// NewEthcallTransactor creates a new write-only instance of Ethcall, bound to a specific deployed contract.
func NewEthcallTransactor(address common.Address, transactor bind.ContractTransactor) (*EthcallTransactor, error) {
	contract, err := bindEthcall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthcallTransactor{contract: contract}, nil
}

// NewEthcallFilterer creates a new log filterer instance of Ethcall, bound to a specific deployed contract.
func NewEthcallFilterer(address common.Address, filterer bind.ContractFilterer) (*EthcallFilterer, error) {
	contract, err := bindEthcall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthcallFilterer{contract: contract}, nil
}

// bindEthcall binds a generic wrapper to an already deployed contract.
func bindEthcall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthcallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethcall *EthcallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethcall.Contract.EthcallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethcall *EthcallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcall.Contract.EthcallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethcall *EthcallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethcall.Contract.EthcallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethcall *EthcallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethcall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethcall *EthcallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethcall *EthcallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethcall.Contract.contract.Transact(opts, method, params...)
}

// CheckETHBalance is a free data retrieval call binding the contract method 0x131b1522.
//
// Solidity: function checkETHBalance() view returns(uint256)
func (_Ethcall *EthcallCaller) CheckETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethcall.contract.Call(opts, &out, "checkETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckETHBalance is a free data retrieval call binding the contract method 0x131b1522.
//
// Solidity: function checkETHBalance() view returns(uint256)
func (_Ethcall *EthcallSession) CheckETHBalance() (*big.Int, error) {
	return _Ethcall.Contract.CheckETHBalance(&_Ethcall.CallOpts)
}

// CheckETHBalance is a free data retrieval call binding the contract method 0x131b1522.
//
// Solidity: function checkETHBalance() view returns(uint256)
func (_Ethcall *EthcallCallerSession) CheckETHBalance() (*big.Int, error) {
	return _Ethcall.Contract.CheckETHBalance(&_Ethcall.CallOpts)
}

// SendETHBack is a paid mutator transaction binding the contract method 0x3b6216d5.
//
// Solidity: function sendETHBack(uint256 amount) returns(uint256)
func (_Ethcall *EthcallTransactor) SendETHBack(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ethcall.contract.Transact(opts, "sendETHBack", amount)
}

// SendETHBack is a paid mutator transaction binding the contract method 0x3b6216d5.
//
// Solidity: function sendETHBack(uint256 amount) returns(uint256)
func (_Ethcall *EthcallSession) SendETHBack(amount *big.Int) (*types.Transaction, error) {
	return _Ethcall.Contract.SendETHBack(&_Ethcall.TransactOpts, amount)
}

// SendETHBack is a paid mutator transaction binding the contract method 0x3b6216d5.
//
// Solidity: function sendETHBack(uint256 amount) returns(uint256)
func (_Ethcall *EthcallTransactorSession) SendETHBack(amount *big.Int) (*types.Transaction, error) {
	return _Ethcall.Contract.SendETHBack(&_Ethcall.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethcall *EthcallTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcall.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethcall *EthcallSession) Receive() (*types.Transaction, error) {
	return _Ethcall.Contract.Receive(&_Ethcall.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethcall *EthcallTransactorSession) Receive() (*types.Transaction, error) {
	return _Ethcall.Contract.Receive(&_Ethcall.TransactOpts)
}
