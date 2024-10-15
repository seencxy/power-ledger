// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package artifacts

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// HomomorphicEncryptionABI is the input ABI used to generate the binding from.
const HomomorphicEncryptionABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"}],\"name\":\"average\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKey\",\"type\":\"uint256\"}],\"name\":\"decrypt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKey\",\"type\":\"uint256\"}],\"name\":\"encrypt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"}],\"name\":\"isGreaterThan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"}],\"name\":\"isGreaterThanOrEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"}],\"name\":\"isLessThan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encryptedA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"encryptedB\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"}],\"name\":\"isLessThanOrEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HomomorphicEncryptionBin is the compiled bytecode used for deploying new contracts.
var HomomorphicEncryptionBin = "0x608060405234801561001057600080fd5b5061065e806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806373556dad1161005b57806373556dad146101125780637373861b146101425780637e0e155714610172578063fef51107146101a25761007d565b8063457e403a1461008257806362f95002146100b2578063632544f1146100e2575b600080fd5b61009c6004803603810190610097919061043c565b6101d2565b6040516100a991906104be565b60405180910390f35b6100cc60048036038101906100c7919061043c565b61021f565b6040516100d991906104be565b60405180910390f35b6100fc60048036038101906100f7919061043c565b61026d565b60405161010991906104be565b60405180910390f35b61012c600480360381019061012791906104d9565b6102bb565b6040516101399190610528565b60405180910390f35b61015c6004803603810190610157919061043c565b6102c8565b6040516101699190610528565b60405180910390f35b61018c600480360381019061018791906104d9565b610349565b6040516101999190610528565b60405180910390f35b6101bc60048036038101906101b7919061043c565b610356565b6040516101c991906104be565b60405180910390f35b60006101f4838373ffffffffffffffffffffffffffffffffffffffff16610349565b610214868673ffffffffffffffffffffffffffffffffffffffff16610349565b109050949350505050565b6000610241838373ffffffffffffffffffffffffffffffffffffffff16610349565b610261868673ffffffffffffffffffffffffffffffffffffffff16610349565b10159050949350505050565b600061028f838373ffffffffffffffffffffffffffffffffffffffff16610349565b6102af868673ffffffffffffffffffffffffffffffffffffffff16610349565b11159050949350505050565b6000818318905092915050565b600061033f60026102ef858573ffffffffffffffffffffffffffffffffffffffff16610349565b61030f888873ffffffffffffffffffffffffffffffffffffffff16610349565b6103199190610572565b61032391906105f7565b8573ffffffffffffffffffffffffffffffffffffffff166102bb565b9050949350505050565b6000818318905092915050565b6000610378838373ffffffffffffffffffffffffffffffffffffffff16610349565b610398868673ffffffffffffffffffffffffffffffffffffffff16610349565b119050949350505050565b600080fd5b6000819050919050565b6103bb816103a8565b81146103c657600080fd5b50565b6000813590506103d8816103b2565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610409826103de565b9050919050565b610419816103fe565b811461042457600080fd5b50565b60008135905061043681610410565b92915050565b60008060008060808587031215610456576104556103a3565b5b6000610464878288016103c9565b945050602061047587828801610427565b9350506040610486878288016103c9565b925050606061049787828801610427565b91505092959194509250565b60008115159050919050565b6104b8816104a3565b82525050565b60006020820190506104d360008301846104af565b92915050565b600080604083850312156104f0576104ef6103a3565b5b60006104fe858286016103c9565b925050602061050f858286016103c9565b9150509250929050565b610522816103a8565b82525050565b600060208201905061053d6000830184610519565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061057d826103a8565b9150610588836103a8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156105bd576105bc610543565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610602826103a8565b915061060d836103a8565b92508261061d5761061c6105c8565b5b82820490509291505056fea264697066735822122005a998184ed629bf530634173b6113ac74289d8681cb367e11cf7d81dd05e19764736f6c634300080b0033"
var HomomorphicEncryptionSMBin = "0x"

