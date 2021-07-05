// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxyreader

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"unsRegistry\",\"type\":\"address\"},{\"internalType\":\"contractICryptoRegistry\",\"name\":\"cnsRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"childIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDataByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getDataByHashForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"keys\",\"type\":\"string[][]\"},{\"internalType\":\"string[][]\",\"name\":\"values\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getDataForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"values\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMany\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getManyByHash\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ownerOfForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"registryOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"resolverOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractSession) NAME() (string, error) {
	return _Contract.Contract.NAME(&_Contract.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractCallerSession) NAME() (string, error) {
	return _Contract.Contract.NAME(&_Contract.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractSession) VERSION() (string, error) {
	return _Contract.Contract.VERSION(&_Contract.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractCallerSession) VERSION() (string, error) {
	return _Contract.Contract.VERSION(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// ChildIdOf is a free data retrieval call binding the contract method 0x68b62d32.
//
// Solidity: function childIdOf(uint256 tokenId, string label) view returns(uint256)
func (_Contract *ContractCaller) ChildIdOf(opts *bind.CallOpts, tokenId *big.Int, label string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "childIdOf", tokenId, label)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildIdOf is a free data retrieval call binding the contract method 0x68b62d32.
//
// Solidity: function childIdOf(uint256 tokenId, string label) view returns(uint256)
func (_Contract *ContractSession) ChildIdOf(tokenId *big.Int, label string) (*big.Int, error) {
	return _Contract.Contract.ChildIdOf(&_Contract.CallOpts, tokenId, label)
}

// ChildIdOf is a free data retrieval call binding the contract method 0x68b62d32.
//
// Solidity: function childIdOf(uint256 tokenId, string label) view returns(uint256)
func (_Contract *ContractCallerSession) ChildIdOf(tokenId *big.Int, label string) (*big.Int, error) {
	return _Contract.Contract.ChildIdOf(&_Contract.CallOpts, tokenId, label)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractSession) Exists(tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Contract *ContractCaller) Get(opts *bind.CallOpts, key string, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get", key, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Contract *ContractSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Contract *ContractCallerSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractCaller) GetByHash(opts *bind.CallOpts, keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getByHash", keyHash, tokenId)

	outstruct := new(struct {
		Key   string
		Value string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Contract.Contract.GetByHash(&_Contract.CallOpts, keyHash, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractCallerSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Contract.Contract.GetByHash(&_Contract.CallOpts, keyHash, tokenId)
}

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractCaller) GetData(opts *bind.CallOpts, keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getData", keys, tokenId)

	outstruct := new(struct {
		Resolver common.Address
		Owner    common.Address
		Values   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([]string)).(*[]string)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractSession) GetData(keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, keys, tokenId)
}

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractCallerSession) GetData(keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, keys, tokenId)
}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractCaller) GetDataByHash(opts *bind.CallOpts, keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataByHash", keyHashes, tokenId)

	outstruct := new(struct {
		Resolver common.Address
		Owner    common.Address
		Keys     []string
		Values   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Keys = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[3], new([]string)).(*[]string)

	return *outstruct, err

}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractSession) GetDataByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	return _Contract.Contract.GetDataByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractCallerSession) GetDataByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	return _Contract.Contract.GetDataByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractCaller) GetDataByHashForMany(opts *bind.CallOpts, keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataByHashForMany", keyHashes, tokenIds)

	outstruct := new(struct {
		Resolvers []common.Address
		Owners    []common.Address
		Keys      [][]string
		Values    [][]string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolvers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Keys = *abi.ConvertType(out[2], new([][]string)).(*[][]string)
	outstruct.Values = *abi.ConvertType(out[3], new([][]string)).(*[][]string)

	return *outstruct, err

}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractSession) GetDataByHashForMany(keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataByHashForMany(&_Contract.CallOpts, keyHashes, tokenIds)
}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractCallerSession) GetDataByHashForMany(keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataByHashForMany(&_Contract.CallOpts, keyHashes, tokenIds)
}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractCaller) GetDataForMany(opts *bind.CallOpts, keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataForMany", keys, tokenIds)

	outstruct := new(struct {
		Resolvers []common.Address
		Owners    []common.Address
		Values    [][]string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolvers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([][]string)).(*[][]string)

	return *outstruct, err

}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractSession) GetDataForMany(keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataForMany(&_Contract.CallOpts, keys, tokenIds)
}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractCallerSession) GetDataForMany(keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataForMany(&_Contract.CallOpts, keys, tokenIds)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Contract *ContractCaller) GetMany(opts *bind.CallOpts, keys []string, tokenId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMany", keys, tokenId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Contract *ContractSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Contract.Contract.GetMany(&_Contract.CallOpts, keys, tokenId)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Contract *ContractCallerSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Contract.Contract.GetMany(&_Contract.CallOpts, keys, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractCaller) GetManyByHash(opts *bind.CallOpts, keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getManyByHash", keyHashes, tokenId)

	outstruct := new(struct {
		Keys   []string
		Values []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Contract.Contract.GetManyByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractCallerSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Contract.Contract.GetManyByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, arg0, arg1)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractCaller) IsApprovedOrOwner(opts *bind.CallOpts, spender common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedOrOwner", spender, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Contract.Contract.IsApprovedOrOwner(&_Contract.CallOpts, spender, tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Contract.Contract.IsApprovedOrOwner(&_Contract.CallOpts, spender, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractCaller) OwnerOfForMany(opts *bind.CallOpts, tokenIds []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOfForMany", tokenIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractSession) OwnerOfForMany(tokenIds []*big.Int) ([]common.Address, error) {
	return _Contract.Contract.OwnerOfForMany(&_Contract.CallOpts, tokenIds)
}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractCallerSession) OwnerOfForMany(tokenIds []*big.Int) ([]common.Address, error) {
	return _Contract.Contract.OwnerOfForMany(&_Contract.CallOpts, tokenIds)
}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) RegistryOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registryOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) RegistryOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.RegistryOf(&_Contract.CallOpts, tokenId)
}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) RegistryOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.RegistryOf(&_Contract.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) ResolverOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolverOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.ResolverOf(&_Contract.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.ResolverOf(&_Contract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}
