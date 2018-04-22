// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// GatekeeperABI is the input ABI used to generate the binding from.
const GatekeeperABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"txNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PayInTrx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"txNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PayOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"Suicide\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PayIn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_txNumber\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GatekeeperBin is the compiled bytecode used for deploying new contracts.
const GatekeeperBin = `0x60606040526000600255341561001457600080fd5b604051602080610613833981016040528080516000805460018054600160a060020a03948516600160a060020a03199182161790915533939093169083168117909216909117905550506105a68061006d6000396000f3006060604052600436106100695763ffffffff60e060020a60003504166341c0e1b5811461006e578063634235fc146100835780638da5cb5b146100a85780639fab56ac146100d7578063add89bb2146100ed578063d942bffa14610117578063f2fde38b1461013c575b600080fd5b341561007957600080fd5b61008161015b565b005b341561008e57600080fd5b610081600160a060020a03600435166024356044356102a0565b34156100b357600080fd5b6100bb6103ea565b604051600160a060020a03909116815260200160405180910390f35b34156100e257600080fd5b6100816004356103f9565b34156100f857600080fd5b6101036004356104c4565b604051901515815260200160405180910390f35b341561012257600080fd5b61012a6104d9565b60405190815260200160405180910390f35b341561014757600080fd5b610081600160a060020a03600435166104df565b6000805433600160a060020a0390811691161461017757600080fd5b600154600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156101c757600080fd5b5af115156101d457600080fd5b5050506040518051600154600054919350600160a060020a03908116925063a9059cbb91168360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561023d57600080fd5b5af1151561024a57600080fd5b50505060405180519050151561025f57600080fd5b7fa1ea9b09ea114021983e9ecf71cf2ffddfd80f5cb4f925e5bf24f9bdb5e55fde4260405190815260200160405180910390a1600054600160a060020a0316ff5b6000805433600160a060020a039081169116146102bc57600080fd5b838284604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902060008181526003602052604090205490915060ff161561031757600080fd5b600154600160a060020a031663a9059cbb858560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561036d57600080fd5b5af1151561037a57600080fd5b50505060405180519050151561038f57600080fd5b60008181526003602052604090819020805460ff1916600117905583908390600160a060020a038716907fafe5bde68f2dccddb6979d04a1dceb6099fe0ceda6bfc2c37e3a40ed8195d90f905160405180910390a450505050565b600054600160a060020a031681565b600154600160a060020a03166323b872dd33308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561045c57600080fd5b5af1151561046957600080fd5b50505060405180519050151561047e57600080fd5b60028054600101908190558190600160a060020a0330167f5bdf8f42703e019ab5248b71701aac7fc28b9c4d232dc684e53cbf1eb503ae3760405160405180910390a450565b60036020526000908152604090205460ff1681565b60025481565b60005433600160a060020a039081169116146104fa57600080fd5b600160a060020a038116151561050f57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820ea90ecaa6f443abb5f064c00a7e77533447ab7a590c62b95358afeaaacbd0f920029`

// DeployGatekeeper deploys a new Ethereum contract, binding an instance of Gatekeeper to it.
func DeployGatekeeper(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Gatekeeper, error) {
	parsed, err := abi.JSON(strings.NewReader(GatekeeperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GatekeeperBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gatekeeper{GatekeeperCaller: GatekeeperCaller{contract: contract}, GatekeeperTransactor: GatekeeperTransactor{contract: contract}, GatekeeperFilterer: GatekeeperFilterer{contract: contract}}, nil
}

// Gatekeeper is an auto generated Go binding around an Ethereum contract.
type Gatekeeper struct {
	GatekeeperCaller     // Read-only binding to the contract
	GatekeeperTransactor // Write-only binding to the contract
	GatekeeperFilterer   // Log filterer for contract events
}

// GatekeeperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatekeeperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatekeeperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatekeeperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatekeeperSession struct {
	Contract     *Gatekeeper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatekeeperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatekeeperCallerSession struct {
	Contract *GatekeeperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GatekeeperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatekeeperTransactorSession struct {
	Contract     *GatekeeperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GatekeeperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatekeeperRaw struct {
	Contract *Gatekeeper // Generic contract binding to access the raw methods on
}

// GatekeeperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatekeeperCallerRaw struct {
	Contract *GatekeeperCaller // Generic read-only contract binding to access the raw methods on
}

// GatekeeperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatekeeperTransactorRaw struct {
	Contract *GatekeeperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatekeeper creates a new instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeper(address common.Address, backend bind.ContractBackend) (*Gatekeeper, error) {
	contract, err := bindGatekeeper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gatekeeper{GatekeeperCaller: GatekeeperCaller{contract: contract}, GatekeeperTransactor: GatekeeperTransactor{contract: contract}, GatekeeperFilterer: GatekeeperFilterer{contract: contract}}, nil
}

// NewGatekeeperCaller creates a new read-only instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperCaller(address common.Address, caller bind.ContractCaller) (*GatekeeperCaller, error) {
	contract, err := bindGatekeeper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatekeeperCaller{contract: contract}, nil
}

// NewGatekeeperTransactor creates a new write-only instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperTransactor(address common.Address, transactor bind.ContractTransactor) (*GatekeeperTransactor, error) {
	contract, err := bindGatekeeper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatekeeperTransactor{contract: contract}, nil
}

// NewGatekeeperFilterer creates a new log filterer instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperFilterer(address common.Address, filterer bind.ContractFilterer) (*GatekeeperFilterer, error) {
	contract, err := bindGatekeeper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatekeeperFilterer{contract: contract}, nil
}

