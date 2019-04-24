// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

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

// NextyGovernanceABI is the input ABI used to generate the binding from.
const NextyGovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_stakeRequire\",\"type\":\"uint256\"},{\"name\":\"_stakeLockHeight\",\"type\":\"uint256\"},{\"name\":\"_signers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"}]"

// NextyGovernanceBin is the compiled bytecode used for deploying new contracts.
const NextyGovernanceBin = `0x608060405234801561001057600080fd5b5060405161116f38038061116f8339810180604052608081101561003357600080fd5b815160208301516040840151606085018051939592949193918301929164010000000081111561006257600080fd5b8201602081018481111561007557600080fd5b815185602082028301116401000000008211171561009257600080fd5b5050600580546001600160a01b0319166001600160a01b03891617905560038690556004859055925060009150505b815181101561024e5760008282815181106100d857fe5b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055815182908290811061012157fe5b60200260200101516001600084848151811061013957fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081818151811061019157fe5b6020026020010151600260008484815181106101a957fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060016002600084848151811061020a57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916600183600481111561024157fe5b02179055506001016100c1565b5050505050610f0d806102626000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638943b62c11610097578063c690873911610066578063c69087391461030c578063d66d9e1914610314578063f8b2cb4f1461031c578063fc0c546a1461034257610100565b80638943b62c1461029b57806397f735d5146102c15780639b212d01146102e7578063b6b55f25146102ef57610100565b806330ccebb5116100d357806330ccebb5146101c45780633ccfd60b146101fc57806355235d471461020457806373b9aa911461022a57610100565b806302a18484146101055780630a9a3f071461013f5780632079fb9a1461018157806328ffe6c81461019e575b600080fd5b61012b6004803603602081101561011b57600080fd5b50356001600160a01b031661034a565b604080519115158252519081900360200190f35b6101656004803603602081101561015557600080fd5b50356001600160a01b03166103d8565b604080516001600160a01b039092168252519081900360200190f35b6101656004803603602081101561019757600080fd5b50356103fa565b61012b600480360360208110156101b457600080fd5b50356001600160a01b0316610421565b6101ea600480360360208110156101da57600080fd5b50356001600160a01b03166107a7565b60408051918252519081900360200190f35b61012b6107cc565b6101ea6004803603602081101561021a57600080fd5b50356001600160a01b031661097d565b6102506004803603602081101561024057600080fd5b50356001600160a01b031661099b565b6040518085600481111561026057fe5b60ff168152602001848152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060405180910390f35b610165600480360360208110156102b157600080fd5b50356001600160a01b03166109d1565b61012b600480360360208110156102d757600080fd5b50356001600160a01b03166109ec565b6101ea610a1e565b61012b6004803603602081101561030557600080fd5b5035610a24565b6101ea610b29565b61012b610b2f565b6101ea6004803603602081101561033257600080fd5b50356001600160a01b0316610ccf565b610165610ced565b600060016001600160a01b03831660009081526002602052604090205460ff16600481111561037557fe5b141580156103aa575060046001600160a01b03831660009081526002602052604090205460ff1660048111156103a757fe5b14155b80156103d057506001600160a01b03821660009081526002602052604090206003015443115b90505b919050565b6001600160a01b03908116600090815260026020819052604090912001541690565b6000818154811061040757fe5b6000918252602090912001546001600160a01b0316905081565b600060043360009081526002602052604090205460ff16600481111561044357fe5b14156104865760408051600160e51b62461bcd0281526020600482015260076024820152600160cd1b6603130b73732b2102604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156104a657fe5b14156104fc5760408051600160e51b62461bcd02815260206004820152600f60248201527f616c7265616479206a6f696e6564200000000000000000000000000000000000604482015290519081900360640190fd5b6003543360009081526002602052604090206001015410156105685760408051600160e51b62461bcd02815260206004820152600e60248201527f6e6f7420656e6f756768206e7466000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03808316600090815260016020526040902054839116156105da5760408051600160e51b62461bcd02815260206004820152601560248201527f636f696e6261736520616c726561647920757365640000000000000000000000604482015290519081900360640190fd5b6001600160a01b0381166106295760408051600160e51b62461bcd02815260206004820152600b6024820152600160a81b6a7369676e6572207a65726f02604482015290519081900360640190fd5b6001600160a01b03811630141561068a5760408051600160e51b62461bcd02815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b6001600160a01b0381163314156106eb5760408051600160e51b62461bcd02815260206004820152601560248201527f73616d652073656e646572277320616464726573730000000000000000000000604482015290519081900360640190fd5b33600090815260026020819052604090912090810180546001600160a01b0319166001600160a01b03861617905580546001919060ff1916828002179055506001600160a01b038316600090815260016020526040902080546001600160a01b0319163317905561075b83610cfc565b604080513381526001600160a01b038516602082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea929181900390910190a150600192915050565b6001600160a01b0381166000908152600260205260408120546103d09060ff16610d4b565b600060043360009081526002602052604090205460ff1660048111156107ee57fe5b14156108315760408051600160e51b62461bcd0281526020600482015260076024820152600160cd1b6603130b73732b2102604482015290519081900360640190fd5b61083a3361034a565b61088e5760408051600160e51b62461bcd02815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b33600081815260026020908152604080832060018101805490859055815460ff19166003179091556005548251600160e01b63a9059cbb028152600481019690965260248601829052915190946001600160a01b039092169363a9059cbb93604480850194919392918390030190829087803b15801561090d57600080fd5b505af1158015610921573d6000803e3d6000fd5b505050506040513d602081101561093757600080fd5b5050604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160019150505b90565b6001600160a01b031660009081526002602052604090206003015490565b6002602081905260009182526040909120805460018201549282015460039092015460ff90911692916001600160a01b03169084565b6001602052600090815260409020546001600160a01b031681565b600060046001600160a01b03831660009081526002602052604090205460ff166004811115610a1757fe5b1492915050565b60045481565b60055460408051600160e01b6323b872dd0281523360048201523060248201526044810184905290516000926001600160a01b0316916323b872dd91606480830192602092919082900301818787803b158015610a8057600080fd5b505af1158015610a94573d6000803e3d6000fd5b505050506040513d6020811015610aaa57600080fd5b505033600090815260026020526040902060010154610acf908363ffffffff610dc416565b3360008181526002602090815260409182902060010193909355805191825291810184905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a1506001919050565b60035481565b600060043360009081526002602052604090205460ff166004811115610b5157fe5b1415610b945760408051600160e51b62461bcd0281526020600482015260076024820152600160cd1b6603130b73732b2102604482015290519081900360640190fd5b60013360009081526002602052604090205460ff166004811115610bb457fe5b14610bfa5760408051600160e51b62461bcd02815260206004820152600b6024820152600160ad1b6a03737ba103537b4b732b2102604482015290519081900360640190fd5b33600090815260026020819052604090912080820180546001600160a01b03198116909155815460ff191690921790556004546001600160a01b0390911690610c439043610dc4565b336000908152600260209081526040808320600301939093556001600160a01b0384168252600190522080546001600160a01b0319169055610c8481610ddd565b604080513381526001600160a01b038316602082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5929181900390910190a1600191505090565b6001600160a01b031660009081526002602052604090206001015490565b6005546001600160a01b031681565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b600080826004811115610d5a57fe5b1415610d68575060006103d3565b6001826004811115610d7657fe5b1415610d84575060016103d3565b6002826004811115610d9257fe5b1415610da0575060026103d3565b6003826004811115610dae57fe5b1415610dbc575060036103d3565b50607f919050565b600082820183811015610dd657600080fd5b9392505050565b60005b600054811015610e955760008181548110610df757fe5b6000918252602090912001546001600160a01b0383811691161415610e8d57600080546000198101908110610e2857fe5b600091825260208220015481546001600160a01b03909116919083908110610e4c57fe5b6000918252602082200180546001600160a01b0319166001600160a01b039390931692909217909155805490610e86906000198301610e9a565b5050610e97565b600101610de0565b505b50565b815481835581811115610ebe57600083815260209020610ebe918101908301610ec3565b505050565b61097a91905b80821115610edd5760008155600101610ec9565b509056fea165627a7a72305820918ee32c204ef7ed68dcc7745897ab7302419a307f61d9e00c9527d1d3cbca2c0029`

