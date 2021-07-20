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

// StorageStoredBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type StorageStoredBlockInfo struct {
	BlockNumber                  uint32
	PriorityOperations           uint64
	PendingOnchainOperationsHash [32]byte
	Timestamp                    *big.Int
	StateHash                    [32]byte
	Commitment                   [32]byte
}

// ZkSyncCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkSyncCommitBlockInfo struct {
	NewStateHash      [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	OnchainOperations []ZkSyncOnchainOperationData
	BlockNumber       uint32
	FeeAccount        uint32
}

// ZkSyncExecuteBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkSyncExecuteBlockInfo struct {
	StoredBlock              StorageStoredBlockInfo
	PendingOnchainOpsPubdata [][]byte
}

// ZkSyncOnchainOperationData is an auto generated low-level Go binding around an user-defined struct.
type ZkSyncOnchainOperationData struct {
	EthWitness       []byte
	PublicDataOffset uint32
}

// ZkSyncProofInput is an auto generated low-level Go binding around an user-defined struct.
type ZkSyncProofInput struct {
	RecursiveInput []*big.Int
	Proof          []*big.Int
	Commitments    []*big.Int
	VkIndexes      []uint8
	SubproofsLimbs [16]*big.Int
}

// ZkSyncABI is the input ABI used to generate the binding from.
const ZkSyncABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkSyncBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ExodusMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fact\",\"type\":\"bytes\"}],\"name\":\"FactAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkSyncBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumOperations.OpType\",\"name\":\"opType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"WithdrawalNFT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"_transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateExodusMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"authFacts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_n\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_depositsPubdata\",\"type\":\"bytes[]\"}],\"name\":\"cancelOutstandingDepositsForExodusMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"ethWitness\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"publicDataOffset\",\"type\":\"uint32\"}],\"internalType\":\"structZkSync.OnchainOperationData[]\",\"name\":\"onchainOperations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeAccount\",\"type\":\"uint32\"}],\"internalType\":\"structZkSync.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"address\",\"name\":\"_zkSyncAddress\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zkSyncAddress\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"storedBlock\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubdata\",\"type\":\"bytes[]\"}],\"internalType\":\"structZkSync.ExecuteBlockInfo[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"}],\"name\":\"executeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exodusMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_committedBlocks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"recursiveInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"proof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"commitments\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"vkIndexes\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[16]\",\"name\":\"subproofsLimbs\",\"type\":\"uint256[16]\"}],\"internalType\":\"structZkSync.ProofInput\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"proveBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_tokenId\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkeyHash\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"setAuthPubkeyHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksExecuted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_tokenId\",\"type\":\"uint32\"}],\"name\":\"withdrawPendingNFTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AuthFacts is a free data retrieval call binding the contract method 0x8ae20dc9.
//
// Solidity: function authFacts(address , uint32 ) view returns(bytes32)
func (_ZkSync *ZkSyncCaller) AuthFacts(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "authFacts", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AuthFacts is a free data retrieval call binding the contract method 0x8ae20dc9.
//
// Solidity: function authFacts(address , uint32 ) view returns(bytes32)
func (_ZkSync *ZkSyncSession) AuthFacts(arg0 common.Address, arg1 uint32) ([32]byte, error) {
	return _ZkSync.Contract.AuthFacts(&_ZkSync.CallOpts, arg0, arg1)
}

// AuthFacts is a free data retrieval call binding the contract method 0x8ae20dc9.
//
// Solidity: function authFacts(address , uint32 ) view returns(bytes32)
func (_ZkSync *ZkSyncCallerSession) AuthFacts(arg0 common.Address, arg1 uint32) ([32]byte, error) {
	return _ZkSync.Contract.AuthFacts(&_ZkSync.CallOpts, arg0, arg1)
}

// ExodusMode is a free data retrieval call binding the contract method 0x264c0912.
//
// Solidity: function exodusMode() view returns(bool)
func (_ZkSync *ZkSyncCaller) ExodusMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "exodusMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExodusMode is a free data retrieval call binding the contract method 0x264c0912.
//
// Solidity: function exodusMode() view returns(bool)
func (_ZkSync *ZkSyncSession) ExodusMode() (bool, error) {
	return _ZkSync.Contract.ExodusMode(&_ZkSync.CallOpts)
}

// ExodusMode is a free data retrieval call binding the contract method 0x264c0912.
//
// Solidity: function exodusMode() view returns(bool)
func (_ZkSync *ZkSyncCallerSession) ExodusMode() (bool, error) {
	return _ZkSync.Contract.ExodusMode(&_ZkSync.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkSync *ZkSyncCaller) GetNoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "getNoticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkSync *ZkSyncSession) GetNoticePeriod() (*big.Int, error) {
	return _ZkSync.Contract.GetNoticePeriod(&_ZkSync.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkSync *ZkSyncCallerSession) GetNoticePeriod() (*big.Int, error) {
	return _ZkSync.Contract.GetNoticePeriod(&_ZkSync.CallOpts)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _token) view returns(uint128)
func (_ZkSync *ZkSyncCaller) GetPendingBalance(opts *bind.CallOpts, _address common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "getPendingBalance", _address, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _token) view returns(uint128)
func (_ZkSync *ZkSyncSession) GetPendingBalance(_address common.Address, _token common.Address) (*big.Int, error) {
	return _ZkSync.Contract.GetPendingBalance(&_ZkSync.CallOpts, _address, _token)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _token) view returns(uint128)
func (_ZkSync *ZkSyncCallerSession) GetPendingBalance(_address common.Address, _token common.Address) (*big.Int, error) {
	return _ZkSync.Contract.GetPendingBalance(&_ZkSync.CallOpts, _address, _token)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkSync *ZkSyncCaller) IsReadyForUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "isReadyForUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkSync *ZkSyncSession) IsReadyForUpgrade() (bool, error) {
	return _ZkSync.Contract.IsReadyForUpgrade(&_ZkSync.CallOpts)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkSync *ZkSyncCallerSession) IsReadyForUpgrade() (bool, error) {
	return _ZkSync.Contract.IsReadyForUpgrade(&_ZkSync.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkSync *ZkSyncCaller) TotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "totalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkSync *ZkSyncSession) TotalBlocksCommitted() (uint32, error) {
	return _ZkSync.Contract.TotalBlocksCommitted(&_ZkSync.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkSync *ZkSyncCallerSession) TotalBlocksCommitted() (uint32, error) {
	return _ZkSync.Contract.TotalBlocksCommitted(&_ZkSync.CallOpts)
}

// TotalBlocksExecuted is a free data retrieval call binding the contract method 0xf2235487.
//
// Solidity: function totalBlocksExecuted() view returns(uint32)
func (_ZkSync *ZkSyncCaller) TotalBlocksExecuted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkSync.contract.Call(opts, &out, "totalBlocksExecuted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksExecuted is a free data retrieval call binding the contract method 0xf2235487.
//
// Solidity: function totalBlocksExecuted() view returns(uint32)
func (_ZkSync *ZkSyncSession) TotalBlocksExecuted() (uint32, error) {
	return _ZkSync.Contract.TotalBlocksExecuted(&_ZkSync.CallOpts)
}

// TotalBlocksExecuted is a free data retrieval call binding the contract method 0xf2235487.
//
// Solidity: function totalBlocksExecuted() view returns(uint32)
func (_ZkSync *ZkSyncCallerSession) TotalBlocksExecuted() (uint32, error) {
	return _ZkSync.Contract.TotalBlocksExecuted(&_ZkSync.CallOpts)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkSync *ZkSyncTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "_transferERC20", _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkSync *ZkSyncSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.TransferERC20(&_ZkSync.TransactOpts, _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkSync *ZkSyncTransactorSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.TransferERC20(&_ZkSync.TransactOpts, _token, _to, _amount, _maxAmount)
}

// ActivateExodusMode is a paid mutator transaction binding the contract method 0xa7e7aacd.
//
// Solidity: function activateExodusMode() returns(bool)
func (_ZkSync *ZkSyncTransactor) ActivateExodusMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "activateExodusMode")
}

// ActivateExodusMode is a paid mutator transaction binding the contract method 0xa7e7aacd.
//
// Solidity: function activateExodusMode() returns(bool)
func (_ZkSync *ZkSyncSession) ActivateExodusMode() (*types.Transaction, error) {
	return _ZkSync.Contract.ActivateExodusMode(&_ZkSync.TransactOpts)
}

// ActivateExodusMode is a paid mutator transaction binding the contract method 0xa7e7aacd.
//
// Solidity: function activateExodusMode() returns(bool)
func (_ZkSync *ZkSyncTransactorSession) ActivateExodusMode() (*types.Transaction, error) {
	return _ZkSync.Contract.ActivateExodusMode(&_ZkSync.TransactOpts)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubdata) returns()
func (_ZkSync *ZkSyncTransactor) CancelOutstandingDepositsForExodusMode(opts *bind.TransactOpts, _n uint64, _depositsPubdata [][]byte) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "cancelOutstandingDepositsForExodusMode", _n, _depositsPubdata)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubdata) returns()
func (_ZkSync *ZkSyncSession) CancelOutstandingDepositsForExodusMode(_n uint64, _depositsPubdata [][]byte) (*types.Transaction, error) {
	return _ZkSync.Contract.CancelOutstandingDepositsForExodusMode(&_ZkSync.TransactOpts, _n, _depositsPubdata)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubdata) returns()