// DeployHomomorphicEncryption deploys a new contract, binding an instance of HomomorphicEncryption to it.
func DeployHomomorphicEncryption(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *HomomorphicEncryption, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphicEncryptionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(HomomorphicEncryptionSMBin)
	} else {
		bytecode = common.FromHex(HomomorphicEncryptionBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, HomomorphicEncryptionABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &HomomorphicEncryption{HomomorphicEncryptionCaller: HomomorphicEncryptionCaller{contract: contract}, HomomorphicEncryptionTransactor: HomomorphicEncryptionTransactor{contract: contract}, HomomorphicEncryptionFilterer: HomomorphicEncryptionFilterer{contract: contract}}, nil
}

func AsyncDeployHomomorphicEncryption(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphicEncryptionABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(HomomorphicEncryptionSMBin)
	} else {
		bytecode = common.FromHex(HomomorphicEncryptionBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, HomomorphicEncryptionABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// HomomorphicEncryption is an auto generated Go binding around a Solidity contract.
type HomomorphicEncryption struct {
	HomomorphicEncryptionCaller     // Read-only binding to the contract
	HomomorphicEncryptionTransactor // Write-only binding to the contract
	HomomorphicEncryptionFilterer   // Log filterer for contract events
}

// HomomorphicEncryptionCaller is an auto generated read-only Go binding around a Solidity contract.
type HomomorphicEncryptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphicEncryptionTransactor is an auto generated write-only Go binding around a Solidity contract.
type HomomorphicEncryptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphicEncryptionFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type HomomorphicEncryptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomomorphicEncryptionSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type HomomorphicEncryptionSession struct {
	Contract     *HomomorphicEncryption // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HomomorphicEncryptionCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type HomomorphicEncryptionCallerSession struct {
	Contract *HomomorphicEncryptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// HomomorphicEncryptionTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type HomomorphicEncryptionTransactorSession struct {
	Contract     *HomomorphicEncryptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// HomomorphicEncryptionRaw is an auto generated low-level Go binding around a Solidity contract.
type HomomorphicEncryptionRaw struct {
	Contract *HomomorphicEncryption // Generic contract binding to access the raw methods on
}

// HomomorphicEncryptionCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type HomomorphicEncryptionCallerRaw struct {
	Contract *HomomorphicEncryptionCaller // Generic read-only contract binding to access the raw methods on
}

// HomomorphicEncryptionTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type HomomorphicEncryptionTransactorRaw struct {
	Contract *HomomorphicEncryptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHomomorphicEncryption creates a new instance of HomomorphicEncryption, bound to a specific deployed contract.
func NewHomomorphicEncryption(address common.Address, backend bind.ContractBackend) (*HomomorphicEncryption, error) {
	contract, err := bindHomomorphicEncryption(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HomomorphicEncryption{HomomorphicEncryptionCaller: HomomorphicEncryptionCaller{contract: contract}, HomomorphicEncryptionTransactor: HomomorphicEncryptionTransactor{contract: contract}, HomomorphicEncryptionFilterer: HomomorphicEncryptionFilterer{contract: contract}}, nil
}

// NewHomomorphicEncryptionCaller creates a new read-only instance of HomomorphicEncryption, bound to a specific deployed contract.
func NewHomomorphicEncryptionCaller(address common.Address, caller bind.ContractCaller) (*HomomorphicEncryptionCaller, error) {
	contract, err := bindHomomorphicEncryption(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomomorphicEncryptionCaller{contract: contract}, nil
}

// NewHomomorphicEncryptionTransactor creates a new write-only instance of HomomorphicEncryption, bound to a specific deployed contract.
func NewHomomorphicEncryptionTransactor(address common.Address, transactor bind.ContractTransactor) (*HomomorphicEncryptionTransactor, error) {
	contract, err := bindHomomorphicEncryption(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomomorphicEncryptionTransactor{contract: contract}, nil
}

// NewHomomorphicEncryptionFilterer creates a new log filterer instance of HomomorphicEncryption, bound to a specific deployed contract.
func NewHomomorphicEncryptionFilterer(address common.Address, filterer bind.ContractFilterer) (*HomomorphicEncryptionFilterer, error) {
	contract, err := bindHomomorphicEncryption(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomomorphicEncryptionFilterer{contract: contract}, nil
}

// bindHomomorphicEncryption binds a generic wrapper to an already deployed contract.
func bindHomomorphicEncryption(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomomorphicEncryptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomomorphicEncryption *HomomorphicEncryptionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HomomorphicEncryption.Contract.HomomorphicEncryptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomomorphicEncryption *HomomorphicEncryptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _HomomorphicEncryption.Contract.HomomorphicEncryptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomomorphicEncryption *HomomorphicEncryptionRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HomomorphicEncryption.Contract.HomomorphicEncryptionTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomomorphicEncryption *HomomorphicEncryptionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HomomorphicEncryption.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomomorphicEncryption *HomomorphicEncryptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _HomomorphicEncryption.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomomorphicEncryption *HomomorphicEncryptionTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HomomorphicEncryption.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Average is a free data retrieval call binding the contract method 0x7373861b.
//
// Solidity: function average(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) Average(opts *bind.CallOpts, encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "average", encryptedA, addrA, encryptedB, addrB)
	return *ret0, err
}

// Average is a free data retrieval call binding the contract method 0x7373861b.
//
// Solidity: function average(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) Average(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Average(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// Average is a free data retrieval call binding the contract method 0x7373861b.
//
// Solidity: function average(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) Average(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Average(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// Decrypt is a free data retrieval call binding the contract method 0x7e0e1557.
//
// Solidity: function decrypt(uint256 encryptedValue, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) Decrypt(opts *bind.CallOpts, encryptedValue *big.Int, publicKey *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "decrypt", encryptedValue, publicKey)
	return *ret0, err
}

// Decrypt is a free data retrieval call binding the contract method 0x7e0e1557.
//
// Solidity: function decrypt(uint256 encryptedValue, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) Decrypt(encryptedValue *big.Int, publicKey *big.Int) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Decrypt(&_HomomorphicEncryption.CallOpts, encryptedValue, publicKey)
}

// Decrypt is a free data retrieval call binding the contract method 0x7e0e1557.
//
// Solidity: function decrypt(uint256 encryptedValue, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) Decrypt(encryptedValue *big.Int, publicKey *big.Int) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Decrypt(&_HomomorphicEncryption.CallOpts, encryptedValue, publicKey)
}

// Encrypt is a free data retrieval call binding the contract method 0x73556dad.
//
// Solidity: function encrypt(uint256 value, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) Encrypt(opts *bind.CallOpts, value *big.Int, publicKey *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "encrypt", value, publicKey)
	return *ret0, err
}

// Encrypt is a free data retrieval call binding the contract method 0x73556dad.
//
// Solidity: function encrypt(uint256 value, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) Encrypt(value *big.Int, publicKey *big.Int) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Encrypt(&_HomomorphicEncryption.CallOpts, value, publicKey)
}

// Encrypt is a free data retrieval call binding the contract method 0x73556dad.
//
// Solidity: function encrypt(uint256 value, uint256 publicKey) constant returns(uint256)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) Encrypt(value *big.Int, publicKey *big.Int) (*big.Int, error) {
	return _HomomorphicEncryption.Contract.Encrypt(&_HomomorphicEncryption.CallOpts, value, publicKey)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xfef51107.
//
// Solidity: function isGreaterThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) IsGreaterThan(opts *bind.CallOpts, encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "isGreaterThan", encryptedA, addrA, encryptedB, addrB)
	return *ret0, err
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xfef51107.
//
// Solidity: function isGreaterThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) IsGreaterThan(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsGreaterThan(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xfef51107.
//
// Solidity: function isGreaterThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) IsGreaterThan(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsGreaterThan(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x62f95002.
//
// Solidity: function isGreaterThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) IsGreaterThanOrEqual(opts *bind.CallOpts, encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "isGreaterThanOrEqual", encryptedA, addrA, encryptedB, addrB)
	return *ret0, err
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x62f95002.
//
// Solidity: function isGreaterThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) IsGreaterThanOrEqual(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsGreaterThanOrEqual(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x62f95002.
//
// Solidity: function isGreaterThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) IsGreaterThanOrEqual(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsGreaterThanOrEqual(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsLessThan is a free data retrieval call binding the contract method 0x457e403a.
//
// Solidity: function isLessThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) IsLessThan(opts *bind.CallOpts, encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "isLessThan", encryptedA, addrA, encryptedB, addrB)
	return *ret0, err
}

// IsLessThan is a free data retrieval call binding the contract method 0x457e403a.
//
// Solidity: function isLessThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) IsLessThan(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsLessThan(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsLessThan is a free data retrieval call binding the contract method 0x457e403a.
//
// Solidity: function isLessThan(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) IsLessThan(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsLessThan(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x632544f1.
//
// Solidity: function isLessThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCaller) IsLessThanOrEqual(opts *bind.CallOpts, encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomomorphicEncryption.contract.Call(opts, out, "isLessThanOrEqual", encryptedA, addrA, encryptedB, addrB)
	return *ret0, err
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x632544f1.
//
// Solidity: function isLessThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionSession) IsLessThanOrEqual(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsLessThanOrEqual(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x632544f1.
//
// Solidity: function isLessThanOrEqual(uint256 encryptedA, address addrA, uint256 encryptedB, address addrB) constant returns(bool)
func (_HomomorphicEncryption *HomomorphicEncryptionCallerSession) IsLessThanOrEqual(encryptedA *big.Int, addrA common.Address, encryptedB *big.Int, addrB common.Address) (bool, error) {
	return _HomomorphicEncryption.Contract.IsLessThanOrEqual(&_HomomorphicEncryption.CallOpts, encryptedA, addrA, encryptedB, addrB)
}
