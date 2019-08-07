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
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

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
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceABI is the input ABI used to generate the binding from.
const NextyGovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"proposalType\",\"type\":\"uint8\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"popProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSealers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFirstProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"}]"

// NextyGovernanceBin is the compiled bytecode used for deploying new contracts.
const NextyGovernanceBin = `0x608060405234801561001057600080fd5b50611e75806100206000396000f3fe608060405234801561001057600080fd5b506004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900480638943b62c116100e0578063c96be4cb11610099578063c96be4cb146104ca578063cceb68f5146104f0578063cfd59172146105d6578063d66d9e19146105de578063f8b2cb4f146105e6578063fc0c546a1461060c5761016a565b80638943b62c1461044957806397f735d51461046f5780639b212d0114610495578063a4950a291461049d578063b6b55f25146104a5578063c6908739146104c25761016a565b80632c655af3116101325780632c655af3146102a157806330ccebb5146102d15780633ccfd60b1461030957806355235d471461031157806373b9aa91146103375780637c619f82146103a85761016a565b8063013cf08b1461016f57806302a18484146101e25780630a9a3f071461021c5780632079fb9a1461025e57806328ffe6c81461027b575b600080fd5b61018c6004803603602081101561018557600080fd5b5035610614565b6040518084600281111561019c57fe5b60ff16815260200183600160a060020a0316600160a060020a0316815260200182600160a060020a0316600160a060020a03168152602001935050505060405180910390f35b610208600480360360208110156101f857600080fd5b5035600160a060020a0316610655565b604080519115158252519081900360200190f35b6102426004803603602081101561023257600080fd5b5035600160a060020a03166106e3565b60408051600160a060020a039092168252519081900360200190f35b6102426004803603602081101561027457600080fd5b5035610705565b6102086004803603602081101561029157600080fd5b5035600160a060020a031661072d565b6102a9610ba0565b60408051938452600160a060020a039283166020850152911682820152519081900360600190f35b6102f7600480360360208110156102e757600080fd5b5035600160a060020a0316610dda565b60408051918252519081900360200190f35b610208610dff565b6102f76004803603602081101561032757600080fd5b5035600160a060020a0316610fdb565b61035d6004803603602081101561034d57600080fd5b5035600160a060020a0316610ff9565b6040518085600481111561036d57fe5b60ff16815260200184815260200183600160a060020a0316600160a060020a0316815260200182815260200194505050505060405180910390f35b6103b061102f565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103f45781810151838201526020016103dc565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561043357818101518382015260200161041b565b5050505090500194505050505060405180910390f35b6102426004803603602081101561045f57600080fd5b5035600160a060020a0316611166565b6102086004803603602081101561048557600080fd5b5035600160a060020a0316611181565b6102f76111b3565b6102086111b9565b610208600480360360208110156104bb57600080fd5b503561120d565b6102f7611328565b610208600480360360208110156104e057600080fd5b5035600160a060020a031661132e565b6104f86115eb565b60405180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015610540578181015183820152602001610528565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561057f578181015183820152602001610567565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156105be5781810151838201526020016105a6565b50505050905001965050505050505060405180910390f35b6102a961179b565b61020861182d565b6102f7600480360360208110156105fc57600080fd5b5035600160a060020a0316611ab8565b610242611ad6565b600680548290811061062257fe5b60009182526020909120600290910201805460019091015460ff82169250600160a060020a036101009092048216911683565b60006001600160a060020a03831660009081526002602052604090205460ff16600481111561068057fe5b141580156106b557506004600160a060020a03831660009081526002602052604090205460ff1660048111156106b257fe5b14155b80156106db5750600160a060020a03821660009081526002602052604090206003015443115b90505b919050565b600160a060020a03908116600090815260026020819052604090912001541690565b600080548290811061071357fe5b600091825260209091200154600160a060020a0316905081565b600060043360009081526002602052604090205460ff16600481111561074f57fe5b14156107a5576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156107c557fe5b141561081b576040805160e560020a62461bcd02815260206004820152600f60248201527f616c7265616479206a6f696e6564200000000000000000000000000000000000604482015290519081900360640190fd5b600354336000908152600260205260409020600101541015610887576040805160e560020a62461bcd02815260206004820152600e60248201527f6e6f7420656e6f756768206e7466000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03808316600090815260016020526040902054839116156108f9576040805160e560020a62461bcd02815260206004820152601560248201527f636f696e6261736520616c726561647920757365640000000000000000000000604482015290519081900360640190fd5b600160a060020a0381161515610959576040805160e560020a62461bcd02815260206004820152600b60248201527f7369676e6572207a65726f000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0381163014156109ba576040805160e560020a62461bcd02815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b600160a060020a038116331415610a1b576040805160e560020a62461bcd02815260206004820152601560248201527f73616d652073656e646572277320616464726573730000000000000000000000604482015290519081900360640190fd5b3360009081526002602081905260409091209081018054600160a060020a031916600160a060020a03861617905580546001919060ff191682800217905550600160a060020a03831660009081526001602052604090208054600160a060020a03191633179055610a8b83611ae5565b604080516060810182526000808252600160a060020a0386166020830152339282019290925260068054600180820180845592909452825160029182027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054939590939192849260ff191691908490811115610b0557fe5b0217905550602082810151825474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a039283160217835560409384015160019093018054600160a060020a0319169382169390931790925582513381529187169082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea93509081900390910190a150600192915050565b600080803315610be45760405160e560020a62461bcd028152600401808060200182810382526022815260200180611e286022913960400191505060405180910390fd5b60065460001015610dc957610bf7611d2d565b600680546000908110610c0657fe5b9060005260206000209060020201606060405190810160405290816000820160009054906101000a900460ff166002811115610c3e57fe5b6002811115610c4957fe5b81528154600160a060020a0361010090910481166020830152600190920154909116604090910152905060005b60065460001901811015610d44576006805460018301908110610c9557fe5b9060005260206000209060020201600682815481101515610cb257fe5b60009182526020909120825460029283029091018054909260ff90921691839160ff1916906001908490811115610ce557fe5b021790555081548154600160a060020a0361010092839004811690920274ffffffffffffffffffffffffffffffffffffffff00199091161782556001928301549183018054600160a060020a0319169290911691909117905501610c76565b50600680546000198101908110610d5757fe5b60009182526020909120600290910201805474ffffffffffffffffffffffffffffffffffffffffff191681556001018054600160a060020a03191690556006805490610da7906000198301611d4f565b508051610db390611b34565b8160200151826040015193509350935050610dd5565b50600019915060009050805b909192565b600160a060020a0381166000908152600260205260408120546106db9060ff16611b92565b600060043360009081526002602052604090205460ff166004811115610e2157fe5b1415610e77576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610e8033610655565b1515610ed6576040805160e560020a62461bcd02815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b33600081815260026020908152604080832060018101805490859055815460ff191660031790915560055482517fa9059cbb00000000000000000000000000000000000000000000000000000000815260048101969096526024860182905291519094600160a060020a039092169363a9059cbb93604480850194919392918390030190829087803b158015610f6b57600080fd5b505af1158015610f7f573d6000803e3d6000fd5b505050506040513d6020811015610f9557600080fd5b5050604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160019150505b90565b600160a060020a031660009081526002602052604090206003015490565b6002602081905260009182526040909120805460018201549282015460039092015460ff9091169291600160a060020a03169084565b6060806060600080549050604051908082528060200260200182016040528015611063578160200160208202803883390190505b5090506060600080549050604051908082528060200260200182016040528015611097578160200160208202803883390190505b50905060005b60005481101561115c5760008054829081106110b557fe5b6000918252602090912001548351600160a060020a03909116908490839081106110db57fe5b600160a060020a03909216602092830290910190910152825160019060009085908490811061110657fe5b6020908102909101810151600160a060020a0390811683529082019290925260400160002054835191169083908390811061113d57fe5b600160a060020a0390921660209283029091019091015260010161109d565b5090925090509091565b600160205260009081526040902054600160a060020a031681565b60006004600160a060020a03831660009081526002602052604090205460ff1660048111156111ac57fe5b1492915050565b60045481565b600033156111fb5760405160e560020a62461bcd028152600401808060200182810382526022815260200180611e286022913960400191505060405180910390fd5b61120760066000611d80565b50600190565b600554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529051600092600160a060020a0316916323b872dd91606480830192602092919082900301818787803b15801561127f57600080fd5b505af1158015611293573d6000803e3d6000fd5b505050506040513d60208110156112a957600080fd5b5050336000908152600260205260409020600101546112ce908363ffffffff611c0b16565b3360008181526002602090815260409182902060010193909355805191825291810184905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a1506001919050565b60035481565b600033156113705760405160e560020a62461bcd028152600401808060200182810382526022815260200180611e286022913960400191505060405180910390fd5b600160a060020a03808316600090815260016020526040902054839116156113e2576040805160e560020a62461bcd02815260206004820152601560248201527f636f696e6261736520616c726561647920757365640000000000000000000000604482015290519081900360640190fd5b600160a060020a0381161515611442576040805160e560020a62461bcd02815260206004820152600b60248201527f7369676e6572207a65726f000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0381163014156114a3576040805160e560020a62461bcd02815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b600160a060020a038116331415611504576040805160e560020a62461bcd02815260206004820152601560248201527f73616d652073656e646572277320616464726573730000000000000000000000604482015290519081900360640190fd5b604080516060810182526002808252600160a060020a03868116602080850182905260009182526001908190528582205490921694840194909452600680548083018083559190955283519483027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054919590939092849260ff191691849081111561158e57fe5b02179055506020820151815474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a039283160217825560409092015160019182018054600160a060020a0319169190931617909155949350505050565b606080606080600680549050604051908082528060200260200182016040528015611620578160200160208202803883390190505b5090506060600680549050604051908082528060200260200182016040528015611654578160200160208202803883390190505b5090506060600680549050604051908082528060200260200182016040528015611688578160200160208202803883390190505b50905060005b60065481101561178e576116c36006828154811015156116aa57fe5b600091825260209091206002909102015460ff16611b34565b84828151811015156116d157fe5b6020908102909101015260068054829081106116e957fe5b906000526020600020906002020160000160019054906101000a9004600160a060020a0316838281518110151561171c57fe5b600160a060020a03909216602092830290910190910152600680548290811061174157fe5b60009182526020909120600160029092020101548251600160a060020a039091169083908390811061176f57fe5b600160a060020a0390921660209283029091019091015260010161168e565b5091945092509050909192565b6000806000806006805490501115610dc9576117c0600660008154811015156116aa57fe5b6006805460009081106117cf57fe5b906000526020600020906002020160000160019054906101000a9004600160a060020a03166006600081548110151561180457fe5b60009182526020909120600160029092020101549194509250600160a060020a03169050610dd5565b600060043360009081526002602052604090205460ff16600481111561184f57fe5b14156118a5576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156118c557fe5b1461191a576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f74206a6f696e656420000000000000000000000000000000000000000000604482015290519081900360640190fd5b3360009081526002602081905260409091208082018054600160a060020a03198116909155815460ff19169092179055600454600160a060020a03909116906119639043611c0b565b33600090815260026020908152604080832060030193909355600160a060020a038416825260019052208054600160a060020a03191690556119a481611c6f565b604080516060810182526001808252600160a060020a0384166020830152339282019290925260068054808401808355600092909252825160029182027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054939590939192849260ff191691908490811115611a1e57fe5b0217905550602082810151825474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a039283160217835560409384015160019093018054600160a060020a0319169382169390931790925582513381529185169082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c593509081900390910190a1600191505090565b600160a060020a031660009081526002602052604090206001015490565b600554600160a060020a031681565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018054600160a060020a031916600160a060020a0392909216919091179055565b600080826002811115611b4357fe5b1415611b51575060006106de565b6001826002811115611b5f57fe5b1415611b6d575060016106de565b6002826002811115611b7b57fe5b1415611b89575060026106de565b50600019919050565b600080826004811115611ba157fe5b1415611baf575060006106de565b6001826004811115611bbd57fe5b1415611bcb575060016106de565b6002826004811115611bd957fe5b1415611be7575060026106de565b6003826004811115611bf557fe5b1415611c03575060036106de565b50607f919050565b600082820183811015611c68576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60005b600054811015611d28576000805482908110611c8a57fe5b600091825260209091200154600160a060020a0383811691161415611d2057600080546000198101908110611cbb57fe5b60009182526020822001548154600160a060020a03909116919083908110611cdf57fe5b600091825260208220018054600160a060020a031916600160a060020a039390931692909217909155805490611d19906000198301611da1565b5050611d2a565b600101611c72565b505b50565b6040805160608101909152806000815260006020820181905260409091015290565b815481835581811115611d7b57600202816002028360005260206000209182019101611d7b9190611dc5565b505050565b5080546000825560020290600052602060002090810190611d2a9190611dc5565b815481835581811115611d7b57600083815260209020611d7b918101908301611e0d565b610fd891905b80821115611e0957805474ffffffffffffffffffffffffffffffffffffffffff19168155600181018054600160a060020a0319169055600201611dcb565b5090565b610fd891905b80821115611e095760008155600101611e1356fe6f6e6c792063616e2063616c6c2066726f6d20636f6e73656e737573206c6576656ca165627a7a7230582046bee31337abeb613c577dc0bc5d2773f9bb35980c8119c5c30e2d207ac6090b0029`

