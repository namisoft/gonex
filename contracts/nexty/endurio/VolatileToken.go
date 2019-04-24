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

// ContractReceiverABI is the input ABI used to generate the binding from.
const ContractReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractReceiverBin is the compiled bytecode used for deploying new contracts.
const ContractReceiverBin = `0x`

// DeployContractReceiver deploys a new Ethereum contract, binding an instance of ContractReceiver to it.
func DeployContractReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractReceiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// ContractReceiver is an auto generated Go binding around an Ethereum contract.
type ContractReceiver struct {
	ContractReceiverCaller     // Read-only binding to the contract
	ContractReceiverTransactor // Write-only binding to the contract
	ContractReceiverFilterer   // Log filterer for contract events
}

// ContractReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractReceiverSession struct {
	Contract     *ContractReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractReceiverCallerSession struct {
	Contract *ContractReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractReceiverTransactorSession struct {
	Contract     *ContractReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractReceiverRaw struct {
	Contract *ContractReceiver // Generic contract binding to access the raw methods on
}

// ContractReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractReceiverCallerRaw struct {
	Contract *ContractReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractReceiverTransactorRaw struct {
	Contract *ContractReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractReceiver creates a new instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiver(address common.Address, backend bind.ContractBackend) (*ContractReceiver, error) {
	contract, err := bindContractReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// NewContractReceiverCaller creates a new read-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverCaller(address common.Address, caller bind.ContractCaller) (*ContractReceiverCaller, error) {
	contract, err := bindContractReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverCaller{contract: contract}, nil
}

// NewContractReceiverTransactor creates a new write-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractReceiverTransactor, error) {
	contract, err := bindContractReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverTransactor{contract: contract}, nil
}

// NewContractReceiverFilterer creates a new log filterer instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractReceiverFilterer, error) {
	contract, err := bindContractReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverFilterer{contract: contract}, nil
}

// bindContractReceiver binds a generic wrapper to an already deployed contract.
func bindContractReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.ContractReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transact(opts, method, params...)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_ContractReceiver *ContractReceiverTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_ContractReceiver *ContractReceiverSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_ContractReceiver *ContractReceiverTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, _from, _value, _data)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x608060405234801561001057600080fd5b506104e0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561020b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610211565b6100b96004803603604081101561013357600080fd5b506001600160a01b038135169060200135610268565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102a4565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356102bf565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356102fb565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610308565b6000610202338484610333565b50600192915050565b60025490565b600061021e8484846103bb565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461025e918691610259908663ffffffff61048616565b610333565b5060019392505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff61049b16565b6001600160a01b031660009081526020819052604090205490565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff61048616565b60006102023384846103bb565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661034657600080fd5b6001600160a01b03831661035957600080fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0382166103ce57600080fd5b6001600160a01b0383166000908152602081905260409020546103f7908263ffffffff61048616565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461042c908263ffffffff61049b16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008282111561049557600080fd5b50900390565b6000828201838110156104ad57600080fd5b939250505056fea165627a7a723058207f61de5b90b43210f5782719b64f3ba0ab1011e122cc72df5e60e644d8a1b5050029`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC223ABI is the input ABI used to generate the binding from.
const ERC223ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnFromOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC223Bin is the compiled bytecode used for deploying new contracts.
const ERC223Bin = `0x608060405234801561001057600080fd5b50610d3f806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638f32d59b11610097578063be45fd6211610066578063be45fd62146102af578063c4d66de81461036a578063dd62ed3e14610390578063f2fde38b146103be57610100565b80638f32d59b14610232578063a457c2d71461023a578063a9059cbb14610266578063b5f07ea11461029257610100565b806348043f0d116100d357806348043f0d146101c157806370a08231146101e0578063715018a6146102065780638da5cb5b1461020e57610100565b8063095ea7b31461010557806318160ddd1461014557806323b872dd1461015f5780633950935114610195575b600080fd5b6101316004803603604081101561011b57600080fd5b506001600160a01b0381351690602001356103e4565b604080519115158252519081900360200190f35b61014d6103fa565b60408051918252519081900360200190f35b6101316004803603606081101561017557600080fd5b506001600160a01b03813581169160208101359091169060400135610400565b610131600480360360408110156101ab57600080fd5b506001600160a01b038135169060200135610458565b6101de600480360360208110156101d757600080fd5b5035610494565b005b61014d600480360360208110156101f657600080fd5b50356001600160a01b03166104b9565b6101de6104d4565b61021661052f565b604080516001600160a01b039092168252519081900360200190f35b61013161053e565b6101316004803603604081101561025057600080fd5b506001600160a01b03813516906020013561054f565b6101316004803603604081101561027c57600080fd5b506001600160a01b03813516906020013561058b565b6101de600480360360208110156102a857600080fd5b5035610598565b610131600480360360608110156102c557600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156102f557600080fd5b82018360208201111561030757600080fd5b8035906020019184600183028401116401000000008311171561032957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105ba945050505050565b6101de6004803603602081101561038057600080fd5b50356001600160a01b03166105e7565b61014d600480360360408110156103a657600080fd5b506001600160a01b03813581169160200135166106da565b6101de600480360360208110156103d457600080fd5b50356001600160a01b0316610705565b60006103f133848461071f565b50600192915050565b60025490565b600061040d8484846107a7565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461044d918691610448908663ffffffff61087216565b61071f565b5060015b9392505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103f1918590610448908663ffffffff61088716565b61049c61053e565b6104a557600080fd5b6104b66104b061052f565b82610899565b50565b6001600160a01b031660009081526020819052604090205490565b6104dc61053e565b6104e557600080fd5b6036546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603680546001600160a01b0319169055565b6036546001600160a01b031690565b6036546001600160a01b0316331490565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103f1918590610448908663ffffffff61087216565b60006103f13384846107a7565b6105a061053e565b6105a957600080fd5b6104b66105b461052f565b82610941565b60006105c5846109e8565b156105dc576105d58484846109ee565b9050610451565b6105d5848484610b9b565b600354610100900460ff16806106005750610600610c70565b8061060e575060035460ff16155b61064c57604051600160e51b62461bcd02815260040180806020018281038252602e815260200180610ce6602e913960400191505060405180910390fd5b60038054600161010061ff00198316811760ff191691909117909255603680546001600160a01b0319166001600160a01b0385811691909117918290556040519390920460ff16929116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600380549115156101000261ff001990921691909117905550565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61070d61053e565b61071657600080fd5b6104b681610c76565b6001600160a01b03821661073257600080fd5b6001600160a01b03831661074557600080fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0382166107ba57600080fd5b6001600160a01b0383166000908152602081905260409020546107e3908263ffffffff61087216565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610818908263ffffffff61088716565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008282111561088157600080fd5b50900390565b60008282018381101561045157600080fd5b6001600160a01b0382166108ac57600080fd5b6002546108bf908263ffffffff61088716565b6002556001600160a01b0382166000908152602081905260409020546108eb908263ffffffff61088716565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b03821661095457600080fd5b600254610967908263ffffffff61087216565b6002556001600160a01b038216600090815260208190526040902054610993908263ffffffff61087216565b6001600160a01b038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b3b151590565b60006109fa3384610941565b610a048484610899565b604051600160e11b63607705c502815233600482018181526024830186905260606044840190815285516064850152855188946001600160a01b0386169463c0ee0b8a9490938a938a9360840190602085019080838360005b83811015610a75578181015183820152602001610a5d565b50505050905090810190601f168015610aa25780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610ac357600080fd5b505af1158015610ad7573d6000803e3d6000fd5b50505050846001600160a01b0316336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b55578181015183820152602001610b3d565b50505050905090810190601f168015610b825780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3506001949350505050565b6000610ba73384610941565b610bb18484610899565b836001600160a01b0316336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c2b578181015183820152602001610c13565b50505050905090810190601f168015610c585780820380516001836020036101000a031916815260200191505b50935050505060405180910390a35060019392505050565b303b1590565b6001600160a01b038116610c8957600080fd5b6036546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603680546001600160a01b0319166001600160a01b039290921691909117905556fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a165627a7a723058201663e1c7cce70adfe6e4c5eaa5d6dcd0db58fae1c767e68be28058c1d7773ff70029`

// DeployERC223 deploys a new Ethereum contract, binding an instance of ERC223 to it.
func DeployERC223(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC223, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC223Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC223{ERC223Caller: ERC223Caller{contract: contract}, ERC223Transactor: ERC223Transactor{contract: contract}, ERC223Filterer: ERC223Filterer{contract: contract}}, nil
}

// ERC223 is an auto generated Go binding around an Ethereum contract.
type ERC223 struct {
	ERC223Caller     // Read-only binding to the contract
	ERC223Transactor // Write-only binding to the contract
	ERC223Filterer   // Log filterer for contract events
}

// ERC223Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC223Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223Session struct {
	Contract     *ERC223           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223CallerSession struct {
	Contract *ERC223Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC223TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223TransactorSession struct {
	Contract     *ERC223Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223Raw struct {
	Contract *ERC223 // Generic contract binding to access the raw methods on
}

// ERC223CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223CallerRaw struct {
	Contract *ERC223Caller // Generic read-only contract binding to access the raw methods on
}

// ERC223TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223TransactorRaw struct {
	Contract *ERC223Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223 creates a new instance of ERC223, bound to a specific deployed contract.
func NewERC223(address common.Address, backend bind.ContractBackend) (*ERC223, error) {
	contract, err := bindERC223(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223{ERC223Caller: ERC223Caller{contract: contract}, ERC223Transactor: ERC223Transactor{contract: contract}, ERC223Filterer: ERC223Filterer{contract: contract}}, nil
}

// NewERC223Caller creates a new read-only instance of ERC223, bound to a specific deployed contract.
func NewERC223Caller(address common.Address, caller bind.ContractCaller) (*ERC223Caller, error) {
	contract, err := bindERC223(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223Caller{contract: contract}, nil
}

// NewERC223Transactor creates a new write-only instance of ERC223, bound to a specific deployed contract.
func NewERC223Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC223Transactor, error) {
	contract, err := bindERC223(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223Transactor{contract: contract}, nil
}

// NewERC223Filterer creates a new log filterer instance of ERC223, bound to a specific deployed contract.
func NewERC223Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC223Filterer, error) {
	contract, err := bindERC223(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC223Filterer{contract: contract}, nil
}

// bindERC223 binds a generic wrapper to an already deployed contract.
func bindERC223(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223 *ERC223Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223.Contract.ERC223Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223 *ERC223Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223.Contract.ERC223Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223 *ERC223Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223.Contract.ERC223Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223 *ERC223CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223 *ERC223TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223 *ERC223TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC223 *ERC223Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC223 *ERC223Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC223.Contract.Allowance(&_ERC223.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC223 *ERC223CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC223.Contract.Allowance(&_ERC223.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223 *ERC223Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223 *ERC223Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC223.Contract.BalanceOf(&_ERC223.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223 *ERC223CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC223.Contract.BalanceOf(&_ERC223.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC223 *ERC223Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC223 *ERC223Session) IsOwner() (bool, error) {
	return _ERC223.Contract.IsOwner(&_ERC223.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC223 *ERC223CallerSession) IsOwner() (bool, error) {
	return _ERC223.Contract.IsOwner(&_ERC223.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC223 *ERC223Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC223 *ERC223Session) Owner() (common.Address, error) {
	return _ERC223.Contract.Owner(&_ERC223.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC223 *ERC223CallerSession) Owner() (common.Address, error) {
	return _ERC223.Contract.Owner(&_ERC223.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223 *ERC223Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223 *ERC223Session) TotalSupply() (*big.Int, error) {
	return _ERC223.Contract.TotalSupply(&_ERC223.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223 *ERC223CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC223.Contract.TotalSupply(&_ERC223.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC223 *ERC223Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC223 *ERC223Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.Approve(&_ERC223.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC223 *ERC223TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.Approve(&_ERC223.TransactOpts, spender, value)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_ERC223 *ERC223Transactor) BurnFromOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "burnFromOwner", _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_ERC223 *ERC223Session) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.BurnFromOwner(&_ERC223.TransactOpts, _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_ERC223 *ERC223TransactorSession) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.BurnFromOwner(&_ERC223.TransactOpts, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC223 *ERC223Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC223 *ERC223Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.DecreaseAllowance(&_ERC223.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC223 *ERC223TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.DecreaseAllowance(&_ERC223.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC223 *ERC223Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC223 *ERC223Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.IncreaseAllowance(&_ERC223.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC223 *ERC223TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.IncreaseAllowance(&_ERC223.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_ERC223 *ERC223Transactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_ERC223 *ERC223Session) Initialize(sender common.Address) (*types.Transaction, error) {
	return _ERC223.Contract.Initialize(&_ERC223.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_ERC223 *ERC223TransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _ERC223.Contract.Initialize(&_ERC223.TransactOpts, sender)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_ERC223 *ERC223Transactor) MintToOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "mintToOwner", _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_ERC223 *ERC223Session) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.MintToOwner(&_ERC223.TransactOpts, _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_ERC223 *ERC223TransactorSession) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.MintToOwner(&_ERC223.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC223 *ERC223Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC223 *ERC223Session) RenounceOwnership() (*types.Transaction, error) {
	return _ERC223.Contract.RenounceOwnership(&_ERC223.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC223 *ERC223TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC223.Contract.RenounceOwnership(&_ERC223.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_ERC223 *ERC223Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_ERC223 *ERC223Session) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223.Contract.Transfer(&_ERC223.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_ERC223 *ERC223TransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223.Contract.Transfer(&_ERC223.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC223 *ERC223Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC223 *ERC223Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.TransferFrom(&_ERC223.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC223 *ERC223TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223.Contract.TransferFrom(&_ERC223.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC223 *ERC223Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC223.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC223 *ERC223Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC223.Contract.TransferOwnership(&_ERC223.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC223 *ERC223TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC223.Contract.TransferOwnership(&_ERC223.TransactOpts, newOwner)
}

// ERC223ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC223 contract.
type ERC223ApprovalIterator struct {
	Event *ERC223Approval // Event containing the contract specifics and raw log

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
func (it *ERC223ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223Approval)
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
		it.Event = new(ERC223Approval)
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
func (it *ERC223ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223Approval represents a Approval event raised by the ERC223 contract.
type ERC223Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC223 *ERC223Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC223ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC223.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC223ApprovalIterator{contract: _ERC223.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC223 *ERC223Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC223Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC223.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223Approval)
				if err := _ERC223.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC223OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC223 contract.
type ERC223OwnershipTransferredIterator struct {
	Event *ERC223OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC223OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223OwnershipTransferred)
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
		it.Event = new(ERC223OwnershipTransferred)
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
func (it *ERC223OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223OwnershipTransferred represents a OwnershipTransferred event raised by the ERC223 contract.
type ERC223OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC223 *ERC223Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC223OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC223.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC223OwnershipTransferredIterator{contract: _ERC223.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC223 *ERC223Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC223OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC223.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223OwnershipTransferred)
				if err := _ERC223.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ERC223TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC223 contract.
type ERC223TransferIterator struct {
	Event *ERC223Transfer // Event containing the contract specifics and raw log

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
func (it *ERC223TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC223Transfer)
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
		it.Event = new(ERC223Transfer)
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
func (it *ERC223TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC223TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC223Transfer represents a Transfer event raised by the ERC223 contract.
type ERC223Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC223 *ERC223Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC223TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC223.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC223TransferIterator{contract: _ERC223.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC223 *ERC223Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC223Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC223.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC223Transfer)
				if err := _ERC223.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IPairExABI is the input ABI used to generate the binding from.
const IPairExABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_stableTokenAmount\",\"type\":\"uint256\"},{\"name\":\"_volatileTokenAmount\",\"type\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stableTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"volatileTokenRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_volatileTokenAmount\",\"type\":\"uint256\"},{\"name\":\"_stableTokenAmount\",\"type\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPairExBin is the compiled bytecode used for deploying new contracts.
const IPairExBin = `0x`

// DeployIPairEx deploys a new Ethereum contract, binding an instance of IPairEx to it.
func DeployIPairEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IPairEx, error) {
	parsed, err := abi.JSON(strings.NewReader(IPairExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IPairExBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IPairEx{IPairExCaller: IPairExCaller{contract: contract}, IPairExTransactor: IPairExTransactor{contract: contract}, IPairExFilterer: IPairExFilterer{contract: contract}}, nil
}

// IPairEx is an auto generated Go binding around an Ethereum contract.
type IPairEx struct {
	IPairExCaller     // Read-only binding to the contract
	IPairExTransactor // Write-only binding to the contract
	IPairExFilterer   // Log filterer for contract events
}

// IPairExCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPairExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPairExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPairExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPairExSession struct {
	Contract     *IPairEx          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPairExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPairExCallerSession struct {
	Contract *IPairExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IPairExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPairExTransactorSession struct {
	Contract     *IPairExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IPairExRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPairExRaw struct {
	Contract *IPairEx // Generic contract binding to access the raw methods on
}

// IPairExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPairExCallerRaw struct {
	Contract *IPairExCaller // Generic read-only contract binding to access the raw methods on
}

// IPairExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPairExTransactorRaw struct {
	Contract *IPairExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPairEx creates a new instance of IPairEx, bound to a specific deployed contract.
func NewIPairEx(address common.Address, backend bind.ContractBackend) (*IPairEx, error) {
	contract, err := bindIPairEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPairEx{IPairExCaller: IPairExCaller{contract: contract}, IPairExTransactor: IPairExTransactor{contract: contract}, IPairExFilterer: IPairExFilterer{contract: contract}}, nil
}

// NewIPairExCaller creates a new read-only instance of IPairEx, bound to a specific deployed contract.
func NewIPairExCaller(address common.Address, caller bind.ContractCaller) (*IPairExCaller, error) {
	contract, err := bindIPairEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPairExCaller{contract: contract}, nil
}

// NewIPairExTransactor creates a new write-only instance of IPairEx, bound to a specific deployed contract.
func NewIPairExTransactor(address common.Address, transactor bind.ContractTransactor) (*IPairExTransactor, error) {
	contract, err := bindIPairEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPairExTransactor{contract: contract}, nil
}

// NewIPairExFilterer creates a new log filterer instance of IPairEx, bound to a specific deployed contract.
func NewIPairExFilterer(address common.Address, filterer bind.ContractFilterer) (*IPairExFilterer, error) {
	contract, err := bindIPairEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPairExFilterer{contract: contract}, nil
}

// bindIPairEx binds a generic wrapper to an already deployed contract.
func bindIPairEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPairExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPairEx *IPairExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPairEx.Contract.IPairExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPairEx *IPairExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairEx.Contract.IPairExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPairEx *IPairExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPairEx.Contract.IPairExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPairEx *IPairExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPairEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPairEx *IPairExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPairEx *IPairExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPairEx.Contract.contract.Transact(opts, method, params...)
}

// Buy is a paid mutator transaction binding the contract method 0x0d7b117e.
//
// Solidity: function buy(uint256 _stableTokenAmount, uint256 _volatileTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExTransactor) Buy(opts *bind.TransactOpts, _stableTokenAmount *big.Int, _volatileTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.contract.Transact(opts, "buy", _stableTokenAmount, _volatileTokenAmount, maker, index)
}

// Buy is a paid mutator transaction binding the contract method 0x0d7b117e.
//
// Solidity: function buy(uint256 _stableTokenAmount, uint256 _volatileTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExSession) Buy(_stableTokenAmount *big.Int, _volatileTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.Contract.Buy(&_IPairEx.TransactOpts, _stableTokenAmount, _volatileTokenAmount, maker, index)
}

// Buy is a paid mutator transaction binding the contract method 0x0d7b117e.
//
// Solidity: function buy(uint256 _stableTokenAmount, uint256 _volatileTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExTransactorSession) Buy(_stableTokenAmount *big.Int, _volatileTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.Contract.Buy(&_IPairEx.TransactOpts, _stableTokenAmount, _volatileTokenAmount, maker, index)
}

// Sell is a paid mutator transaction binding the contract method 0x5c8ffb07.
//
// Solidity: function sell(uint256 _volatileTokenAmount, uint256 _stableTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExTransactor) Sell(opts *bind.TransactOpts, _volatileTokenAmount *big.Int, _stableTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.contract.Transact(opts, "sell", _volatileTokenAmount, _stableTokenAmount, maker, index)
}

// Sell is a paid mutator transaction binding the contract method 0x5c8ffb07.
//
// Solidity: function sell(uint256 _volatileTokenAmount, uint256 _stableTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExSession) Sell(_volatileTokenAmount *big.Int, _stableTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.Contract.Sell(&_IPairEx.TransactOpts, _volatileTokenAmount, _stableTokenAmount, maker, index)
}

// Sell is a paid mutator transaction binding the contract method 0x5c8ffb07.
//
// Solidity: function sell(uint256 _volatileTokenAmount, uint256 _stableTokenAmount, address maker, bytes32 index) returns()
func (_IPairEx *IPairExTransactorSession) Sell(_volatileTokenAmount *big.Int, _stableTokenAmount *big.Int, maker common.Address, index [32]byte) (*types.Transaction, error) {
	return _IPairEx.Contract.Sell(&_IPairEx.TransactOpts, _volatileTokenAmount, _stableTokenAmount, maker, index)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_IPairEx *IPairExTransactor) StableTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairEx.contract.Transact(opts, "stableTokenRegister")
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_IPairEx *IPairExSession) StableTokenRegister() (*types.Transaction, error) {
	return _IPairEx.Contract.StableTokenRegister(&_IPairEx.TransactOpts)
}

// StableTokenRegister is a paid mutator transaction binding the contract method 0x1264a990.
//
// Solidity: function stableTokenRegister() returns()
func (_IPairEx *IPairExTransactorSession) StableTokenRegister() (*types.Transaction, error) {
	return _IPairEx.Contract.StableTokenRegister(&_IPairEx.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_IPairEx *IPairExTransactor) VolatileTokenRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairEx.contract.Transact(opts, "volatileTokenRegister")
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_IPairEx *IPairExSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _IPairEx.Contract.VolatileTokenRegister(&_IPairEx.TransactOpts)
}

// VolatileTokenRegister is a paid mutator transaction binding the contract method 0x4f657bb1.
//
// Solidity: function volatileTokenRegister() returns()
func (_IPairEx *IPairExTransactorSession) VolatileTokenRegister() (*types.Transaction, error) {
	return _IPairEx.Contract.VolatileTokenRegister(&_IPairEx.TransactOpts)
}

// InitializableABI is the input ABI used to generate the binding from.
const InitializableABI = "[]"

// InitializableBin is the compiled bytecode used for deploying new contracts.
const InitializableBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a723058206e5ed1e4c489c047573b0461037df6848019a0c7ee60c8284eb802ebac9585140029`

