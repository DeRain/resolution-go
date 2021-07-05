// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type RegistryForwarderForwardRequest struct {
	From    common.Address
	Gas     *big.Int
	TokenId *big.Int
	Nonce   *big.Int
	Data    []byte
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"NewKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"NewURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"NewURIPrefix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ResetRecords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"valueIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"burnFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"childIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structRegistryForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"hashes\",\"type\":\"uint256[]\"}],\"name\":\"getKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMany\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getManyByHash\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"mintWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reconfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"reconfigureFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"resetFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"resolverOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"safeMintWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMintWithRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"safeTransferFromFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"safeTransferFromFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setByHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setManyByHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setManyFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transferFromFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structRegistryForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
// Solidity: function childIdOf(uint256 tokenId, string label) pure returns(uint256)
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
// Solidity: function childIdOf(uint256 tokenId, string label) pure returns(uint256)
func (_Contract *ContractSession) ChildIdOf(tokenId *big.Int, label string) (*big.Int, error) {
	return _Contract.Contract.ChildIdOf(&_Contract.CallOpts, tokenId, label)
}

// ChildIdOf is a free data retrieval call binding the contract method 0x68b62d32.
//
// Solidity: function childIdOf(uint256 tokenId, string label) pure returns(uint256)
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

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Contract *ContractCaller) GetKey(opts *bind.CallOpts, keyHash *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getKey", keyHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Contract *ContractSession) GetKey(keyHash *big.Int) (string, error) {
	return _Contract.Contract.GetKey(&_Contract.CallOpts, keyHash)
}

// GetKey is a free data retrieval call binding the contract method 0xbb5b27e1.
//
// Solidity: function getKey(uint256 keyHash) view returns(string)
func (_Contract *ContractCallerSession) GetKey(keyHash *big.Int) (string, error) {
	return _Contract.Contract.GetKey(&_Contract.CallOpts, keyHash)
}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Contract *ContractCaller) GetKeys(opts *bind.CallOpts, hashes []*big.Int) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getKeys", hashes)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Contract *ContractSession) GetKeys(hashes []*big.Int) ([]string, error) {
	return _Contract.Contract.GetKeys(&_Contract.CallOpts, hashes)
}