// DeployNextyGovernance deploys a new Ethereum contract, binding an instance of NextyGovernance to it.
func DeployNextyGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NextyGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyGovernanceBin), backend)
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

// GetAllProposals is a free data retrieval call binding the contract method 0xcceb68f5.
//
// Solidity: function getAllProposals() constant returns(int256[], address[], address[])
func (_NextyGovernance *NextyGovernanceCaller) GetAllProposals(opts *bind.CallOpts) ([]*big.Int, []common.Address, []common.Address, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _NextyGovernance.contract.Call(opts, out, "getAllProposals")
	return *ret0, *ret1, *ret2, err
}

// GetAllProposals is a free data retrieval call binding the contract method 0xcceb68f5.
//
// Solidity: function getAllProposals() constant returns(int256[], address[], address[])
func (_NextyGovernance *NextyGovernanceSession) GetAllProposals() ([]*big.Int, []common.Address, []common.Address, error) {
	return _NextyGovernance.Contract.GetAllProposals(&_NextyGovernance.CallOpts)
}

// GetAllProposals is a free data retrieval call binding the contract method 0xcceb68f5.
//
// Solidity: function getAllProposals() constant returns(int256[], address[], address[])
func (_NextyGovernance *NextyGovernanceCallerSession) GetAllProposals() ([]*big.Int, []common.Address, []common.Address, error) {
	return _NextyGovernance.Contract.GetAllProposals(&_NextyGovernance.CallOpts)
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

// GetFirstProposal is a free data retrieval call binding the contract method 0xcfd59172.
//
// Solidity: function getFirstProposal() constant returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceCaller) GetFirstProposal(opts *bind.CallOpts) (*big.Int, common.Address, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _NextyGovernance.contract.Call(opts, out, "getFirstProposal")
	return *ret0, *ret1, *ret2, err
}

