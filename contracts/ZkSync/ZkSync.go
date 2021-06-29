// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZkSync

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

// ZkSyncABI is the input ABI used to generate the binding from.
const ZkSyncABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EMPTY_STRING_KECCAK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_n\",\"type\":\"uint64\"}],\"name\":\"cancelOutstandingDepositsForExodusMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_feeAccount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_newBlockInfo\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"_publicData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_ethWitness\",\"type\":\"bytes\"},{\"internalType\":\"uint32[]\",\"name\":\"_ethWitnessSizes\",\"type\":\"uint32[]\"}],\"name\":\"commitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_n\",\"type\":\"uint32\"}],\"name\":\"completeWithdrawals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"address\",\"name\":\"_franklinAddr\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_franklinAddr\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_tokenId\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256[]\",\"name\":\"_proof\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"fullExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxBlocksToRevert\",\"type\":\"uint32\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey_hash\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"setAuthPubkeyHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"triggerExodusIfNeeded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"_proof\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawalsData\",\"type\":\"bytes\"}],\"name\":\"verifyBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"withdrawERC20Guarded\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZkSync is an auto generated Go binding around an Ethereum contract.
type ZkSync struct {
	ZkSyncCaller     // Read-only binding to the contract
	ZkSyncTransactor // Write-only binding to the contract
	ZkSyncFilterer   // Log filterer for contract events
}

// ZkSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkSyncSession struct {
	Contract     *ZkSync           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkSyncCallerSession struct {
	Contract *ZkSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkSyncTransactorSession struct {
	Contract     *ZkSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkSyncRaw struct {
	Contract *ZkSync // Generic contract binding to access the raw methods on
}

// ZkSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkSyncCallerRaw struct {
	Contract *ZkSyncCaller // Generic read-only contract binding to access the raw methods on
}

// ZkSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkSyncTransactorRaw struct {
	Contract *ZkSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkSync creates a new instance of ZkSync, bound to a specific deployed contract.
func NewZkSync(address common.Address, backend bind.ContractBackend) (*ZkSync, error) {
	contract, err := bindZkSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkSync{ZkSyncCaller: ZkSyncCaller{contract: contract}, ZkSyncTransactor: ZkSyncTransactor{contract: contract}, ZkSyncFilterer: ZkSyncFilterer{contract: contract}}, nil
}

// NewZkSyncCaller creates a new read-only instance of ZkSync, bound to a specific deployed contract.
func NewZkSyncCaller(address common.Address, caller bind.ContractCaller) (*ZkSyncCaller, error) {
	contract, err := bindZkSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkSyncCaller{contract: contract}, nil
}

// NewZkSyncTransactor creates a new write-only instance of ZkSync, bound to a specific deployed contract.
func NewZkSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkSyncTransactor, error) {
	contract, err := bindZkSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkSyncTransactor{contract: contract}, nil
}

// NewZkSyncFilterer creates a new log filterer instance of ZkSync, bound to a specific deployed contract.
func NewZkSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkSyncFilterer, error) {
	contract, err := bindZkSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkSyncFilterer{contract: contract}, nil
}

// bindZkSync binds a generic wrapper to an already deployed contract.
func bindZkSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkSyncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkSync *ZkSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkSync.Contract.ZkSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkSync *ZkSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.Contract.ZkSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkSync *ZkSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkSync.Contract.ZkSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkSync *ZkSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkSync *ZkSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkSync *ZkSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkSync.Contract.contract.Transact(opts, method, params...)
}

// EMPTYSTRINGKECCAK is a free data retrieval call binding the contract method 0x21ae6054.
//
// Solidity: function EMPTY_STRING_KECCAK() view returns(bytes32)
func (_ZkSync *ZkSyncCaller) EMPTYSTRINGKECCAK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "EMPTY_STRING_KECCAK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYSTRINGKECCAK is a free data retrieval call binding the contract method 0x21ae6054.
//
// Solidity: function EMPTY_STRING_KECCAK() view returns(bytes32)
func (_ZkSync *ZkSyncSession) EMPTYSTRINGKECCAK() ([32]byte, error) {
	return _ZkSync.Contract.EMPTYSTRINGKECCAK(&_ZkSync.CallOpts)
}