func (_ZkSync *ZkSyncTransactorSession) CancelOutstandingDepositsForExodusMode(_n uint64, _depositsPubdata [][]byte) (*types.Transaction, error) {
	return _ZkSync.Contract.CancelOutstandingDepositsForExodusMode(&_ZkSync.TransactOpts, _n, _depositsPubdata)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x45269298.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint32)[] _newBlocksData) returns()
func (_ZkSync *ZkSyncTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkSyncCommitBlockInfo) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x45269298.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint32)[] _newBlocksData) returns()
func (_ZkSync *ZkSyncSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkSyncCommitBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.CommitBlocks(&_ZkSync.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x45269298.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint32)[] _newBlocksData) returns()
func (_ZkSync *ZkSyncTransactorSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkSyncCommitBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.CommitBlocks(&_ZkSync.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _zkSyncAddress) returns()
func (_ZkSync *ZkSyncTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "depositERC20", _token, _amount, _zkSyncAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _zkSyncAddress) returns()
func (_ZkSync *ZkSyncSession) DepositERC20(_token common.Address, _amount *big.Int, _zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositERC20(&_ZkSync.TransactOpts, _token, _amount, _zkSyncAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xe17376b5.
//
// Solidity: function depositERC20(address _token, uint104 _amount, address _zkSyncAddress) returns()
func (_ZkSync *ZkSyncTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositERC20(&_ZkSync.TransactOpts, _token, _amount, _zkSyncAddress)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _zkSyncAddress) payable returns()
func (_ZkSync *ZkSyncTransactor) DepositETH(opts *bind.TransactOpts, _zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "depositETH", _zkSyncAddress)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _zkSyncAddress) payable returns()
func (_ZkSync *ZkSyncSession) DepositETH(_zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositETH(&_ZkSync.TransactOpts, _zkSyncAddress)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _zkSyncAddress) payable returns()
func (_ZkSync *ZkSyncTransactorSession) DepositETH(_zkSyncAddress common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.DepositETH(&_ZkSync.TransactOpts, _zkSyncAddress)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xb0705b42.
//
// Solidity: function executeBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocksData) returns()
func (_ZkSync *ZkSyncTransactor) ExecuteBlocks(opts *bind.TransactOpts, _blocksData []ZkSyncExecuteBlockInfo) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "executeBlocks", _blocksData)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xb0705b42.
//
// Solidity: function executeBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocksData) returns()
func (_ZkSync *ZkSyncSession) ExecuteBlocks(_blocksData []ZkSyncExecuteBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.ExecuteBlocks(&_ZkSync.TransactOpts, _blocksData)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xb0705b42.
//
// Solidity: function executeBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocksData) returns()
func (_ZkSync *ZkSyncTransactorSession) ExecuteBlocks(_blocksData []ZkSyncExecuteBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.ExecuteBlocks(&_ZkSync.TransactOpts, _blocksData)
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

// ProveBlocks is a paid mutator transaction binding the contract method 0x83981808.
//
// Solidity: function proveBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _committedBlocks, (uint256[],uint256[],uint256[],uint8[],uint256[16]) _proof) returns()
func (_ZkSync *ZkSyncTransactor) ProveBlocks(opts *bind.TransactOpts, _committedBlocks []StorageStoredBlockInfo, _proof ZkSyncProofInput) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "proveBlocks", _committedBlocks, _proof)
}

// ProveBlocks is a paid mutator transaction binding the contract method 0x83981808.
//
// Solidity: function proveBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _committedBlocks, (uint256[],uint256[],uint256[],uint8[],uint256[16]) _proof) returns()
func (_ZkSync *ZkSyncSession) ProveBlocks(_committedBlocks []StorageStoredBlockInfo, _proof ZkSyncProofInput) (*types.Transaction, error) {
	return _ZkSync.Contract.ProveBlocks(&_ZkSync.TransactOpts, _committedBlocks, _proof)
}

// ProveBlocks is a paid mutator transaction binding the contract method 0x83981808.
//
// Solidity: function proveBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _committedBlocks, (uint256[],uint256[],uint256[],uint8[],uint256[16]) _proof) returns()
func (_ZkSync *ZkSyncTransactorSession) ProveBlocks(_committedBlocks []StorageStoredBlockInfo, _proof ZkSyncProofInput) (*types.Transaction, error) {
	return _ZkSync.Contract.ProveBlocks(&_ZkSync.TransactOpts, _committedBlocks, _proof)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncTransactor) RequestFullExit(opts *bind.TransactOpts, _accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "requestFullExit", _accountId, _token)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncSession) RequestFullExit(_accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.RequestFullExit(&_ZkSync.TransactOpts, _accountId, _token)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountId, address _token) returns()
func (_ZkSync *ZkSyncTransactorSession) RequestFullExit(_accountId uint32, _token common.Address) (*types.Transaction, error) {
	return _ZkSync.Contract.RequestFullExit(&_ZkSync.TransactOpts, _accountId, _token)
}

// RequestFullExitNFT is a paid mutator transaction binding the contract method 0x13d9787b.
//
// Solidity: function requestFullExitNFT(uint32 _accountId, uint32 _tokenId) returns()
func (_ZkSync *ZkSyncTransactor) RequestFullExitNFT(opts *bind.TransactOpts, _accountId uint32, _tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "requestFullExitNFT", _accountId, _tokenId)
}