// GetFirstProposal is a free data retrieval call binding the contract method 0xcfd59172.
//
// Solidity: function getFirstProposal() constant returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceSession) GetFirstProposal() (*big.Int, common.Address, common.Address, error) {
	return _NextyGovernance.Contract.GetFirstProposal(&_NextyGovernance.CallOpts)
}

// GetFirstProposal is a free data retrieval call binding the contract method 0xcfd59172.
//
// Solidity: function getFirstProposal() constant returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceCallerSession) GetFirstProposal() (*big.Int, common.Address, common.Address, error) {
	return _NextyGovernance.Contract.GetFirstProposal(&_NextyGovernance.CallOpts)
}

// GetSealers is a free data retrieval call binding the contract method 0x7c619f82.
//
// Solidity: function getSealers() constant returns(address[], address[])
func (_NextyGovernance *NextyGovernanceCaller) GetSealers(opts *bind.CallOpts) ([]common.Address, []common.Address, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _NextyGovernance.contract.Call(opts, out, "getSealers")
	return *ret0, *ret1, err
}

// GetSealers is a free data retrieval call binding the contract method 0x7c619f82.
//
// Solidity: function getSealers() constant returns(address[], address[])
func (_NextyGovernance *NextyGovernanceSession) GetSealers() ([]common.Address, []common.Address, error) {
	return _NextyGovernance.Contract.GetSealers(&_NextyGovernance.CallOpts)
}