// bindGatekeeper binds a generic wrapper to an already deployed contract.
func bindGatekeeper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatekeeperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gatekeeper *GatekeeperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gatekeeper.Contract.GatekeeperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gatekeeper *GatekeeperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gatekeeper.Contract.GatekeeperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gatekeeper *GatekeeperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gatekeeper.Contract.GatekeeperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gatekeeper *GatekeeperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gatekeeper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gatekeeper *GatekeeperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gatekeeper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gatekeeper *GatekeeperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gatekeeper.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Gatekeeper *GatekeeperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gatekeeper.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Gatekeeper *GatekeeperSession) Owner() (common.Address, error) {
	return _Gatekeeper.Contract.Owner(&_Gatekeeper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Gatekeeper *GatekeeperCallerSession) Owner() (common.Address, error) {
	return _Gatekeeper.Contract.Owner(&_Gatekeeper.CallOpts)
}

// Paid is a free data retrieval call binding the contract method 0xadd89bb2.
//
// Solidity: function paid( bytes32) constant returns(bool)
func (_Gatekeeper *GatekeeperCaller) Paid(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gatekeeper.contract.Call(opts, out, "paid", arg0)
	return *ret0, err
}

// Paid is a free data retrieval call binding the contract method 0xadd89bb2.
//
// Solidity: function paid( bytes32) constant returns(bool)
func (_Gatekeeper *GatekeeperSession) Paid(arg0 [32]byte) (bool, error) {
	return _Gatekeeper.Contract.Paid(&_Gatekeeper.CallOpts, arg0)
}

// Paid is a free data retrieval call binding the contract method 0xadd89bb2.
//
// Solidity: function paid( bytes32) constant returns(bool)
func (_Gatekeeper *GatekeeperCallerSession) Paid(arg0 [32]byte) (bool, error) {
	return _Gatekeeper.Contract.Paid(&_Gatekeeper.CallOpts, arg0)
}

// TransactionAmount is a free data retrieval call binding the contract method 0xd942bffa.
//
// Solidity: function transactionAmount() constant returns(uint256)
func (_Gatekeeper *GatekeeperCaller) TransactionAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gatekeeper.contract.Call(opts, out, "transactionAmount")
	return *ret0, err
}

// TransactionAmount is a free data retrieval call binding the contract method 0xd942bffa.
//
// Solidity: function transactionAmount() constant returns(uint256)
func (_Gatekeeper *GatekeeperSession) TransactionAmount() (*big.Int, error) {
	return _Gatekeeper.Contract.TransactionAmount(&_Gatekeeper.CallOpts)
}

// TransactionAmount is a free data retrieval call binding the contract method 0xd942bffa.
//
// Solidity: function transactionAmount() constant returns(uint256)
func (_Gatekeeper *GatekeeperCallerSession) TransactionAmount() (*big.Int, error) {
	return _Gatekeeper.Contract.TransactionAmount(&_Gatekeeper.CallOpts)
}

// PayIn is a paid mutator transaction binding the contract method 0x9fab56ac.
//
// Solidity: function PayIn(_value uint256) returns()
func (_Gatekeeper *GatekeeperTransactor) PayIn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "PayIn", _value)
}

// PayIn is a paid mutator transaction binding the contract method 0x9fab56ac.
//
// Solidity: function PayIn(_value uint256) returns()
func (_Gatekeeper *GatekeeperSession) PayIn(_value *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.PayIn(&_Gatekeeper.TransactOpts, _value)
}

// PayIn is a paid mutator transaction binding the contract method 0x9fab56ac.
//
// Solidity: function PayIn(_value uint256) returns()
func (_Gatekeeper *GatekeeperTransactorSession) PayIn(_value *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.PayIn(&_Gatekeeper.TransactOpts, _value)
}

