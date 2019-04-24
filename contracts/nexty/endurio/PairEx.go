// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endurio

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

// ABIABI is the input ABI used to generate the binding from.
const ABIABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"encode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ABIBin is the compiled bytecode used for deploying new contracts.
const ABIBin = `0x610272610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80633915d06014610045578063e5c5e9a3146100dd575b600080fd5b6100686004803603604081101561005b57600080fd5b508035906020013561019c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100a257818101518382015260200161008a565b50505050905090810190601f1680156100cf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610183600480360360208110156100f357600080fd5b81019060208101813564010000000081111561010e57600080fd5b82018360208201111561012057600080fd5b8035906020019184600183028401116401000000008311171561014257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506101c2945050505050565b6040805192835260208301919091528051918290030190f35b604080516020810193909352828101919091528051808303820181526060909201905290565b600080825160401461021e5760408051600160e51b62461bcd02815260206004820152600c60248201527f696e76616c696420646174610000000000000000000000000000000000000000604482015290519081900360640190fd5b82806020019051604081101561023357600080fd5b508051602090910151909250905091509156fea165627a7a723058208e18fe5fa25de21e3a404d0adff43db26d13f2a452360226a132b1ea1632ae0d0029`

// DeployABI deploys a new Ethereum contract, binding an instance of ABI to it.
func DeployABI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ABIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ABIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ABI{ABICaller: ABICaller{contract: contract}, ABITransactor: ABITransactor{contract: contract}, ABIFilterer: ABIFilterer{contract: contract}}, nil
}

// ABI is an auto generated Go binding around an Ethereum contract.
type ABI struct {
	ABICaller     // Read-only binding to the contract
	ABITransactor // Write-only binding to the contract
	ABIFilterer   // Log filterer for contract events
}

// ABICaller is an auto generated read-only Go binding around an Ethereum contract.
type ABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ABISession struct {
	Contract     *ABI              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ABICallerSession struct {
	Contract *ABICaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ABITransactorSession struct {
	Contract     *ABITransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type ABIRaw struct {
	Contract *ABI // Generic contract binding to access the raw methods on
}

// ABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ABICallerRaw struct {
	Contract *ABICaller // Generic read-only contract binding to access the raw methods on
}

// ABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ABITransactorRaw struct {
	Contract *ABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewABI creates a new instance of ABI, bound to a specific deployed contract.
func NewABI(address common.Address, backend bind.ContractBackend) (*ABI, error) {
	contract, err := bindABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ABI{ABICaller: ABICaller{contract: contract}, ABITransactor: ABITransactor{contract: contract}, ABIFilterer: ABIFilterer{contract: contract}}, nil
}

// NewABICaller creates a new read-only instance of ABI, bound to a specific deployed contract.
func NewABICaller(address common.Address, caller bind.ContractCaller) (*ABICaller, error) {
	contract, err := bindABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ABICaller{contract: contract}, nil
}

// NewABITransactor creates a new write-only instance of ABI, bound to a specific deployed contract.
func NewABITransactor(address common.Address, transactor bind.ContractTransactor) (*ABITransactor, error) {
	contract, err := bindABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ABITransactor{contract: contract}, nil
}

// NewABIFilterer creates a new log filterer instance of ABI, bound to a specific deployed contract.
func NewABIFilterer(address common.Address, filterer bind.ContractFilterer) (*ABIFilterer, error) {
	contract, err := bindABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ABIFilterer{contract: contract}, nil
}

// bindABI binds a generic wrapper to an already deployed contract.
func bindABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABI *ABIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ABI.Contract.ABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABI *ABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABI.Contract.ABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABI *ABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABI.Contract.ABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABI *ABICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABI *ABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABI *ABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABI.Contract.contract.Transact(opts, method, params...)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes _data) constant returns(uint256, bytes32)
func (_ABI *ABICaller) Decode(opts *bind.CallOpts, _data []byte) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ABI.contract.Call(opts, out, "decode", _data)
	return *ret0, *ret1, err
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes _data) constant returns(uint256, bytes32)
func (_ABI *ABISession) Decode(_data []byte) (*big.Int, [32]byte, error) {
	return _ABI.Contract.Decode(&_ABI.CallOpts, _data)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes _data) constant returns(uint256, bytes32)