// GetSealers is a free data retrieval call binding the contract method 0x7c619f82.
//
// Solidity: function getSealers() constant returns(address[], address[])
func (_NextyGovernance *NextyGovernanceCallerSession) GetSealers() ([]common.Address, []common.Address, error) {
	return _NextyGovernance.Contract.GetSealers(&_NextyGovernance.CallOpts)
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

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(uint8 proposalType, address signer, address beneficiary)
func (_NextyGovernance *NextyGovernanceCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposalType uint8
	Signer       common.Address
	Beneficiary  common.Address
}, error) {
	ret := new(struct {
		ProposalType uint8
		Signer       common.Address
		Beneficiary  common.Address
	})
	out := ret
	err := _NextyGovernance.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(uint8 proposalType, address signer, address beneficiary)
func (_NextyGovernance *NextyGovernanceSession) Proposals(arg0 *big.Int) (struct {
	ProposalType uint8
	Signer       common.Address
	Beneficiary  common.Address
}, error) {
	return _NextyGovernance.Contract.Proposals(&_NextyGovernance.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(uint8 proposalType, address signer, address beneficiary)
func (_NextyGovernance *NextyGovernanceCallerSession) Proposals(arg0 *big.Int) (struct {
	ProposalType uint8
	Signer       common.Address
	Beneficiary  common.Address
}, error) {
	return _NextyGovernance.Contract.Proposals(&_NextyGovernance.CallOpts, arg0)
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

// CleanProposals is a paid mutator transaction binding the contract method 0xa4950a29.
//
// Solidity: function cleanProposals() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) CleanProposals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "cleanProposals")
}

// CleanProposals is a paid mutator transaction binding the contract method 0xa4950a29.
//
// Solidity: function cleanProposals() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) CleanProposals() (*types.Transaction, error) {
	return _NextyGovernance.Contract.CleanProposals(&_NextyGovernance.TransactOpts)
}

// CleanProposals is a paid mutator transaction binding the contract method 0xa4950a29.
//
// Solidity: function cleanProposals() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) CleanProposals() (*types.Transaction, error) {
	return _NextyGovernance.Contract.CleanProposals(&_NextyGovernance.TransactOpts)
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

// PopProposal is a paid mutator transaction binding the contract method 0x2c655af3.
//
// Solidity: function popProposal() returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceTransactor) PopProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "popProposal")
}

// PopProposal is a paid mutator transaction binding the contract method 0x2c655af3.
//
// Solidity: function popProposal() returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceSession) PopProposal() (*types.Transaction, error) {
	return _NextyGovernance.Contract.PopProposal(&_NextyGovernance.TransactOpts)
}