// EMPTYSTRINGKECCAK is a free data retrieval call binding the contract method 0x21ae6054.
//
// Solidity: function EMPTY_STRING_KECCAK() view returns(bytes32)
func (_ZkSync *ZkSyncCallerSession) EMPTYSTRINGKECCAK() ([32]byte, error) {
	return _ZkSync.Contract.EMPTYSTRINGKECCAK(&_ZkSync.CallOpts)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x2f804bd2.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n) returns()
func (_ZkSync *ZkSyncTransactor) CancelOutstandingDepositsForExodusMode(opts *bind.TransactOpts, _n uint64) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "cancelOutstandingDepositsForExodusMode", _n)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x2f804bd2.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n) returns()
func (_ZkSync *ZkSyncSession) CancelOutstandingDepositsForExodusMode(_n uint64) (*types.Transaction, error) {
	return _ZkSync.Contract.CancelOutstandingDepositsForExodusMode(&_ZkSync.TransactOpts, _n)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x2f804bd2.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n) returns()
func (_ZkSync *ZkSyncTransactorSession) CancelOutstandingDepositsForExodusMode(_n uint64) (*types.Transaction, error) {
	return _ZkSync.Contract.CancelOutstandingDepositsForExodusMode(&_ZkSync.TransactOpts, _n)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x4e913cd9.
//
// Solidity: function commitBlock(uint32 _blockNumber, uint32 _feeAccount, bytes32[] _newBlockInfo, bytes _publicData, bytes _ethWitness, uint32[] _ethWitnessSizes) returns()
func (_ZkSync *ZkSyncTransactor) CommitBlock(opts *bind.TransactOpts, _blockNumber uint32, _feeAccount uint32, _newBlockInfo [][32]byte, _publicData []byte, _ethWitness []byte, _ethWitnessSizes []uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "commitBlock", _blockNumber, _feeAccount, _newBlockInfo, _publicData, _ethWitness, _ethWitnessSizes)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x4e913cd9.
//
// Solidity: function commitBlock(uint32 _blockNumber, uint32 _feeAccount, bytes32[] _newBlockInfo, bytes _publicData, bytes _ethWitness, uint32[] _ethWitnessSizes) returns()
func (_ZkSync *ZkSyncSession) CommitBlock(_blockNumber uint32, _feeAccount uint32, _newBlockInfo [][32]byte, _publicData []byte, _ethWitness []byte, _ethWitnessSizes []uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.CommitBlock(&_ZkSync.TransactOpts, _blockNumber, _feeAccount, _newBlockInfo, _publicData, _ethWitness, _ethWitnessSizes)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x4e913cd9.
//
// Solidity: function commitBlock(uint32 _blockNumber, uint32 _feeAccount, bytes32[] _newBlockInfo, bytes _publicData, bytes _ethWitness, uint32[] _ethWitnessSizes) returns()
func (_ZkSync *ZkSyncTransactorSession) CommitBlock(_blockNumber uint32, _feeAccount uint32, _newBlockInfo [][32]byte, _publicData []byte, _ethWitness []byte, _ethWitnessSizes []uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.CommitBlock(&_ZkSync.TransactOpts, _blockNumber, _feeAccount, _newBlockInfo, _publicData, _ethWitness, _ethWitnessSizes)
}

// CompleteWithdrawals is a paid mutator transaction binding the contract method 0x6a387fc9.
//
// Solidity: function completeWithdrawals(uint32 _n) returns()
func (_ZkSync *ZkSyncTransactor) CompleteWithdrawals(opts *bind.TransactOpts, _n uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "completeWithdrawals", _n)
}

// CompleteWithdrawals is a paid mutator transaction binding the contract method 0x6a387fc9.
//
// Solidity: function completeWithdrawals(uint32 _n) returns()
func (_ZkSync *ZkSyncSession) CompleteWithdrawals(_n uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.CompleteWithdrawals(&_ZkSync.TransactOpts, _n)
}

// CompleteWithdrawals is a paid mutator transaction binding the contract method 0x6a387fc9.
//
// Solidity: function completeWithdrawals(uint32 _n) returns()
func (_ZkSync *ZkSyncTransactorSession) CompleteWithdrawals(_n uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.CompleteWithdrawals(&_ZkSync.TransactOpts, _n)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _franklinAddr) returns()
func (_ZkSync *ZkSyncTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "depositERC20", _token, _amount, _franklinAddr)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _franklinAddr) returns()
func (_ZkSync *ZkSyncSession) DepositERC20(_token common.Address, _amount *big.Int, _franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositERC20(&_ZkSync.TransactOpts, _token, _amount, _franklinAddr)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _franklinAddr) returns()
func (_ZkSync *ZkSyncTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositERC20(&_ZkSync.TransactOpts, _token, _amount, _franklinAddr)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _franklinAddr) payable returns()
func (_ZkSync *ZkSyncTransactor) DepositETH(opts *bind.TransactOpts, _franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "depositETH", _franklinAddr)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _franklinAddr) payable returns()
func (_ZkSync *ZkSyncSession) DepositETH(_franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositETH(&_ZkSync.TransactOpts, _franklinAddr)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _franklinAddr) payable returns()
func (_ZkSync *ZkSyncTransactorSession) DepositETH(_franklinAddr common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositETH(&_ZkSync.TransactOpts, _franklinAddr)
}