func (_ABI *ABICallerSession) Decode(_data []byte) (*big.Int, [32]byte, error) {
	return _ABI.Contract.Decode(&_ABI.CallOpts, _data)
}

// Encode is a free data retrieval call binding the contract method 0x3915d060.
//
// Solidity: function encode(uint256 wantAmount, bytes32 assistingID) constant returns(bytes)
func (_ABI *ABICaller) Encode(opts *bind.CallOpts, wantAmount *big.Int, assistingID [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ABI.contract.Call(opts, out, "encode", wantAmount, assistingID)
	return *ret0, err
}

// Encode is a free data retrieval call binding the contract method 0x3915d060.
//
// Solidity: function encode(uint256 wantAmount, bytes32 assistingID) constant returns(bytes)
func (_ABI *ABISession) Encode(wantAmount *big.Int, assistingID [32]byte) ([]byte, error) {
	return _ABI.Contract.Encode(&_ABI.CallOpts, wantAmount, assistingID)
}

// Encode is a free data retrieval call binding the contract method 0x3915d060.
//
// Solidity: function encode(uint256 wantAmount, bytes32 assistingID) constant returns(bytes)
func (_ABI *ABICallerSession) Encode(wantAmount *big.Int, assistingID [32]byte) ([]byte, error) {
	return _ABI.Contract.Encode(&_ABI.CallOpts, wantAmount, assistingID)
}

// BytesConvertABI is the input ABI used to generate the binding from.
const BytesConvertABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"uint256ToBytes\",\"outputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint256\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BytesConvertBin is the compiled bytecode used for deploying new contracts.
const BytesConvertBin = `0x610330610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80635358493914610050578063c7559da41461010a578063cc9de25d1461019c575b600080fd5b6100f86004803603604081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111640100000000831117156100b557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610242915050565b60408051918252519081900360200190f35b6101276004803603602081101561012057600080fd5b5035610290565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610161578181015183820152602001610149565b50505050905090810190601f16801561018e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100f8600480360360208110156101b257600080fd5b8101906020810181356401000000008111156101cd57600080fd5b8201836020820111156101df57600080fd5b8035906020019184600183028401116401000000008311171561020157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102ba945050505050565b600080805b60208110156102885780600802858286018151811061026257fe5b01602001516001600160f81b031960f891821c90911b16901c9190911790600101610247565b509392505050565b60408051602080825281830190925260609160208201818038833950505060208101929092525090565b600080805b83518110156102fd578060010184510360080260020a8482815181106102e157fe5b016020015160f890811c811b901c0291909101906001016102bf565b509291505056fea165627a7a72305820fde388bc0c0b5989ee544d5b88c348beaf81bd9a1efea1dd38ac52bbeb496de40029`

// DeployBytesConvert deploys a new Ethereum contract, binding an instance of BytesConvert to it.
func DeployBytesConvert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesConvert, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesConvertABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesConvertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesConvert{BytesConvertCaller: BytesConvertCaller{contract: contract}, BytesConvertTransactor: BytesConvertTransactor{contract: contract}, BytesConvertFilterer: BytesConvertFilterer{contract: contract}}, nil
}

// BytesConvert is an auto generated Go binding around an Ethereum contract.
type BytesConvert struct {
	BytesConvertCaller     // Read-only binding to the contract
	BytesConvertTransactor // Write-only binding to the contract
	BytesConvertFilterer   // Log filterer for contract events
}

// BytesConvertCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesConvertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesConvertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesConvertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesConvertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesConvertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesConvertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesConvertSession struct {
	Contract     *BytesConvert     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesConvertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesConvertCallerSession struct {
	Contract *BytesConvertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BytesConvertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesConvertTransactorSession struct {
	Contract     *BytesConvertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BytesConvertRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesConvertRaw struct {
	Contract *BytesConvert // Generic contract binding to access the raw methods on
}

// BytesConvertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesConvertCallerRaw struct {
	Contract *BytesConvertCaller // Generic read-only contract binding to access the raw methods on
}

// BytesConvertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesConvertTransactorRaw struct {
	Contract *BytesConvertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesConvert creates a new instance of BytesConvert, bound to a specific deployed contract.
func NewBytesConvert(address common.Address, backend bind.ContractBackend) (*BytesConvert, error) {
	contract, err := bindBytesConvert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesConvert{BytesConvertCaller: BytesConvertCaller{contract: contract}, BytesConvertTransactor: BytesConvertTransactor{contract: contract}, BytesConvertFilterer: BytesConvertFilterer{contract: contract}}, nil
}

// NewBytesConvertCaller creates a new read-only instance of BytesConvert, bound to a specific deployed contract.
func NewBytesConvertCaller(address common.Address, caller bind.ContractCaller) (*BytesConvertCaller, error) {
	contract, err := bindBytesConvert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesConvertCaller{contract: contract}, nil
}

// NewBytesConvertTransactor creates a new write-only instance of BytesConvert, bound to a specific deployed contract.
func NewBytesConvertTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesConvertTransactor, error) {
	contract, err := bindBytesConvert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesConvertTransactor{contract: contract}, nil
}

// NewBytesConvertFilterer creates a new log filterer instance of BytesConvert, bound to a specific deployed contract.
func NewBytesConvertFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesConvertFilterer, error) {
	contract, err := bindBytesConvert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesConvertFilterer{contract: contract}, nil
}

