// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ABI

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ABIErc1155MetaData contains all meta data concerning the ABIErc1155 contract.
var ABIErc1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUrl\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokenData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"selfUse\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buyNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUrl\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_selfUse\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenUrl\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_selfUse\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"mintProxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpcAddress\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_tokenUrl\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"uriBatch\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_tokenUrls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ABIErc1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ABIErc1155MetaData.ABI instead.
var ABIErc1155ABI = ABIErc1155MetaData.ABI

// ABIErc1155 is an auto generated Go binding around an Ethereum contract.
type ABIErc1155 struct {
	ABIErc1155Caller     // Read-only binding to the contract
	ABIErc1155Transactor // Write-only binding to the contract
	ABIErc1155Filterer   // Log filterer for contract events
}

// ABIErc1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ABIErc1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABIErc1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ABIErc1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABIErc1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ABIErc1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABIErc1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ABIErc1155Session struct {
	Contract     *ABIErc1155       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ABIErc1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ABIErc1155CallerSession struct {
	Contract *ABIErc1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ABIErc1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ABIErc1155TransactorSession struct {
	Contract     *ABIErc1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ABIErc1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ABIErc1155Raw struct {
	Contract *ABIErc1155 // Generic contract binding to access the raw methods on
}

// ABIErc1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ABIErc1155CallerRaw struct {
	Contract *ABIErc1155Caller // Generic read-only contract binding to access the raw methods on
}

// ABIErc1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ABIErc1155TransactorRaw struct {
	Contract *ABIErc1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewABIErc1155 creates a new instance of ABIErc1155, bound to a specific deployed contract.
func NewABIErc1155(address common.Address, backend bind.ContractBackend) (*ABIErc1155, error) {
	contract, err := bindABIErc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155{ABIErc1155Caller: ABIErc1155Caller{contract: contract}, ABIErc1155Transactor: ABIErc1155Transactor{contract: contract}, ABIErc1155Filterer: ABIErc1155Filterer{contract: contract}}, nil
}

// NewABIErc1155Caller creates a new read-only instance of ABIErc1155, bound to a specific deployed contract.
func NewABIErc1155Caller(address common.Address, caller bind.ContractCaller) (*ABIErc1155Caller, error) {
	contract, err := bindABIErc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155Caller{contract: contract}, nil
}

// NewABIErc1155Transactor creates a new write-only instance of ABIErc1155, bound to a specific deployed contract.
func NewABIErc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ABIErc1155Transactor, error) {
	contract, err := bindABIErc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155Transactor{contract: contract}, nil
}

// NewABIErc1155Filterer creates a new log filterer instance of ABIErc1155, bound to a specific deployed contract.
func NewABIErc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ABIErc1155Filterer, error) {
	contract, err := bindABIErc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155Filterer{contract: contract}, nil
}

// bindABIErc1155 binds a generic wrapper to an already deployed contract.
func bindABIErc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ABIErc1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABIErc1155 *ABIErc1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ABIErc1155.Contract.ABIErc1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABIErc1155 *ABIErc1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABIErc1155.Contract.ABIErc1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABIErc1155 *ABIErc1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABIErc1155.Contract.ABIErc1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABIErc1155 *ABIErc1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ABIErc1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABIErc1155 *ABIErc1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABIErc1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABIErc1155 *ABIErc1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABIErc1155.Contract.contract.Transact(opts, method, params...)
}