// Exit is a paid mutator transaction binding the contract method 0xd6973fc6.
//
// Solidity: function exit(uint32 _accountId, uint16 _tokenId, uint128 _amount, uint256[] _proof) returns()
func (_ZkSync *ZkSyncTransactor) Exit(opts *bind.TransactOpts, _accountId uint32, _tokenId uint16, _amount *big.Int, _proof []*big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "exit", _accountId, _tokenId, _amount, _proof)
}

// Exit is a paid mutator transaction binding the contract method 0xd6973fc6.
//
// Solidity: function exit(uint32 _accountId, uint16 _tokenId, uint128 _amount, uint256[] _proof) returns()
func (_ZkSync *ZkSyncSession) Exit(_accountId uint32, _tokenId uint16, _amount *big.Int, _proof []*big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.Exit(&_ZkSync.TransactOpts, _accountId, _tokenId, _amount, _proof)
}

// Exit is a paid mutator transaction binding the contract method 0xd6973fc6.
//
// Solidity: function exit(uint32 _accountId, uint16 _tokenId, uint128 _amount, uint256[] _proof) returns()
func (_ZkSync *ZkSyncTransactorSession) Exit(_accountId uint32, _tokenId uint16, _amount *big.Int, _proof []*big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.Exit(&_ZkSync.TransactOpts, _accountId, _tokenId, _amount, _proof)
}

// FullExit is a paid mutator transaction binding the contract method 0x000000e2.
//
// Solidity: function fullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncTransactor) FullExit(opts *bind.TransactOpts, _accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "fullExit", _accountId, _token)
}

// FullExit is a paid mutator transaction binding the contract method 0x000000e2.
//
// Solidity: function fullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncSession) FullExit(_accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.FullExit(&_ZkSync.TransactOpts, _accountId, _token)
}

// FullExit is a paid mutator transaction binding the contract method 0x000000e2.
//
// Solidity: function fullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncTransactorSession) FullExit(_accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.FullExit(&_ZkSync.TransactOpts, _accountId, _token)
}

// GetNoticePeriod is a paid mutator transaction binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() returns(uint256)
func (_ZkSync *ZkSyncTransactor) GetNoticePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "getNoticePeriod")
}

// GetNoticePeriod is a paid mutator transaction binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() returns(uint256)
func (_ZkSync *ZkSyncSession) GetNoticePeriod() (*types.Transaction, error) {
	return _ZkSync.Contract.GetNoticePeriod(&_ZkSync.TransactOpts)
}