// bindBytesConvert binds a generic wrapper to an already deployed contract.
func bindBytesConvert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesConvertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesConvert *BytesConvertRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesConvert.Contract.BytesConvertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesConvert *BytesConvertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesConvert.Contract.BytesConvertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesConvert *BytesConvertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesConvert.Contract.BytesConvertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesConvert *BytesConvertCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesConvert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesConvert *BytesConvertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesConvert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesConvert *BytesConvertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesConvert.Contract.contract.Transact(opts, method, params...)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0x53584939.
//
// Solidity: function bytesToBytes32(bytes b, uint256 offset) constant returns(bytes32)
func (_BytesConvert *BytesConvertCaller) BytesToBytes32(opts *bind.CallOpts, b []byte, offset *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BytesConvert.contract.Call(opts, out, "bytesToBytes32", b, offset)
	return *ret0, err
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0x53584939.
//
// Solidity: function bytesToBytes32(bytes b, uint256 offset) constant returns(bytes32)
func (_BytesConvert *BytesConvertSession) BytesToBytes32(b []byte, offset *big.Int) ([32]byte, error) {
	return _BytesConvert.Contract.BytesToBytes32(&_BytesConvert.CallOpts, b, offset)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0x53584939.
//
// Solidity: function bytesToBytes32(bytes b, uint256 offset) constant returns(bytes32)
func (_BytesConvert *BytesConvertCallerSession) BytesToBytes32(b []byte, offset *big.Int) ([32]byte, error) {
	return _BytesConvert.Contract.BytesToBytes32(&_BytesConvert.CallOpts, b, offset)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0xcc9de25d.
//
// Solidity: function bytesToUint256(bytes b) constant returns(uint256)
func (_BytesConvert *BytesConvertCaller) BytesToUint256(opts *bind.CallOpts, b []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BytesConvert.contract.Call(opts, out, "bytesToUint256", b)
	return *ret0, err
}

// BytesToUint256 is a free data retrieval call binding the contract method 0xcc9de25d.
//
// Solidity: function bytesToUint256(bytes b) constant returns(uint256)
func (_BytesConvert *BytesConvertSession) BytesToUint256(b []byte) (*big.Int, error) {
	return _BytesConvert.Contract.BytesToUint256(&_BytesConvert.CallOpts, b)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0xcc9de25d.
//
// Solidity: function bytesToUint256(bytes b) constant returns(uint256)
func (_BytesConvert *BytesConvertCallerSession) BytesToUint256(b []byte) (*big.Int, error) {
	return _BytesConvert.Contract.BytesToUint256(&_BytesConvert.CallOpts, b)
}

// Uint256ToBytes is a free data retrieval call binding the contract method 0xc7559da4.
//
// Solidity: function uint256ToBytes(uint256 x) constant returns(bytes b)
func (_BytesConvert *BytesConvertCaller) Uint256ToBytes(opts *bind.CallOpts, x *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _BytesConvert.contract.Call(opts, out, "uint256ToBytes", x)
	return *ret0, err
}

// Uint256ToBytes is a free data retrieval call binding the contract method 0xc7559da4.
//
// Solidity: function uint256ToBytes(uint256 x) constant returns(bytes b)
func (_BytesConvert *BytesConvertSession) Uint256ToBytes(x *big.Int) ([]byte, error) {
	return _BytesConvert.Contract.Uint256ToBytes(&_BytesConvert.CallOpts, x)
}

// Uint256ToBytes is a free data retrieval call binding the contract method 0xc7559da4.
//
// Solidity: function uint256ToBytes(uint256 x) constant returns(bytes b)
func (_BytesConvert *BytesConvertCallerSession) Uint256ToBytes(x *big.Int) ([]byte, error) {
	return _BytesConvert.Contract.Uint256ToBytes(&_BytesConvert.CallOpts, x)
}

// DataSetABI is the input ABI used to generate the binding from.
const DataSetABI = "[]"

// DataSetBin is the compiled bytecode used for deploying new contracts.
const DataSetBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a72305820b555761a200b9628e74c59d10420735f8a27a7558564e66dc01874f1113850490029`

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
const IOwnableERC223ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnFromOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IOwnableERC223 *IOwnableERC223Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IOwnableERC223 *IOwnableERC223Session) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Initialize(&_IOwnableERC223.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IOwnableERC223 *IOwnableERC223TransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IOwnableERC223.Contract.Initialize(&_IOwnableERC223.TransactOpts, _owner)
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
const InitializerABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// InitializerBin is the compiled bytecode used for deploying new contracts.
const InitializerBin = `0x608060405234801561001057600080fd5b506101e5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631264a9901461003b5780634f657bb114610045575b600080fd5b61004361004d565b005b610043610105565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b0316156100c85760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b03191633179055565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b03161561017e5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319163317905556fea165627a7a7230582081fc50ad97f15c48c41e333bfe6ab4c05507e8df11be473d29c9c80fc4e6364f0029`

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

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_Initializer *InitializerTransactor) StableTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializer.contract.Transact(opts, "stableTokenRegister")
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_Initializer *InitializerSession) StableTokenRegister() (*types.Transaction, error) {
	return _Initializer.Contract.StableTokenRegister(&_Initializer.TransactOpts)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_Initializer *InitializerTransactorSession) StableTokenRegister() (*types.Transaction, error) {
	return _Initializer.Contract.StableTokenRegister(&_Initializer.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_Initializer *InitializerTransactor) VolatileTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializer.contract.Transact(opts, "volatileTokenRegister")
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_Initializer *InitializerSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _Initializer.Contract.VolatileTokenRegister(&_Initializer.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_Initializer *InitializerTransactorSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _Initializer.Contract.VolatileTokenRegister(&_Initializer.TransactOpts)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582030d2a02244ecb5badf75c4d176c95b0a925cfcc3862733dfd0a2e83b6f1a376f0029`

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
const OrderBookABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"bottom\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"validOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"}],\"name\":\"initOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderType\",\"type\":\"bool\"},{\"name\":\"_newId\",\"type\":\"bytes32\"},{\"name\":\"_oldId\",\"type\":\"bytes32\"}],\"name\":\"betterOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OrderBookBin is the compiled bytecode used for deploying new contracts.
const OrderBookBin = `0x60806040526001600355612710600455610ae98061001e6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa3f897116100715780638aa3f897146101b4578063bc96b872146101d3578063ce11ae661461020d578063ce24388f14610232578063d39ecc9514610257578063e3dd138f14610297576100b4565b806307c399a3146100b95780631264a99014610113578063174a50881461011d5780634f657bb11461014e57806369722e55146101565780637234c79d1461017b575b600080fd5b6100de600480360360408110156100cf57600080fd5b508035151590602001356102c2565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61011b610308565b005b61013c6004803603602081101561013357600080fd5b503515156103c0565b60408051918252519081900360200190f35b61011b6103e2565b61011b6004803603604081101561016c57600080fd5b50803515159060200135610496565b6101a06004803603604081101561019157600080fd5b5080351515906020013561060a565b604080519115158252519081900360200190f35b61013c600480360360208110156101ca57600080fd5b50351515610633565b61013c600480360360808110156101e957600080fd5b506001600160a01b0381351690602081013515159060408101359060600135610655565b61013c6004803603604081101561022357600080fd5b50803515159060200135610826565b61013c6004803603604081101561024857600080fd5b50803515159060200135610848565b61013c600480360360a081101561026d57600080fd5b5080351515906020810135906040810135906001600160a01b03606082013516906080013561086a565b6101a0600480360360608110156102ad57600080fd5b508035151590602081013590604001356109b7565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d546001600160a01b0316156103835760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b03191633179055565b1515600090815260016020908152604080832083805290915290206003015490565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b03161561045b5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b031633146105155760408051600160e51b62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b600480820154600380840180546000908152602087815260408083208701869055925494825282822090930193909355871515835282825280832054855460018701548351600160e01b63a9059cbb0281526001600160a01b03928316978101979097526024870152915191169363a9059cbb9360448083019493928390030190829087803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b505050506040513d60208110156105d157600080fd5b50505060009182526020526040812080546001600160a01b03191681556001810182905560028101829055600381018290556004015550565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b1515600090815260016020908152604080832083805290915290206004015490565b600080831180156106665750600082115b6106ba5760408051600160e51b62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a151586528352818520868652848452825160609790971b8785015260348701919091526054860189905260748087018990528251808803909101815260949096019182905285519095928291908401908083835b602083106107545780518252601f199092019160209182019101610735565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610793573d6000803e3d6000fd5b5050506040513d60208110156107a857600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b6000610876868361060a565b6108ca5760408051600160e51b62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b8515156000908152600160205260408120906108e885898989610655565b9050836108f68983836109b7565b610936575b6109068983836109b7565b610921576000908152602083905260409020600401546108fb565b61092c898383610a48565b5091506109ae9050565b6000908152602083905260409020600301545b8060001a60f81b6001600160f81b0319161580159061096e575061096e8983836109b7565b1561098a57600090815260208390526040902060030154610949565b6000818152602084905260409020600401546109a9908a908490610a48565b509150505b95945050505050565b8215156000908152600160208181526040808420868552918290528084208585529084206004546002820154948301549394929391928692610a109291610a04919063ffffffff610a8f16565b9063ffffffff610a8f16565b90506000610a3960035460045401610a0486600201548660010154610a8f90919063ffffffff16565b90911198975050505050505050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b600082610a9e5750600061062d565b82820282848281610aab57fe5b0414610ab657600080fd5b939250505056fea165627a7a7230582033236f542bc174d13246640d0a5d2e483a11a5efe4238e0c36c3ad56bbaaef4e0029`

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

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookTransactor) Remove(opts *bind.TransactOpts, _orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "remove", _orderType, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookSession) Remove(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Remove(&_OrderBook.TransactOpts, _orderType, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_OrderBook *OrderBookTransactorSession) Remove(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _OrderBook.Contract.Remove(&_OrderBook.TransactOpts, _orderType, _id)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_OrderBook *OrderBookTransactor) StableTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "stableTokenRegister")
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_OrderBook *OrderBookSession) StableTokenRegister() (*types.Transaction, error) {
	return _OrderBook.Contract.StableTokenRegister(&_OrderBook.TransactOpts)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_OrderBook *OrderBookTransactorSession) StableTokenRegister() (*types.Transaction, error) {
	return _OrderBook.Contract.StableTokenRegister(&_OrderBook.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_OrderBook *OrderBookTransactor) VolatileTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "volatileTokenRegister")
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_OrderBook *OrderBookSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _OrderBook.Contract.VolatileTokenRegister(&_OrderBook.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_OrderBook *OrderBookTransactorSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _OrderBook.Contract.VolatileTokenRegister(&_OrderBook.TransactOpts)
}

// PairExABI is the input ABI used to generate the binding from.
const PairExABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"bottom\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrderType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"validOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_stableTokenTarget\",\"type\":\"uint256\"}],\"name\":\"orderToFill\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"}],\"name\":\"initOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"bool\"},{\"name\":\"_haveAmount\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_maker\",\"type\":\"address\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderType\",\"type\":\"bool\"},{\"name\":\"_newId\",\"type\":\"bytes32\"},{\"name\":\"_oldId\",\"type\":\"bytes32\"}],\"name\":\"betterOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PairExBin is the compiled bytecode used for deploying new contracts.
const PairExBin = `0x6080604052600160035561271060045534801561001b57600080fd5b50610024610219565b600160408201908152600080527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602090815282517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f780546001600160a01b03199081166001600160a01b039384161790915582850180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f85584517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664f9556060860180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fa556080870180517fe5d06582d467054dda5404b9e1ec93f72b608a4970ba970773776c69ca5664fb557fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f90955295517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957c80549093169316929092179055517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957d5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957e5590517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274957f55517f09d41a60c7eb9e1f3f38bbee2eea2761087cd398a4df0eb22dbaa4eaa274958055610247565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b6112b5806102566000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637e01b4f111610097578063ce11ae6611610066578063ce11ae6614610319578063ce24388f1461033e578063d39ecc9514610363578063e3dd138f146103a3576100f5565b80637e01b4f1146101fd5780638aa3f8971461023b578063bc96b8721461025a578063c0ee0b8a14610294576100f5565b80634f657bb1116100d35780634f657bb11461018f57806355ebd76d1461019757806369722e55146101b35780637234c79d146101d8576100f5565b806307c399a3146100fa5780631264a99014610154578063174a50881461015e575b600080fd5b61011f6004803603604081101561011057600080fd5b508035151590602001356103ce565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61015c610414565b005b61017d6004803603602081101561017457600080fd5b503515156104a8565b60408051918252519081900360200190f35b61015c6104ca565b61019f61057e565b604080519115158252519081900360200190f35b61015c600480360360408110156101c957600080fd5b50803515159060200135610651565b61019f600480360360408110156101ee57600080fd5b508035151590602001356107c5565b6102226004803603604081101561021357600080fd5b508035151590602001356107ee565b6040805192835260208301919091528051918290030190f35b61017d6004803603602081101561025157600080fd5b50351515610934565b61017d6004803603608081101561027057600080fd5b506001600160a01b0381351690602081013515159060408101359060600135610956565b61015c600480360360608110156102aa57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156102da57600080fd5b8201836020820111156102ec57600080fd5b8035906020019184600183028401116401000000008311171561030e57600080fd5b509092509050610b27565b61017d6004803603604081101561032f57600080fd5b50803515159060200135610b9e565b61017d6004803603604081101561035457600080fd5b50803515159060200135610bc0565b61017d600480360360a081101561037957600080fd5b5080351515906020810135906040810135906001600160a01b036060820135169060800135610be2565b61019f600480360360608110156103b957600080fd5b50803515159060208101359060400135610d2f565b9015156000908152600160208181526040808420948452939052919020805491810154600282015460038301546004909301546001600160a01b03909416949193909291565b6001600090815260205260008051602061123f833981519152546001600160a01b03161561047d5760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b6001600090815260205260008051602061123f83398151915280546001600160a01b03191633179055565b1515600090815260016020908152604080832083805290915290206003015490565b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b0316156105435760408051600160e51b62461bcd02815260206004820152600b6024820152600160aa1b6a185b1c9958591e481cd95d02604482015290519081900360640190fd5b60008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055565b60016000908152602081905260008051602061123f8339815191525433906001600160a01b03168114806105e6575060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5546001600160a01b038281169116145b61062457604051600160e51b62461bcd02815260040180806020018281038252602b81526020018061125f602b913960400191505060405180910390fd5b6001600090815260205260008051602061123f833981519152546001600160a01b03918216911614905090565b81151560009081526001602090815260408083208484529182905290912080546001600160a01b031633146106d05760408051600160e51b62461bcd02815260206004820152601060248201527f6f6e6c79206f72646572206f776e657200000000000000000000000000000000604482015290519081900360640190fd5b600480820154600380840180546000908152602087815260408083208701869055925494825282822090930193909355871515835282825280832054855460018701548351600160e01b63a9059cbb0281526001600160a01b03928316978101979097526024870152915191169363a9059cbb9360448083019493928390030190829087803b15801561076257600080fd5b505af1158015610776573d6000803e3d6000fd5b505050506040513d602081101561078c57600080fd5b50505060009182526020526040812080546001600160a01b03191681556001810182905560028101829055600381018290556004015550565b811515600090815260016020908152604080832084845290915290206002015415155b92915050565b8115156000908152600160205260408120819081808061080d88610934565b90505b8060001a60f81b6001600160f81b0319161580159061082e57508683105b156109265760008181526020859052604081209089610851578160020154610857565b81600101545b905060008a61086a578260010154610870565b82600201545b905089610883878463ffffffff610dc016565b11156108f357600061089b8b8863ffffffff610dd916565b905060006108bf846108b3858563ffffffff610dee16565b9063ffffffff610e1516565b90506108d1878263ffffffff610dc016565b6108e1898463ffffffff610dc016565b9a509a5050505050505050505061092d565b610903858263ffffffff610dc016565b9450610915868363ffffffff610dc016565b955082600401549350505050610810565b5093509150505b9250929050565b1515600090815260016020908152604080832083805290915290206004015490565b600080831180156109675750600082115b6109bb5760408051600160e51b62461bcd02815260206004820152600e60248201527f7361766520796f75722074696d65000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038516600081815260026020818152604080842080546001908101918290558a151586528352818520868652848452825160609790971b8785015260348701919091526054860189905260748087018990528251808803909101815260949096019182905285519095928291908401908083835b60208310610a555780518252601f199092019160209182019101610a36565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610a94573d6000803e3d6000fd5b5050506040513d6020811015610aa957600080fd5b50516040805160a0810182526001600160a01b03998a1681526020808201988952818301978852600060608301818152608084018281528683529790925292909220905181546001600160a01b0319169a16999099178955955160018901559351600288015550925160038601555090516004909301929092555090565b836000610b3261057e565b90508460008060208614610b5d5786866040811015610b5057600080fd5b5080359060200135610b72565b86866020811015610b6d57600080fd5b503560005b90925090506000610b868585858986610be2565b9050610b928582610e37565b50505050505050505050565b9015156000908152600160209081526040808320938352929052206003015490565b9015156000908152600160209081526040808320938352929052206004015490565b6000610bee86836107c5565b610c425760408051600160e51b62461bcd02815260206004820152600d60248201527f7361766520796f75722067617300000000000000000000000000000000000000604482015290519081900360640190fd5b851515600090815260016020526040812090610c6085898989610956565b905083610c6e898383610d2f565b610cae575b610c7e898383610d2f565b610c9957600090815260208390526040902060040154610c73565b610ca489838361113d565b509150610d269050565b6000908152602083905260409020600301545b8060001a60f81b6001600160f81b03191615801590610ce65750610ce6898383610d2f565b15610d0257600090815260208390526040902060030154610cc1565b600081815260208490526040902060040154610d21908a90849061113d565b509150505b95945050505050565b8215156000908152600160208181526040808420868552918290528084208585529084206004546002820154948301549394929391928692610d889291610d7c919063ffffffff610dee16565b9063ffffffff610dee16565b90506000610db160035460045401610d7c86600201548660010154610dee90919063ffffffff16565b90911198975050505050505050565b600082820183811015610dd257600080fd5b9392505050565b600082821115610de857600080fd5b50900390565b600082610dfd575060006107e8565b82820282848281610e0a57fe5b0414610dd257600080fd5b6000808211610e2357600080fd5b6000828481610e2e57fe5b04949350505050565b8115801560009081526001602081815260408084208685528083528185208686529390925283209092610e6985610934565b90505b8060001a60f81b6001600160f81b0319161561113357600081815260208390526040902060028082015490850154610ea99163ffffffff610dee16565b60018083015490860154610ec29163ffffffff610dee16565b1015610ed357505050505050611139565b6000610ee785600101548360020154611184565b6001860154909150610f17906108b3610f06828563ffffffff610dd916565b60028901549063ffffffff610dee16565b60028601556001850154610f31908263ffffffff610dd916565b85600101819055506000610f5a83600201546108b3848660010154610dee90919063ffffffff16565b6001840154909150610f8a906108b3610f79828563ffffffff610dd916565b60028701549063ffffffff610dee16565b60028401556001830154610fa4908263ffffffff610dd916565b60018401558715156000908152602081815260408083205489548251600160e01b63a9059cbb0281526001600160a01b03918216600482015260248101879052925191169363a9059cbb93604480850194919392918390030190829087803b15801561100f57600080fd5b505af1158015611023573d6000803e3d6000fd5b505050506040513d602081101561103957600080fd5b50508915156000908152602081815260408083205486548251600160e01b63a9059cbb0281526001600160a01b03918216600482015260248101889052925191169363a9059cbb93604480850194919392918390030190829087803b1580156110a157600080fd5b505af11580156110b5573d6000803e3d6000fd5b505050506040513d60208110156110cb57600080fd5b5050600183015415806110e057506002830154155b156110ef576110ef888561119a565b6110f888610934565b935085600101546000148061110f57506002860154155b1561112b5761111e8a8a61119a565b5050505050505050611139565b505050610e6c565b50505050505b5050565b911515600090815260016020908152604080832085845290915280822060039081018054858552838520928301819055600492830196909655849055938252902090910155565b60008183106111935781610dd2565b5090919050565b8115156000818152600160208181526040808420868552808352818520600480820154600380840180548a52868a208401839055549189528589200155968652858452828620548154958201548451600160e01b63a9059cbb0281526001600160a01b03978816998101999099526024890152925191969095949092169363a9059cbb93604480850194919392918390030190829087803b15801561076257600080fdfeada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d6f6e6c7920566f6c6174696c65546f6b656e20616e6420537461626c65546f6b656e206163636570746564a165627a7a72305820855d82b0e881e54ef0cbcdf079a3eddec2ffc6cfd0b14b542a1501e763d1c5b50029`