// PopProposal is a paid mutator transaction binding the contract method 0x2c655af3.
//
// Solidity: function popProposal() returns(int256, address, address)
func (_NextyGovernance *NextyGovernanceTransactorSession) PopProposal() (*types.Transaction, error) {
	return _NextyGovernance.Contract.PopProposal(&_NextyGovernance.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Slash(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "slash", _signer)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Slash(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Slash(&_NextyGovernance.TransactOpts, _signer)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Slash(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Slash(&_NextyGovernance.TransactOpts, _signer)
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

// ParseBanned is a log parse operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseBanned(log types.Log) (*NextyGovernanceBanned, error) {
	event := new(NextyGovernanceBanned)
	if err := _NextyGovernance.contract.UnpackLog(event, "Banned", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) ParseDeposited(log types.Log) (*NextyGovernanceDeposited, error) {
	event := new(NextyGovernanceDeposited)
	if err := _NextyGovernance.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseJoined is a log parse operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseJoined(log types.Log) (*NextyGovernanceJoined, error) {
	event := new(NextyGovernanceJoined)
	if err := _NextyGovernance.contract.UnpackLog(event, "Joined", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseLeft is a log parse operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseLeft(log types.Log) (*NextyGovernanceLeft, error) {
	event := new(NextyGovernanceLeft)
	if err := _NextyGovernance.contract.UnpackLog(event, "Left", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseUnbanned is a log parse operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseUnbanned(log types.Log) (*NextyGovernanceUnbanned, error) {
	event := new(NextyGovernanceUnbanned)
	if err := _NextyGovernance.contract.UnpackLog(event, "Unbanned", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) ParseWithdrawn(log types.Log) (*NextyGovernanceWithdrawn, error) {
	event := new(NextyGovernanceWithdrawn)
	if err := _NextyGovernance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582044074257881defe4431a17c4e75fcc39e01f692e4f1342b8a506aec4590e5e5d0029`

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