// DeployInitializable deploys a new Ethereum contract, binding an instance of Initializable to it.
func DeployInitializable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Initializable, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InitializableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b50610349806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063715018a61461005c5780638da5cb5b146100665780638f32d59b1461008a578063c4d66de8146100a6578063f2fde38b146100cc575b600080fd5b6100646100f2565b005b61006e61014d565b604080516001600160a01b039092168252519081900360200190f35b61009261015c565b604080519115158252519081900360200190f35b610064600480360360208110156100bc57600080fd5b50356001600160a01b031661016d565b610064600480360360208110156100e257600080fd5b50356001600160a01b031661025d565b6100fa61015c565b61010357600080fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b600054610100900460ff1680610186575061018661027a565b80610194575060005460ff16155b6101d257604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806102f0602e913960400191505060405180910390fd5b60008054600161010061ff00198316811760ff1916919091178355603380546001600160a01b0319166001600160a01b0386811691909117918290556040519290930460ff1693921691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600080549115156101000261ff001990921691909117905550565b61026561015c565b61026e57600080fd5b61027781610280565b50565b303b1590565b6001600160a01b03811661029357600080fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b039290921691909117905556fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a165627a7a723058207b5d40f765083775dc20e397aba8a5949777370cdbdf0a84885c5d522af0a3280029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Ownable *OwnableTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Ownable *OwnableSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.Initialize(&_Ownable.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Ownable *OwnableTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.Initialize(&_Ownable.TransactOpts, sender)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// VolatileTokenABI is the input ABI used to generate the binding from.
const VolatileTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnFromOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_wantAmount\",\"type\":\"uint256\"},{\"name\":\"_assistingID\",\"type\":\"bytes32\"}],\"name\":\"simpleBuy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"sellTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_orderbook\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// VolatileTokenBin is the compiled bytecode used for deploying new contracts.
const VolatileTokenBin = `0x608060405234801561001057600080fd5b50604051602080620013eb8339810180604052602081101561003157600080fd5b5051606980546001600160a01b0319166001600160a01b038084169190911791829055604080517f4f657bb100000000000000000000000000000000000000000000000000000000815290519290911691634f657bb19160048082019260009290919082900301818387803b1580156100a957600080fd5b505af11580156100bd573d6000803e3d6000fd5b505050506100d0816100d660201b60201c565b506101ed565b600354610100900460ff16806100f557506100f56101e760201b60201c565b80610103575060035460ff16155b610159576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180620013bd602e913960400191505060405180910390fd5b60038054600161010061ff00198316811760ff191691909117909255603680546001600160a01b0319166001600160a01b0385811691909117918290556040519390920460ff16929116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600380549115156101000261ff001990921691909117905550565b303b1590565b6111c080620001fd6000396000f3fe6080604052600436106101665760003560e01c80638f32d59b116100d1578063be45fd621161008a578063dd62ed3e11610064578063dd62ed3e1461067a578063e4849b32146106b5578063f2fde38b146106df578063f542c8671461071257610166565b8063be45fd6214610556578063c4582fa11461061e578063c4d66de81461064757610166565b80638f32d59b1461048857806395d89b411461049d578063a457c2d7146104b2578063a6f2ae3a146104eb578063a9059cbb146104f3578063b5f07ea11461052c57610166565b806348043f0d1161012357806348043f0d146103105780636f0b51801461033c57806370a0823114610362578063715018a61461039557806382189551146103aa5780638da5cb5b1461045757610166565b806306fdde031461016b578063095ea7b3146101f557806318160ddd1461024257806323b872dd14610269578063313ce567146102ac57806339509351146102d7575b600080fd5b34801561017757600080fd5b5061018061074b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ba5781810151838201526020016101a2565b50505050905090810190601f1680156101e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020157600080fd5b5061022e6004803603604081101561021857600080fd5b506001600160a01b03813516906020013561076e565b604080519115158252519081900360200190f35b34801561024e57600080fd5b50610257610784565b60408051918252519081900360200190f35b34801561027557600080fd5b5061022e6004803603606081101561028c57600080fd5b506001600160a01b0381358116916020810135909116906040013561078a565b3480156102b857600080fd5b506102c16107e2565b6040805160ff9092168252519081900360200190f35b3480156102e357600080fd5b5061022e600480360360408110156102fa57600080fd5b506001600160a01b0381351690602001356107e7565b34801561031c57600080fd5b5061033a6004803603602081101561033357600080fd5b5035610823565b005b61022e6004803603602081101561035257600080fd5b50356001600160a01b0316610848565b34801561036e57600080fd5b506102576004803603602081101561038557600080fd5b50356001600160a01b0316610855565b3480156103a157600080fd5b5061033a610870565b61033a600480360360408110156103c057600080fd5b813591908101906040810160208201356401000000008111156103e257600080fd5b8201836020820111156103f457600080fd5b8035906020019184600183028401116401000000008311171561041657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108cb945050505050565b34801561046357600080fd5b5061046c6108eb565b604080516001600160a01b039092168252519081900360200190f35b34801561049457600080fd5b5061022e6108fa565b3480156104a957600080fd5b5061018061090b565b3480156104be57600080fd5b5061022e600480360360408110156104d557600080fd5b506001600160a01b038135169060200135610932565b61022e61096e565b3480156104ff57600080fd5b5061022e6004803603604081101561051657600080fd5b506001600160a01b03813516906020013561097d565b34801561053857600080fd5b5061033a6004803603602081101561054f57600080fd5b503561098a565b34801561056257600080fd5b5061022e6004803603606081101561057957600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156105a957600080fd5b8201836020820111156105bb57600080fd5b803590602001918460018302840111640100000000831117156105dd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109ac945050505050565b61033a6004803603606081101561063457600080fd5b50803590602081013590604001356109d9565b34801561065357600080fd5b5061033a6004803603602081101561066a57600080fd5b50356001600160a01b0316610a0b565b34801561068657600080fd5b506102576004803603604081101561069d57600080fd5b506001600160a01b0381358116916020013516610afe565b3480156106c157600080fd5b5061022e600480360360208110156106d857600080fd5b5035610b29565b3480156106eb57600080fd5b5061033a6004803603602081101561070257600080fd5b50356001600160a01b0316610b3b565b34801561071e57600080fd5b5061022e6004803603604081101561073557600080fd5b50803590602001356001600160a01b0316610b55565b604051806040016040528060048152602001600160e01b634d4e54590281525081565b600061077b338484610ba0565b50600192915050565b60025490565b6000610797848484610c28565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546107d79186916107d2908663ffffffff610cf316565b610ba0565b5060015b9392505050565b601881565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161077b9185906107d2908663ffffffff610d0816565b61082b6108fa565b61083457600080fd5b61084561083f6108eb565b82610d1a565b50565b60003461077b8382610d1a565b6001600160a01b031660009081526020819052604090205490565b6108786108fa565b61088157600080fd5b6036546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603680546001600160a01b0319169055565b6108d361096e565b506108e66108df6108eb565b83836109ac565b505050565b6036546001600160a01b031690565b6036546001600160a01b0316331490565b604051806040016040528060088152602001600160c01b674d656761204e54590281525081565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161077b9185906107d2908663ffffffff610cf316565b600061097933610848565b5090565b600061077b338484610c28565b6109926108fa565b61099b57600080fd5b6108456109a66108eb565b82610dc2565b60006109b784610e69565b156109ce576109c7848484610e6f565b90506107db565b6109c784848461101c565b60408051602081018490528082018390528151808203830181526060909101909152610a0584826108cb565b50505050565b600354610100900460ff1680610a245750610a246110f1565b80610a32575060035460ff16155b610a7057604051600160e51b62461bcd02815260040180806020018281038252602e815260200180611167602e913960400191505060405180910390fd5b60038054600161010061ff00198316811760ff191691909117909255603680546001600160a01b0319166001600160a01b0385811691909117918290556040519390920460ff16929116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600380549115156101000261ff001990921691909117905550565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000610b358233610b55565b50919050565b610b436108fa565b610b4c57600080fd5b610845816110f7565b600033610b628185610dc2565b6040516001600160a01b0384169085156108fc029086906000818181858888f19350505050158015610b98573d6000803e3d6000fd5b505092915050565b6001600160a01b038216610bb357600080fd5b6001600160a01b038316610bc657600080fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038216610c3b57600080fd5b6001600160a01b038316600090815260208190526040902054610c64908263ffffffff610cf316565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610c99908263ffffffff610d0816565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610d0257600080fd5b50900390565b6000828201838110156107db57600080fd5b6001600160a01b038216610d2d57600080fd5b600254610d40908263ffffffff610d0816565b6002556001600160a01b038216600090815260208190526040902054610d6c908263ffffffff610d0816565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b038216610dd557600080fd5b600254610de8908263ffffffff610cf316565b6002556001600160a01b038216600090815260208190526040902054610e14908263ffffffff610cf316565b6001600160a01b038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b3b151590565b6000610e7b3384610dc2565b610e858484610d1a565b604051600160e11b63607705c502815233600482018181526024830186905260606044840190815285516064850152855188946001600160a01b0386169463c0ee0b8a9490938a938a9360840190602085019080838360005b83811015610ef6578181015183820152602001610ede565b50505050905090810190601f168015610f235780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610f4457600080fd5b505af1158015610f58573d6000803e3d6000fd5b50505050846001600160a01b0316336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fd6578181015183820152602001610fbe565b50505050905090810190601f1680156110035780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3506001949350505050565b60006110283384610dc2565b6110328484610d1a565b836001600160a01b0316336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156110ac578181015183820152602001611094565b50505050905090810190601f1680156110d95780820380516001836020036101000a031916815260200191505b50935050505060405180910390a35060019392505050565b303b1590565b6001600160a01b03811661110a57600080fd5b6036546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603680546001600160a01b0319166001600160a01b039290921691909117905556fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a165627a7a72305820b3a7262b332d2b9537c42f2692ecb91f1a36089b61d57cc6556efb4825f9273d0029436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564`

// DeployVolatileToken deploys a new Ethereum contract, binding an instance of VolatileToken to it.
func DeployVolatileToken(auth *bind.TransactOpts, backend bind.ContractBackend, _orderbook common.Address) (common.Address, *types.Transaction, *VolatileToken, error) {
	parsed, err := abi.JSON(strings.NewReader(VolatileTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VolatileTokenBin), backend, _orderbook)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VolatileToken{VolatileTokenCaller: VolatileTokenCaller{contract: contract}, VolatileTokenTransactor: VolatileTokenTransactor{contract: contract}, VolatileTokenFilterer: VolatileTokenFilterer{contract: contract}}, nil
}

// VolatileToken is an auto generated Go binding around an Ethereum contract.
type VolatileToken struct {
	VolatileTokenCaller     // Read-only binding to the contract
	VolatileTokenTransactor // Write-only binding to the contract
	VolatileTokenFilterer   // Log filterer for contract events
}

// VolatileTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolatileTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolatileTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolatileTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolatileTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolatileTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolatileTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolatileTokenSession struct {
	Contract     *VolatileToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolatileTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolatileTokenCallerSession struct {
	Contract *VolatileTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VolatileTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolatileTokenTransactorSession struct {
	Contract     *VolatileTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VolatileTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolatileTokenRaw struct {
	Contract *VolatileToken // Generic contract binding to access the raw methods on
}

// VolatileTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolatileTokenCallerRaw struct {
	Contract *VolatileTokenCaller // Generic read-only contract binding to access the raw methods on
}

// VolatileTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolatileTokenTransactorRaw struct {
	Contract *VolatileTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolatileToken creates a new instance of VolatileToken, bound to a specific deployed contract.
func NewVolatileToken(address common.Address, backend bind.ContractBackend) (*VolatileToken, error) {
	contract, err := bindVolatileToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolatileToken{VolatileTokenCaller: VolatileTokenCaller{contract: contract}, VolatileTokenTransactor: VolatileTokenTransactor{contract: contract}, VolatileTokenFilterer: VolatileTokenFilterer{contract: contract}}, nil
}

// NewVolatileTokenCaller creates a new read-only instance of VolatileToken, bound to a specific deployed contract.
func NewVolatileTokenCaller(address common.Address, caller bind.ContractCaller) (*VolatileTokenCaller, error) {
	contract, err := bindVolatileToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenCaller{contract: contract}, nil
}

// NewVolatileTokenTransactor creates a new write-only instance of VolatileToken, bound to a specific deployed contract.
func NewVolatileTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*VolatileTokenTransactor, error) {
	contract, err := bindVolatileToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenTransactor{contract: contract}, nil
}

// NewVolatileTokenFilterer creates a new log filterer instance of VolatileToken, bound to a specific deployed contract.
func NewVolatileTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*VolatileTokenFilterer, error) {
	contract, err := bindVolatileToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenFilterer{contract: contract}, nil
}