// GetNoticePeriod is a paid mutator transaction binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() returns(uint256)
func (_ZkSync *ZkSyncTransactorSession) GetNoticePeriod() (*types.Transaction, error) {
	return _ZkSync.Contract.GetNoticePeriod(&_ZkSync.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkSync *ZkSyncTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkSync *ZkSyncSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.Initialize(&_ZkSync.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkSync *ZkSyncTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.Initialize(&_ZkSync.TransactOpts, initializationParameters)
}

// IsReadyForUpgrade is a paid mutator transaction binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() returns(bool)
func (_ZkSync *ZkSyncTransactor) IsReadyForUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "isReadyForUpgrade")
}

// IsReadyForUpgrade is a paid mutator transaction binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() returns(bool)
func (_ZkSync *ZkSyncSession) IsReadyForUpgrade() (*types.Transaction, error) {
	return _ZkSync.Contract.IsReadyForUpgrade(&_ZkSync.TransactOpts)
}

// IsReadyForUpgrade is a paid mutator transaction binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() returns(bool)
func (_ZkSync *ZkSyncTransactorSession) IsReadyForUpgrade() (*types.Transaction, error) {
	return _ZkSync.Contract.IsReadyForUpgrade(&_ZkSync.TransactOpts)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xa6289e5a.
//
// Solidity: function revertBlocks(uint32 _maxBlocksToRevert) returns()
func (_ZkSync *ZkSyncTransactor) RevertBlocks(opts *bind.TransactOpts, _maxBlocksToRevert uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "revertBlocks", _maxBlocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xa6289e5a.
//
// Solidity: function revertBlocks(uint32 _maxBlocksToRevert) returns()
func (_ZkSync *ZkSyncSession) RevertBlocks(_maxBlocksToRevert uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.RevertBlocks(&_ZkSync.TransactOpts, _maxBlocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xa6289e5a.
//
// Solidity: function revertBlocks(uint32 _maxBlocksToRevert) returns()
func (_ZkSync *ZkSyncTransactorSession) RevertBlocks(_maxBlocksToRevert uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.RevertBlocks(&_ZkSync.TransactOpts, _maxBlocksToRevert)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkey_hash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncTransactor) SetAuthPubkeyHash(opts *bind.TransactOpts, _pubkey_hash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "setAuthPubkeyHash", _pubkey_hash, _nonce)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkey_hash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncSession) SetAuthPubkeyHash(_pubkey_hash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.SetAuthPubkeyHash(&_ZkSync.TransactOpts, _pubkey_hash, _nonce)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkey_hash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncTransactorSession) SetAuthPubkeyHash(_pubkey_hash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.SetAuthPubkeyHash(&_ZkSync.TransactOpts, _pubkey_hash, _nonce)
}

// TriggerExodusIfNeeded is a paid mutator transaction binding the contract method 0x6b27a044.
//
// Solidity: function triggerExodusIfNeeded() returns(bool)
func (_ZkSync *ZkSyncTransactor) TriggerExodusIfNeeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "triggerExodusIfNeeded")
}

// TriggerExodusIfNeeded is a paid mutator transaction binding the contract method 0x6b27a044.
//
// Solidity: function triggerExodusIfNeeded() returns(bool)
func (_ZkSync *ZkSyncSession) TriggerExodusIfNeeded() (*types.Transaction, error) {
	return _ZkSync.Contract.TriggerExodusIfNeeded(&_ZkSync.TransactOpts)
}

// TriggerExodusIfNeeded is a paid mutator transaction binding the contract method 0x6b27a044.
//
// Solidity: function triggerExodusIfNeeded() returns(bool)
func (_ZkSync *ZkSyncTransactorSession) TriggerExodusIfNeeded() (*types.Transaction, error) {
	return _ZkSync.Contract.TriggerExodusIfNeeded(&_ZkSync.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkSync *ZkSyncTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkSync *ZkSyncSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.Upgrade(&_ZkSync.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkSync *ZkSyncTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.Upgrade(&_ZkSync.TransactOpts, upgradeParameters)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkSync *ZkSyncTransactor) UpgradeCanceled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "upgradeCanceled")
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkSync *ZkSyncSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeCanceled(&_ZkSync.TransactOpts)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkSync *ZkSyncTransactorSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeCanceled(&_ZkSync.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkSync *ZkSyncTransactor) UpgradeFinishes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "upgradeFinishes")
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkSync *ZkSyncSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeFinishes(&_ZkSync.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkSync *ZkSyncTransactorSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeFinishes(&_ZkSync.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkSync *ZkSyncTransactor) UpgradeNoticePeriodStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "upgradeNoticePeriodStarted")
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkSync *ZkSyncSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeNoticePeriodStarted(&_ZkSync.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkSync *ZkSyncTransactorSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradeNoticePeriodStarted(&_ZkSync.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkSync *ZkSyncTransactor) UpgradePreparationStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "upgradePreparationStarted")
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkSync *ZkSyncSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradePreparationStarted(&_ZkSync.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkSync *ZkSyncTransactorSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZkSync.Contract.UpgradePreparationStarted(&_ZkSync.TransactOpts)
}

// VerifyBlock is a paid mutator transaction binding the contract method 0x0231c02c.
//
// Solidity: function verifyBlock(uint32 _blockNumber, uint256[] _proof, bytes _withdrawalsData) returns()
func (_ZkSync *ZkSyncTransactor) VerifyBlock(opts *bind.TransactOpts, _blockNumber uint32, _proof []*big.Int, _withdrawalsData []byte) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "verifyBlock", _blockNumber, _proof, _withdrawalsData)
}

