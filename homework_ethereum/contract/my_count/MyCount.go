// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package my_count

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyCountMetaData contains all meta data concerning the MyCount contract.
var MyCountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102818061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806345f43dd8146100385780634f0cd27b14610054575b5f5ffd5b610052600480360381019061004d919061017b565b610084565b005b61006e6004803603810190610069919061017b565b6100d8565b60405161007b91906101be565b60405180910390f35b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8154809291906100d090610204565b919050555050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014a82610121565b9050919050565b61015a81610140565b8114610164575f5ffd5b50565b5f8135905061017581610151565b92915050565b5f602082840312156101905761018f61011d565b5b5f61019d84828501610167565b91505092915050565b5f819050919050565b6101b8816101a6565b82525050565b5f6020820190506101d15f8301846101af565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61020e826101a6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036102405761023f6101d7565b5b60018201905091905056fea2646970667358221220ff329c73411425478a6812ee58d2a80a71101b7817942c48910c3b56c2b485f364736f6c634300081e0033",
}

// MyCountABI is the input ABI used to generate the binding from.
// Deprecated: Use MyCountMetaData.ABI instead.
var MyCountABI = MyCountMetaData.ABI

// MyCountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyCountMetaData.Bin instead.
var MyCountBin = MyCountMetaData.Bin

// DeployMyCount deploys a new Ethereum contract, binding an instance of MyCount to it.
func DeployMyCount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyCount, error) {
	parsed, err := MyCountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyCountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyCount{MyCountCaller: MyCountCaller{contract: contract}, MyCountTransactor: MyCountTransactor{contract: contract}, MyCountFilterer: MyCountFilterer{contract: contract}}, nil
}

// MyCount is an auto generated Go binding around an Ethereum contract.
type MyCount struct {
	MyCountCaller     // Read-only binding to the contract
	MyCountTransactor // Write-only binding to the contract
	MyCountFilterer   // Log filterer for contract events
}

// MyCountCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyCountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyCountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyCountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyCountSession struct {
	Contract     *MyCount          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyCountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyCountCallerSession struct {
	Contract *MyCountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyCountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyCountTransactorSession struct {
	Contract     *MyCountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyCountRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyCountRaw struct {
	Contract *MyCount // Generic contract binding to access the raw methods on
}

// MyCountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyCountCallerRaw struct {
	Contract *MyCountCaller // Generic read-only contract binding to access the raw methods on
}

// MyCountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyCountTransactorRaw struct {
	Contract *MyCountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyCount creates a new instance of MyCount, bound to a specific deployed contract.
func NewMyCount(address common.Address, backend bind.ContractBackend) (*MyCount, error) {
	contract, err := bindMyCount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyCount{MyCountCaller: MyCountCaller{contract: contract}, MyCountTransactor: MyCountTransactor{contract: contract}, MyCountFilterer: MyCountFilterer{contract: contract}}, nil
}

// NewMyCountCaller creates a new read-only instance of MyCount, bound to a specific deployed contract.
func NewMyCountCaller(address common.Address, caller bind.ContractCaller) (*MyCountCaller, error) {
	contract, err := bindMyCount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyCountCaller{contract: contract}, nil
}

// NewMyCountTransactor creates a new write-only instance of MyCount, bound to a specific deployed contract.
func NewMyCountTransactor(address common.Address, transactor bind.ContractTransactor) (*MyCountTransactor, error) {
	contract, err := bindMyCount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyCountTransactor{contract: contract}, nil
}

// NewMyCountFilterer creates a new log filterer instance of MyCount, bound to a specific deployed contract.
func NewMyCountFilterer(address common.Address, filterer bind.ContractFilterer) (*MyCountFilterer, error) {
	contract, err := bindMyCount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyCountFilterer{contract: contract}, nil
}

// bindMyCount binds a generic wrapper to an already deployed contract.
func bindMyCount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyCountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyCount *MyCountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyCount.Contract.MyCountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyCount *MyCountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyCount.Contract.MyCountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyCount *MyCountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyCount.Contract.MyCountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyCount *MyCountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyCount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyCount *MyCountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyCount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyCount *MyCountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyCount.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(address userAddress) view returns(uint256)
func (_MyCount *MyCountCaller) GetCount(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyCount.contract.Call(opts, &out, "getCount", userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(address userAddress) view returns(uint256)
func (_MyCount *MyCountSession) GetCount(userAddress common.Address) (*big.Int, error) {
	return _MyCount.Contract.GetCount(&_MyCount.CallOpts, userAddress)
}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(address userAddress) view returns(uint256)
func (_MyCount *MyCountCallerSession) GetCount(userAddress common.Address) (*big.Int, error) {
	return _MyCount.Contract.GetCount(&_MyCount.CallOpts, userAddress)
}

// Increment is a paid mutator transaction binding the contract method 0x45f43dd8.
//
// Solidity: function increment(address userAddress) returns()
func (_MyCount *MyCountTransactor) Increment(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _MyCount.contract.Transact(opts, "increment", userAddress)
}

// Increment is a paid mutator transaction binding the contract method 0x45f43dd8.
//
// Solidity: function increment(address userAddress) returns()
func (_MyCount *MyCountSession) Increment(userAddress common.Address) (*types.Transaction, error) {
	return _MyCount.Contract.Increment(&_MyCount.TransactOpts, userAddress)
}

// Increment is a paid mutator transaction binding the contract method 0x45f43dd8.
//
// Solidity: function increment(address userAddress) returns()
func (_MyCount *MyCountTransactorSession) Increment(userAddress common.Address) (*types.Transaction, error) {
	return _MyCount.Contract.Increment(&_MyCount.TransactOpts, userAddress)
}