// DeployPairEx deploys a new Ethereum contract, binding an instance of PairEx to it.
func DeployPairEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PairEx, error) {
	parsed, err := abi.JSON(strings.NewReader(PairExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairExBin), backend)
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

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExTransactor) Remove(opts *bind.TransactOpts, _orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "remove", _orderType, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExSession) Remove(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Remove(&_PairEx.TransactOpts, _orderType, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x69722e55.
//
// Solidity: function remove(bool _orderType, bytes32 _id) returns()
func (_PairEx *PairExTransactorSession) Remove(_orderType bool, _id [32]byte) (*types.Transaction, error) {
	return _PairEx.Contract.Remove(&_PairEx.TransactOpts, _orderType, _id)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_PairEx *PairExTransactor) StableTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "stableTokenRegister")
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_PairEx *PairExSession) StableTokenRegister() (*types.Transaction, error) {
	return _PairEx.Contract.StableTokenRegister(&_PairEx.TransactOpts)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_PairEx *PairExTransactorSession) StableTokenRegister() (*types.Transaction, error) {
	return _PairEx.Contract.StableTokenRegister(&_PairEx.TransactOpts)
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

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_PairEx *PairExTransactor) VolatileTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairEx.contract.Transact(opts, "volatileTokenRegister")
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_PairEx *PairExSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _PairEx.Contract.VolatileTokenRegister(&_PairEx.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_PairEx *PairExTransactorSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _PairEx.Contract.VolatileTokenRegister(&_PairEx.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058207aa0f872b3fb9f91df8848f3481038e0c504933eee23cfb378f54b01ace0a2030029`

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
