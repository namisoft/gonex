// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pairex

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataSetABI is the input ABI used to generate the binding from.
const DataSetABI = "[]"

// DataSetBin is the compiled bytecode used for deploying new contracts.
const DataSetBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a72305820b558898f0688da45e5673b02f86defe591ce6eb549dbd4b076da0a8317a5f2c20029`

// DeployDataSet deploys a new Ethereum contract, binding an instance of DataSet to it.
func DeployDataSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataSet, error) {
	parsed, err := abi.JSON(strings.NewReader(DataSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataSet{DataSetCaller: DataSetCaller{contract: contract}, DataSetTransactor: DataSetTransactor{contract: contract}, DataSetFilterer: DataSetFilterer{contract: contract}}, nil
}

// DataSet is an auto generated Go binding around an Ethereum contract.
type DataSet struct {
	DataSetCaller     // Read-only binding to the contract
	DataSetTransactor // Write-only binding to the contract
	DataSetFilterer   // Log filterer for contract events
}

// DataSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSetSession struct {
	Contract     *DataSet          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataSetCallerSession struct {
	Contract *DataSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DataSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataSetTransactorSession struct {
	Contract     *DataSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DataSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataSetRaw struct {
	Contract *DataSet // Generic contract binding to access the raw methods on
}

// DataSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataSetCallerRaw struct {
	Contract *DataSetCaller // Generic read-only contract binding to access the raw methods on
}

// DataSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataSetTransactorRaw struct {
	Contract *DataSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataSet creates a new instance of DataSet, bound to a specific deployed contract.
func NewDataSet(address common.Address, backend bind.ContractBackend) (*DataSet, error) {
	contract, err := bindDataSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataSet{DataSetCaller: DataSetCaller{contract: contract}, DataSetTransactor: DataSetTransactor{contract: contract}, DataSetFilterer: DataSetFilterer{contract: contract}}, nil
}

// NewDataSetCaller creates a new read-only instance of DataSet, bound to a specific deployed contract.
func NewDataSetCaller(address common.Address, caller bind.ContractCaller) (*DataSetCaller, error) {
	contract, err := bindDataSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataSetCaller{contract: contract}, nil
}

// NewDataSetTransactor creates a new write-only instance of DataSet, bound to a specific deployed contract.
func NewDataSetTransactor(address common.Address, transactor bind.ContractTransactor) (*DataSetTransactor, error) {
	contract, err := bindDataSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataSetTransactor{contract: contract}, nil
}

// NewDataSetFilterer creates a new log filterer instance of DataSet, bound to a specific deployed contract.
func NewDataSetFilterer(address common.Address, filterer bind.ContractFilterer) (*DataSetFilterer, error) {
	contract, err := bindDataSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataSetFilterer{contract: contract}, nil
}

// bindDataSet binds a generic wrapper to an already deployed contract.
func bindDataSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataSet *DataSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataSet.Contract.DataSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataSet *DataSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataSet.Contract.DataSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataSet *DataSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataSet.Contract.DataSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataSet *DataSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataSet *DataSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataSet *DataSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataSet.Contract.contract.Transact(opts, method, params...)
}

// IOwnableERC223ABI is the input ABI used to generate the binding from.
const IOwnableERC223ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderbook\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnFromOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOwnableERC223Bin is the compiled bytecode used for deploying new contracts.
const IOwnableERC223Bin = `0x`

// DeployIOwnableERC223 deploys a new Ethereum contract, binding an instance of IOwnableERC223 to it.
func DeployIOwnableERC223(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IOwnableERC223, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnableERC223ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IOwnableERC223Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IOwnableERC223{IOwnableERC223Caller: IOwnableERC223Caller{contract: contract}, IOwnableERC223Transactor: IOwnableERC223Transactor{contract: contract}, IOwnableERC223Filterer: IOwnableERC223Filterer{contract: contract}}, nil
}

// IOwnableERC223 is an auto generated Go binding around an Ethereum contract.
type IOwnableERC223 struct {
	IOwnableERC223Caller     // Read-only binding to the contract
	IOwnableERC223Transactor // Write-only binding to the contract
	IOwnableERC223Filterer   // Log filterer for contract events
}

// IOwnableERC223Caller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnableERC223Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableERC223Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnableERC223Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableERC223Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnableERC223Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableERC223Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnableERC223Session struct {
	Contract     *IOwnableERC223   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnableERC223CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnableERC223CallerSession struct {
	Contract *IOwnableERC223Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOwnableERC223TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnableERC223TransactorSession struct {
	Contract     *IOwnableERC223Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOwnableERC223Raw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnableERC223Raw struct {
	Contract *IOwnableERC223 // Generic contract binding to access the raw methods on
}

// IOwnableERC223CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnableERC223CallerRaw struct {
	Contract *IOwnableERC223Caller // Generic read-only contract binding to access the raw methods on
}

// IOwnableERC223TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnableERC223TransactorRaw struct {
	Contract *IOwnableERC223Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwnableERC223 creates a new instance of IOwnableERC223, bound to a specific deployed contract.
func NewIOwnableERC223(address common.Address, backend bind.ContractBackend) (*IOwnableERC223, error) {
	contract, err := bindIOwnableERC223(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwnableERC223{IOwnableERC223Caller: IOwnableERC223Caller{contract: contract}, IOwnableERC223Transactor: IOwnableERC223Transactor{contract: contract}, IOwnableERC223Filterer: IOwnableERC223Filterer{contract: contract}}, nil
}

// NewIOwnableERC223Caller creates a new read-only instance of IOwnableERC223, bound to a specific deployed contract.
func NewIOwnableERC223Caller(address common.Address, caller bind.ContractCaller) (*IOwnableERC223Caller, error) {
	contract, err := bindIOwnableERC223(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableERC223Caller{contract: contract}, nil
}

// NewIOwnableERC223Transactor creates a new write-only instance of IOwnableERC223, bound to a specific deployed contract.
func NewIOwnableERC223Transactor(address common.Address, transactor bind.ContractTransactor) (*IOwnableERC223Transactor, error) {
	contract, err := bindIOwnableERC223(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableERC223Transactor{contract: contract}, nil
}

// NewIOwnableERC223Filterer creates a new log filterer instance of IOwnableERC223, bound to a specific deployed contract.
func NewIOwnableERC223Filterer(address common.Address, filterer bind.ContractFilterer) (*IOwnableERC223Filterer, error) {
	contract, err := bindIOwnableERC223(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnableERC223Filterer{contract: contract}, nil
}

// bindIOwnableERC223 binds a generic wrapper to an already deployed contract.
func bindIOwnableERC223(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnableERC223ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnableERC223 *IOwnableERC223Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOwnableERC223.Contract.IOwnableERC223Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnableERC223 *IOwnableERC223Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.IOwnableERC223Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnableERC223 *IOwnableERC223Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.IOwnableERC223Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnableERC223 *IOwnableERC223CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOwnableERC223.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnableERC223 *IOwnableERC223TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnableERC223 *IOwnableERC223TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.contract.Transact(opts, method, params...)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223Transactor) BurnFromOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.contract.Transact(opts, "burnFromOwner", _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223Session) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.BurnFromOwner(&_IOwnableERC223.TransactOpts, _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223TransactorSession) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.BurnFromOwner(&_IOwnableERC223.TransactOpts, _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223Transactor) MintToOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.contract.Transact(opts, "mintToOwner", _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223Session) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.MintToOwner(&_IOwnableERC223.TransactOpts, _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_IOwnableERC223 *IOwnableERC223TransactorSession) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.MintToOwner(&_IOwnableERC223.TransactOpts, _amount)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address _orderbook) returns()
func (_IOwnableERC223 *IOwnableERC223Transactor) Setup(opts *bind.TransactOpts, _orderbook common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.contract.Transact(opts, "setup", _orderbook)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address _orderbook) returns()
func (_IOwnableERC223 *IOwnableERC223Session) Setup(_orderbook common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Setup(&_IOwnableERC223.TransactOpts, _orderbook)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address _orderbook) returns()
func (_IOwnableERC223 *IOwnableERC223TransactorSession) Setup(_orderbook common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Setup(&_IOwnableERC223.TransactOpts, _orderbook)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IOwnableERC223 *IOwnableERC223Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IOwnableERC223 *IOwnableERC223Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Transfer(&_IOwnableERC223.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IOwnableERC223 *IOwnableERC223TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Transfer(&_IOwnableERC223.TransactOpts, to, value)
}

// InitializerABI is the input ABI used to generate the binding from.
const InitializerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"Stable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Volatile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// InitializerBin is the compiled bytecode used for deploying new contracts.
const InitializerBin = `0x608060405234801561001057600080fd5b50610367806100206000396000f3fe608060405234801561001057600080fd5b506004361061007e577c010000000000000000000000000000000000000000000000000000000060003504639157772a8114610083578063974c86b514610083578063aea90cb81461009f578063c0409a5c146100a7578063f2a946b41461009f578063fdeb1779146100dc575b600080fd5b61008b61010f565b604080519115158252519081900360200190f35b61008b610114565b6100da600480360360208110156100bd57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610119565b005b6100da600480360360208110156100f257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610228565b600181565b600081565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb55473ffffffffffffffffffffffffffffffffffffffff16156101c557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d5473ffffffffffffffffffffffffffffffffffffffff16156102d657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff9290921691909117905556fea165627a7a72305820a3f33bde2f34d6bd7497786f0199ed04782f54799d924e89596e0e226e6037ef0029`

// DeployInitializer deploys a new Ethereum contract, binding an instance of Initializer to it.
func DeployInitializer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Initializer, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InitializerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Initializer{InitializerCaller: InitializerCaller{contract: contract}, InitializerTransactor: InitializerTransactor{contract: contract}, InitializerFilterer: InitializerFilterer{contract: contract}}, nil
}

// Initializer is an auto generated Go binding around an Ethereum contract.
type Initializer struct {
	InitializerCaller     // Read-only binding to the contract
	InitializerTransactor // Write-only binding to the contract
	InitializerFilterer   // Log filterer for contract events
}

// InitializerCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializerSession struct {
	Contract     *Initializer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializerCallerSession struct {
	Contract *InitializerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InitializerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializerTransactorSession struct {
	Contract     *InitializerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InitializerRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializerRaw struct {
	Contract *Initializer // Generic contract binding to access the raw methods on
}

// InitializerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializerCallerRaw struct {
	Contract *InitializerCaller // Generic read-only contract binding to access the raw methods on
}

// InitializerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializerTransactorRaw struct {
	Contract *InitializerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializer creates a new instance of Initializer, bound to a specific deployed contract.
func NewInitializer(address common.Address, backend bind.ContractBackend) (*Initializer, error) {
	contract, err := bindInitializer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializer{InitializerCaller: InitializerCaller{contract: contract}, InitializerTransactor: InitializerTransactor{contract: contract}, InitializerFilterer: InitializerFilterer{contract: contract}}, nil
}

// NewInitializerCaller creates a new read-only instance of Initializer, bound to a specific deployed contract.
func NewInitializerCaller(address common.Address, caller bind.ContractCaller) (*InitializerCaller, error) {
	contract, err := bindInitializer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializerCaller{contract: contract}, nil
}

// NewInitializerTransactor creates a new write-only instance of Initializer, bound to a specific deployed contract.
func NewInitializerTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializerTransactor, error) {
	contract, err := bindInitializer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializerTransactor{contract: contract}, nil
}

// NewInitializerFilterer creates a new log filterer instance of Initializer, bound to a specific deployed contract.
func NewInitializerFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializerFilterer, error) {
	contract, err := bindInitializer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializerFilterer{contract: contract}, nil
}

// bindInitializer binds a generic wrapper to an already deployed contract.
func bindInitializer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializer *InitializerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializer.Contract.InitializerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializer *InitializerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializer.Contract.InitializerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializer *InitializerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializer.Contract.InitializerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializer *InitializerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializer *InitializerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializer *InitializerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializer.Contract.contract.Transact(opts, method, params...)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_Initializer *InitializerCaller) Buy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Initializer.contract.Call(opts, out, "Buy")
	return *ret0, err
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_Initializer *InitializerSession) Buy() (bool, error) {
	return _Initializer.Contract.Buy(&_Initializer.CallOpts)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_Initializer *InitializerCallerSession) Buy() (bool, error) {
	return _Initializer.Contract.Buy(&_Initializer.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_Initializer *InitializerCaller) Sell(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Initializer.contract.Call(opts, out, "Sell")
	return *ret0, err
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_Initializer *InitializerSession) Sell() (bool, error) {
	return _Initializer.Contract.Sell(&_Initializer.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_Initializer *InitializerCallerSession) Sell() (bool, error) {
	return _Initializer.Contract.Sell(&_Initializer.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_Initializer *InitializerCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Initializer.contract.Call(opts, out, "Stable")
	return *ret0, err
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_Initializer *InitializerSession) Stable() (bool, error) {
	return _Initializer.Contract.Stable(&_Initializer.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_Initializer *InitializerCallerSession) Stable() (bool, error) {
	return _Initializer.Contract.Stable(&_Initializer.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_Initializer *InitializerCaller) Volatile(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Initializer.contract.Call(opts, out, "Volatile")
	return *ret0, err
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_Initializer *InitializerSession) Volatile() (bool, error) {
	return _Initializer.Contract.Volatile(&_Initializer.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_Initializer *InitializerCallerSession) Volatile() (bool, error) {
	return _Initializer.Contract.Volatile(&_Initializer.CallOpts)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_Initializer *InitializerTransactor) StableTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Initializer.contract.Transact(opts, "stableTokenRegister", _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_Initializer *InitializerSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _Initializer.Contract.StableTokenRegister(&_Initializer.TransactOpts, _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_Initializer *InitializerTransactorSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _Initializer.Contract.StableTokenRegister(&_Initializer.TransactOpts, _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_Initializer *InitializerTransactor) VolatileTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Initializer.contract.Transact(opts, "volatileTokenRegister", _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_Initializer *InitializerSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _Initializer.Contract.VolatileTokenRegister(&_Initializer.TransactOpts, _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_Initializer *InitializerTransactorSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _Initializer.Contract.VolatileTokenRegister(&_Initializer.TransactOpts, _address)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582047c688b08c825b2396b83695b6c494ce2ab0cbb4e720c36885697c8b613f9ed80029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OrderBookABI is the input ABI used to generate the binding from.
const OrderBookABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"bottom\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"validOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Stable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"}],\"name\":\"initOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INPUTS_MAX\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderType\",\"type\":\"bool\"},{\"name\":\"_newId\",\"type\":\"bytes32\"},{\"name\":\"_oldId\",\"type\":\"bytes32\"}],\"name\":\"betterOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Volatile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrderBookBin is the compiled bytecode used for deploying new contracts.
const OrderBookBin = `0x60806040526000600e556001600f55610e3f8061001d6000396000f3fe608060405234801561001057600080fd5b5060043610610128576000357c010000000000000000000000000000000000000000000000000000000090048063bc96b872116100bf578063d39ecc951161008e578063d39ecc95146102f1578063d4dcff8214610331578063e3dd138f14610339578063f2a946b41461023f578063fdeb17791461036457610128565b8063bc96b87214610247578063c0409a5c14610281578063ce11ae66146102a7578063ce24388f146102cc57610128565b80638aa3f897116100fb5780638aa3f897146102185780639157772a14610237578063974c86b514610237578063aea90cb81461023f57610128565b806307c399a31461012d578063174a50881461018757806343271d79146101b85780637234c79d146101df575b600080fd5b6101526004803603604081101561014357600080fd5b5080351515906020013561038a565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101a66004803603602081101561019d57600080fd5b503515156103d0565b60408051918252519081900360200190f35b6101dd600480360360408110156101ce57600080fd5b508035151590602001356103f2565b005b610204600480360360408110156101f557600080fd5b50803515159060200135610484565b604080519115158252519081900360200190f35b6101a66004803603602081101561022e57600080fd5b503515156104ac565b6102046104ce565b6102046104d3565b6101a66004803603608081101561025d57600080fd5b50600160a060020a03813516906020810135151590604081013590606001356104d8565b6101dd6004803603602081101561029757600080fd5b5035600160a060020a0316610748565b6101a6600480360360408110156102bd57600080fd5b50803515159060200135610826565b6101a6600480360360408110156102e257600080fd5b50803515159060200135610848565b6101a6600480360360a081101561030757600080fd5b508035151590602081013590604081013590600160a060020a03606082013516906080013561086a565b6101a66109c5565b6102046004803603606081101561034f57600080fd5b508035151590602081013590604001356109d9565b6101dd6004803603602081101561037a57600080fd5b5035600160a060020a0316610a65565b9015156000908152600c60209081526040808320938352929052208054600182015460028301546003840154600490940154600160a060020a0390931694919390929091565b15156000908152600c6020908152604080832083805290915290206003015490565b8115156000908152600c60209081526040808320848452918290529091208054600160a060020a03163314610471576040805160e560020a62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b61047d84846000610b47565b5050505050565b8115156000908152600c60209081526040808320848452909152812060020154115b92915050565b15156000908152600c6020908152604080832083805290915290206004015490565b600181565b600081565b600080831180156104e95750600082115b151561053f576040805160e560020a62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b6f010000000000000000000000000000008310801561056d57506f0100000000000000000000000000000082105b15156105c3576040805160e560020a62461bcd02815260206004820152601460248201527f67726561746572207468616e20737570706c793f000000000000000000000000604482015290519081900360640190fd5b600160a060020a0385166000818152600d602081815260408084208054600101908190558915158552600c835281852086865293835281516c01000000000000000000000000909602868401526034860152605485018890526074808601889052815180870390910181526094909501908190528451929460029390928291908401908083835b602083106106695780518252601f19909201916020918201910161064a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156106a8573d6000803e3d6000fd5b5050506040513d60208110156106bd57600080fd5b50516040805160a081018252600160a060020a03998a16815260208082019889528183019788526000606083018181526080840182815286835297909252929092209051815473ffffffffffffffffffffffffffffffffffffffff19169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb554600160a060020a0316156107d0576040805160e560020a62461bcd02815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b9015156000908152600c60209081526040808320938352929052206003015490565b9015156000908152600c60209081526040808320938352929052206004015490565b60006108768683610484565b15156108cc576040805160e560020a62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b8515156000908152600c60205260408120906108ea858989896104d8565b9050600084156108fa5784610903565b610903896104ac565b90506109108983836109d9565b1515610954575b6109228983836109d9565b151561093f57600090815260208390526040902060040154610917565b61094a898383610d44565b5091506109bc9050565b6000908152602083905260409020600301545b801580159061097c575061097c8983836109d9565b1561099857600090815260208390526040902060030154610967565b6000818152602084905260409020600401546109b7908a908490610d44565b509150505b95945050505050565b6f0100000000000000000000000000000081565b8215156000908152600c6020908152604080832085845291829052808320848452908320600f54600282015460018401548692610a2d929091610a219163ffffffff610d8b16565b9063ffffffff610d8b16565b90506000610a56600e54600f5401610a2186600201548660010154610d8b90919063ffffffff16565b90911198975050505050505050565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d54600160a060020a031615610aef576040805160e560020a62461bcd02815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b8215156000908152600c6020908152604080832085845291829052808320600480820154600380840180548852858820909301829055915490865292852001919091558315610c43578515600090815260208181526040808320548454600286015483517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610c1157600080fd5b505af1158015610c25573d6000803e3d6000fd5b505050506040513d6020811015610c3b57600080fd5b50610cf29050565b851515600090815260208181526040808320548454600186015483517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610cc557600080fd5b505af1158015610cd9573d6000803e3d6000fd5b505050506040513d6020811015610cef57600080fd5b50505b600490810154600095865260209290925260408520805473ffffffffffffffffffffffffffffffffffffffff191681556001810186905560028101869055600381018690550193909355509092915050565b9115156000908152600c6020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b6000821515610d9c575060006104a6565b828202828482811515610dab57fe5b0414610deb5760405160e560020a62461bcd028152600401808060200182810382526021815260200180610df36021913960400191505060405180910390fd5b939250505056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a165627a7a72305820ec4905c86d658fccc6875d14136992aafb609ea60f99b0ee7e9fe73aee1945bd0029`

// DeployOrderBook deploys a new Ethereum contract, binding an instance of OrderBook to it.
func DeployOrderBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderBook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderBookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// OrderBook is an auto generated Go binding around an Ethereum contract.
type OrderBook struct {
	OrderBookCaller     // Read-only binding to the contract
	OrderBookTransactor // Write-only binding to the contract
	OrderBookFilterer   // Log filterer for contract events
}

// OrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBookSession struct {
	Contract     *OrderBook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBookCallerSession struct {
	Contract *OrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBookTransactorSession struct {
	Contract     *OrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBookRaw struct {
	Contract *OrderBook // Generic contract binding to access the raw methods on
}

// OrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBookCallerRaw struct {
	Contract *OrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBookTransactorRaw struct {
	Contract *OrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBook creates a new instance of OrderBook, bound to a specific deployed contract.
func NewOrderBook(address common.Address, backend bind.ContractBackend) (*OrderBook, error) {
	contract, err := bindOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// NewOrderBookCaller creates a new read-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookCaller(address common.Address, caller bind.ContractCaller) (*OrderBookCaller, error) {
	contract, err := bindOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookCaller{contract: contract}, nil
}

// NewOrderBookTransactor creates a new write-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBookTransactor, error) {
	contract, err := bindOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookTransactor{contract: contract}, nil
}

// NewOrderBookFilterer creates a new log filterer instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBookFilterer, error) {
	contract, err := bindOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBookFilterer{contract: contract}, nil
}

// bindOrderBook binds a generic wrapper to an already deployed contract.
func bindOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderBookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.OrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transact(opts, method, params...)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_OrderBook *OrderBookCaller) Buy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "Buy")
	return *ret0, err
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_OrderBook *OrderBookSession) Buy() (bool, error) {
	return _OrderBook.Contract.Buy(&_OrderBook.CallOpts)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_OrderBook *OrderBookCallerSession) Buy() (bool, error) {
	return _OrderBook.Contract.Buy(&_OrderBook.CallOpts)
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_OrderBook *OrderBookCaller) INPUTSMAX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "INPUTS_MAX")
	return *ret0, err
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_OrderBook *OrderBookSession) INPUTSMAX() (*big.Int, error) {
	return _OrderBook.Contract.INPUTSMAX(&_OrderBook.CallOpts)
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_OrderBook *OrderBookCallerSession) INPUTSMAX() (*big.Int, error) {
	return _OrderBook.Contract.INPUTSMAX(&_OrderBook.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_OrderBook *OrderBookCaller) Sell(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "Sell")
	return *ret0, err
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_OrderBook *OrderBookSession) Sell() (bool, error) {
	return _OrderBook.Contract.Sell(&_OrderBook.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_OrderBook *OrderBookCallerSession) Sell() (bool, error) {
	return _OrderBook.Contract.Sell(&_OrderBook.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_OrderBook *OrderBookCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "Stable")
	return *ret0, err
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_OrderBook *OrderBookSession) Stable() (bool, error) {
	return _OrderBook.Contract.Stable(&_OrderBook.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_OrderBook *OrderBookCallerSession) Stable() (bool, error) {
	return _OrderBook.Contract.Stable(&_OrderBook.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_OrderBook *OrderBookCaller) Volatile(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "Volatile")
	return *ret0, err
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_OrderBook *OrderBookSession) Volatile() (bool, error) {
	return _OrderBook.Contract.Volatile(&_OrderBook.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_OrderBook *OrderBookCallerSession) Volatile() (bool, error) {
	return _OrderBook.Contract.Volatile(&_OrderBook.CallOpts)
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_OrderBook *OrderBookCaller) BetterOrder(opts *bind.CallOpts, orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "betterOrder", orderType, _newId, _oldId)
	return *ret0, err
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_OrderBook *OrderBookSession) BetterOrder(orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	return _OrderBook.Contract.BetterOrder(&_OrderBook.CallOpts, orderType, _newId, _oldId)
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_OrderBook *OrderBookCallerSession) BetterOrder(orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	return _OrderBook.Contract.BetterOrder(&_OrderBook.CallOpts, orderType, _newId, _oldId)
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookCaller) Bottom(opts *bind.CallOpts, _orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "bottom", _orderType)
	return *ret0, err
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookSession) Bottom(_orderType bool) ([32]byte, error) {
	return _OrderBook.Contract.Bottom(&_OrderBook.CallOpts, _orderType)
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookCallerSession) Bottom(_orderType bool) ([32]byte, error) {
	return _OrderBook.Contract.Bottom(&_OrderBook.CallOpts, _orderType)
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookCaller) GetNext(opts *bind.CallOpts, _orderType bool, _id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "getNext", _orderType, _id)
	return *ret0, err
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookSession) GetNext(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _OrderBook.Contract.GetNext(&_OrderBook.CallOpts, _orderType, _id)
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookCallerSession) GetNext(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _OrderBook.Contract.GetNext(&_OrderBook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_OrderBook *OrderBookCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _OrderBook.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_OrderBook *OrderBookSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _OrderBook.Contract.GetOrder(&_OrderBook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_OrderBook *OrderBookCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _OrderBook.Contract.GetOrder(&_OrderBook.CallOpts, _orderType, _id)
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookCaller) GetPrev(opts *bind.CallOpts, _orderType bool, _id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "getPrev", _orderType, _id)
	return *ret0, err
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookSession) GetPrev(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _OrderBook.Contract.GetPrev(&_OrderBook.CallOpts, _orderType, _id)
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_OrderBook *OrderBookCallerSession) GetPrev(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _OrderBook.Contract.GetPrev(&_OrderBook.CallOpts, _orderType, _id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookCaller) Top(opts *bind.CallOpts, _orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "top", _orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookSession) Top(_orderType bool) ([32]byte, error) {
	return _OrderBook.Contract.Top(&_OrderBook.CallOpts, _orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_OrderBook *OrderBookCallerSession) Top(_orderType bool) ([32]byte, error) {
	return _OrderBook.Contract.Top(&_OrderBook.CallOpts, _orderType)
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_OrderBook *OrderBookCaller) ValidOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrderBook.contract.Call(opts, out, "validOrder", _orderType, _id)
	return *ret0, err
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_OrderBook *OrderBookSession) ValidOrder(_orderType bool, _id [32]byte) (bool, error) {
	return _OrderBook.Contract.ValidOrder(&_OrderBook.CallOpts, _orderType, _id)
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_OrderBook *OrderBookCallerSession) ValidOrder(_orderType bool, _id [32]byte) (bool, error) {
	return _OrderBook.Contract.ValidOrder(&_OrderBook.CallOpts, _orderType, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookTransactor) Cancel(opts *bind.TransactOpts, _orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancel", _orderType, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookSession) Cancel(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Cancel(&_OrderBook.TransactOpts, _orderType, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookTransactorSession) Cancel(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Cancel(&_OrderBook.TransactOpts, _orderType, _id)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_OrderBook *OrderBookTransactor) InitOrder(opts *bind.TransactOpts, _maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "initOrder", _maker, _orderType, _haveAmount, _wantAmount)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_OrderBook *OrderBookSession) InitOrder(_maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.InitOrder(&_OrderBook.TransactOpts, _maker, _orderType, _haveAmount, _wantAmount)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_OrderBook *OrderBookTransactorSession) InitOrder(_maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.InitOrder(&_OrderBook.TransactOpts, _maker, _orderType, _haveAmount, _wantAmount)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_OrderBook *OrderBookTransactor) Insert(opts *bind.TransactOpts, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "insert", _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_OrderBook *OrderBookSession) Insert(_orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Insert(&_OrderBook.TransactOpts, _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_OrderBook *OrderBookTransactorSession) Insert(_orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Insert(&_OrderBook.TransactOpts, _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_OrderBook *OrderBookTransactor) StableTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "stableTokenRegister", _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_OrderBook *OrderBookSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.StableTokenRegister(&_OrderBook.TransactOpts, _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_OrderBook *OrderBookTransactorSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.StableTokenRegister(&_OrderBook.TransactOpts, _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_OrderBook *OrderBookTransactor) VolatileTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "volatileTokenRegister", _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_OrderBook *OrderBookSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.VolatileTokenRegister(&_OrderBook.TransactOpts, _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_OrderBook *OrderBookTransactorSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.VolatileTokenRegister(&_OrderBook.TransactOpts, _address)
}

// PairExABI is the input ABI used to generate the binding from.
const PairExABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"bottom\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_volatileTokenAddress\",\"type\":\"address\"},{\"name\":\"_stableTokenAddress\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrderType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"validOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlockPremptive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_stableTokenTarget\",\"type\":\"uint256\"}],\"name\":\"orderToFill\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasActivePremptive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Stable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inflate\",\"type\":\"bool\"},{\"name\":\"_stableTokenTarget\",\"type\":\"uint256\"}],\"name\":\"absorb\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"}],\"name\":\"initOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inflate\",\"type\":\"bool\"},{\"name\":\"_totalSTB\",\"type\":\"uint256\"},{\"name\":\"_totalVOL\",\"type\":\"uint256\"}],\"name\":\"premptiveAbsorb\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INPUTS_MAX\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPremptive\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderType\",\"type\":\"bool\"},{\"name\":\"_newId\",\"type\":\"bytes32\"},{\"name\":\"_oldId\",\"type\":\"bytes32\"}],\"name\":\"betterOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activatePremptive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Volatile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_volatileTokenAddress\",\"type\":\"address\"},{\"name\":\"_stableTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_initator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PremptiveActivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"PremptiveUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_stb\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_vol\",\"type\":\"uint256\"}],\"name\":\"PremptiveAbsorption\",\"type\":\"event\"}]"

// PairExBin is the compiled bytecode used for deploying new contracts.
const PairExBin = `0x60806040526000600e556001600f553480156200001b57600080fd5b5060405160408062002bf9833981018060405260408110156200003d57600080fd5b508051602090910151600160a060020a038216156200006a576200006a82640100000000620000a9810204565b600160a060020a038116156200008e576200008e8164010000000062000170810204565b620000a16401000000006200023b810204565b505062000465565b600080805260205260008051602062002bd983398151915254600160a060020a0316156200013857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080805260205260008051602062002bd98339815191528054600160a060020a031916600160a060020a0392909216919091179055565b6001600090815260205260008051602062002bb983398151915254600160a060020a0316156200020157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600090815260205260008051602062002bb98339815191528054600160a060020a031916600160a060020a0392909216919091179055565b6200024562000437565b600160408201908152600080527f13649b2456f1b42fef0f0040b3aaeabcd21a76a0f3f5defd4f583839455116e8602090815282517f18ea5137d4b638f7150ef5d67284baa6628dc5dcf2cde6ec4377e9be2774c55f8054600160a060020a0319908116600160a060020a039384161790915582850180517f18ea5137d4b638f7150ef5d67284baa6628dc5dcf2cde6ec4377e9be2774c5605584517f18ea5137d4b638f7150ef5d67284baa6628dc5dcf2cde6ec4377e9be2774c561556060860180517f18ea5137d4b638f7150ef5d67284baa6628dc5dcf2cde6ec4377e9be2774c562556080870180517f18ea5137d4b638f7150ef5d67284baa6628dc5dcf2cde6ec4377e9be2774c563557fd421a5181c571bba3f01190c922c3b2a896fc1d84e86c9f17ac10e67ebef8b5c90955295517fcc045a98e72e59344d4f7091f80bbb561222da6a20eec7c35a2c5e8d6ed5fd8380549093169316929092179055517fcc045a98e72e59344d4f7091f80bbb561222da6a20eec7c35a2c5e8d6ed5fd845590517fcc045a98e72e59344d4f7091f80bbb561222da6a20eec7c35a2c5e8d6ed5fd855590517fcc045a98e72e59344d4f7091f80bbb561222da6a20eec7c35a2c5e8d6ed5fd8655517fcc045a98e72e59344d4f7091f80bbb561222da6a20eec7c35a2c5e8d6ed5fd8755565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b61274480620004756000396000f3fe6080604052600436106101ea576000357c0100000000000000000000000000000000000000000000000000000000900480639b7c4e0111610114578063ce24388f116100b2578063e3dd138f11610081578063e3dd138f1461073b578063ebc59a5514610773578063f2a946b4146104ca578063fdeb177914610788576101ea565b8063ce24388f14610655578063d39ecc9514610687578063d4dcff82146106d4578063d951baed146106e9576101ea565b8063c0409a5c116100ee578063c0409a5c14610526578063c0ee0b8a14610559578063c2b9cb4d146105eb578063ce11ae6614610623576101ea565b80639b7c4e0114610498578063aea90cb8146104ca578063bc96b872146104df576101ea565b806355ebd76d1161018c57806383576ec11161015b57806383576ec1146104425780638aa3f897146104575780639157772a14610483578063974c86b514610483576101ea565b806355ebd76d146103875780637234c79d146103b05780637919c445146103e25780637e01b4f1146103f7576101ea565b80632d34ba79116101c85780632d34ba79146102b357806341ec6870146102ee57806343271d791461032b57806345bc4d101461035d576101ea565b806307c399a3146101ef578063174a50881461025657806329870c7514610294575b600080fd5b3480156101fb57600080fd5b506102216004803603604081101561021257600080fd5b508035151590602001356107bb565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561026257600080fd5b506102826004803603602081101561027957600080fd5b50351515610801565b60408051918252519081900360200190f35b6102b1600480360360208110156102aa57600080fd5b5035610827565b005b3480156102bf57600080fd5b506102b1600480360360408110156102d657600080fd5b50600160a060020a038135811691602001351661093b565b3480156102fa57600080fd5b50610303610a69565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34801561033757600080fd5b506102b16004803603604081101561034e57600080fd5b50803515159060200135610a83565b34801561036957600080fd5b506102826004803603602081101561038057600080fd5b5035610b0e565b34801561039357600080fd5b5061039c610bd8565b604080519115158252519081900360200190f35b3480156103bc57600080fd5b5061039c600480360360408110156103d357600080fd5b50803515159060200135610c9c565b3480156103ee57600080fd5b506102b1610cc4565b34801561040357600080fd5b506104296004803603604081101561041a57600080fd5b50803515159060200135610df3565b6040805192835260208301919091528051918290030190f35b34801561044e57600080fd5b5061039c610f1d565b34801561046357600080fd5b506102826004803603602081101561047a57600080fd5b50351515610f4e565b34801561048f57600080fd5b5061039c610f70565b3480156104a457600080fd5b50610429600480360360408110156104bb57600080fd5b50803515159060200135610f75565b3480156104d657600080fd5b5061039c611449565b3480156104eb57600080fd5b506102826004803603608081101561050257600080fd5b50600160a060020a038135169060208101351515906040810135906060013561144e565b34801561053257600080fd5b506102b16004803603602081101561054957600080fd5b5035600160a060020a03166116b1565b34801561056557600080fd5b506102b16004803603606081101561057c57600080fd5b600160a060020a03823516916020810135918101906060810160408201356401000000008111156105ac57600080fd5b8201836020820111156105be57600080fd5b803590602001918460018302840111640100000000831117156105e057600080fd5b50909250905061175e565b3480156105f757600080fd5b506102b16004803603606081101561060e57600080fd5b508035151590602081013590604001356117d5565b34801561062f57600080fd5b506102826004803603604081101561064657600080fd5b50803515159060200135611a69565b34801561066157600080fd5b506102826004803603604081101561067857600080fd5b50803515159060200135611a8b565b34801561069357600080fd5b50610282600480360360a08110156106aa57600080fd5b508035151590602081013590604081013590600160a060020a036060820135169060800135611aad565b3480156106e057600080fd5b50610282611c08565b3480156106f557600080fd5b506106fe611c1c565b60408051600160a060020a0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b34801561074757600080fd5b5061039c6004803603606081101561075e57600080fd5b50803515159060208101359060400135611c42565b34801561077f57600080fd5b5061039c611cce565b34801561079457600080fd5b506102b1600480360360208110156107ab57600080fd5b5035600160a060020a0316611e65565b9015156000908152600c60209081526040808320938352929052208054600182015460028301546003840154600490940154600160a060020a0390931694919390929091565b8015156000908152600c602090815260408083208380529091529020600301545b919050565b600061083282611f16565b600154600354600254929350600160a060020a039091169160009061085690611f16565b905061088a61087d606461087184600a63ffffffff611f3016565b9063ffffffff611f9716565b829063ffffffff61200616565b8410156108cb5760405160e560020a62461bcd02815260040180806020018281038252603e8152602001806126db603e913960400191505060405180910390fd5b60018054600160a060020a0319163317905560028590553460035561025860045562127500600555600082111561093457604051600160a060020a0384169083156108fc029084906000818181858888f19350505050158015610932573d6000803e3d6000fd5b505b5050505050565b610944826116b1565b61094d81611e65565b600080805260208190526000805160206126bb83398151915254604080517f66d382030000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a03909216926366d382039260248084019382900301818387803b1580156109c357600080fd5b505af11580156109d7573d6000803e3d6000fd5b50506001600090815260208190526000805160206125e283398151915254604080517f66d382030000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a0390921694506366d382039350602480820193929182900301818387803b158015610a5557600080fd5b505af1158015610932573d6000803e3d6000fd5b600154600254600354600160a060020a0390921691909192565b8115156000908152600c60209081526040808320848452918290529091208054600160a060020a03163314610b02576040805160e560020a62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b61093484846000612063565b6000333014610b55576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612670833981519152604482015290519081900360640190fd5b610b5d610f1d565b1515610bb3576040805160e560020a62461bcd02815260206004820152601560248201527f6e6f206f6e2d676f696e67207072656d70746976650000000000000000000000604482015290519081900360640190fd5b60095482811015610bca5760006009559050610822565b505060098054829003905590565b6001600090815260208190526000805160206125e2833981519152543390600160a060020a0316811480610c2e575060008080526020526000805160206126bb83398151915254600160a060020a038281169116145b1515610c6e5760405160e560020a62461bcd02815260040180806020018281038252602b815260200180612690602b913960400191505060405180910390fd5b600160009081526020526000805160206125e283398151915254600160a060020a0391821691161490505b90565b8115156000908152600c60209081526040808320848452909152812060020154115b92915050565b6007544311610d075760405160e560020a62461bcd0281526004018080602001828103825260288152602001806126276028913960400191505060405180910390fd5b60098054600680546040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a001819052600160a060020a031982169092556007829055600882905592819055600a819055600b8190559091600160a060020a031690821115610db057604051600160a060020a0382169083156108fc029084906000818181858888f19350505050158015610dae573d6000803e3d6000fd5b505b604080518381529051600160a060020a038316917f0684fde3a3fbbb0e152c72ee32e5607f2e86047a4e977d4317cb0fcd8393a082919081900360200190a25050565b8115156000908152600c602052604081208190818080610e1288610f4e565b90505b8015801590610e2357508683105b15610f0f5760008181526020859052604081209089610e46578160020154610e4c565b81600101545b905060008a610e5f578260010154610e65565b82600201545b905089610e78878463ffffffff61200616565b1115610edc576000610e908b8863ffffffff61222716565b90506000610ea884610871858563ffffffff611f3016565b9050610eba878263ffffffff61200616565b610eca898463ffffffff61200616565b9a509a50505050505050505050610f16565b610eec858263ffffffff61200616565b9450610efe868363ffffffff61200616565b955082600401549350505050610e15565b5093509150505b9250929050565b600654600090600160a060020a0316151580610f3b57506007546000105b15610f4857506001610c99565b50600090565b15156000908152600c6020908152604080832083805290915290206004015490565b600181565b600080333014610fbd576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612670833981519152604482015290519081900360640190fd5b600084610fcb576001610fce565b60005b8015156000908152600c602052604081209192508080610fed85610f4e565b90505b8015801590610ffe57508782105b1561143b576000818152602085905260408120908a611021578160020154611027565b81600101545b905060008b61103a578260010154611040565b82600201545b90508a611053868363ffffffff61200616565b116111ac57611068868363ffffffff61200616565b955061107a858263ffffffff61200616565b8c1560009081526020819052604080822054600187015482517fb5f07ea100000000000000000000000000000000000000000000000000000000815260048101919091529151939850600160a060020a03169263b5f07ea19260248084019391929182900301818387803b1580156110f157600080fd5b505af1158015611105573d6000803e3d6000fd5b505050508b151560009081526020819052604080822054600286015482517f48043f0d00000000000000000000000000000000000000000000000000000000815260048101919091529151600160a060020a03909116926348043f0d926024808201939182900301818387803b15801561117e57600080fd5b505af1158015611192573d6000803e3d6000fd5b505050506111a288856001612063565b9350505050610ff0565b600080806111c08e8963ffffffff61222716565b905060006111d885610871888563ffffffff611f3016565b90506111ea8a8263ffffffff61200616565b99506111fc898363ffffffff61200616565b98508f611209578161120b565b805b93508f611218578061121a565b815b925050506000808f1515151515815260200190815260200160002060009054906101000a9004600160a060020a0316600160a060020a031663b5f07ea1836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156112aa57600080fd5b505af11580156112be573d6000803e3d6000fd5b505050508d15156000908152602081905260408082205481517f48043f0d000000000000000000000000000000000000000000000000000000008152600481018590529151600160a060020a03909116926348043f0d926024808201939182900301818387803b15801561133157600080fd5b505af1158015611345573d6000803e3d6000fd5b505050508d1515600090815260208181526040808320548854825160e060020a63a9059cbb028152600160a060020a03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b1580156113af57600080fd5b505af11580156113c3573d6000803e3d6000fd5b505050506040513d60208110156113d957600080fd5b505060018501546113f0908363ffffffff61222716565b6001860155600285015461140a908263ffffffff61222716565b60028601556001850154158061142257506002850154155b15611435576114338a876000612063565b505b50505050505b509097909650945050505050565b600081565b6000808311801561145f5750600082115b15156114b5576040805160e560020a62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b6f01000000000000000000000000000000831080156114e357506f0100000000000000000000000000000082105b1515611539576040805160e560020a62461bcd02815260206004820152601460248201527f67726561746572207468616e20737570706c793f000000000000000000000000604482015290519081900360640190fd5b600160a060020a0385166000818152600d602081815260408084208054600101908190558915158552600c835281852086865293835281516c01000000000000000000000000909602868401526034860152605485018890526074808601889052815180870390910181526094909501908190528451929460029390928291908401908083835b602083106115df5780518252601f1990920191602091820191016115c0565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561161e573d6000803e3d6000fd5b5050506040513d602081101561163357600080fd5b50516040805160a081018252600160a060020a03998a168152602080820198895281830197885260006060830181815260808401828152868352979092529290922090518154600160a060020a0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b60008080526020526000805160206126bb83398151915254600160a060020a031615611727576040805160e560020a62461bcd02815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008080526020526000805160206126bb8339815191528054600160a060020a031916600160a060020a0392909216919091179055565b836000611769610bd8565b90508460008060208614611794578686604081101561178757600080fd5b50803590602001356117a9565b868660208110156117a457600080fd5b503560005b909250905060006117bd8585858986611aad565b90506117c98582612287565b50505050505050505050565b33301461181a576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612670833981519152604482015290519081900360640190fd5b611822610f1d565b1515611878576040805160e560020a62461bcd02815260206004820152601560248201527f6e6f206f6e2d676f696e67207072656d70746976650000000000000000000000604482015290519081900360640190fd5b8215156000908152602081905260408082205481517f48043f0d000000000000000000000000000000000000000000000000000000008152600481018690529151600160a060020a03909116926348043f0d926024808201939182900301818387803b1580156118e757600080fd5b505af11580156118fb573d6000803e3d6000fd5b50505083156000908152602081905260408082205481517fb5f07ea1000000000000000000000000000000000000000000000000000000008152600481018690529151600160a060020a03909116935063b5f07ea19260248084019391929182900301818387803b15801561196f57600080fd5b505af1158015611983573d6000803e3d6000fd5b5050505082151560009081526020818152604080832054600654825160e060020a63a9059cbb028152600160a060020a03918216600482015260248101889052925191169363a9059cbb93604480850194919392918390030190829087803b1580156119ee57600080fd5b505af1158015611a02573d6000803e3d6000fd5b505050506040513d6020811015611a1857600080fd5b505060065460408051848152602081018490528151600160a060020a03909316927fd756774999251f248bfb9b50f86e8cce31c30a63ade2e5e7ebe3185b27309e10929181900390910190a2505050565b9015156000908152600c60209081526040808320938352929052206003015490565b9015156000908152600c60209081526040808320938352929052206004015490565b6000611ab98683610c9c565b1515611b0f576040805160e560020a62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b8515156000908152600c6020526040812090611b2d8589898961144e565b905060008415611b3d5784611b46565b611b4689610f4e565b9050611b53898383611c42565b1515611b97575b611b65898383611c42565b1515611b8257600090815260208390526040902060040154611b5a565b611b8d898383612584565b509150611bff9050565b6000908152602083905260409020600301545b8015801590611bbf5750611bbf898383611c42565b15611bdb57600090815260208390526040902060030154611baa565b600081815260208490526040902060040154611bfa908a908490612584565b509150505b95945050505050565b6f0100000000000000000000000000000081565b600654600854600954600754600a54600b54600160a060020a0390951694909192939495565b8215156000908152600c6020908152604080832085845291829052808320848452908320600f54600282015460018401548692611c96929091611c8a9163ffffffff611f3016565b9063ffffffff611f3016565b90506000611cbf600e54600f5401611c8a86600201548660010154611f3090919063ffffffff16565b90911198975050505050505050565b600654600090600160a060020a0316158015611cea5750600754155b1515611d2a5760405160e560020a62461bcd0281526004018080602001828103825260258152602001806126026025913960400191505060405180910390fd5b333014611d6f576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612670833981519152604482015290519081900360640190fd5b600154600160a060020a03161580611d875750600254155b15611d9457506000610c99565b6001805460068054600160a060020a03808416600160a060020a031992831617928390556002805460088190556003805460098190556005805443810160075560048054600a55600b919091556040805160a0810182526000808252602080830182905282840182905260608301829052608090920181905299909b16909b5594879055918690559285905593909355855192835293820193909352835191909216927ff378fdea5eb2049f2d75b012428cd37d03fae41412c5b5e7a70f10f79c06731a928290030190a250600190565b600160009081526020526000805160206125e283398151915254600160a060020a031615611edd576040805160e560020a62461bcd02815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160009081526020526000805160206125e28339815191528054600160a060020a031916600160a060020a0392909216919091179055565b600080821215611f2c5781600019029050610822565b5090565b6000821515611f4157506000610cbe565b828202828482811515611f5057fe5b0414611f905760405160e560020a62461bcd02815260040180806020018281038252602181526020018061264f6021913960400191505060405180910390fd5b9392505050565b6000808211611ff0576040805160e560020a62461bcd02815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b60008284811515611ffd57fe5b04949350505050565b600082820183811015611f90576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b8215156000908152600c60209081526040808320858452918290528083206004808201546003808401805488528588209093018290559154908652928520019190915583156121495785156000908152602081815260408083205484546002860154835160e060020a63a9059cbb028152600160a060020a0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b15801561211757600080fd5b505af115801561212b573d6000803e3d6000fd5b505050506040513d602081101561214157600080fd5b506121e29050565b8515156000908152602081815260408083205484546001860154835160e060020a63a9059cbb028152600160a060020a0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b1580156121b557600080fd5b505af11580156121c9573d6000803e3d6000fd5b505050506040513d60208110156121df57600080fd5b50505b6004908101546000958652602092909252604085208054600160a060020a03191681556001810186905560028101869055600381018690550193909355509092915050565b600082821115612281576040805160e560020a62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b811580156000908152600c6020818152604080842086855280835281852086865293909252832090926122b985610f4e565b90505b801561257a576000818152602083905260409020600280820154908501546122e99163ffffffff611f3016565b600180830154908601546123029163ffffffff611f3016565b101561231357505050505050612580565b6000612327856001015483600201546125cb565b600186015490915061235790610871612346828563ffffffff61222716565b60028901549063ffffffff611f3016565b60028601556001850154612371908263ffffffff61222716565b8560010181905550600061239a8360020154610871848660010154611f3090919063ffffffff16565b60018401549091506123ca906108716123b9828563ffffffff61222716565b60028701549063ffffffff611f3016565b600284015560018301546123e4908263ffffffff61222716565b6001840155871515600090815260208181526040808320548954825160e060020a63a9059cbb028152600160a060020a03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b15801561244f57600080fd5b505af1158015612463573d6000803e3d6000fd5b505050506040513d602081101561247957600080fd5b5050891515600090815260208181526040808320548654825160e060020a63a9059cbb028152600160a060020a03918216600482015260248101889052925191169363a9059cbb93604480850194919392918390030190829087803b1580156124e157600080fd5b505af11580156124f5573d6000803e3d6000fd5b505050506040513d602081101561250b57600080fd5b50506001830154158061252057506002830154155b156125335761253188856000612063565b505b61253c88610f4e565b935085600101546000148061255357506002860154155b15612572576125648a8a6000612063565b505050505050505050612580565b5050506122bc565b50505050505b5050565b9115156000908152600c6020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b60008183106125da5781611f90565b509091905056feada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d63757272656e74207072656d7074697665206973206e6f742066696e697368656420796574697320696e206c6f636b206475726174696f6e206f662063757272656e74207072656d7074697665536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77636f6e73656e737573206f6e6c790000000000000000000000000000000000006f6e6c7920566f6c6174696c65546f6b656e20616e6420537461626c65546f6b656e206163636570746564ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5796f7572207072656d7074697665206973206e6f7420676f6f6420656e6f75676820636f6d7061726520746f2063757272656e74207072656d7074697665a165627a7a7230582037a37e80125e8592370590fd6671a834ed2b6b63372842bc20a18b9c41ed25230029ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5`

// DeployPairEx deploys a new Ethereum contract, binding an instance of PairEx to it.
func DeployPairEx(auth *bind.TransactOpts, backend bind.ContractBackend, _volatileTokenAddress common.Address, _stableTokenAddress common.Address) (common.Address, *types.Transaction, *PairEx, error) {
	parsed, err := abi.JSON(strings.NewReader(PairExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairExBin), backend, _volatileTokenAddress, _stableTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PairEx{PairExCaller: PairExCaller{contract: contract}, PairExTransactor: PairExTransactor{contract: contract}, PairExFilterer: PairExFilterer{contract: contract}}, nil
}

// PairEx is an auto generated Go binding around an Ethereum contract.
type PairEx struct {
	PairExCaller     // Read-only binding to the contract
	PairExTransactor // Write-only binding to the contract
	PairExFilterer   // Log filterer for contract events
}

// PairExCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairExSession struct {
	Contract     *PairEx           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairExCallerSession struct {
	Contract *PairExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PairExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairExTransactorSession struct {
	Contract     *PairExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairExRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairExRaw struct {
	Contract *PairEx // Generic contract binding to access the raw methods on
}

// PairExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairExCallerRaw struct {
	Contract *PairExCaller // Generic read-only contract binding to access the raw methods on
}

// PairExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairExTransactorRaw struct {
	Contract *PairExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairEx creates a new instance of PairEx, bound to a specific deployed contract.
func NewPairEx(address common.Address, backend bind.ContractBackend) (*PairEx, error) {
	contract, err := bindPairEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairEx{PairExCaller: PairExCaller{contract: contract}, PairExTransactor: PairExTransactor{contract: contract}, PairExFilterer: PairExFilterer{contract: contract}}, nil
}

// NewPairExCaller creates a new read-only instance of PairEx, bound to a specific deployed contract.
func NewPairExCaller(address common.Address, caller bind.ContractCaller) (*PairExCaller, error) {
	contract, err := bindPairEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairExCaller{contract: contract}, nil
}

// NewPairExTransactor creates a new write-only instance of PairEx, bound to a specific deployed contract.
func NewPairExTransactor(address common.Address, transactor bind.ContractTransactor) (*PairExTransactor, error) {
	contract, err := bindPairEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairExTransactor{contract: contract}, nil
}

// NewPairExFilterer creates a new log filterer instance of PairEx, bound to a specific deployed contract.
func NewPairExFilterer(address common.Address, filterer bind.ContractFilterer) (*PairExFilterer, error) {
	contract, err := bindPairEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairExFilterer{contract: contract}, nil
}

// bindPairEx binds a generic wrapper to an already deployed contract.
func bindPairEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairEx *PairExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairEx.Contract.PairExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairEx *PairExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.Contract.PairExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairEx *PairExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairEx.Contract.PairExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairEx *PairExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairEx *PairExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairEx *PairExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairEx.Contract.contract.Transact(opts, method, params...)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_PairEx *PairExCaller) Buy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "Buy")
	return *ret0, err
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_PairEx *PairExSession) Buy() (bool, error) {
	return _PairEx.Contract.Buy(&_PairEx.CallOpts)
}

// Buy is a free data retrieval call binding the contract method 0x974c86b5.
//
// Solidity: function Buy() constant returns(bool)
func (_PairEx *PairExCallerSession) Buy() (bool, error) {
	return _PairEx.Contract.Buy(&_PairEx.CallOpts)
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_PairEx *PairExCaller) INPUTSMAX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "INPUTS_MAX")
	return *ret0, err
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_PairEx *PairExSession) INPUTSMAX() (*big.Int, error) {
	return _PairEx.Contract.INPUTSMAX(&_PairEx.CallOpts)
}

// INPUTSMAX is a free data retrieval call binding the contract method 0xd4dcff82.
//
// Solidity: function INPUTS_MAX() constant returns(uint256)
func (_PairEx *PairExCallerSession) INPUTSMAX() (*big.Int, error) {
	return _PairEx.Contract.INPUTSMAX(&_PairEx.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_PairEx *PairExCaller) Sell(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "Sell")
	return *ret0, err
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_PairEx *PairExSession) Sell() (bool, error) {
	return _PairEx.Contract.Sell(&_PairEx.CallOpts)
}

// Sell is a free data retrieval call binding the contract method 0xaea90cb8.
//
// Solidity: function Sell() constant returns(bool)
func (_PairEx *PairExCallerSession) Sell() (bool, error) {
	return _PairEx.Contract.Sell(&_PairEx.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_PairEx *PairExCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "Stable")
	return *ret0, err
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_PairEx *PairExSession) Stable() (bool, error) {
	return _PairEx.Contract.Stable(&_PairEx.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x9157772a.
//
// Solidity: function Stable() constant returns(bool)
func (_PairEx *PairExCallerSession) Stable() (bool, error) {
	return _PairEx.Contract.Stable(&_PairEx.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_PairEx *PairExCaller) Volatile(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "Volatile")
	return *ret0, err
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_PairEx *PairExSession) Volatile() (bool, error) {
	return _PairEx.Contract.Volatile(&_PairEx.CallOpts)
}

// Volatile is a free data retrieval call binding the contract method 0xf2a946b4.
//
// Solidity: function Volatile() constant returns(bool)
func (_PairEx *PairExCallerSession) Volatile() (bool, error) {
	return _PairEx.Contract.Volatile(&_PairEx.CallOpts)
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_PairEx *PairExCaller) BetterOrder(opts *bind.CallOpts, orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "betterOrder", orderType, _newId, _oldId)
	return *ret0, err
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_PairEx *PairExSession) BetterOrder(orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	return _PairEx.Contract.BetterOrder(&_PairEx.CallOpts, orderType, _newId, _oldId)
}

// BetterOrder is a free data retrieval call binding the contract method 0xe3dd138f.
//
// Solidity: function betterOrder(bool orderType, bytes32 _newId, bytes32 _oldId) constant returns(bool)
func (_PairEx *PairExCallerSession) BetterOrder(orderType bool, _newId [32]byte, _oldId [32]byte) (bool, error) {
	return _PairEx.Contract.BetterOrder(&_PairEx.CallOpts, orderType, _newId, _oldId)
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExCaller) Bottom(opts *bind.CallOpts, _orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "bottom", _orderType)
	return *ret0, err
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExSession) Bottom(_orderType bool) ([32]byte, error) {
	return _PairEx.Contract.Bottom(&_PairEx.CallOpts, _orderType)
}

// Bottom is a free data retrieval call binding the contract method 0x174a5088.
//
// Solidity: function bottom(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExCallerSession) Bottom(_orderType bool) ([32]byte, error) {
	return _PairEx.Contract.Bottom(&_PairEx.CallOpts, _orderType)
}

// GetCurrentPremptive is a free data retrieval call binding the contract method 0xd951baed.
//
// Solidity: function getCurrentPremptive() constant returns(address, int256, uint256, uint256, uint256, uint256)
func (_PairEx *PairExCaller) GetCurrentPremptive(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _PairEx.contract.Call(opts, out, "getCurrentPremptive")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentPremptive is a free data retrieval call binding the contract method 0xd951baed.
//
// Solidity: function getCurrentPremptive() constant returns(address, int256, uint256, uint256, uint256, uint256)
func (_PairEx *PairExSession) GetCurrentPremptive() (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PairEx.Contract.GetCurrentPremptive(&_PairEx.CallOpts)
}

// GetCurrentPremptive is a free data retrieval call binding the contract method 0xd951baed.
//
// Solidity: function getCurrentPremptive() constant returns(address, int256, uint256, uint256, uint256, uint256)
func (_PairEx *PairExCallerSession) GetCurrentPremptive() (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PairEx.Contract.GetCurrentPremptive(&_PairEx.CallOpts)
}

// GetCurrentProposal is a free data retrieval call binding the contract method 0x41ec6870.
//
// Solidity: function getCurrentProposal() constant returns(address, int256, uint256)
func (_PairEx *PairExCaller) GetCurrentProposal(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PairEx.contract.Call(opts, out, "getCurrentProposal")
	return *ret0, *ret1, *ret2, err
}

// GetCurrentProposal is a free data retrieval call binding the contract method 0x41ec6870.
//
// Solidity: function getCurrentProposal() constant returns(address, int256, uint256)
func (_PairEx *PairExSession) GetCurrentProposal() (common.Address, *big.Int, *big.Int, error) {
	return _PairEx.Contract.GetCurrentProposal(&_PairEx.CallOpts)
}

// GetCurrentProposal is a free data retrieval call binding the contract method 0x41ec6870.
//
// Solidity: function getCurrentProposal() constant returns(address, int256, uint256)
func (_PairEx *PairExCallerSession) GetCurrentProposal() (common.Address, *big.Int, *big.Int, error) {
	return _PairEx.Contract.GetCurrentProposal(&_PairEx.CallOpts)
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExCaller) GetNext(opts *bind.CallOpts, _orderType bool, _id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "getNext", _orderType, _id)
	return *ret0, err
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExSession) GetNext(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _PairEx.Contract.GetNext(&_PairEx.CallOpts, _orderType, _id)
}

// GetNext is a free data retrieval call binding the contract method 0xce24388f.
//
// Solidity: function getNext(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExCallerSession) GetNext(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _PairEx.Contract.GetNext(&_PairEx.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_PairEx *PairExCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _PairEx.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_PairEx *PairExSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _PairEx.Contract.GetOrder(&_PairEx.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_PairEx *PairExCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _PairEx.Contract.GetOrder(&_PairEx.CallOpts, _orderType, _id)
}

// GetOrderType is a free data retrieval call binding the contract method 0x55ebd76d.
//
// Solidity: function getOrderType() constant returns(bool)
func (_PairEx *PairExCaller) GetOrderType(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "getOrderType")
	return *ret0, err
}

// GetOrderType is a free data retrieval call binding the contract method 0x55ebd76d.
//
// Solidity: function getOrderType() constant returns(bool)
func (_PairEx *PairExSession) GetOrderType() (bool, error) {
	return _PairEx.Contract.GetOrderType(&_PairEx.CallOpts)
}

// GetOrderType is a free data retrieval call binding the contract method 0x55ebd76d.
//
// Solidity: function getOrderType() constant returns(bool)
func (_PairEx *PairExCallerSession) GetOrderType() (bool, error) {
	return _PairEx.Contract.GetOrderType(&_PairEx.CallOpts)
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExCaller) GetPrev(opts *bind.CallOpts, _orderType bool, _id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "getPrev", _orderType, _id)
	return *ret0, err
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExSession) GetPrev(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _PairEx.Contract.GetPrev(&_PairEx.CallOpts, _orderType, _id)
}

// GetPrev is a free data retrieval call binding the contract method 0xce11ae66.
//
// Solidity: function getPrev(bool _orderType, bytes32 _id) constant returns(bytes32)
func (_PairEx *PairExCallerSession) GetPrev(_orderType bool, _id [32]byte) ([32]byte, error) {
	return _PairEx.Contract.GetPrev(&_PairEx.CallOpts, _orderType, _id)
}

// HasActivePremptive is a free data retrieval call binding the contract method 0x83576ec1.
//
// Solidity: function hasActivePremptive() constant returns(bool)
func (_PairEx *PairExCaller) HasActivePremptive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "hasActivePremptive")
	return *ret0, err
}

// HasActivePremptive is a free data retrieval call binding the contract method 0x83576ec1.
//
// Solidity: function hasActivePremptive() constant returns(bool)
func (_PairEx *PairExSession) HasActivePremptive() (bool, error) {
	return _PairEx.Contract.HasActivePremptive(&_PairEx.CallOpts)
}

// HasActivePremptive is a free data retrieval call binding the contract method 0x83576ec1.
//
// Solidity: function hasActivePremptive() constant returns(bool)
func (_PairEx *PairExCallerSession) HasActivePremptive() (bool, error) {
	return _PairEx.Contract.HasActivePremptive(&_PairEx.CallOpts)
}

// OrderToFill is a free data retrieval call binding the contract method 0x7e01b4f1.
//
// Solidity: function orderToFill(bool _orderType, uint256 _stableTokenTarget) constant returns(uint256, uint256)
func (_PairEx *PairExCaller) OrderToFill(opts *bind.CallOpts, _orderType bool, _stableTokenTarget *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PairEx.contract.Call(opts, out, "orderToFill", _orderType, _stableTokenTarget)
	return *ret0, *ret1, err
}

// OrderToFill is a free data retrieval call binding the contract method 0x7e01b4f1.
//
// Solidity: function orderToFill(bool _orderType, uint256 _stableTokenTarget) constant returns(uint256, uint256)
func (_PairEx *PairExSession) OrderToFill(_orderType bool, _stableTokenTarget *big.Int) (*big.Int, *big.Int, error) {
	return _PairEx.Contract.OrderToFill(&_PairEx.CallOpts, _orderType, _stableTokenTarget)
}

// OrderToFill is a free data retrieval call binding the contract method 0x7e01b4f1.
//
// Solidity: function orderToFill(bool _orderType, uint256 _stableTokenTarget) constant returns(uint256, uint256)
func (_PairEx *PairExCallerSession) OrderToFill(_orderType bool, _stableTokenTarget *big.Int) (*big.Int, *big.Int, error) {
	return _PairEx.Contract.OrderToFill(&_PairEx.CallOpts, _orderType, _stableTokenTarget)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExCaller) Top(opts *bind.CallOpts, _orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "top", _orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExSession) Top(_orderType bool) ([32]byte, error) {
	return _PairEx.Contract.Top(&_PairEx.CallOpts, _orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool _orderType) constant returns(bytes32)
func (_PairEx *PairExCallerSession) Top(_orderType bool) ([32]byte, error) {
	return _PairEx.Contract.Top(&_PairEx.CallOpts, _orderType)
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_PairEx *PairExCaller) ValidOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairEx.contract.Call(opts, out, "validOrder", _orderType, _id)
	return *ret0, err
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_PairEx *PairExSession) ValidOrder(_orderType bool, _id [32]byte) (bool, error) {
	return _PairEx.Contract.ValidOrder(&_PairEx.CallOpts, _orderType, _id)
}

// ValidOrder is a free data retrieval call binding the contract method 0x7234c79d.
//
// Solidity: function validOrder(bool _orderType, bytes32 _id) constant returns(bool)
func (_PairEx *PairExCallerSession) ValidOrder(_orderType bool, _id [32]byte) (bool, error) {
	return _PairEx.Contract.ValidOrder(&_PairEx.CallOpts, _orderType, _id)
}

// Absorb is a paid mutator transaction binding the contract method 0x9b7c4e01.
//
// Solidity: function absorb(bool _inflate, uint256 _stableTokenTarget) returns(uint256, uint256)
func (_PairEx *PairExTransactor) Absorb(opts *bind.TransactOpts, _inflate bool, _stableTokenTarget *big.Int) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "absorb", _inflate, _stableTokenTarget)
}

// Absorb is a paid mutator transaction binding the contract method 0x9b7c4e01.
//
// Solidity: function absorb(bool _inflate, uint256 _stableTokenTarget) returns(uint256, uint256)
func (_PairEx *PairExSession) Absorb(_inflate bool, _stableTokenTarget *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Absorb(&_PairEx.TransactOpts, _inflate, _stableTokenTarget)
}

// Absorb is a paid mutator transaction binding the contract method 0x9b7c4e01.
//
// Solidity: function absorb(bool _inflate, uint256 _stableTokenTarget) returns(uint256, uint256)
func (_PairEx *PairExTransactorSession) Absorb(_inflate bool, _stableTokenTarget *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Absorb(&_PairEx.TransactOpts, _inflate, _stableTokenTarget)
}

// ActivatePremptive is a paid mutator transaction binding the contract method 0xebc59a55.
//
// Solidity: function activatePremptive() returns(bool)
func (_PairEx *PairExTransactor) ActivatePremptive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "activatePremptive")
}

// ActivatePremptive is a paid mutator transaction binding the contract method 0xebc59a55.
//
// Solidity: function activatePremptive() returns(bool)
func (_PairEx *PairExSession) ActivatePremptive() (*types.Transaction, error) {
	return _PairEx.Contract.ActivatePremptive(&_PairEx.TransactOpts)
}

// ActivatePremptive is a paid mutator transaction binding the contract method 0xebc59a55.
//
// Solidity: function activatePremptive() returns(bool)
func (_PairEx *PairExTransactorSession) ActivatePremptive() (*types.Transaction, error) {
	return _PairEx.Contract.ActivatePremptive(&_PairEx.TransactOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExTransactor) Cancel(opts *bind.TransactOpts, _orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "cancel", _orderType, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExSession) Cancel(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Cancel(&_PairEx.TransactOpts, _orderType, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExTransactorSession) Cancel(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Cancel(&_PairEx.TransactOpts, _orderType, _id)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_PairEx *PairExTransactor) InitOrder(opts *bind.TransactOpts, _maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "initOrder", _maker, _orderType, _haveAmount, _wantAmount)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_PairEx *PairExSession) InitOrder(_maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.InitOrder(&_PairEx.TransactOpts, _maker, _orderType, _haveAmount, _wantAmount)
}

// InitOrder is a paid mutator transaction binding the contract method 0xbc96b872.
//
// Solidity: function initOrder(address _maker, bool _orderType, uint256 _haveAmount, uint256 _wantAmount) returns(bytes32)
func (_PairEx *PairExTransactorSession) InitOrder(_maker common.Address, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.InitOrder(&_PairEx.TransactOpts, _maker, _orderType, _haveAmount, _wantAmount)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_PairEx *PairExTransactor) Insert(opts *bind.TransactOpts, _orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "insert", _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_PairEx *PairExSession) Insert(_orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Insert(&_PairEx.TransactOpts, _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// Insert is a paid mutator transaction binding the contract method 0xd39ecc95.
//
// Solidity: function insert(bool _orderType, uint256 _haveAmount, uint256 _wantAmount, address _maker, bytes32 _assistingID) returns(bytes32)
func (_PairEx *PairExTransactorSession) Insert(_orderType bool, _haveAmount *big.Int, _wantAmount *big.Int, _maker common.Address, _assistingID [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Insert(&_PairEx.TransactOpts, _orderType, _haveAmount, _wantAmount, _maker, _assistingID)
}

// PremptiveAbsorb is a paid mutator transaction binding the contract method 0xc2b9cb4d.
//
// Solidity: function premptiveAbsorb(bool _inflate, uint256 _totalSTB, uint256 _totalVOL) returns()
func (_PairEx *PairExTransactor) PremptiveAbsorb(opts *bind.TransactOpts, _inflate bool, _totalSTB *big.Int, _totalVOL *big.Int) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "premptiveAbsorb", _inflate, _totalSTB, _totalVOL)
}

// PremptiveAbsorb is a paid mutator transaction binding the contract method 0xc2b9cb4d.
//
// Solidity: function premptiveAbsorb(bool _inflate, uint256 _totalSTB, uint256 _totalVOL) returns()
func (_PairEx *PairExSession) PremptiveAbsorb(_inflate bool, _totalSTB *big.Int, _totalVOL *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.PremptiveAbsorb(&_PairEx.TransactOpts, _inflate, _totalSTB, _totalVOL)
}

// PremptiveAbsorb is a paid mutator transaction binding the contract method 0xc2b9cb4d.
//
// Solidity: function premptiveAbsorb(bool _inflate, uint256 _totalSTB, uint256 _totalVOL) returns()
func (_PairEx *PairExTransactorSession) PremptiveAbsorb(_inflate bool, _totalSTB *big.Int, _totalVOL *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.PremptiveAbsorb(&_PairEx.TransactOpts, _inflate, _totalSTB, _totalVOL)
}

// Propose is a paid mutator transaction binding the contract method 0x29870c75.
//
// Solidity: function propose(int256 amount) returns()
func (_PairEx *PairExTransactor) Propose(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "propose", amount)
}

// Propose is a paid mutator transaction binding the contract method 0x29870c75.
//
// Solidity: function propose(int256 amount) returns()
func (_PairEx *PairExSession) Propose(amount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Propose(&_PairEx.TransactOpts, amount)
}

// Propose is a paid mutator transaction binding the contract method 0x29870c75.
//
// Solidity: function propose(int256 amount) returns()
func (_PairEx *PairExTransactorSession) Propose(amount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Propose(&_PairEx.TransactOpts, amount)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _volatileTokenAddress, address _stableTokenAddress) returns()
func (_PairEx *PairExTransactor) Setup(opts *bind.TransactOpts, _volatileTokenAddress common.Address, _stableTokenAddress common.Address) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "setup", _volatileTokenAddress, _stableTokenAddress)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _volatileTokenAddress, address _stableTokenAddress) returns()
func (_PairEx *PairExSession) Setup(_volatileTokenAddress common.Address, _stableTokenAddress common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.Setup(&_PairEx.TransactOpts, _volatileTokenAddress, _stableTokenAddress)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _volatileTokenAddress, address _stableTokenAddress) returns()
func (_PairEx *PairExTransactorSession) Setup(_volatileTokenAddress common.Address, _stableTokenAddress common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.Setup(&_PairEx.TransactOpts, _volatileTokenAddress, _stableTokenAddress)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 _amount) returns(uint256)
func (_PairEx *PairExTransactor) Slash(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "slash", _amount)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 _amount) returns(uint256)
func (_PairEx *PairExSession) Slash(_amount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Slash(&_PairEx.TransactOpts, _amount)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 _amount) returns(uint256)
func (_PairEx *PairExTransactorSession) Slash(_amount *big.Int) (*types.Transaction, error) {
	return _PairEx.Contract.Slash(&_PairEx.TransactOpts, _amount)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_PairEx *PairExTransactor) StableTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "stableTokenRegister", _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_PairEx *PairExSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.StableTokenRegister(&_PairEx.TransactOpts, _address)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0xfdeb1779.
//
// Solidity: function stableTokenRegister(address _address) returns()
func (_PairEx *PairExTransactorSession) StableTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.StableTokenRegister(&_PairEx.TransactOpts, _address)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_PairEx *PairExTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_PairEx *PairExSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _PairEx.Contract.TokenFallback(&_PairEx.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_PairEx *PairExTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _PairEx.Contract.TokenFallback(&_PairEx.TransactOpts, _from, _value, _data)
}

// UnlockPremptive is a paid mutator transaction binding the contract method 0x7919c445.
//
// Solidity: function unlockPremptive() returns()
func (_PairEx *PairExTransactor) UnlockPremptive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "unlockPremptive")
}

// UnlockPremptive is a paid mutator transaction binding the contract method 0x7919c445.
//
// Solidity: function unlockPremptive() returns()
func (_PairEx *PairExSession) UnlockPremptive() (*types.Transaction, error) {
	return _PairEx.Contract.UnlockPremptive(&_PairEx.TransactOpts)
}

// UnlockPremptive is a paid mutator transaction binding the contract method 0x7919c445.
//
// Solidity: function unlockPremptive() returns()
func (_PairEx *PairExTransactorSession) UnlockPremptive() (*types.Transaction, error) {
	return _PairEx.Contract.UnlockPremptive(&_PairEx.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_PairEx *PairExTransactor) VolatileTokenRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "volatileTokenRegister", _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_PairEx *PairExSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.VolatileTokenRegister(&_PairEx.TransactOpts, _address)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0xc0409a5c.
//
// Solidity: function volatileTokenRegister(address _address) returns()
func (_PairEx *PairExTransactorSession) VolatileTokenRegister(_address common.Address) (*types.Transaction, error) {
	return _PairEx.Contract.VolatileTokenRegister(&_PairEx.TransactOpts, _address)
}

// PairExPremptiveAbsorptionIterator is returned from FilterPremptiveAbsorption and is used to iterate over the raw logs and unpacked data for PremptiveAbsorption events raised by the PairEx contract.
type PairExPremptiveAbsorptionIterator struct {
	Event *PairExPremptiveAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PairExPremptiveAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairExPremptiveAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PairExPremptiveAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PairExPremptiveAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairExPremptiveAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairExPremptiveAbsorption represents a PremptiveAbsorption event raised by the PairEx contract.
type PairExPremptiveAbsorption struct {
	Initiator common.Address
	Stb       *big.Int
	Vol       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPremptiveAbsorption is a free log retrieval operation binding the contract event 0xd756774999251f248bfb9b50f86e8cce31c30a63ade2e5e7ebe3185b27309e10.
//
// Solidity: event PremptiveAbsorption(address indexed _initiator, uint256 _stb, uint256 _vol)
func (_PairEx *PairExFilterer) FilterPremptiveAbsorption(opts *bind.FilterOpts, _initiator []common.Address) (*PairExPremptiveAbsorptionIterator, error) {

	var _initiatorRule []interface{}
	for _, _initiatorItem := range _initiator {
		_initiatorRule = append(_initiatorRule, _initiatorItem)
	}

	logs, sub, err := _PairEx.contract.FilterLogs(opts, "PremptiveAbsorption", _initiatorRule)
	if err != nil {
		return nil, err
	}
	return &PairExPremptiveAbsorptionIterator{contract: _PairEx.contract, event: "PremptiveAbsorption", logs: logs, sub: sub}, nil
}

// WatchPremptiveAbsorption is a free log subscription operation binding the contract event 0xd756774999251f248bfb9b50f86e8cce31c30a63ade2e5e7ebe3185b27309e10.
//
// Solidity: event PremptiveAbsorption(address indexed _initiator, uint256 _stb, uint256 _vol)
func (_PairEx *PairExFilterer) WatchPremptiveAbsorption(opts *bind.WatchOpts, sink chan<- *PairExPremptiveAbsorption, _initiator []common.Address) (event.Subscription, error) {

	var _initiatorRule []interface{}
	for _, _initiatorItem := range _initiator {
		_initiatorRule = append(_initiatorRule, _initiatorItem)
	}

	logs, sub, err := _PairEx.contract.WatchLogs(opts, "PremptiveAbsorption", _initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairExPremptiveAbsorption)
				if err := _PairEx.contract.UnpackLog(event, "PremptiveAbsorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PairExPremptiveActivationIterator is returned from FilterPremptiveActivation and is used to iterate over the raw logs and unpacked data for PremptiveActivation events raised by the PairEx contract.
type PairExPremptiveActivationIterator struct {
	Event *PairExPremptiveActivation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PairExPremptiveActivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairExPremptiveActivation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PairExPremptiveActivation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PairExPremptiveActivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairExPremptiveActivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairExPremptiveActivation represents a PremptiveActivation event raised by the PairEx contract.
type PairExPremptiveActivation struct {
	Initator common.Address
	Amount   *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPremptiveActivation is a free log retrieval operation binding the contract event 0xf378fdea5eb2049f2d75b012428cd37d03fae41412c5b5e7a70f10f79c06731a.
//
// Solidity: event PremptiveActivation(address indexed _initator, int256 _amount, uint256 _value)
func (_PairEx *PairExFilterer) FilterPremptiveActivation(opts *bind.FilterOpts, _initator []common.Address) (*PairExPremptiveActivationIterator, error) {

	var _initatorRule []interface{}
	for _, _initatorItem := range _initator {
		_initatorRule = append(_initatorRule, _initatorItem)
	}

	logs, sub, err := _PairEx.contract.FilterLogs(opts, "PremptiveActivation", _initatorRule)
	if err != nil {
		return nil, err
	}
	return &PairExPremptiveActivationIterator{contract: _PairEx.contract, event: "PremptiveActivation", logs: logs, sub: sub}, nil
}

// WatchPremptiveActivation is a free log subscription operation binding the contract event 0xf378fdea5eb2049f2d75b012428cd37d03fae41412c5b5e7a70f10f79c06731a.
//
// Solidity: event PremptiveActivation(address indexed _initator, int256 _amount, uint256 _value)
func (_PairEx *PairExFilterer) WatchPremptiveActivation(opts *bind.WatchOpts, sink chan<- *PairExPremptiveActivation, _initator []common.Address) (event.Subscription, error) {

	var _initatorRule []interface{}
	for _, _initatorItem := range _initator {
		_initatorRule = append(_initatorRule, _initatorItem)
	}

	logs, sub, err := _PairEx.contract.WatchLogs(opts, "PremptiveActivation", _initatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairExPremptiveActivation)
				if err := _PairEx.contract.UnpackLog(event, "PremptiveActivation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PairExPremptiveUnlockIterator is returned from FilterPremptiveUnlock and is used to iterate over the raw logs and unpacked data for PremptiveUnlock events raised by the PairEx contract.
type PairExPremptiveUnlockIterator struct {
	Event *PairExPremptiveUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PairExPremptiveUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairExPremptiveUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PairExPremptiveUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PairExPremptiveUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairExPremptiveUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairExPremptiveUnlock represents a PremptiveUnlock event raised by the PairEx contract.
type PairExPremptiveUnlock struct {
	Initiator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPremptiveUnlock is a free log retrieval operation binding the contract event 0x0684fde3a3fbbb0e152c72ee32e5607f2e86047a4e977d4317cb0fcd8393a082.
//
// Solidity: event PremptiveUnlock(address indexed _initiator, uint256 _amount)
func (_PairEx *PairExFilterer) FilterPremptiveUnlock(opts *bind.FilterOpts, _initiator []common.Address) (*PairExPremptiveUnlockIterator, error) {

	var _initiatorRule []interface{}
	for _, _initiatorItem := range _initiator {
		_initiatorRule = append(_initiatorRule, _initiatorItem)
	}

	logs, sub, err := _PairEx.contract.FilterLogs(opts, "PremptiveUnlock", _initiatorRule)
	if err != nil {
		return nil, err
	}
	return &PairExPremptiveUnlockIterator{contract: _PairEx.contract, event: "PremptiveUnlock", logs: logs, sub: sub}, nil
}

// WatchPremptiveUnlock is a free log subscription operation binding the contract event 0x0684fde3a3fbbb0e152c72ee32e5607f2e86047a4e977d4317cb0fcd8393a082.
//
// Solidity: event PremptiveUnlock(address indexed _initiator, uint256 _amount)
func (_PairEx *PairExFilterer) WatchPremptiveUnlock(opts *bind.WatchOpts, sink chan<- *PairExPremptiveUnlock, _initiator []common.Address) (event.Subscription, error) {

	var _initiatorRule []interface{}
	for _, _initiatorItem := range _initiator {
		_initiatorRule = append(_initiatorRule, _initiatorItem)
	}

	logs, sub, err := _PairEx.contract.WatchLogs(opts, "PremptiveUnlock", _initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairExPremptiveUnlock)
				if err := _PairEx.contract.UnpackLog(event, "PremptiveUnlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582019dc6bba8288aab3e918a673ffb7cfb739a764a6685690baaff566e5de1b63420029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
