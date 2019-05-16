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
const DataSetBin = `0x6080604052348015600f57600080fd5b50606380601d6000396000f3fe6080604052600080fdfea265627a7a7230582087916a625d35fe56faed3747e7a2c5f1537d52af14d6b7b0ca4073d04e5e853f64736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057`

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
const InitializerBin = `0x608060405234801561001057600080fd5b506102c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80639157772a14610067578063974c86b514610067578063aea90cb814610083578063c0409a5c1461008b578063f2a946b414610083578063fdeb1779146100b3575b600080fd5b61006f6100d9565b604080519115158252519081900360200190f35b61006f6100de565b6100b1600480360360208110156100a157600080fd5b50356001600160a01b03166100e3565b005b6100b1600480360360208110156100c957600080fd5b50356001600160a01b03166101a5565b600181565b600081565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b03161561015c5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b0392909216919091179055565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b0316156102205760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723058204901482737c54aa795049232552c78bf90d7c698136c287dfc7d6905729c418e64736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057`

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
const MathBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058206525c1ec3b64d9ac81932c4ae41636b82617e580901cddc4bd1609bdc17508f064736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057`

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
const OrderBookBin = `0x608060405260006003556001600455610d2c8061001d6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063bc96b872116100a2578063d39ecc9511610071578063d39ecc95146102d4578063d4dcff8214610314578063e3dd138f1461031c578063f2a946b414610222578063fdeb1779146103475761010b565b8063bc96b8721461022a578063c0409a5c14610264578063ce11ae661461028a578063ce24388f146102af5761010b565b80638aa3f897116100de5780638aa3f897146101fb5780639157772a1461021a578063974c86b51461021a578063aea90cb8146102225761010b565b806307c399a314610110578063174a50881461016a57806343271d791461019b5780637234c79d146101c2575b600080fd5b6101356004803603604081101561012657600080fd5b5080351515906020013561036d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101896004803603602081101561018057600080fd5b503515156103b3565b60408051918252519081900360200190f35b6101c0600480360360408110156101b157600080fd5b508035151590602001356103d5565b005b6101e7600480360360408110156101d857600080fd5b50803515159060200135610467565b604080519115158252519081900360200190f35b6101896004803603602081101561021157600080fd5b50351515610490565b6101e76104b2565b6101e76104b7565b6101896004803603608081101561024057600080fd5b506001600160a01b03813516906020810135151590604081013590606001356104bc565b6101c06004803603602081101561027a57600080fd5b50356001600160a01b03166106f7565b610189600480360360408110156102a057600080fd5b508035151590602001356107b9565b610189600480360360408110156102c557600080fd5b508035151590602001356107db565b610189600480360360a08110156102ea57600080fd5b5080351515906020810135906040810135906001600160a01b0360608201351690608001356107fd565b61018961093a565b6101e76004803603606081101561033257600080fd5b50803515159060208101359060400135610942565b6101c06004803603602081101561035d57600080fd5b50356001600160a01b03166109d3565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b1515600090815260016020908152604080832083805290915290206003015490565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b031633146104545760408051600160e51b62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b61046084846000610a99565b5050505050565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b1515600090815260016020908152604080832083805290915290206004015490565b600181565b600081565b600080831180156104cd5750600082115b6105215760408051600160e51b62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b600160781b831080156105375750600160781b82105b61058b5760408051600160e51b62461bcd02815260206004820152601460248201527f67726561746572207468616e20737570706c793f000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a151586528352818520868652848452825160609790971b8785015260348701919091526054860189905260748087018990528251808803909101815260949096019182905285519095928291908401908083835b602083106106255780518252601f199092019160209182019101610606565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610664573d6000803e3d6000fd5b5050506040513d602081101561067957600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b0316156107705760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b0392909216919091179055565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b60006108098683610467565b61085d5760408051600160e51b62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b85151560009081526001602052604081209061087b858989896104bc565b905083610889898383610942565b6108c9575b610899898383610942565b6108b45760009081526020839052604090206004015461088e565b6108bf898383610c5d565b5091506109319050565b6000908152602083905260409020600301545b80158015906108f157506108f1898383610942565b1561090d576000908152602083905260409020600301546108dc565b60008181526020849052604090206004015461092c908a908490610c5d565b509150505b95945050505050565b600160781b81565b821515600090815260016020818152604080842086855291829052808420858552908420600454600282015494830154939492939192869261099b929161098f919063ffffffff610ca416565b9063ffffffff610ca416565b905060006109c46003546004540161098f86600201548660010154610ca490919063ffffffff16565b90911198975050505050505050565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b031615610a4e5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b0319166001600160a01b0392909216919091179055565b821515600090815260016020908152604080832085845291829052808320600480820154600380840180548852858820909301829055915490865292852001919091558315610b7f57851560009081526020818152604080832054845460028601548351600160e01b63a9059cbb0281526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d6020811015610b7757600080fd5b50610c189050565b85151560009081526020818152604080832054845460018601548351600160e01b63a9059cbb0281526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610beb57600080fd5b505af1158015610bff573d6000803e3d6000fd5b505050506040513d6020811015610c1557600080fd5b50505b60049081015460009586526020929092526040852080546001600160a01b03191681556001810186905560028101869055600381018690550193909355509092915050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b600082610cb35750600061048a565b82820282848281610cc057fe5b0414610ccb57600080fd5b939250505056fea265627a7a72305820665e7b22c1244de0b63d72f443a9fb9bf08a9a63966c05c02e0828f820c23c3464736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057`

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
const PairExBin = `0x6080604052600060035560016004553480156200001b57600080fd5b5060405160408062001eac833981018060405260408110156200003d57600080fd5b5080516020909101516001600160a01b0382161562000067576200006782620000a060201b60201c565b6001600160a01b03811615620000885762000088816200016760201b60201c565b620000986200023260201b60201c565b50506200045c565b600080805260205260008051602062001e8c833981519152546001600160a01b0316156200012f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080805260205260008051602062001e8c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600090815260205260008051602062001e6c833981519152546001600160a01b031615620001f857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f616c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600090815260205260008051602062001e6c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6200023c6200042e565b600160408201908152600080527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602090815282517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f780546001600160a01b03199081166001600160a01b039384161790915582850180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f85584517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f9556060860180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fa556080870180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fb557fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f90955295517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957c80549093169316929092179055517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957d5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957e5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957f55517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274958055565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b611a00806200046c6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063aea90cb8116100b8578063ce24388f1161007c578063ce24388f14610404578063d39ecc9514610429578063d4dcff8214610469578063e3dd138f14610471578063f2a946b4146102f2578063fdeb17791461049c57610142565b8063aea90cb8146102f2578063bc96b872146102fa578063c0409a5c14610334578063c0ee0b8a1461035a578063ce11ae66146103df57610142565b80637234c79d1161010a5780637234c79d146102435780637e01b4f1146102685780638aa3f897146102a65780639157772a146102c5578063974c86b5146102c55780639b7c4e01146102cd57610142565b806307c399a314610147578063174a5088146101a15780632d34ba79146101d257806343271d791461020257806355ebd76d14610227575b600080fd5b61016c6004803603604081101561015d57600080fd5b508035151590602001356104c2565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101c0600480360360208110156101b757600080fd5b50351515610508565b60408051918252519081900360200190f35b610200600480360360408110156101e857600080fd5b506001600160a01b038135811691602001351661052a565b005b6102006004803603604081101561021857600080fd5b50803515159060200135610634565b61022f6106c6565b604080519115158252519081900360200190f35b61022f6004803603604081101561025957600080fd5b50803515159060200135610787565b61028d6004803603604081101561027e57600080fd5b508035151590602001356107b0565b6040805192835260208301919091528051918290030190f35b6101c0600480360360208110156102bc57600080fd5b503515156108e6565b61022f610908565b61028d600480360360408110156102e357600080fd5b5080351515906020013561090d565b61022f610d95565b6101c06004803603608081101561031057600080fd5b506001600160a01b0381351690602081013515159060408101359060600135610d9a565b6102006004803603602081101561034a57600080fd5b50356001600160a01b0316610fd5565b6102006004803603606081101561037057600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156103a057600080fd5b8201836020820111156103b257600080fd5b803590602001918460018302840111640100000000831117156103d457600080fd5b509092509050611073565b6101c0600480360360408110156103f557600080fd5b508035151590602001356110ea565b6101c06004803603604081101561041a57600080fd5b5080351515906020013561110c565b6101c0600480360360a081101561043f57600080fd5b5080351515906020810135906040810135906001600160a01b03606082013516906080013561112e565b6101c061126b565b61022f6004803603606081101561048757600080fd5b50803515159060208101359060400135611273565b610200600480360360208110156104b257600080fd5b50356001600160a01b0316611304565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b1515600090815260016020908152604080832083805290915290206003015490565b61053382610fd5565b61053c81611304565b600080805260208190526000805160206119878339815191525460408051600160e01b6366d3820302815230600482015290516001600160a01b03909216926366d382039260248084019382900301818387803b15801561059c57600080fd5b505af11580156105b0573d6000803e3d6000fd5b505060016000908152602081905260008051602061193c8339815191525460408051600160e01b6366d3820302815230600482015290516001600160a01b0390921694506366d382039350602480820193929182900301818387803b15801561061857600080fd5b505af115801561062c573d6000803e3d6000fd5b505050505050565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b031633146106b35760408051600160e51b62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b6106bf848460006113a6565b5050505050565b60016000908152602081905260008051602061193c8339815191525433906001600160a01b031681148061071c57506000808052602052600080516020611987833981519152546001600160a01b038281169116145b61075a57604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061195c602b913960400191505060405180910390fd5b6001600090815260205260008051602061193c833981519152546001600160a01b03918216911614905090565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b811515600090815260016020526040812081908180806107cf886108e6565b90505b80158015906107e057508683105b156108d85760008181526020859052604081209089610803578160020154610809565b81600101545b905060008a61081c578260010154610822565b82600201545b905089610835878463ffffffff61156a16565b11156108a557600061084d8b8863ffffffff61158316565b9050600061087184610865858563ffffffff61159816565b9063ffffffff6115bf16565b9050610883878263ffffffff61156a16565b610893898463ffffffff61156a16565b9a509a505050505050505050506108df565b6108b5858263ffffffff61156a16565b94506108c7868363ffffffff61156a16565b9550826004015493505050506107d2565b5093509150505b9250929050565b1515600090815260016020908152604080832083805290915290206004015490565b600181565b6000803330146109675760408051600160e51b62461bcd02815260206004820152600e60248201527f636f6e73656e737573206f6e6c79000000000000000000000000000000000000604482015290519081900360640190fd5b600084610975576001610978565b60005b80151560009081526001602052604081209192508080610997856108e6565b90505b80158015906109a857508782105b15610d87576000818152602085905260408120908a6109cb5781600201546109d1565b81600101545b905060008b6109e45782600101546109ea565b82600201545b90508a6109fd868363ffffffff61156a16565b11610b2a57610a12868363ffffffff61156a16565b9550610a24858263ffffffff61156a16565b8c156000908152602081905260408082205460018701548251600160e01b63b5f07ea1028152600481019190915291519398506001600160a01b03169263b5f07ea19260248084019391929182900301818387803b158015610a8557600080fd5b505af1158015610a99573d6000803e3d6000fd5b505050508b15156000908152602081905260408082205460028601548251600160e01b6348043f0d028152600481019190915291516001600160a01b03909116926348043f0d926024808201939182900301818387803b158015610afc57600080fd5b505af1158015610b10573d6000803e3d6000fd5b50505050610b20888560016113a6565b935050505061099a565b60008080610b3e8e8963ffffffff61158316565b90506000610b5685610865888563ffffffff61159816565b9050610b688a8263ffffffff61156a16565b9950610b7a898363ffffffff61156a16565b98508f610b875781610b89565b805b93508f610b965780610b98565b815b925050506000808f1515151515815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b031663b5f07ea1836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610c0c57600080fd5b505af1158015610c20573d6000803e3d6000fd5b505050508d1515600090815260208190526040808220548151600160e01b6348043f0d0281526004810185905291516001600160a01b03909116926348043f0d926024808201939182900301818387803b158015610c7d57600080fd5b505af1158015610c91573d6000803e3d6000fd5b505050508d15156000908152602081815260408083205488548251600160e01b63a9059cbb0281526001600160a01b03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b158015610cfb57600080fd5b505af1158015610d0f573d6000803e3d6000fd5b505050506040513d6020811015610d2557600080fd5b50506001850154610d3c908363ffffffff61158316565b60018601556002850154610d56908263ffffffff61158316565b600286015560018501541580610d6e57506002850154155b15610d8157610d7f8a8760006113a6565b505b50505050505b509097909650945050505050565b600081565b60008083118015610dab5750600082115b610dff5760408051600160e51b62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b600160781b83108015610e155750600160781b82105b610e695760408051600160e51b62461bcd02815260206004820152601460248201527f67726561746572207468616e20737570706c793f000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a151586528352818520868652848452825160609790971b8785015260348701919091526054860189905260748087018990528251808803909101815260949096019182905285519095928291908401908083835b60208310610f035780518252601f199092019160209182019101610ee4565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610f42573d6000803e3d6000fd5b5050506040513d6020811015610f5757600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b6000808052602052600080516020611987833981519152546001600160a01b03161561103c5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b600080805260205260008051602061198783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b83600061107e6106c6565b905084600080602086146110a9578686604081101561109c57600080fd5b50803590602001356110be565b868660208110156110b957600080fd5b503560005b909250905060006110d2858585898661112e565b90506110de85826115e1565b50505050505050505050565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b600061113a8683610787565b61118e5760408051600160e51b62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b8515156000908152600160205260408120906111ac85898989610d9a565b9050836111ba898383611273565b6111fa575b6111ca898383611273565b6111e5576000908152602083905260409020600401546111bf565b6111f08983836118de565b5091506112629050565b6000908152602083905260409020600301545b80158015906112225750611222898383611273565b1561123e5760009081526020839052604090206003015461120d565b60008181526020849052604090206004015461125d908a9084906118de565b509150505b95945050505050565b600160781b81565b82151560009081526001602081815260408084208685529182905280842085855290842060045460028201549483015493949293919286926112cc92916112c0919063ffffffff61159816565b9063ffffffff61159816565b905060006112f5600354600454016112c08660020154866001015461159890919063ffffffff16565b90911198975050505050505050565b6001600090815260205260008051602061193c833981519152546001600160a01b03161561136d5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b6001600090815260205260008051602061193c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b82151560009081526001602090815260408083208584529182905280832060048082015460038084018054885285882090930182905591549086529285200191909155831561148c57851560009081526020818152604080832054845460028601548351600160e01b63a9059cbb0281526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b15801561145a57600080fd5b505af115801561146e573d6000803e3d6000fd5b505050506040513d602081101561148457600080fd5b506115259050565b85151560009081526020818152604080832054845460018601548351600160e01b63a9059cbb0281526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b1580156114f857600080fd5b505af115801561150c573d6000803e3d6000fd5b505050506040513d602081101561152257600080fd5b50505b60049081015460009586526020929092526040852080546001600160a01b03191681556001810186905560028101869055600381018690550193909355509092915050565b60008282018381101561157c57600080fd5b9392505050565b60008282111561159257600080fd5b50900390565b6000826115a7575060006107aa565b828202828482816115b457fe5b041461157c57600080fd5b60008082116115cd57600080fd5b60008284816115d857fe5b04949350505050565b8115801560009081526001602081815260408084208685528083528185208686529390925283209092611613856108e6565b90505b80156118d4576000818152602083905260409020600280820154908501546116439163ffffffff61159816565b6001808301549086015461165c9163ffffffff61159816565b101561166d575050505050506118da565b600061168185600101548360020154611925565b60018601549091506116b1906108656116a0828563ffffffff61158316565b60028901549063ffffffff61159816565b600286015560018501546116cb908263ffffffff61158316565b856001018190555060006116f4836002015461086584866001015461159890919063ffffffff16565b600184015490915061172490610865611713828563ffffffff61158316565b60028701549063ffffffff61159816565b6002840155600183015461173e908263ffffffff61158316565b60018401558715156000908152602081815260408083205489548251600160e01b63a9059cbb0281526001600160a01b03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b1580156117a957600080fd5b505af11580156117bd573d6000803e3d6000fd5b505050506040513d60208110156117d357600080fd5b50508915156000908152602081815260408083205486548251600160e01b63a9059cbb0281526001600160a01b03918216600482015260248101889052925191169363a9059cbb93604480850194919392918390030190829087803b15801561183b57600080fd5b505af115801561184f573d6000803e3d6000fd5b505050506040513d602081101561186557600080fd5b50506001830154158061187a57506002830154155b1561188d5761188b888560006113a6565b505b611896886108e6565b93508560010154600014806118ad57506002860154155b156118cc576118be8a8a60006113a6565b5050505050505050506118da565b505050611616565b50505050505b5050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b6000818310611934578161157c565b509091905056feada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d6f6e6c7920566f6c6174696c65546f6b656e20616e6420537461626c65546f6b656e206163636570746564ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a723058203374c44f5671c08ede448b914f7e6bcd55eed760362e58ff1571e78e05d214b164736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5`

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
const SafeMathBin = `0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582011615a337b3035a376b7c2e7299126cde2b3f5ab3ea772763fe16568f6196e9864736f6c637827302e352e392d646576656c6f702e323031392e352e31342b636f6d6d69742e63386464343132300057`

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