// RequestFullExitNFT is a paid mutator transaction binding the contract method 0x13d9787b.
//
// Solidity: function requestFullExitNFT(uint32 _accountId, uint32 _tokenId) returns()
func (_ZkSync *ZkSyncSession) RequestFullExitNFT(_accountId uint32, _tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.RequestFullExitNFT(&_ZkSync.TransactOpts, _accountId, _tokenId)
}

// RequestFullExitNFT is a paid mutator transaction binding the contract method 0x13d9787b.
//
// Solidity: function requestFullExitNFT(uint32 _accountId, uint32 _tokenId) returns()
func (_ZkSync *ZkSyncTransactorSession) RequestFullExitNFT(_accountId uint32, _tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.RequestFullExitNFT(&_ZkSync.TransactOpts, _accountId, _tokenId)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkSync *ZkSyncTransactor) RevertBlocks(opts *bind.TransactOpts, _blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "revertBlocks", _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkSync *ZkSyncSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.RevertBlocks(&_ZkSync.TransactOpts, _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkSync *ZkSyncTransactorSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkSync.Contract.RevertBlocks(&_ZkSync.TransactOpts, _blocksToRevert)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkeyHash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncTransactor) SetAuthPubkeyHash(opts *bind.TransactOpts, _pubkeyHash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "setAuthPubkeyHash", _pubkeyHash, _nonce)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkeyHash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncSession) SetAuthPubkeyHash(_pubkeyHash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.SetAuthPubkeyHash(&_ZkSync.TransactOpts, _pubkeyHash, _nonce)
}