// DeployNextyGovernance deploys a new Ethereum contract, binding an instance of NextyGovernance to it.
func DeployNextyGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _stakeRequire *big.Int, _stakeLockHeight *big.Int, _signers []common.Address) (common.Address, *types.Transaction, *NextyGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyGovernanceBin), backend, _token, _stakeRequire, _stakeLockHeight, _signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NextyGovernance is an auto generated Go binding around an Ethereum contract.
type NextyGovernance struct {
	NextyGovernanceCaller     // Read-only binding to the contract
	NextyGovernanceTransactor // Write-only binding to the contract
	NextyGovernanceFilterer   // Log filterer for contract events
}

// NextyGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NextyGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NextyGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NextyGovernanceSession struct {
	Contract     *NextyGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NextyGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NextyGovernanceCallerSession struct {
	Contract *NextyGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NextyGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NextyGovernanceTransactorSession struct {
	Contract     *NextyGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NextyGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NextyGovernanceRaw struct {
	Contract *NextyGovernance // Generic contract binding to access the raw methods on
}

// NextyGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NextyGovernanceCallerRaw struct {
	Contract *NextyGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// NextyGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactorRaw struct {
	Contract *NextyGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNextyGovernance creates a new instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernance(address common.Address, backend bind.ContractBackend) (*NextyGovernance, error) {
	contract, err := bindNextyGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NewNextyGovernanceCaller creates a new read-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceCaller(address common.Address, caller bind.ContractCaller) (*NextyGovernanceCaller, error) {
	contract, err := bindNextyGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceCaller{contract: contract}, nil
}

// NewNextyGovernanceTransactor creates a new write-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*NextyGovernanceTransactor, error) {
	contract, err := bindNextyGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceTransactor{contract: contract}, nil
}

// NewNextyGovernanceFilterer creates a new log filterer instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*NextyGovernanceFilterer, error) {
	contract, err := bindNextyGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceFilterer{contract: contract}, nil
}

// bindNextyGovernance binds a generic wrapper to an already deployed contract.
func bindNextyGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.NextyGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) constant returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceCaller) Account(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	ret := new(struct {
		Status       uint8
		Balance      *big.Int
		Signer       common.Address
		UnlockHeight *big.Int
	})
	out := ret
	err := _NextyGovernance.contract.Call(opts, out, "account", arg0)
	return *ret, err
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) constant returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceSession) Account(arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.CallOpts, arg0)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) constant returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceCallerSession) Account(arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetBalance(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getBalance", _address)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetBalance(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.CallOpts, _address)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetBalance(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.CallOpts, _address)
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) GetCoinbase(opts *bind.CallOpts, _address common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getCoinbase", _address)
	return *ret0, err
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) GetCoinbase(_address common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.CallOpts, _address)
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) GetCoinbase(_address common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.CallOpts, _address)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetStatus(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getStatus", _address)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetStatus(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.CallOpts, _address)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetStatus(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.CallOpts, _address)
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetUnlockHeight(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getUnlockHeight", _address)
	return *ret0, err
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetUnlockHeight(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.CallOpts, _address)
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetUnlockHeight(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.CallOpts, _address)
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCaller) IsBanned(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "isBanned", _address)
	return *ret0, err
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsBanned(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.CallOpts, _address)
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCallerSession) IsBanned(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.CallOpts, _address)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCaller) IsWithdrawable(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "isWithdrawable", _address)
	return *ret0, err
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsWithdrawable(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.CallOpts, _address)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCallerSession) IsWithdrawable(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.CallOpts, _address)
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) SignerCoinbase(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "signerCoinbase", arg0)
	return *ret0, err
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) SignerCoinbase(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.CallOpts, arg0)
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) SignerCoinbase(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "signers", arg0)
	return *ret0, err
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.CallOpts, arg0)
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) StakeLockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "stakeLockHeight")
	return *ret0, err
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeLockHeight() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.CallOpts)
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) StakeLockHeight() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.CallOpts)
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) StakeRequire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "stakeRequire")
	return *ret0, err
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeRequire() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.CallOpts)
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) StakeRequire() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) Token() (common.Address, error) {
	return _NextyGovernance.Contract.Token(&_NextyGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) Token() (common.Address, error) {
	return _NextyGovernance.Contract.Token(&_NextyGovernance.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts, _amount)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Join(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "join", _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// NextyGovernanceBannedIterator is returned from FilterBanned and is used to iterate over the raw logs and unpacked data for Banned events raised by the NextyGovernance contract.
type NextyGovernanceBannedIterator struct {
	Event *NextyGovernanceBanned // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceBanned)
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
		it.Event = new(NextyGovernanceBanned)
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
func (it *NextyGovernanceBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceBanned represents a Banned event raised by the NextyGovernance contract.
type NextyGovernanceBanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBanned is a free log retrieval operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterBanned(opts *bind.FilterOpts) (*NextyGovernanceBannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceBannedIterator{contract: _NextyGovernance.contract, event: "Banned", logs: logs, sub: sub}, nil
}

// WatchBanned is a free log subscription operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchBanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceBanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceBanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Banned", log); err != nil {
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

// NextyGovernanceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the NextyGovernance contract.
type NextyGovernanceDepositedIterator struct {
	Event *NextyGovernanceDeposited // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceDeposited)
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
		it.Event = new(NextyGovernanceDeposited)
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
func (it *NextyGovernanceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceDeposited represents a Deposited event raised by the NextyGovernance contract.
type NextyGovernanceDeposited struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) FilterDeposited(opts *bind.FilterOpts) (*NextyGovernanceDepositedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceDepositedIterator{contract: _NextyGovernance.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NextyGovernanceDeposited) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceDeposited)
				if err := _NextyGovernance.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// NextyGovernanceJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the NextyGovernance contract.
type NextyGovernanceJoinedIterator struct {
	Event *NextyGovernanceJoined // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceJoined)
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
		it.Event = new(NextyGovernanceJoined)
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
func (it *NextyGovernanceJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceJoined represents a Joined event raised by the NextyGovernance contract.
type NextyGovernanceJoined struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterJoined(opts *bind.FilterOpts) (*NextyGovernanceJoinedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceJoinedIterator{contract: _NextyGovernance.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *NextyGovernanceJoined) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceJoined)
				if err := _NextyGovernance.contract.UnpackLog(event, "Joined", log); err != nil {
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

// NextyGovernanceLeftIterator is returned from FilterLeft and is used to iterate over the raw logs and unpacked data for Left events raised by the NextyGovernance contract.
type NextyGovernanceLeftIterator struct {
	Event *NextyGovernanceLeft // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceLeft)
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
		it.Event = new(NextyGovernanceLeft)
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
func (it *NextyGovernanceLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceLeft represents a Left event raised by the NextyGovernance contract.
type NextyGovernanceLeft struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterLeft(opts *bind.FilterOpts) (*NextyGovernanceLeftIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceLeftIterator{contract: _NextyGovernance.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchLeft(opts *bind.WatchOpts, sink chan<- *NextyGovernanceLeft) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceLeft)
				if err := _NextyGovernance.contract.UnpackLog(event, "Left", log); err != nil {
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

// NextyGovernanceUnbannedIterator is returned from FilterUnbanned and is used to iterate over the raw logs and unpacked data for Unbanned events raised by the NextyGovernance contract.
type NextyGovernanceUnbannedIterator struct {
	Event *NextyGovernanceUnbanned // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceUnbannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceUnbanned)
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
		it.Event = new(NextyGovernanceUnbanned)
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
func (it *NextyGovernanceUnbannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceUnbannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceUnbanned represents a Unbanned event raised by the NextyGovernance contract.
type NextyGovernanceUnbanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbanned is a free log retrieval operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterUnbanned(opts *bind.FilterOpts) (*NextyGovernanceUnbannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceUnbannedIterator{contract: _NextyGovernance.contract, event: "Unbanned", logs: logs, sub: sub}, nil
}

// WatchUnbanned is a free log subscription operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchUnbanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceUnbanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceUnbanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Unbanned", log); err != nil {
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

// NextyGovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the NextyGovernance contract.
type NextyGovernanceWithdrawnIterator struct {
	Event *NextyGovernanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceWithdrawn)
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
		it.Event = new(NextyGovernanceWithdrawn)
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
func (it *NextyGovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceWithdrawn represents a Withdrawn event raised by the NextyGovernance contract.
type NextyGovernanceWithdrawn struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*NextyGovernanceWithdrawnIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceWithdrawnIterator{contract: _NextyGovernance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NextyGovernanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceWithdrawn)
				if err := _NextyGovernance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
const SafeMathBin = `0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820a35d459709bba78a5afb18397e0e35a477bdd916d5a9dccc089b6e92893c882c0029`

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