// Payout is a paid mutator transaction binding the contract method 0x634235fc.
//
// Solidity: function Payout(_to address, _value uint256, _txNumber uint256) returns()
func (_Gatekeeper *GatekeeperTransactor) Payout(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _txNumber *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "Payout", _to, _value, _txNumber)
}

// Payout is a paid mutator transaction binding the contract method 0x634235fc.
//
// Solidity: function Payout(_to address, _value uint256, _txNumber uint256) returns()
func (_Gatekeeper *GatekeeperSession) Payout(_to common.Address, _value *big.Int, _txNumber *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.Payout(&_Gatekeeper.TransactOpts, _to, _value, _txNumber)
}

// Payout is a paid mutator transaction binding the contract method 0x634235fc.
//
// Solidity: function Payout(_to address, _value uint256, _txNumber uint256) returns()
func (_Gatekeeper *GatekeeperTransactorSession) Payout(_to common.Address, _value *big.Int, _txNumber *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.Payout(&_Gatekeeper.TransactOpts, _to, _value, _txNumber)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Gatekeeper *GatekeeperTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Gatekeeper *GatekeeperSession) Kill() (*types.Transaction, error) {
	return _Gatekeeper.Contract.Kill(&_Gatekeeper.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Gatekeeper *GatekeeperTransactorSession) Kill() (*types.Transaction, error) {
	return _Gatekeeper.Contract.Kill(&_Gatekeeper.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Gatekeeper *GatekeeperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Gatekeeper *GatekeeperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gatekeeper.Contract.TransferOwnership(&_Gatekeeper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Gatekeeper *GatekeeperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gatekeeper.Contract.TransferOwnership(&_Gatekeeper.TransactOpts, newOwner)
}

// GatekeeperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gatekeeper contract.
type GatekeeperOwnershipTransferredIterator struct {
	Event *GatekeeperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatekeeperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatekeeperOwnershipTransferred)
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
		it.Event = new(GatekeeperOwnershipTransferred)
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
func (it *GatekeeperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatekeeperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatekeeperOwnershipTransferred represents a OwnershipTransferred event raised by the Gatekeeper contract.
type GatekeeperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Gatekeeper *GatekeeperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatekeeperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gatekeeper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatekeeperOwnershipTransferredIterator{contract: _Gatekeeper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Gatekeeper *GatekeeperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatekeeperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gatekeeper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatekeeperOwnershipTransferred)
				if err := _Gatekeeper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// GatekeeperPayInTrxIterator is returned from FilterPayInTrx and is used to iterate over the raw logs and unpacked data for PayInTrx events raised by the Gatekeeper contract.
type GatekeeperPayInTrxIterator struct {
	Event *GatekeeperPayInTrx // Event containing the contract specifics and raw log

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
func (it *GatekeeperPayInTrxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatekeeperPayInTrx)
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
		it.Event = new(GatekeeperPayInTrx)
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
func (it *GatekeeperPayInTrxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatekeeperPayInTrxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatekeeperPayInTrx represents a PayInTrx event raised by the Gatekeeper contract.
type GatekeeperPayInTrx struct {
	From     common.Address
	TxNumber *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayInTrx is a free log retrieval operation binding the contract event 0x5bdf8f42703e019ab5248b71701aac7fc28b9c4d232dc684e53cbf1eb503ae37.
//
// Solidity: event PayInTrx(from indexed address, txNumber indexed uint256, value indexed uint256)
func (_Gatekeeper *GatekeeperFilterer) FilterPayInTrx(opts *bind.FilterOpts, from []common.Address, txNumber []*big.Int, value []*big.Int) (*GatekeeperPayInTrxIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var txNumberRule []interface{}
	for _, txNumberItem := range txNumber {
		txNumberRule = append(txNumberRule, txNumberItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Gatekeeper.contract.FilterLogs(opts, "PayInTrx", fromRule, txNumberRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &GatekeeperPayInTrxIterator{contract: _Gatekeeper.contract, event: "PayInTrx", logs: logs, sub: sub}, nil
}

// WatchPayInTrx is a free log subscription operation binding the contract event 0x5bdf8f42703e019ab5248b71701aac7fc28b9c4d232dc684e53cbf1eb503ae37.
//
// Solidity: event PayInTrx(from indexed address, txNumber indexed uint256, value indexed uint256)
func (_Gatekeeper *GatekeeperFilterer) WatchPayInTrx(opts *bind.WatchOpts, sink chan<- *GatekeeperPayInTrx, from []common.Address, txNumber []*big.Int, value []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var txNumberRule []interface{}
	for _, txNumberItem := range txNumber {
		txNumberRule = append(txNumberRule, txNumberItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Gatekeeper.contract.WatchLogs(opts, "PayInTrx", fromRule, txNumberRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatekeeperPayInTrx)
				if err := _Gatekeeper.contract.UnpackLog(event, "PayInTrx", log); err != nil {
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

// GatekeeperPayOutIterator is returned from FilterPayOut and is used to iterate over the raw logs and unpacked data for PayOut events raised by the Gatekeeper contract.
type GatekeeperPayOutIterator struct {
	Event *GatekeeperPayOut // Event containing the contract specifics and raw log

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
func (it *GatekeeperPayOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatekeeperPayOut)
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
		it.Event = new(GatekeeperPayOut)
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
func (it *GatekeeperPayOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatekeeperPayOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatekeeperPayOut represents a PayOut event raised by the Gatekeeper contract.
type GatekeeperPayOut struct {
	From     common.Address
	TxNumber *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayOut is a free log retrieval operation binding the contract event 0xafe5bde68f2dccddb6979d04a1dceb6099fe0ceda6bfc2c37e3a40ed8195d90f.
//
// Solidity: event PayOut(from indexed address, txNumber indexed uint256, value indexed uint256)
func (_Gatekeeper *GatekeeperFilterer) FilterPayOut(opts *bind.FilterOpts, from []common.Address, txNumber []*big.Int, value []*big.Int) (*GatekeeperPayOutIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var txNumberRule []interface{}
	for _, txNumberItem := range txNumber {
		txNumberRule = append(txNumberRule, txNumberItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Gatekeeper.contract.FilterLogs(opts, "PayOut", fromRule, txNumberRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &GatekeeperPayOutIterator{contract: _Gatekeeper.contract, event: "PayOut", logs: logs, sub: sub}, nil
}

// WatchPayOut is a free log subscription operation binding the contract event 0xafe5bde68f2dccddb6979d04a1dceb6099fe0ceda6bfc2c37e3a40ed8195d90f.
//
// Solidity: event PayOut(from indexed address, txNumber indexed uint256, value indexed uint256)
func (_Gatekeeper *GatekeeperFilterer) WatchPayOut(opts *bind.WatchOpts, sink chan<- *GatekeeperPayOut, from []common.Address, txNumber []*big.Int, value []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var txNumberRule []interface{}
	for _, txNumberItem := range txNumber {
		txNumberRule = append(txNumberRule, txNumberItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Gatekeeper.contract.WatchLogs(opts, "PayOut", fromRule, txNumberRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatekeeperPayOut)
				if err := _Gatekeeper.contract.UnpackLog(event, "PayOut", log); err != nil {
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

// GatekeeperSuicideIterator is returned from FilterSuicide and is used to iterate over the raw logs and unpacked data for Suicide events raised by the Gatekeeper contract.
type GatekeeperSuicideIterator struct {
	Event *GatekeeperSuicide // Event containing the contract specifics and raw log

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
func (it *GatekeeperSuicideIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatekeeperSuicide)
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
		it.Event = new(GatekeeperSuicide)
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
func (it *GatekeeperSuicideIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatekeeperSuicideIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatekeeperSuicide represents a Suicide event raised by the Gatekeeper contract.
type GatekeeperSuicide struct {
	Block *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSuicide is a free log retrieval operation binding the contract event 0xa1ea9b09ea114021983e9ecf71cf2ffddfd80f5cb4f925e5bf24f9bdb5e55fde.
//
// Solidity: event Suicide(block uint256)
func (_Gatekeeper *GatekeeperFilterer) FilterSuicide(opts *bind.FilterOpts) (*GatekeeperSuicideIterator, error) {

	logs, sub, err := _Gatekeeper.contract.FilterLogs(opts, "Suicide")
	if err != nil {
		return nil, err
	}
	return &GatekeeperSuicideIterator{contract: _Gatekeeper.contract, event: "Suicide", logs: logs, sub: sub}, nil
}

// WatchSuicide is a free log subscription operation binding the contract event 0xa1ea9b09ea114021983e9ecf71cf2ffddfd80f5cb4f925e5bf24f9bdb5e55fde.
//
// Solidity: event Suicide(block uint256)
func (_Gatekeeper *GatekeeperFilterer) WatchSuicide(opts *bind.WatchOpts, sink chan<- *GatekeeperSuicide) (event.Subscription, error) {

	logs, sub, err := _Gatekeeper.contract.WatchLogs(opts, "Suicide")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatekeeperSuicide)
				if err := _Gatekeeper.contract.UnpackLog(event, "Suicide", log); err != nil {
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
