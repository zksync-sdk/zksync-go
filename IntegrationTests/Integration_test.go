package IntegrationTests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/zksync-sdk/zksync-go"
	"math/big"
	"testing"
	"time"
)

var (
	masterPkHex             = "0945012b971f943073f6e066581f513c8cd81660bb5a64306d5d092b8df9dd3f"
	ethNode                 = "https://rinkeby.infura.io/v3/511e4b90f71747658a167e9737ed544a"
	testAmount        int64 = 2000000000000000            // in wei (0.002 ETH)
	testAmountUSDC    int64 = 2000000                     // 2 USDC
	depositAmount     int64 = 1000000000000000            // in wei (0.001 ETH)
	depositAmountUSDC int64 = 1000000                     // 1 USDC
	transferAmount          = big.NewInt(100000000000000) // in wei (0.0001 ETH)
	withdrawAmount          = big.NewInt(100000000000000) // in wei (0.0001 ETH)
	swapAmount              = big.NewInt(500000)          // 0.5 USDC
	txWaitTimeout           = time.Minute * 5
	txCheckInterval         = time.Second * 2
	rinkebyUSDCsc           = "0xeb8f08a975ab53e34d8a0330e0d34de942c95926" // USDC smart-contract address on Rinkeby

	err       error
	ethClient *ethclient.Client
)