// VerifyBlock is a paid mutator transaction binding the contract method 0x0231c02c.
//
// Solidity: function verifyBlock(uint32 _blockNumber, uint256[] _proof, bytes _withdrawalsData) returns()
func (_ZkSync *ZkSyncSession) VerifyBlock(_blockNumber uint32, _proof []*big.Int, _withdrawalsData []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.VerifyBlock(&_ZkSync.TransactOpts, _blockNumber, _proof, _withdrawalsData)
}

// VerifyBlock is a paid mutator transaction binding the contract method 0x0231c02c.
//
// Solidity: function verifyBlock(uint32 _blockNumber, uint256[] _proof, bytes _withdrawalsData) returns()
func (_ZkSync *ZkSyncTransactorSession) VerifyBlock(_blockNumber uint32, _proof []*big.Int, _withdrawalsData []byte) (*types.Transaction, error) {
	return _ZkSync.Contract.VerifyBlock(&_ZkSync.TransactOpts, _blockNumber, _proof, _withdrawalsData)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xc94c5b7c.
//
// Solidity: function withdrawERC20(address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "withdrawERC20", _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xc94c5b7c.
//
// Solidity: function withdrawERC20(address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawERC20(&_ZkSync.TransactOpts, _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xc94c5b7c.
//
// Solidity: function withdrawERC20(address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawERC20(&_ZkSync.TransactOpts, _token, _amount)
}

// WithdrawERC20Guarded is a paid mutator transaction binding the contract method 0x9a83400d.
//
// Solidity: function withdrawERC20Guarded(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128)
func (_ZkSync *ZkSyncTransactor) WithdrawERC20Guarded(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "withdrawERC20Guarded", _token, _to, _amount, _maxAmount)
}

// WithdrawERC20Guarded is a paid mutator transaction binding the contract method 0x9a83400d.
//
// Solidity: function withdrawERC20Guarded(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128)
func (_ZkSync *ZkSyncSession) WithdrawERC20Guarded(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawERC20Guarded(&_ZkSync.TransactOpts, _token, _to, _amount, _maxAmount)
}

// WithdrawERC20Guarded is a paid mutator transaction binding the contract method 0x9a83400d.
//
// Solidity: function withdrawERC20Guarded(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128)
func (_ZkSync *ZkSyncTransactorSession) WithdrawERC20Guarded(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawERC20Guarded(&_ZkSync.TransactOpts, _token, _to, _amount, _maxAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc488a09c.
//
// Solidity: function withdrawETH(uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactor) WithdrawETH(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "withdrawETH", _amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc488a09c.
//
// Solidity: function withdrawETH(uint128 _amount) returns()
func (_ZkSync *ZkSyncSession) WithdrawETH(_amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawETH(&_ZkSync.TransactOpts, _amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc488a09c.
//
// Solidity: function withdrawETH(uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactorSession) WithdrawETH(_amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawETH(&_ZkSync.TransactOpts, _amount)
}