// bindVolatileToken binds a generic wrapper to an already deployed contract.
func bindVolatileToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VolatileTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolatileToken *VolatileTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VolatileToken.Contract.VolatileTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolatileToken *VolatileTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolatileToken.Contract.VolatileTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolatileToken *VolatileTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolatileToken.Contract.VolatileTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolatileToken *VolatileTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VolatileToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolatileToken *VolatileTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolatileToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolatileToken *VolatileTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolatileToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VolatileToken *VolatileTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VolatileToken *VolatileTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolatileToken.Contract.Allowance(&_VolatileToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VolatileToken *VolatileTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolatileToken.Contract.Allowance(&_VolatileToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VolatileToken *VolatileTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VolatileToken *VolatileTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolatileToken.Contract.BalanceOf(&_VolatileToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VolatileToken *VolatileTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolatileToken.Contract.BalanceOf(&_VolatileToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VolatileToken *VolatileTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VolatileToken *VolatileTokenSession) Decimals() (uint8, error) {
	return _VolatileToken.Contract.Decimals(&_VolatileToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VolatileToken *VolatileTokenCallerSession) Decimals() (uint8, error) {
	return _VolatileToken.Contract.Decimals(&_VolatileToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VolatileToken *VolatileTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VolatileToken *VolatileTokenSession) IsOwner() (bool, error) {
	return _VolatileToken.Contract.IsOwner(&_VolatileToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VolatileToken *VolatileTokenCallerSession) IsOwner() (bool, error) {
	return _VolatileToken.Contract.IsOwner(&_VolatileToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VolatileToken *VolatileTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VolatileToken *VolatileTokenSession) Name() (string, error) {
	return _VolatileToken.Contract.Name(&_VolatileToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VolatileToken *VolatileTokenCallerSession) Name() (string, error) {
	return _VolatileToken.Contract.Name(&_VolatileToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VolatileToken *VolatileTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VolatileToken *VolatileTokenSession) Owner() (common.Address, error) {
	return _VolatileToken.Contract.Owner(&_VolatileToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VolatileToken *VolatileTokenCallerSession) Owner() (common.Address, error) {
	return _VolatileToken.Contract.Owner(&_VolatileToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VolatileToken *VolatileTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VolatileToken *VolatileTokenSession) Symbol() (string, error) {
	return _VolatileToken.Contract.Symbol(&_VolatileToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VolatileToken *VolatileTokenCallerSession) Symbol() (string, error) {
	return _VolatileToken.Contract.Symbol(&_VolatileToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VolatileToken *VolatileTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VolatileToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VolatileToken *VolatileTokenSession) TotalSupply() (*big.Int, error) {
	return _VolatileToken.Contract.TotalSupply(&_VolatileToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VolatileToken *VolatileTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _VolatileToken.Contract.TotalSupply(&_VolatileToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.Approve(&_VolatileToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.Approve(&_VolatileToken.TransactOpts, spender, value)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenTransactor) BurnFromOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "burnFromOwner", _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenSession) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.BurnFromOwner(&_VolatileToken.TransactOpts, _amount)
}

// BurnFromOwner is a paid mutator transaction binding the contract method 0xb5f07ea1.
//
// Solidity: function burnFromOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenTransactorSession) BurnFromOwner(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.BurnFromOwner(&_VolatileToken.TransactOpts, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns(bool)
func (_VolatileToken *VolatileTokenTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns(bool)
func (_VolatileToken *VolatileTokenSession) Buy() (*types.Transaction, error) {
	return _VolatileToken.Contract.Buy(&_VolatileToken.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) Buy() (*types.Transaction, error) {
	return _VolatileToken.Contract.Buy(&_VolatileToken.TransactOpts)
}

// BuyFor is a paid mutator transaction binding the contract method 0x6f0b5180.
//
// Solidity: function buyFor(address _to) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) BuyFor(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "buyFor", _to)
}

// BuyFor is a paid mutator transaction binding the contract method 0x6f0b5180.
//
// Solidity: function buyFor(address _to) returns(bool)
func (_VolatileToken *VolatileTokenSession) BuyFor(_to common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.BuyFor(&_VolatileToken.TransactOpts, _to)
}

// BuyFor is a paid mutator transaction binding the contract method 0x6f0b5180.
//
// Solidity: function buyFor(address _to) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) BuyFor(_to common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.BuyFor(&_VolatileToken.TransactOpts, _to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolatileToken *VolatileTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.DecreaseAllowance(&_VolatileToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.DecreaseAllowance(&_VolatileToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolatileToken *VolatileTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.IncreaseAllowance(&_VolatileToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.IncreaseAllowance(&_VolatileToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_VolatileToken *VolatileTokenTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_VolatileToken *VolatileTokenSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.Initialize(&_VolatileToken.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_VolatileToken *VolatileTokenTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.Initialize(&_VolatileToken.TransactOpts, sender)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenTransactor) MintToOwner(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "mintToOwner", _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenSession) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.MintToOwner(&_VolatileToken.TransactOpts, _amount)
}

// MintToOwner is a paid mutator transaction binding the contract method 0x48043f0d.
//
// Solidity: function mintToOwner(uint256 _amount) returns()
func (_VolatileToken *VolatileTokenTransactorSession) MintToOwner(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.MintToOwner(&_VolatileToken.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolatileToken *VolatileTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolatileToken *VolatileTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolatileToken.Contract.RenounceOwnership(&_VolatileToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolatileToken *VolatileTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolatileToken.Contract.RenounceOwnership(&_VolatileToken.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _amount) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) Sell(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "sell", _amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _amount) returns(bool)
func (_VolatileToken *VolatileTokenSession) Sell(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.Sell(&_VolatileToken.TransactOpts, _amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _amount) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) Sell(_amount *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.Sell(&_VolatileToken.TransactOpts, _amount)
}

// SellTo is a paid mutator transaction binding the contract method 0xf542c867.
//
// Solidity: function sellTo(uint256 _amount, address _to) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) SellTo(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "sellTo", _amount, _to)
}

// SellTo is a paid mutator transaction binding the contract method 0xf542c867.
//
// Solidity: function sellTo(uint256 _amount, address _to) returns(bool)
func (_VolatileToken *VolatileTokenSession) SellTo(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.SellTo(&_VolatileToken.TransactOpts, _amount, _to)
}

// SellTo is a paid mutator transaction binding the contract method 0xf542c867.
//
// Solidity: function sellTo(uint256 _amount, address _to) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) SellTo(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.SellTo(&_VolatileToken.TransactOpts, _amount, _to)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0xc4582fa1.
//
// Solidity: function simpleBuy(uint256 _value, uint256 _wantAmount, bytes32 _assistingID) returns()
func (_VolatileToken *VolatileTokenTransactor) SimpleBuy(opts *bind.TransactOpts, _value *big.Int, _wantAmount *big.Int, _assistingID [32]byte) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "simpleBuy", _value, _wantAmount, _assistingID)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0xc4582fa1.
//
// Solidity: function simpleBuy(uint256 _value, uint256 _wantAmount, bytes32 _assistingID) returns()
func (_VolatileToken *VolatileTokenSession) SimpleBuy(_value *big.Int, _wantAmount *big.Int, _assistingID [32]byte) (*types.Transaction, error) {
	return _VolatileToken.Contract.SimpleBuy(&_VolatileToken.TransactOpts, _value, _wantAmount, _assistingID)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0xc4582fa1.
//
// Solidity: function simpleBuy(uint256 _value, uint256 _wantAmount, bytes32 _assistingID) returns()
func (_VolatileToken *VolatileTokenTransactorSession) SimpleBuy(_value *big.Int, _wantAmount *big.Int, _assistingID [32]byte) (*types.Transaction, error) {
	return _VolatileToken.Contract.SimpleBuy(&_VolatileToken.TransactOpts, _value, _wantAmount, _assistingID)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_VolatileToken *VolatileTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_VolatileToken *VolatileTokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VolatileToken.Contract.Transfer(&_VolatileToken.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address _to, uint256 _value, bytes _data) returns(bool success)
func (_VolatileToken *VolatileTokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VolatileToken.Contract.Transfer(&_VolatileToken.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.TransferFrom(&_VolatileToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VolatileToken *VolatileTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VolatileToken.Contract.TransferFrom(&_VolatileToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolatileToken *VolatileTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolatileToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolatileToken *VolatileTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.TransferOwnership(&_VolatileToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolatileToken *VolatileTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolatileToken.Contract.TransferOwnership(&_VolatileToken.TransactOpts, newOwner)
}

// VolatileTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VolatileToken contract.
type VolatileTokenApprovalIterator struct {
	Event *VolatileTokenApproval // Event containing the contract specifics and raw log

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
func (it *VolatileTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolatileTokenApproval)
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
		it.Event = new(VolatileTokenApproval)
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
func (it *VolatileTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolatileTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolatileTokenApproval represents a Approval event raised by the VolatileToken contract.
type VolatileTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolatileToken *VolatileTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VolatileTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolatileToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenApprovalIterator{contract: _VolatileToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolatileToken *VolatileTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VolatileTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolatileToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolatileTokenApproval)
				if err := _VolatileToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// VolatileTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolatileToken contract.
type VolatileTokenOwnershipTransferredIterator struct {
	Event *VolatileTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolatileTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolatileTokenOwnershipTransferred)
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
		it.Event = new(VolatileTokenOwnershipTransferred)
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
func (it *VolatileTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolatileTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolatileTokenOwnershipTransferred represents a OwnershipTransferred event raised by the VolatileToken contract.
type VolatileTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolatileToken *VolatileTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolatileTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolatileToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenOwnershipTransferredIterator{contract: _VolatileToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolatileToken *VolatileTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolatileTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolatileToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolatileTokenOwnershipTransferred)
				if err := _VolatileToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// VolatileTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VolatileToken contract.
type VolatileTokenTransferIterator struct {
	Event *VolatileTokenTransfer // Event containing the contract specifics and raw log

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
func (it *VolatileTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolatileTokenTransfer)
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
		it.Event = new(VolatileTokenTransfer)
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
func (it *VolatileTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolatileTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolatileTokenTransfer represents a Transfer event raised by the VolatileToken contract.
type VolatileTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolatileToken *VolatileTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VolatileTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolatileToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VolatileTokenTransferIterator{contract: _VolatileToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolatileToken *VolatileTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VolatileTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolatileToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolatileTokenTransfer)
				if err := _VolatileToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