// TokenData is a free data retrieval call binding the contract method 0x2c2cdd60.
//
// Solidity: function _tokenData(uint256 ) view returns(uint256 price, bool selfUse, string tokenUri, uint256 limit, uint256 startTime, uint256 endTime)
func (_ABIErc1155 *ABIErc1155Caller) TokenData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price     *big.Int
	SelfUse   bool
	TokenUri  string
	Limit     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "_tokenData", arg0)

	outstruct := new(struct {
		Price     *big.Int
		SelfUse   bool
		TokenUri  string
		Limit     *big.Int
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SelfUse = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TokenUri = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Limit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenData is a free data retrieval call binding the contract method 0x2c2cdd60.
//
// Solidity: function _tokenData(uint256 ) view returns(uint256 price, bool selfUse, string tokenUri, uint256 limit, uint256 startTime, uint256 endTime)
func (_ABIErc1155 *ABIErc1155Session) TokenData(arg0 *big.Int) (struct {
	Price     *big.Int
	SelfUse   bool
	TokenUri  string
	Limit     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _ABIErc1155.Contract.TokenData(&_ABIErc1155.CallOpts, arg0)
}

// TokenData is a free data retrieval call binding the contract method 0x2c2cdd60.
//
// Solidity: function _tokenData(uint256 ) view returns(uint256 price, bool selfUse, string tokenUri, uint256 limit, uint256 startTime, uint256 endTime)
func (_ABIErc1155 *ABIErc1155CallerSession) TokenData(arg0 *big.Int) (struct {
	Price     *big.Int
	SelfUse   bool
	TokenUri  string
	Limit     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _ABIErc1155.Contract.TokenData(&_ABIErc1155.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ABIErc1155 *ABIErc1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ABIErc1155 *ABIErc1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ABIErc1155.Contract.BalanceOf(&_ABIErc1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ABIErc1155 *ABIErc1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ABIErc1155.Contract.BalanceOf(&_ABIErc1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ABIErc1155 *ABIErc1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ABIErc1155 *ABIErc1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ABIErc1155.Contract.BalanceOfBatch(&_ABIErc1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ABIErc1155 *ABIErc1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ABIErc1155.Contract.BalanceOfBatch(&_ABIErc1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ABIErc1155 *ABIErc1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ABIErc1155 *ABIErc1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ABIErc1155.Contract.IsApprovedForAll(&_ABIErc1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ABIErc1155 *ABIErc1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ABIErc1155.Contract.IsApprovedForAll(&_ABIErc1155.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ABIErc1155 *ABIErc1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ABIErc1155 *ABIErc1155Session) Owner() (common.Address, error) {
	return _ABIErc1155.Contract.Owner(&_ABIErc1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ABIErc1155 *ABIErc1155CallerSession) Owner() (common.Address, error) {
	return _ABIErc1155.Contract.Owner(&_ABIErc1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ABIErc1155 *ABIErc1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ABIErc1155 *ABIErc1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ABIErc1155.Contract.SupportsInterface(&_ABIErc1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ABIErc1155 *ABIErc1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ABIErc1155.Contract.SupportsInterface(&_ABIErc1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string _tokenUrl)
func (_ABIErc1155 *ABIErc1155Caller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string _tokenUrl)
func (_ABIErc1155 *ABIErc1155Session) Uri(_id *big.Int) (string, error) {
	return _ABIErc1155.Contract.Uri(&_ABIErc1155.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string _tokenUrl)
func (_ABIErc1155 *ABIErc1155CallerSession) Uri(_id *big.Int) (string, error) {
	return _ABIErc1155.Contract.Uri(&_ABIErc1155.CallOpts, _id)
}

// UriBatch is a free data retrieval call binding the contract method 0x96d7e229.
//
// Solidity: function uriBatch(uint256[] _tokenIds) view returns(string[] _tokenUrls)
func (_ABIErc1155 *ABIErc1155Caller) UriBatch(opts *bind.CallOpts, _tokenIds []*big.Int) ([]string, error) {
	var out []interface{}
	err := _ABIErc1155.contract.Call(opts, &out, "uriBatch", _tokenIds)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// UriBatch is a free data retrieval call binding the contract method 0x96d7e229.
//
// Solidity: function uriBatch(uint256[] _tokenIds) view returns(string[] _tokenUrls)
func (_ABIErc1155 *ABIErc1155Session) UriBatch(_tokenIds []*big.Int) ([]string, error) {
	return _ABIErc1155.Contract.UriBatch(&_ABIErc1155.CallOpts, _tokenIds)
}

// UriBatch is a free data retrieval call binding the contract method 0x96d7e229.
//
// Solidity: function uriBatch(uint256[] _tokenIds) view returns(string[] _tokenUrls)
func (_ABIErc1155 *ABIErc1155CallerSession) UriBatch(_tokenIds []*big.Int) ([]string, error) {
	return _ABIErc1155.Contract.UriBatch(&_ABIErc1155.CallOpts, _tokenIds)
}

// BuyNFTs is a paid mutator transaction binding the contract method 0x406bb06c.
//
// Solidity: function buyNFTs(uint256 tokenId, address _account, uint256 _amount) returns()
func (_ABIErc1155 *ABIErc1155Transactor) BuyNFTs(opts *bind.TransactOpts, tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "buyNFTs", tokenId, _account, _amount)
}

// BuyNFTs is a paid mutator transaction binding the contract method 0x406bb06c.
//
// Solidity: function buyNFTs(uint256 tokenId, address _account, uint256 _amount) returns()
func (_ABIErc1155 *ABIErc1155Session) BuyNFTs(tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.BuyNFTs(&_ABIErc1155.TransactOpts, tokenId, _account, _amount)
}

// BuyNFTs is a paid mutator transaction binding the contract method 0x406bb06c.
//
// Solidity: function buyNFTs(uint256 tokenId, address _account, uint256 _amount) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) BuyNFTs(tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.BuyNFTs(&_ABIErc1155.TransactOpts, tokenId, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x58390413.
//
// Solidity: function mint(uint256 _amount, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155Transactor) Mint(opts *bind.TransactOpts, _amount *big.Int, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "mint", _amount, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// Mint is a paid mutator transaction binding the contract method 0x58390413.
//
// Solidity: function mint(uint256 _amount, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155Session) Mint(_amount *big.Int, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.Mint(&_ABIErc1155.TransactOpts, _amount, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// Mint is a paid mutator transaction binding the contract method 0x58390413.
//
// Solidity: function mint(uint256 _amount, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155TransactorSession) Mint(_amount *big.Int, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.Mint(&_ABIErc1155.TransactOpts, _amount, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// MintProxy is a paid mutator transaction binding the contract method 0x1d208c11.
//
// Solidity: function mintProxy(uint256 _amount, address _address, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155Transactor) MintProxy(opts *bind.TransactOpts, _amount *big.Int, _address common.Address, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "mintProxy", _amount, _address, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// MintProxy is a paid mutator transaction binding the contract method 0x1d208c11.
//
// Solidity: function mintProxy(uint256 _amount, address _address, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155Session) MintProxy(_amount *big.Int, _address common.Address, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.MintProxy(&_ABIErc1155.TransactOpts, _amount, _address, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// MintProxy is a paid mutator transaction binding the contract method 0x1d208c11.
//
// Solidity: function mintProxy(uint256 _amount, address _address, string _tokenUrl, uint256 _tokenPrice, bool _selfUse, uint256 limit, uint256 startTime, uint256 endTime) returns(uint256 _mintTokenId)
func (_ABIErc1155 *ABIErc1155TransactorSession) MintProxy(_amount *big.Int, _address common.Address, _tokenUrl string, _tokenPrice *big.Int, _selfUse bool, limit *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _ABIErc1155.Contract.MintProxy(&_ABIErc1155.TransactOpts, _amount, _address, _tokenUrl, _tokenPrice, _selfUse, limit, startTime, endTime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ABIErc1155 *ABIErc1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ABIErc1155 *ABIErc1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _ABIErc1155.Contract.RenounceOwnership(&_ABIErc1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ABIErc1155.Contract.RenounceOwnership(&_ABIErc1155.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ABIErc1155 *ABIErc1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ABIErc1155 *ABIErc1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SafeBatchTransferFrom(&_ABIErc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SafeBatchTransferFrom(&_ABIErc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ABIErc1155 *ABIErc1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ABIErc1155 *ABIErc1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SafeTransferFrom(&_ABIErc1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SafeTransferFrom(&_ABIErc1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ABIErc1155 *ABIErc1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ABIErc1155 *ABIErc1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SetApprovalForAll(&_ABIErc1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SetApprovalForAll(&_ABIErc1155.TransactOpts, operator, approved)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address _cpcAddress) returns()
func (_ABIErc1155 *ABIErc1155Transactor) SetContract(opts *bind.TransactOpts, _cpcAddress common.Address) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "setContract", _cpcAddress)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address _cpcAddress) returns()
func (_ABIErc1155 *ABIErc1155Session) SetContract(_cpcAddress common.Address) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SetContract(&_ABIErc1155.TransactOpts, _cpcAddress)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address _cpcAddress) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) SetContract(_cpcAddress common.Address) (*types.Transaction, error) {
	return _ABIErc1155.Contract.SetContract(&_ABIErc1155.TransactOpts, _cpcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ABIErc1155 *ABIErc1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ABIErc1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ABIErc1155 *ABIErc1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ABIErc1155.Contract.TransferOwnership(&_ABIErc1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ABIErc1155 *ABIErc1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ABIErc1155.Contract.TransferOwnership(&_ABIErc1155.TransactOpts, newOwner)
}

// ABIErc1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ABIErc1155 contract.
type ABIErc1155ApprovalForAllIterator struct {
	Event *ABIErc1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ABIErc1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ABIErc1155ApprovalForAll)
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
		it.Event = new(ABIErc1155ApprovalForAll)
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
func (it *ABIErc1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ABIErc1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ABIErc1155ApprovalForAll represents a ApprovalForAll event raised by the ABIErc1155 contract.
type ABIErc1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ABIErc1155 *ABIErc1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ABIErc1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ABIErc1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155ApprovalForAllIterator{contract: _ABIErc1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ABIErc1155 *ABIErc1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ABIErc1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ABIErc1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ABIErc1155ApprovalForAll)
				if err := _ABIErc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ABIErc1155 *ABIErc1155Filterer) ParseApprovalForAll(log types.Log) (*ABIErc1155ApprovalForAll, error) {
	event := new(ABIErc1155ApprovalForAll)
	if err := _ABIErc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ABIErc1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ABIErc1155 contract.
type ABIErc1155OwnershipTransferredIterator struct {
	Event *ABIErc1155OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ABIErc1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ABIErc1155OwnershipTransferred)
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
		it.Event = new(ABIErc1155OwnershipTransferred)
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
func (it *ABIErc1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ABIErc1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ABIErc1155OwnershipTransferred represents a OwnershipTransferred event raised by the ABIErc1155 contract.
type ABIErc1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ABIErc1155 *ABIErc1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ABIErc1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ABIErc1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155OwnershipTransferredIterator{contract: _ABIErc1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ABIErc1155 *ABIErc1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ABIErc1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ABIErc1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ABIErc1155OwnershipTransferred)
				if err := _ABIErc1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ABIErc1155 *ABIErc1155Filterer) ParseOwnershipTransferred(log types.Log) (*ABIErc1155OwnershipTransferred, error) {
	event := new(ABIErc1155OwnershipTransferred)
	if err := _ABIErc1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ABIErc1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ABIErc1155 contract.
type ABIErc1155TransferBatchIterator struct {
	Event *ABIErc1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *ABIErc1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ABIErc1155TransferBatch)
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
		it.Event = new(ABIErc1155TransferBatch)
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
func (it *ABIErc1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ABIErc1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ABIErc1155TransferBatch represents a TransferBatch event raised by the ABIErc1155 contract.
type ABIErc1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ABIErc1155 *ABIErc1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ABIErc1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ABIErc1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155TransferBatchIterator{contract: _ABIErc1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ABIErc1155 *ABIErc1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ABIErc1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ABIErc1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ABIErc1155TransferBatch)
				if err := _ABIErc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ABIErc1155 *ABIErc1155Filterer) ParseTransferBatch(log types.Log) (*ABIErc1155TransferBatch, error) {
	event := new(ABIErc1155TransferBatch)
	if err := _ABIErc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ABIErc1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ABIErc1155 contract.
type ABIErc1155TransferSingleIterator struct {
	Event *ABIErc1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *ABIErc1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ABIErc1155TransferSingle)
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
		it.Event = new(ABIErc1155TransferSingle)
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
func (it *ABIErc1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ABIErc1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ABIErc1155TransferSingle represents a TransferSingle event raised by the ABIErc1155 contract.
type ABIErc1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ABIErc1155 *ABIErc1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ABIErc1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ABIErc1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155TransferSingleIterator{contract: _ABIErc1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ABIErc1155 *ABIErc1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ABIErc1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ABIErc1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ABIErc1155TransferSingle)
				if err := _ABIErc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ABIErc1155 *ABIErc1155Filterer) ParseTransferSingle(log types.Log) (*ABIErc1155TransferSingle, error) {
	event := new(ABIErc1155TransferSingle)
	if err := _ABIErc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ABIErc1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ABIErc1155 contract.
type ABIErc1155URIIterator struct {
	Event *ABIErc1155URI // Event containing the contract specifics and raw log

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
func (it *ABIErc1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ABIErc1155URI)
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
		it.Event = new(ABIErc1155URI)
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
func (it *ABIErc1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ABIErc1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ABIErc1155URI represents a URI event raised by the ABIErc1155 contract.
type ABIErc1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ABIErc1155 *ABIErc1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ABIErc1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ABIErc1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ABIErc1155URIIterator{contract: _ABIErc1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ABIErc1155 *ABIErc1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ABIErc1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ABIErc1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ABIErc1155URI)
				if err := _ABIErc1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ABIErc1155 *ABIErc1155Filterer) ParseURI(log types.Log) (*ABIErc1155URI, error) {
	event := new(ABIErc1155URI)
	if err := _ABIErc1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