// SetAuthPubkeyHash is a paid mutator transaction binding the contract method 0x595a5ebc.
//
// Solidity: function setAuthPubkeyHash(bytes _pubkeyHash, uint32 _nonce) returns()
func (_ZkSync *ZkSyncTransactorSession) SetAuthPubkeyHash(_pubkeyHash []byte, _nonce uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.SetAuthPubkeyHash(&_ZkSync.TransactOpts, _pubkeyHash, _nonce)
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

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactor) WithdrawPendingBalance(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "withdrawPendingBalance", _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawPendingBalance(&_ZkSync.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkSync *ZkSyncTransactorSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawPendingBalance(&_ZkSync.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x505a7573.
//
// Solidity: function withdrawPendingNFTBalance(uint32 _tokenId) returns()
func (_ZkSync *ZkSyncTransactor) WithdrawPendingNFTBalance(opts *bind.TransactOpts, _tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.contract.Transact(opts, "withdrawPendingNFTBalance", _tokenId)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x505a7573.
//
// Solidity: function withdrawPendingNFTBalance(uint32 _tokenId) returns()
func (_ZkSync *ZkSyncSession) WithdrawPendingNFTBalance(_tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawPendingNFTBalance(&_ZkSync.TransactOpts, _tokenId)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x505a7573.
//
// Solidity: function withdrawPendingNFTBalance(uint32 _tokenId) returns()
func (_ZkSync *ZkSyncTransactorSession) WithdrawPendingNFTBalance(_tokenId uint32) (*types.Transaction, error) {
	return _ZkSync.Contract.WithdrawPendingNFTBalance(&_ZkSync.TransactOpts, _tokenId)
}

// ZkSyncBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the ZkSync contract.
type ZkSyncBlockCommitIterator struct {
	Event *ZkSyncBlockCommit // Event containing the contract specifics and raw log

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
func (it *ZkSyncBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncBlockCommit)
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
		it.Event = new(ZkSyncBlockCommit)
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
func (it *ZkSyncBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncBlockCommit represents a BlockCommit event raised by the ZkSync contract.
type ZkSyncBlockCommit struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) FilterBlockCommit(opts *bind.FilterOpts, blockNumber []uint32) (*ZkSyncBlockCommitIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "BlockCommit", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncBlockCommitIterator{contract: _ZkSync.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *ZkSyncBlockCommit, blockNumber []uint32) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "BlockCommit", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncBlockCommit)
				if err := _ZkSync.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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

// ParseBlockCommit is a log parse operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) ParseBlockCommit(log types.Log) (*ZkSyncBlockCommit, error) {
	event := new(ZkSyncBlockCommit)
	if err := _ZkSync.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncBlockVerificationIterator is returned from FilterBlockVerification and is used to iterate over the raw logs and unpacked data for BlockVerification events raised by the ZkSync contract.
type ZkSyncBlockVerificationIterator struct {
	Event *ZkSyncBlockVerification // Event containing the contract specifics and raw log

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
func (it *ZkSyncBlockVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncBlockVerification)
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
		it.Event = new(ZkSyncBlockVerification)
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
func (it *ZkSyncBlockVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncBlockVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncBlockVerification represents a BlockVerification event raised by the ZkSync contract.
type ZkSyncBlockVerification struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockVerification is a free log retrieval operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) FilterBlockVerification(opts *bind.FilterOpts, blockNumber []uint32) (*ZkSyncBlockVerificationIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "BlockVerification", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncBlockVerificationIterator{contract: _ZkSync.contract, event: "BlockVerification", logs: logs, sub: sub}, nil
}

// WatchBlockVerification is a free log subscription operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) WatchBlockVerification(opts *bind.WatchOpts, sink chan<- *ZkSyncBlockVerification, blockNumber []uint32) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "BlockVerification", blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncBlockVerification)
				if err := _ZkSync.contract.UnpackLog(event, "BlockVerification", log); err != nil {
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

// ParseBlockVerification is a log parse operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 indexed blockNumber)
func (_ZkSync *ZkSyncFilterer) ParseBlockVerification(log types.Log) (*ZkSyncBlockVerification, error) {
	event := new(ZkSyncBlockVerification)
	if err := _ZkSync.contract.UnpackLog(event, "BlockVerification", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the ZkSync contract.
type ZkSyncBlocksRevertIterator struct {
	Event *ZkSyncBlocksRevert // Event containing the contract specifics and raw log

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
func (it *ZkSyncBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncBlocksRevert)
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
		it.Event = new(ZkSyncBlocksRevert)
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
func (it *ZkSyncBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncBlocksRevert represents a BlocksRevert event raised by the ZkSync contract.
type ZkSyncBlocksRevert struct {
	TotalBlocksVerified  uint32
	TotalBlocksCommitted uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkSync *ZkSyncFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*ZkSyncBlocksRevertIterator, error) {

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &ZkSyncBlocksRevertIterator{contract: _ZkSync.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkSync *ZkSyncFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *ZkSyncBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncBlocksRevert)
				if err := _ZkSync.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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

// ParseBlocksRevert is a log parse operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkSync *ZkSyncFilterer) ParseBlocksRevert(log types.Log) (*ZkSyncBlocksRevert, error) {
	event := new(ZkSyncBlocksRevert)
	if err := _ZkSync.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZkSync contract.
type ZkSyncDepositIterator struct {
	Event *ZkSyncDeposit // Event containing the contract specifics and raw log

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
func (it *ZkSyncDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncDeposit)
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
		it.Event = new(ZkSyncDeposit)
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
func (it *ZkSyncDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncDeposit represents a Deposit event raised by the ZkSync contract.
type ZkSyncDeposit struct {
	TokenId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8f5f51448394699ad6a3b80cdadf4ec68c5d724c8c3fea09bea55b3c2d0e2dd0.
//
// Solidity: event Deposit(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) FilterDeposit(opts *bind.FilterOpts, tokenId []uint16) (*ZkSyncDepositIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "Deposit", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncDepositIterator{contract: _ZkSync.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8f5f51448394699ad6a3b80cdadf4ec68c5d724c8c3fea09bea55b3c2d0e2dd0.
//
// Solidity: event Deposit(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZkSyncDeposit, tokenId []uint16) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "Deposit", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncDeposit)
				if err := _ZkSync.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x8f5f51448394699ad6a3b80cdadf4ec68c5d724c8c3fea09bea55b3c2d0e2dd0.
//
// Solidity: event Deposit(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) ParseDeposit(log types.Log) (*ZkSyncDeposit, error) {
	event := new(ZkSyncDeposit)
	if err := _ZkSync.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncDepositCommitIterator is returned from FilterDepositCommit and is used to iterate over the raw logs and unpacked data for DepositCommit events raised by the ZkSync contract.
type ZkSyncDepositCommitIterator struct {
	Event *ZkSyncDepositCommit // Event containing the contract specifics and raw log

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
func (it *ZkSyncDepositCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncDepositCommit)
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
		it.Event = new(ZkSyncDepositCommit)
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
func (it *ZkSyncDepositCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncDepositCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncDepositCommit represents a DepositCommit event raised by the ZkSync contract.
type ZkSyncDepositCommit struct {
	ZkSyncBlockId uint32
	AccountId     uint32
	Owner         common.Address
	TokenId       uint16
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositCommit is a free log retrieval operation binding the contract event 0xc4e73a5b67a0594d06ea2b5c311c2aa44aa340dd4dd9ec5a1a718dc391b64470.
//
// Solidity: event DepositCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) FilterDepositCommit(opts *bind.FilterOpts, zkSyncBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZkSyncDepositCommitIterator, error) {

	var zkSyncBlockIdRule []interface{}
	for _, zkSyncBlockIdItem := range zkSyncBlockId {
		zkSyncBlockIdRule = append(zkSyncBlockIdRule, zkSyncBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "DepositCommit", zkSyncBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncDepositCommitIterator{contract: _ZkSync.contract, event: "DepositCommit", logs: logs, sub: sub}, nil
}

// WatchDepositCommit is a free log subscription operation binding the contract event 0xc4e73a5b67a0594d06ea2b5c311c2aa44aa340dd4dd9ec5a1a718dc391b64470.
//
// Solidity: event DepositCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) WatchDepositCommit(opts *bind.WatchOpts, sink chan<- *ZkSyncDepositCommit, zkSyncBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zkSyncBlockIdRule []interface{}
	for _, zkSyncBlockIdItem := range zkSyncBlockId {
		zkSyncBlockIdRule = append(zkSyncBlockIdRule, zkSyncBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "DepositCommit", zkSyncBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncDepositCommit)
				if err := _ZkSync.contract.UnpackLog(event, "DepositCommit", log); err != nil {
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

// ParseDepositCommit is a log parse operation binding the contract event 0xc4e73a5b67a0594d06ea2b5c311c2aa44aa340dd4dd9ec5a1a718dc391b64470.
//
// Solidity: event DepositCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) ParseDepositCommit(log types.Log) (*ZkSyncDepositCommit, error) {
	event := new(ZkSyncDepositCommit)
	if err := _ZkSync.contract.UnpackLog(event, "DepositCommit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncExodusModeIterator is returned from FilterExodusMode and is used to iterate over the raw logs and unpacked data for ExodusMode events raised by the ZkSync contract.
type ZkSyncExodusModeIterator struct {
	Event *ZkSyncExodusMode // Event containing the contract specifics and raw log

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
func (it *ZkSyncExodusModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncExodusMode)
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
		it.Event = new(ZkSyncExodusMode)
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
func (it *ZkSyncExodusModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncExodusModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncExodusMode represents a ExodusMode event raised by the ZkSync contract.
type ZkSyncExodusMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExodusMode is a free log retrieval operation binding the contract event 0xc71028c67eb0ef128ea270a59a674629e767d51c1af44ed6753fd2fad2c7b677.
//
// Solidity: event ExodusMode()
func (_ZkSync *ZkSyncFilterer) FilterExodusMode(opts *bind.FilterOpts) (*ZkSyncExodusModeIterator, error) {

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "ExodusMode")
	if err != nil {
		return nil, err
	}
	return &ZkSyncExodusModeIterator{contract: _ZkSync.contract, event: "ExodusMode", logs: logs, sub: sub}, nil
}

// WatchExodusMode is a free log subscription operation binding the contract event 0xc71028c67eb0ef128ea270a59a674629e767d51c1af44ed6753fd2fad2c7b677.
//
// Solidity: event ExodusMode()
func (_ZkSync *ZkSyncFilterer) WatchExodusMode(opts *bind.WatchOpts, sink chan<- *ZkSyncExodusMode) (event.Subscription, error) {

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "ExodusMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncExodusMode)
				if err := _ZkSync.contract.UnpackLog(event, "ExodusMode", log); err != nil {
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

// ParseExodusMode is a log parse operation binding the contract event 0xc71028c67eb0ef128ea270a59a674629e767d51c1af44ed6753fd2fad2c7b677.
//
// Solidity: event ExodusMode()
func (_ZkSync *ZkSyncFilterer) ParseExodusMode(log types.Log) (*ZkSyncExodusMode, error) {
	event := new(ZkSyncExodusMode)
	if err := _ZkSync.contract.UnpackLog(event, "ExodusMode", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncFactAuthIterator is returned from FilterFactAuth and is used to iterate over the raw logs and unpacked data for FactAuth events raised by the ZkSync contract.
type ZkSyncFactAuthIterator struct {
	Event *ZkSyncFactAuth // Event containing the contract specifics and raw log

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
func (it *ZkSyncFactAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncFactAuth)
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
		it.Event = new(ZkSyncFactAuth)
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
func (it *ZkSyncFactAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncFactAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncFactAuth represents a FactAuth event raised by the ZkSync contract.
type ZkSyncFactAuth struct {
	Sender common.Address
	Nonce  uint32
	Fact   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFactAuth is a free log retrieval operation binding the contract event 0x9ea39b45a0cc96a2139996ec8dd30326216111249750781e563ae27c31ae8766.
//
// Solidity: event FactAuth(address indexed sender, uint32 nonce, bytes fact)
func (_ZkSync *ZkSyncFilterer) FilterFactAuth(opts *bind.FilterOpts, sender []common.Address) (*ZkSyncFactAuthIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "FactAuth", senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncFactAuthIterator{contract: _ZkSync.contract, event: "FactAuth", logs: logs, sub: sub}, nil
}

// WatchFactAuth is a free log subscription operation binding the contract event 0x9ea39b45a0cc96a2139996ec8dd30326216111249750781e563ae27c31ae8766.
//
// Solidity: event FactAuth(address indexed sender, uint32 nonce, bytes fact)
func (_ZkSync *ZkSyncFilterer) WatchFactAuth(opts *bind.WatchOpts, sink chan<- *ZkSyncFactAuth, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "FactAuth", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncFactAuth)
				if err := _ZkSync.contract.UnpackLog(event, "FactAuth", log); err != nil {
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

// ParseFactAuth is a log parse operation binding the contract event 0x9ea39b45a0cc96a2139996ec8dd30326216111249750781e563ae27c31ae8766.
//
// Solidity: event FactAuth(address indexed sender, uint32 nonce, bytes fact)
func (_ZkSync *ZkSyncFilterer) ParseFactAuth(log types.Log) (*ZkSyncFactAuth, error) {
	event := new(ZkSyncFactAuth)
	if err := _ZkSync.contract.UnpackLog(event, "FactAuth", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncFullExitCommitIterator is returned from FilterFullExitCommit and is used to iterate over the raw logs and unpacked data for FullExitCommit events raised by the ZkSync contract.
type ZkSyncFullExitCommitIterator struct {
	Event *ZkSyncFullExitCommit // Event containing the contract specifics and raw log

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
func (it *ZkSyncFullExitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncFullExitCommit)
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
		it.Event = new(ZkSyncFullExitCommit)
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
func (it *ZkSyncFullExitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncFullExitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncFullExitCommit represents a FullExitCommit event raised by the ZkSync contract.
type ZkSyncFullExitCommit struct {
	ZkSyncBlockId uint32
	AccountId     uint32
	Owner         common.Address
	TokenId       uint16
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFullExitCommit is a free log retrieval operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) FilterFullExitCommit(opts *bind.FilterOpts, zkSyncBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZkSyncFullExitCommitIterator, error) {

	var zkSyncBlockIdRule []interface{}
	for _, zkSyncBlockIdItem := range zkSyncBlockId {
		zkSyncBlockIdRule = append(zkSyncBlockIdRule, zkSyncBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "FullExitCommit", zkSyncBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncFullExitCommitIterator{contract: _ZkSync.contract, event: "FullExitCommit", logs: logs, sub: sub}, nil
}

// WatchFullExitCommit is a free log subscription operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) WatchFullExitCommit(opts *bind.WatchOpts, sink chan<- *ZkSyncFullExitCommit, zkSyncBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zkSyncBlockIdRule []interface{}
	for _, zkSyncBlockIdItem := range zkSyncBlockId {
		zkSyncBlockIdRule = append(zkSyncBlockIdRule, zkSyncBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "FullExitCommit", zkSyncBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncFullExitCommit)
				if err := _ZkSync.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
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

// ParseFullExitCommit is a log parse operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkSyncBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) ParseFullExitCommit(log types.Log) (*ZkSyncFullExitCommit, error) {
	event := new(ZkSyncFullExitCommit)
	if err := _ZkSync.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the ZkSync contract.
type ZkSyncNewPriorityRequestIterator struct {
	Event *ZkSyncNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *ZkSyncNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncNewPriorityRequest)
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
		it.Event = new(ZkSyncNewPriorityRequest)
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
func (it *ZkSyncNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncNewPriorityRequest represents a NewPriorityRequest event raised by the ZkSync contract.
type ZkSyncNewPriorityRequest struct {
	Sender          common.Address
	SerialId        uint64
	OpType          uint8
	PubData         []byte
	ExpirationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 opType, bytes pubData, uint256 expirationBlock)
func (_ZkSync *ZkSyncFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*ZkSyncNewPriorityRequestIterator, error) {

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &ZkSyncNewPriorityRequestIterator{contract: _ZkSync.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 opType, bytes pubData, uint256 expirationBlock)
func (_ZkSync *ZkSyncFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *ZkSyncNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncNewPriorityRequest)
				if err := _ZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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

// ParseNewPriorityRequest is a log parse operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 opType, bytes pubData, uint256 expirationBlock)
func (_ZkSync *ZkSyncFilterer) ParseNewPriorityRequest(log types.Log) (*ZkSyncNewPriorityRequest, error) {
	event := new(ZkSyncNewPriorityRequest)
	if err := _ZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZkSync contract.
type ZkSyncWithdrawalIterator struct {
	Event *ZkSyncWithdrawal // Event containing the contract specifics and raw log

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
func (it *ZkSyncWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncWithdrawal)
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
		it.Event = new(ZkSyncWithdrawal)
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
func (it *ZkSyncWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncWithdrawal represents a Withdrawal event raised by the ZkSync contract.
type ZkSyncWithdrawal struct {
	TokenId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) FilterWithdrawal(opts *bind.FilterOpts, tokenId []uint16) (*ZkSyncWithdrawalIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "Withdrawal", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncWithdrawalIterator{contract: _ZkSync.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZkSyncWithdrawal, tokenId []uint16) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "Withdrawal", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncWithdrawal)
				if err := _ZkSync.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 indexed tokenId, uint128 amount)
func (_ZkSync *ZkSyncFilterer) ParseWithdrawal(log types.Log) (*ZkSyncWithdrawal, error) {
	event := new(ZkSyncWithdrawal)
	if err := _ZkSync.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZkSyncWithdrawalNFTIterator is returned from FilterWithdrawalNFT and is used to iterate over the raw logs and unpacked data for WithdrawalNFT events raised by the ZkSync contract.
type ZkSyncWithdrawalNFTIterator struct {
	Event *ZkSyncWithdrawalNFT // Event containing the contract specifics and raw log

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
func (it *ZkSyncWithdrawalNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncWithdrawalNFT)
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
		it.Event = new(ZkSyncWithdrawalNFT)
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
func (it *ZkSyncWithdrawalNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncWithdrawalNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncWithdrawalNFT represents a WithdrawalNFT event raised by the ZkSync contract.
type ZkSyncWithdrawalNFT struct {
	TokenId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalNFT is a free log retrieval operation binding the contract event 0x0b9f3586023bf754b8d962232407f7ac4d90fd19a1c4756c6619927abf067560.
//
// Solidity: event WithdrawalNFT(uint32 indexed tokenId)
func (_ZkSync *ZkSyncFilterer) FilterWithdrawalNFT(opts *bind.FilterOpts, tokenId []uint32) (*ZkSyncWithdrawalNFTIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.FilterLogs(opts, "WithdrawalNFT", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkSyncWithdrawalNFTIterator{contract: _ZkSync.contract, event: "WithdrawalNFT", logs: logs, sub: sub}, nil
}

// WatchWithdrawalNFT is a free log subscription operation binding the contract event 0x0b9f3586023bf754b8d962232407f7ac4d90fd19a1c4756c6619927abf067560.
//
// Solidity: event WithdrawalNFT(uint32 indexed tokenId)
func (_ZkSync *ZkSyncFilterer) WatchWithdrawalNFT(opts *bind.WatchOpts, sink chan<- *ZkSyncWithdrawalNFT, tokenId []uint32) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkSync.contract.WatchLogs(opts, "WithdrawalNFT", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncWithdrawalNFT)
				if err := _ZkSync.contract.UnpackLog(event, "WithdrawalNFT", log); err != nil {
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

// ParseWithdrawalNFT is a log parse operation binding the contract event 0x0b9f3586023bf754b8d962232407f7ac4d90fd19a1c4756c6619927abf067560.
//
// Solidity: event WithdrawalNFT(uint32 indexed tokenId)
func (_ZkSync *ZkSyncFilterer) ParseWithdrawalNFT(log types.Log) (*ZkSyncWithdrawalNFT, error) {
	event := new(ZkSyncWithdrawalNFT)
	if err := _ZkSync.contract.UnpackLog(event, "WithdrawalNFT", log); err != nil {
		return nil, err
	}
	return event, nil
}
