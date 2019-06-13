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
const DataSetBin = `0x6080604052348015600f57600080fd5b50606380601d6000396000f3fe6080604052600080fdfea265627a7a72305820f37703a94d69652cb8ea5626f13e73c3e4e5371f66d31c5832bffcf59c1e8c5364736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057`

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
const InitializerBin = `0x608060405234801561001057600080fd5b506102b9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80639157772a14610067578063974c86b514610067578063aea90cb814610083578063c0409a5c1461008b578063f2a946b414610083578063fdeb1779146100b3575b600080fd5b61006f6100d9565b604080519115158252519081900360200190f35b61006f6100de565b6100b1600480360360208110156100a157600080fd5b50356001600160a01b03166100e3565b005b6100b1600480360360208110156100c957600080fd5b50356001600160a01b031661019f565b600181565b600081565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b031615610156576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b0392909216919091179055565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b031615610214576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723058207545895e905b2705d9e978795ce4268f98b8358507891bef929f45557398642464736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057`

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
const MathBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058206a83166a8ac40b7a7d02722775edbc24b3b4bc9d767410293175817aa99d7ec864736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057`

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
const OrderBookBin = `0x608060405260006003556001600455610cfe8061001d6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063bc96b872116100a2578063d39ecc9511610071578063d39ecc95146102d4578063d4dcff8214610314578063e3dd138f1461031c578063f2a946b414610222578063fdeb1779146103475761010b565b8063bc96b8721461022a578063c0409a5c14610264578063ce11ae661461028a578063ce24388f146102af5761010b565b80638aa3f897116100de5780638aa3f897146101fb5780639157772a1461021a578063974c86b51461021a578063aea90cb8146102225761010b565b806307c399a314610110578063174a50881461016a57806343271d791461019b5780637234c79d146101c2575b600080fd5b6101356004803603604081101561012657600080fd5b5080351515906020013561036d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101896004803603602081101561018057600080fd5b503515156103b3565b60408051918252519081900360200190f35b6101c0600480360360408110156101b157600080fd5b508035151590602001356103d5565b005b6101e7600480360360408110156101d857600080fd5b50803515159060200135610457565b604080519115158252519081900360200190f35b6101896004803603602081101561021157600080fd5b50351515610480565b6101e76104a2565b6101e76104a7565b6101896004803603608081101561024057600080fd5b506001600160a01b03813516906020810135151590604081013590606001356104ac565b6101c06004803603602081101561027a57600080fd5b50356001600160a01b03166106d6565b610189600480360360408110156102a057600080fd5b50803515159060200135610792565b610189600480360360408110156102c557600080fd5b508035151590602001356107b4565b610189600480360360a08110156102ea57600080fd5b5080351515906020810135906040810135906001600160a01b0360608201351690608001356107d6565b610189610918565b6101e76004803603606081101561033257600080fd5b50803515159060208101359060400135610920565b6101c06004803603602081101561035d57600080fd5b50356001600160a01b03166109b1565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b1515600090815260016020908152604080832083805290915290206003015490565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b03163314610444576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91037bbb732b960811b604482015290519081900360640190fd5b61045084846000610a71565b5050505050565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b1515600090815260016020908152604080832083805290915290206004015490565b600181565b600081565b600080831180156104bd5750600082115b6104ff576040805162461bcd60e51b815260206004820152600e60248201526d7361766520796f75722074696d6560901b604482015290519081900360640190fd5b600160781b831080156105155750600160781b82105b61055d576040805162461bcd60e51b815260206004820152601460248201527367726561746572207468616e20737570706c793f60601b604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a15158652835281852095855283835281516bffffffffffffffffffffffff1960608d901b1681850152603481019190915260548101899052607480820189905282518083039091018152609490910191829052805190928291908401908083835b602083106106045780518252601f1990920191602091820191016105e5565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610643573d6000803e3d6000fd5b5050506040513d602081101561065857600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b031615610749576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b0392909216919091179055565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b60006107e28683610457565b610823576040805162461bcd60e51b815260206004820152600d60248201526c7361766520796f75722067617360981b604482015290519081900360640190fd5b851515600090815260016020526040812090610841858989896104ac565b905060008415610851578461085a565b61085a89610480565b9050610867898383610920565b6108a7575b610877898383610920565b6108925760009081526020839052604090206004015461086c565b61089d898383610c2f565b50915061090f9050565b6000908152602083905260409020600301545b80158015906108cf57506108cf898383610920565b156108eb576000908152602083905260409020600301546108ba565b60008181526020849052604090206004015461090a908a908490610c2f565b509150505b95945050505050565b600160781b81565b8215156000908152600160208181526040808420868552918290528084208585529084206004546002820154948301549394929391928692610979929161096d919063ffffffff610c7616565b9063ffffffff610c7616565b905060006109a26003546004540161096d86600201548660010154610c7690919063ffffffff16565b90911198975050505050505050565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b031615610a26576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b0319166001600160a01b0392909216919091179055565b821515600090815260016020908152604080832085845291829052808320600480820154600380840180548852858820909301829055915490865292852001919091558315610b545785156000908152602081815260408083205484546002860154835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610b2257600080fd5b505af1158015610b36573d6000803e3d6000fd5b505050506040513d6020811015610b4c57600080fd5b50610bea9050565b8515156000908152602081815260408083205484546001860154835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610bbd57600080fd5b505af1158015610bd1573d6000803e3d6000fd5b505050506040513d6020811015610be757600080fd5b50505b60049081015460009586526020929092526040852080546001600160a01b03191681556001810186905560028101869055600381018690550193909355509092915050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b600082610c855750600061047a565b82820282848281610c9257fe5b0414610c9d57600080fd5b939250505056fea265627a7a723058200bb8dd94e3a16b02cf816fe44194d9b3ffd7a23dd1343eb5c8965fab51eaecc064736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057`

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
const PairExABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"bottom\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_volatileTokenAddress\",\"type\":\"address\"},{\"name\":\"_stableTokenAddress\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrderType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"validOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_stableTokenTarget\",\"type\":\"uint256\"}],\"name\":\"orderToFill\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Stable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inflate\",\"type\":\"bool\"},{\"name\":\"_stableTokenTarget\",\"type\":\"uint256\"}],\"name\":\"absorb\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"}],\"name\":\"initOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INPUTS_MAX\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderType\",\"type\":\"bool\"},{\"name\":\"_newId\",\"type\":\"bytes32\"},{\"name\":\"_oldId\",\"type\":\"bytes32\"}],\"name\":\"betterOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Volatile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_volatileTokenAddress\",\"type\":\"address\"},{\"name\":\"_stableTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PairExBin is the compiled bytecode used for deploying new contracts.
const PairExBin = `0x6080604052600060035560016004553480156200001b57600080fd5b5060405162001e5e38038062001e5e833981810160405260408110156200004157600080fd5b5080516020909101516001600160a01b038216156200006e576200006e826001600160e01b03620000ad16565b6001600160a01b03811615620000925762000092816001600160e01b036200017416565b620000a56001600160e01b036200023f16565b505062000469565b600080805260205260008051602062001e3e833981519152546001600160a01b0316156200013c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080805260205260008051602062001e3e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600090815260205260008051602062001e1e833981519152546001600160a01b0316156200020557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600090815260205260008051602062001e1e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b620002496200043b565b600160408201908152600080527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602090815282517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f780546001600160a01b03199081166001600160a01b039384161790915582850180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f85584517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f9556060860180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fa556080870180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fb557fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f90955295517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957c80549093169316929092179055517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957d5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957e5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957f55517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274958055565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b6119a580620004796000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063aea90cb8116100b8578063ce24388f1161007c578063ce24388f14610404578063d39ecc9514610429578063d4dcff8214610469578063e3dd138f14610471578063f2a946b4146102f2578063fdeb17791461049c57610142565b8063aea90cb8146102f2578063bc96b872146102fa578063c0409a5c14610334578063c0ee0b8a1461035a578063ce11ae66146103df57610142565b80637234c79d1161010a5780637234c79d146102435780637e01b4f1146102685780638aa3f897146102a65780639157772a146102c5578063974c86b5146102c55780639b7c4e01146102cd57610142565b806307c399a314610147578063174a5088146101a15780632d34ba79146101d257806343271d791461020257806355ebd76d14610227575b600080fd5b61016c6004803603604081101561015d57600080fd5b508035151590602001356104c2565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101c0600480360360208110156101b757600080fd5b50351515610508565b60408051918252519081900360200190f35b610200600480360360408110156101e857600080fd5b506001600160a01b038135811691602001351661052a565b005b6102006004803603604081101561021857600080fd5b5080351515906020013561062e565b61022f6106b0565b604080519115158252519081900360200190f35b61022f6004803603604081101561025957600080fd5b5080351515906020013561076e565b61028d6004803603604081101561027e57600080fd5b50803515159060200135610797565b6040805192835260208301919091528051918290030190f35b6101c0600480360360208110156102bc57600080fd5b503515156108cd565b61022f6108ef565b61028d600480360360408110156102e357600080fd5b508035151590602001356108f4565b61022f610d5e565b6101c06004803603608081101561031057600080fd5b506001600160a01b0381351690602081013515159060408101359060600135610d63565b6102006004803603602081101561034a57600080fd5b50356001600160a01b0316610f8d565b6102006004803603606081101561037057600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156103a057600080fd5b8201836020820111156103b257600080fd5b803590602001918460018302840111640100000000831117156103d457600080fd5b509092509050611025565b6101c0600480360360408110156103f557600080fd5b5080351515906020013561109c565b6101c06004803603604081101561041a57600080fd5b508035151590602001356110be565b6101c0600480360360a081101561043f57600080fd5b5080351515906020810135906040810135906001600160a01b0360608201351690608001356110e0565b6101c0611222565b61022f6004803603606081101561048757600080fd5b5080351515906020810135906040013561122a565b610200600480360360208110156104b257600080fd5b50356001600160a01b03166112bb565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b1515600090815260016020908152604080832083805290915290206003015490565b61053382610f8d565b61053c816112bb565b6000808052602081905260008051602061192c83398151915254604080516366d3820360e01b815230600482015290516001600160a01b03909216926366d382039260248084019382900301818387803b15801561059957600080fd5b505af11580156105ad573d6000803e3d6000fd5b50506001600090815260208190526000805160206118e183398151915254604080516366d3820360e01b815230600482015290516001600160a01b0390921694506366d382039350602480820193929182900301818387803b15801561061257600080fd5b505af1158015610626573d6000803e3d6000fd5b505050505050565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b0316331461069d576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91037bbb732b960811b604482015290519081900360640190fd5b6106a984846000611357565b5050505050565b6001600090815260208190526000805160206118e18339815191525433906001600160a01b03168114806107065750600080805260205260008051602061192c833981519152546001600160a01b038281169116145b6107415760405162461bcd60e51b815260040180806020018281038252602b815260200180611901602b913960400191505060405180910390fd5b600160009081526020526000805160206118e1833981519152546001600160a01b03918216911614905090565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b811515600090815260016020526040812081908180806107b6886108cd565b90505b80158015906107c757508683105b156108bf57600081815260208590526040812090896107ea5781600201546107f0565b81600101545b905060008a610803578260010154610809565b82600201545b90508961081c878463ffffffff61151516565b111561088c5760006108348b8863ffffffff61152e16565b905060006108588461084c858563ffffffff61154316565b9063ffffffff61156a16565b905061086a878263ffffffff61151516565b61087a898463ffffffff61151516565b9a509a505050505050505050506108c6565b61089c858263ffffffff61151516565b94506108ae868363ffffffff61151516565b9550826004015493505050506107b9565b5093509150505b9250929050565b1515600090815260016020908152604080832083805290915290206004015490565b600181565b60008033301461093c576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60008461094a57600161094d565b60005b8015156000908152600160205260408120919250808061096c856108cd565b90505b801580159061097d57508782105b15610d50576000818152602085905260408120908a6109a05781600201546109a6565b81600101545b905060008b6109b95782600101546109bf565b82600201545b90508a6109d2868363ffffffff61151516565b11610af9576109e7868363ffffffff61151516565b95506109f9858263ffffffff61151516565b8c15600090815260208190526040808220546001870154825163b5f07ea160e01b8152600481019190915291519398506001600160a01b03169263b5f07ea19260248084019391929182900301818387803b158015610a5757600080fd5b505af1158015610a6b573d6000803e3d6000fd5b505050508b151560009081526020819052604080822054600286015482516348043f0d60e01b8152600481019190915291516001600160a01b03909116926348043f0d926024808201939182900301818387803b158015610acb57600080fd5b505af1158015610adf573d6000803e3d6000fd5b50505050610aef88856001611357565b935050505061096f565b60008080610b0d8e8963ffffffff61152e16565b90506000610b258561084c888563ffffffff61154316565b9050610b378a8263ffffffff61151516565b9950610b49898363ffffffff61151516565b98508f610b565781610b58565b805b93508f610b655780610b67565b815b925050506000808f1515151515815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b031663b5f07ea1836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610bdb57600080fd5b505af1158015610bef573d6000803e3d6000fd5b505050508d15156000908152602081905260408082205481516348043f0d60e01b81526004810185905291516001600160a01b03909116926348043f0d926024808201939182900301818387803b158015610c4957600080fd5b505af1158015610c5d573d6000803e3d6000fd5b505050508d1515600090815260208181526040808320548854825163a9059cbb60e01b81526001600160a01b03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b158015610cc457600080fd5b505af1158015610cd8573d6000803e3d6000fd5b505050506040513d6020811015610cee57600080fd5b50506001850154610d05908363ffffffff61152e16565b60018601556002850154610d1f908263ffffffff61152e16565b600286015560018501541580610d3757506002850154155b15610d4a57610d488a876000611357565b505b50505050505b509097909650945050505050565b600081565b60008083118015610d745750600082115b610db6576040805162461bcd60e51b815260206004820152600e60248201526d7361766520796f75722074696d6560901b604482015290519081900360640190fd5b600160781b83108015610dcc5750600160781b82105b610e14576040805162461bcd60e51b815260206004820152601460248201527367726561746572207468616e20737570706c793f60601b604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a15158652835281852095855283835281516bffffffffffffffffffffffff1960608d901b1681850152603481019190915260548101899052607480820189905282518083039091018152609490910191829052805190928291908401908083835b60208310610ebb5780518252601f199092019160209182019101610e9c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610efa573d6000803e3d6000fd5b5050506040513d6020811015610f0f57600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b600080805260205260008051602061192c833981519152546001600160a01b031615610fee576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600080805260205260008051602061192c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b8360006110306106b0565b9050846000806020861461105b578686604081101561104e57600080fd5b5080359060200135611070565b8686602081101561106b57600080fd5b503560005b9092509050600061108485858589866110e0565b9050611090858261158c565b50505050505050505050565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b60006110ec868361076e565b61112d576040805162461bcd60e51b815260206004820152600d60248201526c7361766520796f75722067617360981b604482015290519081900360640190fd5b85151560009081526001602052604081209061114b85898989610d63565b90506000841561115b5784611164565b611164896108cd565b905061117189838361122a565b6111b1575b61118189838361122a565b61119c57600090815260208390526040902060040154611176565b6111a7898383611883565b5091506112199050565b6000908152602083905260409020600301545b80158015906111d957506111d989838361122a565b156111f5576000908152602083905260409020600301546111c4565b600081815260208490526040902060040154611214908a908490611883565b509150505b95945050505050565b600160781b81565b82151560009081526001602081815260408084208685529182905280842085855290842060045460028201549483015493949293919286926112839291611277919063ffffffff61154316565b9063ffffffff61154316565b905060006112ac600354600454016112778660020154866001015461154390919063ffffffff16565b90911198975050505050505050565b600160009081526020526000805160206118e1833981519152546001600160a01b03161561131e576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600160009081526020526000805160206118e183398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b82151560009081526001602090815260408083208584529182905280832060048082015460038084018054885285882090930182905591549086529285200191909155831561143a5785156000908152602081815260408083205484546002860154835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b15801561140857600080fd5b505af115801561141c573d6000803e3d6000fd5b505050506040513d602081101561143257600080fd5b506114d09050565b8515156000908152602081815260408083205484546001860154835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b1580156114a357600080fd5b505af11580156114b7573d6000803e3d6000fd5b505050506040513d60208110156114cd57600080fd5b50505b60049081015460009586526020929092526040852080546001600160a01b03191681556001810186905560028101869055600381018690550193909355509092915050565b60008282018381101561152757600080fd5b9392505050565b60008282111561153d57600080fd5b50900390565b60008261155257506000610791565b8282028284828161155f57fe5b041461152757600080fd5b600080821161157857600080fd5b600082848161158357fe5b04949350505050565b81158015600090815260016020818152604080842086855280835281852086865293909252832090926115be856108cd565b90505b8015611879576000818152602083905260409020600280820154908501546115ee9163ffffffff61154316565b600180830154908601546116079163ffffffff61154316565b10156116185750505050505061187f565b600061162c856001015483600201546118ca565b600186015490915061165c9061084c61164b828563ffffffff61152e16565b60028901549063ffffffff61154316565b60028601556001850154611676908263ffffffff61152e16565b8560010181905550600061169f836002015461084c84866001015461154390919063ffffffff16565b60018401549091506116cf9061084c6116be828563ffffffff61152e16565b60028701549063ffffffff61154316565b600284015560018301546116e9908263ffffffff61152e16565b6001840155871515600090815260208181526040808320548954825163a9059cbb60e01b81526001600160a01b03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b15801561175157600080fd5b505af1158015611765573d6000803e3d6000fd5b505050506040513d602081101561177b57600080fd5b5050891515600090815260208181526040808320548654825163a9059cbb60e01b81526001600160a01b03918216600482015260248101889052925191169363a9059cbb93604480850194919392918390030190829087803b1580156117e057600080fd5b505af11580156117f4573d6000803e3d6000fd5b505050506040513d602081101561180a57600080fd5b50506001830154158061181f57506002830154155b156118325761183088856000611357565b505b61183b886108cd565b935085600101546000148061185257506002860154155b15611871576118638a8a6000611357565b50505050505050505061187f565b5050506115c1565b50505050505b5050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b60008183106118d95781611527565b509091905056feada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d6f6e6c7920566f6c6174696c65546f6b656e20616e6420537461626c65546f6b656e206163636570746564ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72305820e91619e5601ce0a3dae87450db899cf2e4a562fa3651487d3f5add4b1df6781464736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5`

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820bf274441fabcf0136282726c8598d0b5ae6ee760145de173920f9c5ce325ef4364736f6c637827302e352e392d646576656c6f702e323031392e352e31362b636f6d6d69742e65323066626433380057`

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