func TestFullFlow(t *testing.T) {
	ethClient, err = ethclient.Dial(ethNode)
	require.NoError(t, err)
	require.NotNil(t, ethClient)
	var w1 *zksync.Wallet
	var w2 *zksync.Wallet
	var ep1 zksync.EthProvider

	t.Run("wallet 1", func(t *testing.T) {
		w, zs := newWallet(t, "")
		w1 = w
		ep, err := w.CreateEthereumProvider(ethClient)
		require.NoError(t, err)
		require.NotNil(t, ep)
		ep1 = ep

		//// 0 - fulfill new wallet with some test balance
		fulfillment(t, w.GetAddress())
		// also send some USDC
		fulfillmentUSDC(t, w.GetAddress())

		// 1 - deposit amount to zkSync
		deposit(t, w, ep)
		waitZkAccount(t, w.GetProvider(), w.GetAddress())

		// deposit USDC also
		allTokens, err := w1.GetProvider().GetTokens()
		require.NoError(t, err)
		require.NotNil(t, allTokens)
		usdcToken, err := allTokens.GetToken("USDC")
		require.NoError(t, err)
		require.NotNil(t, usdcToken)
		depositUSDC(t, w1, ep1, usdcToken)

		// 2 - SetAuthPubkeyHash for next ChangePubKeyOnchain
		state, err := w.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		isPkhSet, err := ep.IsOnChainAuthPubkeyHashSet(state.Committed.Nonce)
		require.NoError(t, err)
		require.False(t, isPkhSet)
		tx, err := ep.SetAuthPubkeyHash(zs.GetPublicKeyHash(), state.Committed.Nonce, nil)
		require.NoError(t, err)
		require.NotNil(t, tx)
		waitEthTx(t, context.Background(), tx.Hash(), "setAuthPubkeyHash")

		// 3 - ChangePubKeyOnchain
		is := w.IsSigningKeySet()
		require.False(t, is)
		fee, err := w.GetProvider().GetTransactionFee(zksync.TransactionTypeChangePubKeyOnchain, w.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		state, err = w.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		txHash, err := w.SetSigningKey(fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, true, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w.GetProvider(), txHash, "ChangePubKeyOnchain")
		require.NoError(t, err)
	})

	t.Run("wallet 2", func(t *testing.T) {
		w, _ := newWallet(t, "")
		w2 = w
		ep, err := w.CreateEthereumProvider(ethClient)
		require.NoError(t, err)
		require.NotNil(t, ep)

		//// 0 - fulfill new wallet with some test balance
		fulfillment(t, w.GetAddress())

		// 1 - deposit amount to zkSync
		deposit(t, w, ep)
		waitZkAccount(t, w.GetProvider(), w.GetAddress())

		// 2 - ChangePubKeyECDSA
		state, err := w.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		is := w.IsSigningKeySet()
		require.False(t, is)
		fee, err := w.GetProvider().GetTransactionFee(zksync.TransactionTypeChangePubKeyECDSA, w.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w.SetSigningKey(fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, false, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w.GetProvider(), txHash, "ChangePubKeyECDSA")
		require.NoError(t, err)
	})

	t.Run("transfer", func(t *testing.T) {
		state, err := w1.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		fee, err := w1.GetProvider().GetTransactionFee(zksync.TransactionTypeTransfer, w1.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w1.SyncTransfer(w2.GetAddress(), transferAmount, fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w1.GetProvider(), txHash, "SyncTransfer")
		require.NoError(t, err)
	})

	t.Run("withdraw", func(t *testing.T) {
		state, err := w2.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		fee, err := w2.GetProvider().GetTransactionFee(zksync.TransactionTypeWithdraw, w2.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w2.SyncWithdraw(w2.GetAddress(), withdrawAmount, fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, false, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w2.GetProvider(), txHash, "SyncWithdraw")
		require.NoError(t, err)
	})

	t.Run("mint NFT 1", func(t *testing.T) {
		state, err := w1.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		fee, err := w1.GetProvider().GetTransactionFee(zksync.TransactionTypeMintNFT, w1.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w1.SyncMintNFT(w1.GetAddress(), common.HexToHash("1111"), fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce)
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w1.GetProvider(), txHash, "SyncMintNFT")
		require.NoError(t, err)
	})

	t.Run("mint NFT 2", func(t *testing.T) {
		state, err := w2.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		fee, err := w2.GetProvider().GetTransactionFee(zksync.TransactionTypeMintNFT, w2.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w2.SyncMintNFT(w2.GetAddress(), common.HexToHash("2222"), fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce)
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w2.GetProvider(), txHash, "SyncMintNFT")
		require.NoError(t, err)
	})

	// can't test transfer of just minted NFT because mint tx must be verified (too much time to wait)
	//t.Run("transfer NFT", func(t *testing.T) {
	//	state, err := w1.GetState()
	//	require.NoError(t, err)
	//	require.NotNil(t, state)
	//	require.NotNil(t, state.Committed)
	//	require.NotNil(t, state.Verified)
	//	require.NotNil(t, state.Verified.Nfts)
	//	require.GreaterOrEqual(t, len(state.Verified.Nfts), 1)
	//	var nft *zksync.NFT
	//	for _, nft = range state.Verified.Nfts {
	//		break // get some first NFT
	//	}
	//	fee, err := w1.GetProvider().GetTransactionsBatchFee(
	//		[]zksync.TransactionType{zksync.TransactionTypeTransfer, zksync.TransactionTypeTransfer},
	//		[]common.Address{w1.GetAddress(), w1.GetAddress()}, zksync.CreateETH())
	//	require.NoError(t, err)
	//	require.NotNil(t, fee)
	//
	//	txHashes, err := w1.SyncTransferNFT(w2.GetAddress(), nft, fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, zksync.DefaultTimeRange())
	//	require.NoError(t, err)
	//	require.NotEmpty(t, txHashes)
	//	require.Equal(t, len(txHashes), 2)
	//	err = waitZkTx(w1.GetProvider(), txHashes[0], "SyncTransferNFT")
	//  require.NoError(t, err)
	//})

	t.Run("full exit NFT 1", func(t *testing.T) {
		state, err := w1.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		require.NotNil(t, state.Committed.Nfts)
		require.GreaterOrEqual(t, len(state.Committed.Nfts), 1)
		var nft *zksync.NFT
		for _, nft = range state.Committed.Nfts {
			break // get some first NFT
		}
		tx, err := ep1.FullExitNFT(nft, 0, &zksync.GasOptions{GasLimit: 300000})
		require.NoError(t, err)
		require.NotNil(t, tx)
		waitEthTx(t, context.Background(), tx.Hash(), "FullExitNFT")
	})

	t.Run("withdraw NFT 2", func(t *testing.T) {
		state, err := w2.GetState()
		require.NoError(t, err)
		require.NotNil(t, state)
		require.NotNil(t, state.Committed)
		require.NotNil(t, state.Committed.Nfts)
		require.GreaterOrEqual(t, len(state.Committed.Nfts), 1)
		var nft *zksync.NFT
		for _, nft = range state.Committed.Nfts {
			break // get some first NFT
		}
		fee, err := w2.GetProvider().GetTransactionFee(zksync.TransactionTypeWithdrawNFT, w2.GetAddress(), zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)
		txHash, err := w2.SyncWithdrawNFT(w2.GetAddress(), nft, fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w2.GetProvider(), txHash, "SyncWithdrawNFT")
		require.NoError(t, err)
	})

	t.Run("swap", func(t *testing.T) {
		allTokens, err := w1.GetProvider().GetTokens()
		require.NoError(t, err)
		require.NotNil(t, allTokens)
		usdcToken, err := allTokens.GetToken("USDC")
		require.NoError(t, err)
		require.NotNil(t, usdcToken)

		fee, err := w1.GetProvider().GetTransactionsBatchFee(
			[]zksync.TransactionType{zksync.TransactionTypeSwap},
			[]common.Address{w1.GetAddress()}, zksync.CreateETH())
		require.NoError(t, err)
		require.NotNil(t, fee)

		state1, err := w1.GetState()
		require.NoError(t, err)
		require.NotNil(t, state1)
		require.NotNil(t, state1.Committed)
		o1, err := w1.BuildSignedOrder(w1.GetAddress(), usdcToken, zksync.CreateETH(), []*big.Int{big.NewInt(1), big.NewInt(10)}, swapAmount, state1.Committed.Nonce, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotNil(t, o1)

		state2, err := w2.GetState()
		require.NoError(t, err)
		require.NotNil(t, state2)
		require.NotNil(t, state2.Committed)
		o2, err := w2.BuildSignedOrder(w2.GetAddress(), zksync.CreateETH(), usdcToken, []*big.Int{big.NewInt(10), big.NewInt(1)}, big.NewInt(0).Mul(swapAmount, big.NewInt(10)), state2.Committed.Nonce, zksync.DefaultTimeRange())
		require.NoError(t, err)
		require.NotNil(t, o2)

		txHash, err := w1.SyncSwap(o1, o2, o1.Amount, o2.Amount, fee.GetTxFee(zksync.CreateETH()), state1.Committed.Nonce)
		require.NoError(t, err)
		require.NotEmpty(t, txHash)
		err = waitZkTx(w1.GetProvider(), txHash, "SyncSwap")
		require.NoError(t, err)
	})

	// can be tested due to requirements: Target account exists less than required minimum amount (1 hours)
	//t.Run("forced exit", func(t *testing.T) {
	//	state, err := w1.GetState()
	//	require.NoError(t, err)
	//	require.NotNil(t, state)
	//	require.NotNil(t, state.Committed)
	//	fee, err := w1.GetProvider().GetTransactionFee(zksync.TransactionTypeForcedExit, w1.GetAddress(), zksync.CreateETH())
	//	require.NoError(t, err)
	//	require.NotNil(t, fee)
	//	txHash, err := w1.SyncForcedExit(w1.GetAddress(), fee.GetTxFee(zksync.CreateETH()), state.Committed.Nonce, zksync.DefaultTimeRange())
	//	require.NoError(t, err)
	//	require.NotEmpty(t, txHash)
	//	waitZkTx(t, w1.GetProvider(), txHash, "ForcedExit")
	//})

	t.Run("full exit", func(t *testing.T) {
		tx, err := ep1.FullExit(zksync.CreateETH(), 0, &zksync.GasOptions{GasLimit: 300000})
		require.NoError(t, err)
		require.NotNil(t, tx)
		waitEthTx(t, context.Background(), tx.Hash(), "FullExit")
	})

}

func newWallet(t *testing.T, mnemonic string) (*zksync.Wallet, *zksync.ZkSigner) {
	var err error
	if len(mnemonic) == 0 {
		mnemonic, err = hdwallet.NewMnemonic(160)
		require.NoError(t, err)
		require.NotEmpty(t, mnemonic)
		//mnemonic = "actor feature blade risk rocket behind wide indicate frequent upset session crane tape dentist hundred"
	}
	fmt.Printf("mnemonic - %s\n", mnemonic)

	seed, err := hdwallet.NewSeedFromMnemonic(mnemonic)
	require.NoError(t, err)
	require.NotEmpty(t, seed)

	// create zkSync signer from seed
	zs, err := zksync.NewZkSignerFromSeed(seed)
	require.NoError(t, err)
	require.NotNil(t, zs.GetPublicKeyHash())
	//fmt.Printf("zkSigner pubkey - %+v\n", zs.GetPublicKey())
	//fmt.Printf("zkSigner pkh - %+v\n", zs.GetPublicKeyHash())

	// create ethereum signer from mnemonic
	es, err := zksync.NewEthSignerFromMnemonic(mnemonic)
	require.NoError(t, err)
	require.NotNil(t, es)
	fmt.Printf("ethSigner address - %+v\n", es.GetAddress().String())

	// create zkSync provider for specified chainId
	zp, err := zksync.NewDefaultProviderFor(zksync.ChainIdRinkeby)
	require.NoError(t, err)
	require.NotNil(t, zp)

	// create wallet
	w, err := zksync.NewWallet(es, zs, zp)
	require.NoError(t, err)
	require.NotNil(t, w)
	return w, zs
}

func fulfillment(t *testing.T, toAddress common.Address) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	masterPrivateKey, err := crypto.HexToECDSA(masterPkHex)
	require.NoError(t, err)
	masterPublicKey := masterPrivateKey.Public()
	masterPublicKeyECDSA, ok := masterPublicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	masterAddress := crypto.PubkeyToAddress(*masterPublicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(ctx, masterAddress)
	require.NoError(t, err)

	value := big.NewInt(testAmount)
	gasLimit := uint64(21000) // in units
	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	require.NoError(t, err)

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := ethClient.NetworkID(ctx)
	require.NoError(t, err)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), masterPrivateKey)
	require.NoError(t, err)

	err = ethClient.SendTransaction(ctx, signedTx)
	require.NoError(t, err)

	waitEthTx(t, ctx, signedTx.Hash(), "fulfillment")
}

func fulfillmentUSDC(t *testing.T, toAddress common.Address) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	masterPrivateKey, err := crypto.HexToECDSA(masterPkHex)
	require.NoError(t, err)
	masterPublicKey := masterPrivateKey.Public()
	masterPublicKeyECDSA, ok := masterPublicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	masterAddress := crypto.PubkeyToAddress(*masterPublicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(ctx, masterAddress)
	require.NoError(t, err)

	value := big.NewInt(0) // 0 ETH
	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	require.NoError(t, err)

	transferFnSignature := []byte("transfer(address,uint256)")
	methodID := crypto.Keccak256(transferFnSignature)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	valueUSDC := big.NewInt(testAmountUSDC)
	paddedAmount := common.LeftPadBytes(valueUSDC.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, common.HexToAddress(rinkebyUSDCsc), value, uint64(75000), gasPrice, data)

	chainID, err := ethClient.NetworkID(ctx)
	require.NoError(t, err)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), masterPrivateKey)
	require.NoError(t, err)

	err = ethClient.SendTransaction(ctx, signedTx)
	require.NoError(t, err)

	waitEthTx(t, ctx, signedTx.Hash(), "fulfillment USDC")
}

func deposit(t *testing.T, w *zksync.Wallet, ep zksync.EthProvider) {
	tx, err := ep.Deposit(zksync.CreateETH(), big.NewInt(depositAmount), w.GetAddress(), nil)
	require.NoError(t, err)
	require.NotNil(t, tx)

	waitEthTx(t, context.Background(), tx.Hash(), "deposit")
}

func depositUSDC(t *testing.T, w *zksync.Wallet, ep zksync.EthProvider, token *zksync.Token) {
	tx1, err := ep.ApproveDeposits(token, big.NewInt(testAmountUSDC), &zksync.GasOptions{GasLimit: 300000})
	require.NoError(t, err)
	require.NotNil(t, tx1)
	waitEthTx(t, context.Background(), tx1.Hash(), "approve deposit USDC")

	tx2, err := ep.Deposit(token, big.NewInt(depositAmountUSDC), w.GetAddress(), &zksync.GasOptions{GasLimit: 300000})
	require.NoError(t, err)
	require.NotNil(t, tx2)
	waitEthTx(t, context.Background(), tx2.Hash(), "deposit USDC")
}

func waitEthTx(t *testing.T, ctx context.Context, txHash common.Hash, title string) {
	fmt.Print("Waiting for ", title, " Eth Tx ", txHash.String())
	ctx2, cancel := context.WithTimeout(ctx, txWaitTimeout)
	defer cancel()
	isPending := true
	retries := 10
	for isPending && err == nil && retries > 0 {
		select {
		case <-time.After(txCheckInterval):
			fmt.Print(".")
			_, isPending, err = ethClient.TransactionByHash(ctx2, txHash)
			if err == ethereum.NotFound {
				// sometimes it returns this false error, so try again
				retries--
				continue
			}
			retries = 10
		case <-ctx2.Done():
			err = errors.New("context timeout")
		}
	}
	require.NoError(t, err)
	tr, err := ethClient.TransactionReceipt(ctx2, txHash)
	require.NoError(t, err)
	require.EqualValues(t, 1, tr.Status)
	fmt.Print("DONE\n")
}

func waitZkTx(zp zksync.Provider, txHash string, title string) error {
	fmt.Print("Waiting for ", title, " ZkSync Tx ", txHash)
	ctx, cancel := context.WithTimeout(context.Background(), txWaitTimeout)
	defer cancel()
	var err error
	var td *zksync.TransactionDetails
	var isCommitted bool
	for !isCommitted && err == nil {
		select {
		case <-time.After(txCheckInterval):
			fmt.Print(".")
			td, err = zp.GetTransactionDetails(txHash)
			if err != nil {
				break
			}
			if td.Block != nil && td.Block.Committed {
				isCommitted = true
			}
		case <-ctx.Done():
			err = errors.New("context timeout")
		}
	}
	if err == nil && td != nil && !td.Success && len(td.FailReason) > 0 {
		err = errors.New(td.FailReason)
	}
	if err != nil {
		fmt.Println("FAIL")
	} else {
		fmt.Println("DONE")
	}
	return err
}

func waitZkAccount(t *testing.T, zp zksync.Provider, address common.Address) {
	fmt.Print("Waiting for new ZkSync account ", address)
	ctx, cancel := context.WithTimeout(context.Background(), txWaitTimeout)
	defer cancel()
	var err error
	var as *zksync.AccountState
	var isCreated bool
	for !isCreated && err == nil {
		select {
		case <-time.After(txCheckInterval):
			fmt.Print(".")
			as, err = zp.GetState(address)
			require.NoError(t, err)
			if as.Id > 0 {
				isCreated = true
			}
		case <-ctx.Done():
			err = errors.New("context timeout")
		}
	}
	require.NoError(t, err)
	fmt.Print("DONE\n")
}