// GetKeys is a free data retrieval call binding the contract method 0xf5c1f76e.
//
// Solidity: function getKeys(uint256[] hashes) view returns(string[] values)
func (_Contract *ContractCallerSession) GetKeys(hashes []*big.Int) ([]string, error) {
	return _Contract.Contract.GetKeys(&_Contract.CallOpts, hashes)
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
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
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

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Contract *ContractCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Contract *ContractSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Contract.Contract.IsTrustedForwarder(&_Contract.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Contract *ContractCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Contract.Contract.IsTrustedForwarder(&_Contract.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractCaller) NonceOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonceOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Contract.Contract.NonceOf(&_Contract.CallOpts, tokenId)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractCallerSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Contract.Contract.NonceOf(&_Contract.CallOpts, tokenId)
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

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Contract *ContractCaller) Root(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Contract *ContractSession) Root() (*big.Int, error) {
	return _Contract.Contract.Root(&_Contract.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() pure returns(uint256)
func (_Contract *ContractCallerSession) Root() (*big.Int, error) {
	return _Contract.Contract.Root(&_Contract.CallOpts)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
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

// Verify is a free data retrieval call binding the contract method 0x1796e180.
//
// Solidity: function verify((address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Contract *ContractCaller) Verify(opts *bind.CallOpts, req RegistryForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x1796e180.
//
// Solidity: function verify((address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Contract *ContractSession) Verify(req RegistryForwarderForwardRequest, signature []byte) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0x1796e180.
//
// Solidity: function verify((address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_Contract *ContractCallerSession) Verify(req RegistryForwarderForwardRequest, signature []byte) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, req, signature)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Contract *ContractTransactor) AddKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addKey", key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Contract *ContractSession) AddKey(key string) (*types.Transaction, error) {
	return _Contract.Contract.AddKey(&_Contract.TransactOpts, key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Contract *ContractTransactorSession) AddKey(key string) (*types.Transaction, error) {
	return _Contract.Contract.AddKey(&_Contract.TransactOpts, key)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// BurnFor is a paid mutator transaction binding the contract method 0x61603dd9.
//
// Solidity: function burnFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) BurnFor(opts *bind.TransactOpts, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burnFor", tokenId, signature)
}

// BurnFor is a paid mutator transaction binding the contract method 0x61603dd9.
//
// Solidity: function burnFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) BurnFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.BurnFor(&_Contract.TransactOpts, tokenId, signature)
}

// BurnFor is a paid mutator transaction binding the contract method 0x61603dd9.
//
// Solidity: function burnFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) BurnFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.BurnFor(&_Contract.TransactOpts, tokenId, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x185be4e5.
//
// Solidity: function execute((address,uint256,uint256,uint256,bytes) req, bytes signature) returns(bool, bytes)
func (_Contract *ContractTransactor) Execute(opts *bind.TransactOpts, req RegistryForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x185be4e5.
//
// Solidity: function execute((address,uint256,uint256,uint256,bytes) req, bytes signature) returns(bool, bytes)
func (_Contract *ContractSession) Execute(req RegistryForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.Execute(&_Contract.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x185be4e5.
//
// Solidity: function execute((address,uint256,uint256,uint256,bytes) req, bytes signature) returns(bool, bytes)
func (_Contract *ContractTransactorSession) Execute(req RegistryForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.Execute(&_Contract.TransactOpts, req, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address mintingManager) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, mintingManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", mintingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address mintingManager) returns()
func (_Contract *ContractSession) Initialize(mintingManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, mintingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address mintingManager) returns()
func (_Contract *ContractTransactorSession) Initialize(mintingManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, mintingManager)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mint", to, tokenId, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractSession) Mint(to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, to, tokenId, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractTransactorSession) Mint(to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, to, tokenId, uri)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xb0f59177.
//
// Solidity: function mintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractTransactor) MintWithRecords(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintWithRecords", to, tokenId, uri, keys, values)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xb0f59177.
//
// Solidity: function mintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractSession) MintWithRecords(to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.Contract.MintWithRecords(&_Contract.TransactOpts, to, tokenId, uri, keys, values)
}

// MintWithRecords is a paid mutator transaction binding the contract method 0xb0f59177.
//
// Solidity: function mintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractTransactorSession) MintWithRecords(to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.Contract.MintWithRecords(&_Contract.TransactOpts, to, tokenId, uri, keys, values)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Reconfigure(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reconfigure", keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) ReconfigureFor(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reconfigureFor", keys, values, tokenId, signature)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) ReconfigureFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ReconfigureFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) ReconfigureFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ReconfigureFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractTransactor) Reset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reset", tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, tokenId)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) ResetFor(opts *bind.TransactOpts, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resetFor", tokenId, signature)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) ResetFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ResetFor(&_Contract.TransactOpts, tokenId, signature)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) ResetFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ResetFor(&_Contract.TransactOpts, tokenId, signature)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri, bytes _data) returns()
func (_Contract *ContractTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeMint", to, tokenId, uri, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri, bytes _data) returns()
func (_Contract *ContractSession) SafeMint(to common.Address, tokenId *big.Int, uri string, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeMint(&_Contract.TransactOpts, to, tokenId, uri, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri, bytes _data) returns()
func (_Contract *ContractTransactorSession) SafeMint(to common.Address, tokenId *big.Int, uri string, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeMint(&_Contract.TransactOpts, to, tokenId, uri, _data)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xcd279c7c.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractTransactor) SafeMint0(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeMint0", to, tokenId, uri)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xcd279c7c.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractSession) SafeMint0(to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SafeMint0(&_Contract.TransactOpts, to, tokenId, uri)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xcd279c7c.
//
// Solidity: function safeMint(address to, uint256 tokenId, string uri) returns()
func (_Contract *ContractTransactorSession) SafeMint0(to common.Address, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SafeMint0(&_Contract.TransactOpts, to, tokenId, uri)
}

// SafeMintWithRecords is a paid mutator transaction binding the contract method 0x4348b562.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractTransactor) SafeMintWithRecords(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeMintWithRecords", to, tokenId, uri, keys, values)
}

// SafeMintWithRecords is a paid mutator transaction binding the contract method 0x4348b562.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractSession) SafeMintWithRecords(to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.Contract.SafeMintWithRecords(&_Contract.TransactOpts, to, tokenId, uri, keys, values)
}

// SafeMintWithRecords is a paid mutator transaction binding the contract method 0x4348b562.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values) returns()
func (_Contract *ContractTransactorSession) SafeMintWithRecords(to common.Address, tokenId *big.Int, uri string, keys []string, values []string) (*types.Transaction, error) {
	return _Contract.Contract.SafeMintWithRecords(&_Contract.TransactOpts, to, tokenId, uri, keys, values)
}

// SafeMintWithRecords0 is a paid mutator transaction binding the contract method 0xefda4d3e.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values, bytes _data) returns()
func (_Contract *ContractTransactor) SafeMintWithRecords0(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, uri string, keys []string, values []string, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeMintWithRecords0", to, tokenId, uri, keys, values, _data)
}

// SafeMintWithRecords0 is a paid mutator transaction binding the contract method 0xefda4d3e.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values, bytes _data) returns()
func (_Contract *ContractSession) SafeMintWithRecords0(to common.Address, tokenId *big.Int, uri string, keys []string, values []string, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeMintWithRecords0(&_Contract.TransactOpts, to, tokenId, uri, keys, values, _data)
}

// SafeMintWithRecords0 is a paid mutator transaction binding the contract method 0xefda4d3e.
//
// Solidity: function safeMintWithRecords(address to, uint256 tokenId, string uri, string[] keys, string[] values, bytes _data) returns()
func (_Contract *ContractTransactorSession) SafeMintWithRecords0(to common.Address, tokenId *big.Int, uri string, keys []string, values []string, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeMintWithRecords0(&_Contract.TransactOpts, to, tokenId, uri, keys, values, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFromFor is a paid mutator transaction binding the contract method 0x280d9b05.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes _data, bytes signature) returns()
func (_Contract *ContractTransactor) SafeTransferFromFor(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFromFor", from, to, tokenId, _data, signature)
}

// SafeTransferFromFor is a paid mutator transaction binding the contract method 0x280d9b05.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes _data, bytes signature) returns()
func (_Contract *ContractSession) SafeTransferFromFor(from common.Address, to common.Address, tokenId *big.Int, _data []byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFromFor(&_Contract.TransactOpts, from, to, tokenId, _data, signature)
}

// SafeTransferFromFor is a paid mutator transaction binding the contract method 0x280d9b05.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes _data, bytes signature) returns()
func (_Contract *ContractTransactorSession) SafeTransferFromFor(from common.Address, to common.Address, tokenId *big.Int, _data []byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFromFor(&_Contract.TransactOpts, from, to, tokenId, _data, signature)
}

// SafeTransferFromFor0 is a paid mutator transaction binding the contract method 0x6debcb8d.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) SafeTransferFromFor0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFromFor0", from, to, tokenId, signature)
}

// SafeTransferFromFor0 is a paid mutator transaction binding the contract method 0x6debcb8d.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) SafeTransferFromFor0(from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFromFor0(&_Contract.TransactOpts, from, to, tokenId, signature)
}

// SafeTransferFromFor0 is a paid mutator transaction binding the contract method 0x6debcb8d.
//
// Solidity: function safeTransferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) SafeTransferFromFor0(from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFromFor0(&_Contract.TransactOpts, from, to, tokenId, signature)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Set(opts *bind.TransactOpts, key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set", key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Set(&_Contract.TransactOpts, key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Set(&_Contract.TransactOpts, key, value, tokenId)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SetByHash(opts *bind.TransactOpts, keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setByHash", keyHash, value, tokenId)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Contract *ContractSession) SetByHash(keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetByHash(&_Contract.TransactOpts, keyHash, value, tokenId)
}

// SetByHash is a paid mutator transaction binding the contract method 0x4a72584d.
//
// Solidity: function setByHash(uint256 keyHash, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SetByHash(keyHash *big.Int, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetByHash(&_Contract.TransactOpts, keyHash, value, tokenId)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) SetFor(opts *bind.TransactOpts, key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFor", key, value, tokenId, signature)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) SetFor(key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetFor(&_Contract.TransactOpts, key, value, tokenId, signature)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetFor(key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetFor(&_Contract.TransactOpts, key, value, tokenId, signature)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SetMany(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMany", keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMany(&_Contract.TransactOpts, keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMany(&_Contract.TransactOpts, keys, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SetManyByHash(opts *bind.TransactOpts, keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setManyByHash", keyHashes, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) SetManyByHash(keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetManyByHash(&_Contract.TransactOpts, keyHashes, values, tokenId)
}

// SetManyByHash is a paid mutator transaction binding the contract method 0x27f18975.
//
// Solidity: function setManyByHash(uint256[] keyHashes, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SetManyByHash(keyHashes []*big.Int, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetManyByHash(&_Contract.TransactOpts, keyHashes, values, tokenId)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) SetManyFor(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setManyFor", keys, values, tokenId, signature)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) SetManyFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetManyFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetManyFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetManyFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SetOwner(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOwner", to, tokenId)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SetOwner(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, to, tokenId)
}

// SetOwner is a paid mutator transaction binding the contract method 0xab3b87fe.
//
// Solidity: function setOwner(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SetOwner(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, to, tokenId)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Contract *ContractTransactor) SetTokenURIPrefix(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTokenURIPrefix", prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Contract *ContractSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenURIPrefix(&_Contract.TransactOpts, prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_Contract *ContractTransactorSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenURIPrefix(&_Contract.TransactOpts, prefix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFromFor is a paid mutator transaction binding the contract method 0xef2c3088.
//
// Solidity: function transferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) TransferFromFor(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFromFor", from, to, tokenId, signature)
}

// TransferFromFor is a paid mutator transaction binding the contract method 0xef2c3088.
//
// Solidity: function transferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) TransferFromFor(from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.TransferFromFor(&_Contract.TransactOpts, from, to, tokenId, signature)
}

// TransferFromFor is a paid mutator transaction binding the contract method 0xef2c3088.
//
// Solidity: function transferFromFor(address from, address to, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) TransferFromFor(from common.Address, to common.Address, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.TransferFromFor(&_Contract.TransactOpts, from, to, tokenId, signature)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
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
		it.Event = new(ContractApprovalForAll)
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
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewKeyIterator is returned from FilterNewKey and is used to iterate over the raw logs and unpacked data for NewKey events raised by the Contract contract.
type ContractNewKeyIterator struct {
	Event *ContractNewKey // Event containing the contract specifics and raw log

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
func (it *ContractNewKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewKey)
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
		it.Event = new(ContractNewKey)
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
func (it *ContractNewKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewKey represents a NewKey event raised by the Contract contract.
type ContractNewKey struct {
	TokenId  *big.Int
	KeyIndex common.Hash
	Key      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewKey is a free log retrieval operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) FilterNewKey(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string) (*ContractNewKeyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewKeyIterator{contract: _Contract.contract, event: "NewKey", logs: logs, sub: sub}, nil
}

// WatchNewKey is a free log subscription operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) WatchNewKey(opts *bind.WatchOpts, sink chan<- *ContractNewKey, tokenId []*big.Int, keyIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewKey)
				if err := _Contract.contract.UnpackLog(event, "NewKey", log); err != nil {
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

// ParseNewKey is a log parse operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) ParseNewKey(log types.Log) (*ContractNewKey, error) {
	event := new(ContractNewKey)
	if err := _Contract.contract.UnpackLog(event, "NewKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewURIIterator is returned from FilterNewURI and is used to iterate over the raw logs and unpacked data for NewURI events raised by the Contract contract.
type ContractNewURIIterator struct {
	Event *ContractNewURI // Event containing the contract specifics and raw log

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
func (it *ContractNewURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewURI)
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
		it.Event = new(ContractNewURI)
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
func (it *ContractNewURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewURI represents a NewURI event raised by the Contract contract.
type ContractNewURI struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewURI is a free log retrieval operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Contract *ContractFilterer) FilterNewURI(opts *bind.FilterOpts, tokenId []*big.Int) (*ContractNewURIIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewURIIterator{contract: _Contract.contract, event: "NewURI", logs: logs, sub: sub}, nil
}

// WatchNewURI is a free log subscription operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Contract *ContractFilterer) WatchNewURI(opts *bind.WatchOpts, sink chan<- *ContractNewURI, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewURI)
				if err := _Contract.contract.UnpackLog(event, "NewURI", log); err != nil {
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

// ParseNewURI is a log parse operation binding the contract event 0xc5beef08f693b11c316c0c8394a377a0033c9cf701b8cd8afd79cecef60c3952.
//
// Solidity: event NewURI(uint256 indexed tokenId, string uri)
func (_Contract *ContractFilterer) ParseNewURI(log types.Log) (*ContractNewURI, error) {
	event := new(ContractNewURI)
	if err := _Contract.contract.UnpackLog(event, "NewURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewURIPrefixIterator is returned from FilterNewURIPrefix and is used to iterate over the raw logs and unpacked data for NewURIPrefix events raised by the Contract contract.
type ContractNewURIPrefixIterator struct {
	Event *ContractNewURIPrefix // Event containing the contract specifics and raw log

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
func (it *ContractNewURIPrefixIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewURIPrefix)
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
		it.Event = new(ContractNewURIPrefix)
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
func (it *ContractNewURIPrefixIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewURIPrefixIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewURIPrefix represents a NewURIPrefix event raised by the Contract contract.
type ContractNewURIPrefix struct {
	Prefix string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewURIPrefix is a free log retrieval operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Contract *ContractFilterer) FilterNewURIPrefix(opts *bind.FilterOpts) (*ContractNewURIPrefixIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewURIPrefix")
	if err != nil {
		return nil, err
	}
	return &ContractNewURIPrefixIterator{contract: _Contract.contract, event: "NewURIPrefix", logs: logs, sub: sub}, nil
}

// WatchNewURIPrefix is a free log subscription operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Contract *ContractFilterer) WatchNewURIPrefix(opts *bind.WatchOpts, sink chan<- *ContractNewURIPrefix) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewURIPrefix")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewURIPrefix)
				if err := _Contract.contract.UnpackLog(event, "NewURIPrefix", log); err != nil {
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

// ParseNewURIPrefix is a log parse operation binding the contract event 0x4b120d6a959a84a520fa48f5f937cca0e79129423487af7901213b5d2e89313b.
//
// Solidity: event NewURIPrefix(string prefix)
func (_Contract *ContractFilterer) ParseNewURIPrefix(log types.Log) (*ContractNewURIPrefix, error) {
	event := new(ContractNewURIPrefix)
	if err := _Contract.contract.UnpackLog(event, "NewURIPrefix", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResetRecordsIterator is returned from FilterResetRecords and is used to iterate over the raw logs and unpacked data for ResetRecords events raised by the Contract contract.
type ContractResetRecordsIterator struct {
	Event *ContractResetRecords // Event containing the contract specifics and raw log

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
func (it *ContractResetRecordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResetRecords)
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
		it.Event = new(ContractResetRecords)
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
func (it *ContractResetRecordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResetRecordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResetRecords represents a ResetRecords event raised by the Contract contract.
type ContractResetRecords struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetRecords is a free log retrieval operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterResetRecords(opts *bind.FilterOpts, tokenId []*big.Int) (*ContractResetRecordsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractResetRecordsIterator{contract: _Contract.contract, event: "ResetRecords", logs: logs, sub: sub}, nil
}

// WatchResetRecords is a free log subscription operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchResetRecords(opts *bind.WatchOpts, sink chan<- *ContractResetRecords, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResetRecords)
				if err := _Contract.contract.UnpackLog(event, "ResetRecords", log); err != nil {
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

// ParseResetRecords is a log parse operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseResetRecords(log types.Log) (*ContractResetRecords, error) {
	event := new(ContractResetRecords)
	if err := _Contract.contract.UnpackLog(event, "ResetRecords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Contract contract.
type ContractSetIterator struct {
	Event *ContractSet // Event containing the contract specifics and raw log

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
func (it *ContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSet)
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
		it.Event = new(ContractSet)
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
func (it *ContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSet represents a Set event raised by the Contract contract.
type ContractSet struct {
	TokenId    *big.Int
	KeyIndex   common.Hash
	ValueIndex common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) FilterSet(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string, valueIndex []string) (*ContractSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetIterator{contract: _Contract.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ContractSet, tokenId []*big.Int, keyIndex []string, valueIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSet)
				if err := _Contract.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) ParseSet(log types.Log) (*ContractSet, error) {
	event := new(ContractSet)
	if err := _Contract.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
