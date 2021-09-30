package zksync

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	providerMock *ProviderMock
	wallet       *Wallet

	expAccountId = uint32(185689)
	toAccountId  = uint32(189095)
	nonce        = uint32(8)
	nonce2       = uint32(3)
	accountState = &AccountState{
		Address: expAddress,
		Id:      expAccountId,
		Depositing: &DepositingState{
			Balances: map[string]*DepositingBalance{},
		},
		Committed: &State{
			Balances:   map[string]string{"ETH": "98501901000000000"},
			Nonce:      nonce,
			PubKeyHash: "", //expPubKeyHash,
			Nfts:       map[string]*NFT{},
			MintedNfts: map[string]*NFT{},
		},
		Verified: &State{
			Balances:   map[string]string{"ETH": "98501901000000000"},
			Nonce:      nonce,
			PubKeyHash: "", //expPubKeyHash,
			Nfts:       map[string]*NFT{},
			MintedNfts: map[string]*NFT{},
		},
	}
	account2State = &AccountState{
		Address: toAddress,
		Id:      toAccountId,
		Depositing: &DepositingState{
			Balances: map[string]*DepositingBalance{},
		},
		Committed: &State{
			Balances:   map[string]string{"ETH": "197846750000000000"},
			Nonce:      nonce2,
			PubKeyHash: "", //expPubKeyHash,
			Nfts:       map[string]*NFT{},
			MintedNfts: map[string]*NFT{},
		},
		Verified: &State{
			Balances:   map[string]string{"ETH": "197846750000000000"},
			Nonce:      nonce2,
			PubKeyHash: "", //expPubKeyHash,
			Nfts:       map[string]*NFT{},
			MintedNfts: map[string]*NFT{},
		},
	}

	tokens = &Tokens{
		Tokens: map[string]*Token{
			"ETH": {
				Id:       0,
				Address:  "0x0000000000000000000000000000000000000000",
				Symbol:   "ETH",
				Decimals: 18,
				IsNft:    false,
			},
		},
	}

	defaultTimeRange = DefaultTimeRange()

	totalFee     = "45400000000000"
	txFeeDetails = &TransactionFeeDetails{
		GasTxAmount: "13900",
		GasPriceWei: "1000000011",
		GasFee:      "18070000194600",
		ZkpFee:      "27404918495870",
		TotalFee:    totalFee,
	}

	NFTContentHash = common.HexToHash("5555")
	expNFT         = &NFT{
		Id:             70231,
		Symbol:         "NFT-70231",
		CreatorId:      185689,
		ContentHash:    NFTContentHash,
		CreatorAddress: common.HexToAddress(expAddress),
		SerialId:       0,
		Address:        "0xb75ec5d9b1671ede6159ac5b8c16f26ae8335abd",
	}
)

func createWallet(t *testing.T) {
	if wallet == nil {
		// create EthSigner
		ethSigner, err := NewEthSignerFromMnemonic(mnemonic)
		require.NoError(t, err)
		assert.NotNil(t, ethSigner)
		// create ZkSigner
		seed, _ := hex.DecodeString(seedHex)
		zkSigner, err := NewZkSignerFromSeed(seed)
		require.NoError(t, err)
		assert.NotNil(t, zkSigner)
		// setup Provider mock
		providerMock = &ProviderMock{}
		providerMock.On("GetState", common.HexToAddress(expAddress)).
			Return(accountState, nil).Once()

		wallet, err = NewWallet(ethSigner, zkSigner, providerMock)
		require.NoError(t, err)
		assert.NotNil(t, wallet)

		providerMock.On("GetTokens").Return(tokens, nil)
	}
}

func createWallet2(t *testing.T) *Wallet {
	// create EthSigner
	ethSigner, err := NewEthSignerFromMnemonic(mnemonic2)
	require.NoError(t, err)
	assert.NotNil(t, ethSigner)
	// create ZkSigner
	seed2, _ := hex.DecodeString(seed2Hex)
	zkSigner, err := NewZkSignerFromSeed(seed2)
	require.NoError(t, err)
	assert.NotNil(t, zkSigner)
	// setup Provider mock
	providerMock.On("GetState", common.HexToAddress(toAddress)).
		Return(account2State, nil).Once()

	wallet2, err := NewWallet(ethSigner, zkSigner, providerMock)
	require.NoError(t, err)
	assert.NotNil(t, wallet2)

	return wallet2
}

func TestNewWallet(t *testing.T) {
	createWallet(t)

	address := wallet.GetAddress()
	assert.NotNil(t, address)
	assert.IsType(t, common.Address{}, address)
	assert.EqualValues(t, expAddress, address.String())

	accountId, err := wallet.GetAccountId()
	require.NoError(t, err)
	assert.EqualValues(t, expAccountId, accountId)

	providerMock.On("GetState", common.HexToAddress(expAddress)).
		Return(accountState, nil).Once()
	state, err := wallet.GetState()
	require.NoError(t, err)
	assert.Equal(t, accountState, state)

	pr := wallet.GetProvider()
	require.NoError(t, err)
	assert.EqualValues(t, providerMock, pr)
}

func TestChangePubKey(t *testing.T) {
	createWallet(t)

	t.Run("ECDSA", func(t *testing.T) {
		// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
		providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.ChangePubKey"), mock.AnythingOfType("*zksync.EthSignature"), false).
			Run(func(args mock.Arguments) {
				// assert arguments
				assert.EqualValues(t, &ChangePubKey{
					Type:      "ChangePubKey",
					AccountId: expAccountId,
					Account:   common.HexToAddress(expAddress),
					NewPkHash: expPubKeyHash,
					FeeToken:  0,
					Fee:       totalFee,
					Nonce:     nonce,
					Signature: &Signature{
						PubKey:    expPubKey,
						Signature: "3807301038f4edeaad14c39cd27e2f112a1edcdc69f7de714acd48301c340a1447606e546246167e4a6c5b3a0da740525c6c6698989c6529c749898037333c05",
					},
					EthAuthData: &ChangePubKeyECDSA{
						Type:         "ECDSA",
						EthSignature: "0x50904be19b8dd300fd60d73325c183293b9ee21ee8dced7b462d257368bd60180666dac29b4b83fac9efa6b69619e181ad82b3ea9255a50589b2460d3173db311b",
						BatchHash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
					},
					TimeRange: defaultTimeRange,
				}, args.Get(0))
				assert.EqualValues(t, (*EthSignature)(nil), args.Get(1))
				assert.EqualValues(t, false, args.Get(2))
			}).
			Return("txHash", nil).Once()
		txHash, err := wallet.SetSigningKey(txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, false, defaultTimeRange)
		require.NoError(t, err)
		assert.EqualValues(t, "txHash", txHash)
	})

	t.Run("OnChain", func(t *testing.T) {
		// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
		providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.ChangePubKey"), mock.AnythingOfType("*zksync.EthSignature"), false).
			Run(func(args mock.Arguments) {
				// assert arguments
				assert.EqualValues(t, &ChangePubKey{
					Type:      "ChangePubKey",
					AccountId: expAccountId,
					Account:   common.HexToAddress(expAddress),
					NewPkHash: expPubKeyHash,
					FeeToken:  0,
					Fee:       totalFee,
					Nonce:     nonce,
					Signature: &Signature{
						PubKey:    expPubKey,
						Signature: "3807301038f4edeaad14c39cd27e2f112a1edcdc69f7de714acd48301c340a1447606e546246167e4a6c5b3a0da740525c6c6698989c6529c749898037333c05",
					},
					EthAuthData: &ChangePubKeyOnchain{
						Type: "Onchain",
					},
					TimeRange: defaultTimeRange,
				}, args.Get(0))
				assert.EqualValues(t, (*EthSignature)(nil), args.Get(1))
				assert.EqualValues(t, false, args.Get(2))
			}).
			Return("txHash", nil).Once()
		txHash, err := wallet.SetSigningKey(txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, true, defaultTimeRange)
		require.NoError(t, err)
		assert.EqualValues(t, "txHash", txHash)
	})
}

func TestTransfer(t *testing.T) {
	createWallet(t)

	amount := big.NewInt(12312124)
	// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
	providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.Transfer"), mock.AnythingOfType("*zksync.EthSignature"), false).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &Transfer{
				Type:      "Transfer",
				AccountId: expAccountId,
				From:      common.HexToAddress(expAddress),
				To:        common.HexToAddress(toAddress),
				Amount:    amount,
				Token:     CreateETH(),
				TokenId:   0,
				Fee:       totalFee,
				Nonce:     nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "8a215f7ce352471542a1de78d011558765ead7fdfb172a65b06fdadfac23309e9f1957d51090bc390fbc62791bbb5691f3b3078de3d5426712d5d2c570bc5000",
				},
				TimeRange: defaultTimeRange,
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0xfe1af278836c4556fd08a7b1b56c77a2f9e7ee85f34b1498067c1277c752268213bab629c4cb3edb25818b005545accd04f4d03e05afd4fd730b4c4cc1c27cd91b",
			}, args.Get(1))
			assert.EqualValues(t, false, args.Get(2))
		}).
		Return("txHash", nil).Once()
	txHash, err := wallet.SyncTransfer(common.HexToAddress(toAddress), amount, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}

func TestWithdraw(t *testing.T) {
	createWallet(t)

	amount := big.NewInt(12312124)
	// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
	providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.Withdraw"), mock.AnythingOfType("*zksync.EthSignature"), false).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &Withdraw{
				Type:      "Withdraw",
				AccountId: expAccountId,
				From:      common.HexToAddress(expAddress),
				To:        common.HexToAddress(toAddress),
				Amount:    amount,
				TokenId:   0,
				Fee:       totalFee,
				Nonce:     nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "e3ee5a13c147cd70579c17b261e057a41ed4719932e782eb69483aba0a79eb09b4517a16b1e1dbe6dc97f31f1fd38ae211d4b9afdcd3be59d0f9e4df65f88901",
				},
				TimeRange: defaultTimeRange,
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0x3da4e0fb87b97886099113c88e9f15858455a5047e22ec8bd043d559da8d1e9f492751bd33ccdd1fa51ebe95f0d633cbbcaeb8591a0ec4aaf8e70c7bbae529fb1c",
			}, args.Get(1))
			assert.EqualValues(t, false, args.Get(2))
		}).
		Return("txHash", nil).Once()
	txHash, err := wallet.SyncWithdraw(common.HexToAddress(toAddress), amount, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, false, defaultTimeRange)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}

func TestForcedExit(t *testing.T) {
	createWallet(t)

	// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
	providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.ForcedExit"), mock.AnythingOfType("*zksync.EthSignature"), false).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &ForcedExit{
				Type:      "ForcedExit",
				AccountId: expAccountId,
				Target:    common.HexToAddress(toAddress),
				Amount:    nil,
				TokenId:   0,
				Fee:       totalFee,
				Nonce:     nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "cf72d635211d6d9105c9da6d249d325134278b865d64889f5d586722d8bfa92c8bd6b2a865ab881bf6cd63089b7cb875927f12c6469189f455562901f57ea202",
				},
				TimeRange: defaultTimeRange,
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0x09f55237cb0056f4fc056b76568b97f17742d3ff7e82598c88e533a0dd118df544d625682ad8a1570f6571aac94aadc4ed50531248473292292abaf5950ca4fa1b",
			}, args.Get(1))
			assert.EqualValues(t, false, args.Get(2))
		}).
		Return("txHash", nil).Once()
	txHash, err := wallet.SyncForcedExit(common.HexToAddress(toAddress), txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}

func TestMintNFT(t *testing.T) {
	createWallet(t)

	// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
	providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.MintNFT"), mock.AnythingOfType("*zksync.EthSignature"), false).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &MintNFT{
				Type:           "MintNFT",
				CreatorId:      expAccountId,
				CreatorAddress: common.HexToAddress(expAddress),
				Recipient:      common.HexToAddress(toAddress),
				ContentHash:    NFTContentHash,
				Fee:            totalFee,
				Nonce:          nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "908eb8f6337676b0561b99d6ae83086f3e93d4900cbdf2f8b9234038dd08801590a4d301f71e5c77680db7a2ab10c65dee18945267a373c75015eb67879a5003",
				},
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0x1651b964f061df09f7d710020d0bbfd05c631216a1e29db2abfc1431a52dc6696a5cb736a4bdc0e45f22be438f97ed71ffd47a46e3dc6710549962e7be31b5e71b",
			}, args.Get(1))
			assert.EqualValues(t, false, args.Get(2))
		}).
		Return("txHash", nil).Once()
	txHash, err := wallet.SyncMintNFT(common.HexToAddress(toAddress), NFTContentHash, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}

func TestWithdrawNFT(t *testing.T) {
	createWallet(t)

	// expect for invoking of mocked provider.SubmitTx() method with properly prepared tx and valid signatures
	providerMock.On("SubmitTx", mock.AnythingOfType("*zksync.WithdrawNFT"), mock.AnythingOfType("*zksync.EthSignature"), false).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &WithdrawNFT{
				Type:      "WithdrawNFT",
				AccountId: expAccountId,
				From:      common.HexToAddress(expAddress),
				To:        common.HexToAddress(toAddress),
				Token:     expNFT.Id,
				FeeToken:  0,
				Fee:       totalFee,
				Nonce:     nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "f8b164dd54e2be0239bcefc9bea69b26e74a6b4126ae186da2277c0cf76a43a2cb6002a86cc3e2fbce62646550ef2f7753d787e68e8d0e3baf1f6d52508de500",
				},
				TimeRange: defaultTimeRange,
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0xe1c842e393bc6db5cf45fc2e99144c66e2db6e79b1715e554f1304bf67d0e8220b84dcd724b382ff4d8d5ba3e7d5cbfc72b1c11f31bd2c1317ec8f6b958754691c",
			}, args.Get(1))
			assert.EqualValues(t, false, args.Get(2))
		}).
		Return("txHash", nil).Once()
	txHash, err := wallet.SyncWithdrawNFT(common.HexToAddress(toAddress), expNFT, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}

func TestTransferNFT(t *testing.T) {
	createWallet(t)

	// expect for invoking of provider.SubmitTxsBatch() method with properly prepared txs and valid signatures
	providerMock.On("SubmitTxsBatch", mock.AnythingOfType("[]*zksync.SignedTransaction"), mock.AnythingOfType("*zksync.EthSignature")).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, []*SignedTransaction{
				{
					Transaction: &Transfer{
						Type:      "Transfer",
						AccountId: expAccountId,
						From:      common.HexToAddress(expAddress),
						To:        common.HexToAddress(toAddress),
						Token:     expNFT.ToToken(),
						TokenId:   expNFT.Id,
						Amount:    big.NewInt(1),
						Nonce:     nonce,
						Fee:       big.NewInt(0).String(),
						Signature: &Signature{
							PubKey:    expPubKey,
							Signature: "8d89c37a4586215622b27a46561f313779e0ab4df0a3370c5c032ede15bedb9ef92d6aa92d135bd1e305d24b0b24cccc6939c2fe901a36f312796ae05b8c4103",
						},
						TimeRange: defaultTimeRange,
					},
				},
				{
					Transaction: &Transfer{
						Type:      "Transfer",
						AccountId: expAccountId,
						From:      common.HexToAddress(expAddress),
						To:        common.HexToAddress(expAddress),
						Token:     CreateETH(),
						TokenId:   0,
						Amount:    big.NewInt(0),
						Nonce:     nonce + 1,
						Fee:       totalFee,
						Signature: &Signature{
							PubKey:    expPubKey,
							Signature: "cb4ce668c30b5b65a8cdfa8bb4260c619ba975e8593d6d169dd57ad3a6554b119f3d07d8125814b5b1970ee22d69e9ed30b9410e65df5a333cda1daf0da79704",
						},
						TimeRange: defaultTimeRange,
					},
				},
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0x6006560738fe3b90022a7dac270c35ab79c224021d04fa18ec130d266be131dd0d54aac3663ef0a3f4e6ec69dd904725fe1bb051e87b4bdcf5b30fd2f168ec1f1c",
			}, args.Get(1))

		}).
		Return([]string{"tx1Hash", "tx2Hash"}, nil).Once()
	txs, err := wallet.SyncTransferNFT(common.HexToAddress(toAddress), expNFT, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(txs))
	assert.EqualValues(t, "tx1Hash", txs[0])
	assert.EqualValues(t, "tx2Hash", txs[1])
}

func TestSwap(t *testing.T) {
	createWallet(t)
	amount := big.NewInt(100000000000000)

	order1, err := wallet.BuildSignedOrder(common.HexToAddress(expAddress),
		CreateETH(), CreateETH(), []*big.Int{big.NewInt(1), big.NewInt(1)}, amount,
		accountState.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	require.NotNil(t, order1)
	assert.EqualValues(t, &Order{
		AccountId:        expAccountId,
		RecipientAddress: common.HexToAddress(expAddress),
		Nonce:            nonce,
		TokenBuy:         0,
		TokenSell:        0,
		Ratio:            []*big.Int{big.NewInt(1), big.NewInt(1)},
		Amount:           amount,
		Signature: &Signature{
			PubKey:    expPubKey,
			Signature: "0021f5b88420ee5fba86ae31ae03b22c56b61cf13fc14a73ce04fee40e46b68e2e086271513844059a0d079ad37e5ad1b7dd22d19290c8d26d22d480fb070805",
		},
		EthereumSignature: &EthSignature{
			Type:      "EthereumSignature",
			Signature: "0xc101d1eecf8f35b74efc9d296f5d236389178f917a33416a15a85795f3b2c88d4e2306edde31106b16f7ffdc008451f57c9334a38bdbda81e608dc01b3c2a6831c",
		},
		TimeRange: defaultTimeRange,
	}, order1)

	wallet2 := createWallet2(t)
	order2, err := wallet2.BuildSignedOrder(common.HexToAddress(toAddress),
		CreateETH(), CreateETH(), []*big.Int{big.NewInt(1), big.NewInt(1)}, amount,
		account2State.Committed.Nonce, defaultTimeRange)
	require.NoError(t, err)
	require.NotNil(t, order2)
	assert.EqualValues(t, &Order{
		AccountId:        toAccountId,
		RecipientAddress: common.HexToAddress(toAddress),
		Nonce:            nonce2,
		TokenBuy:         0,
		TokenSell:        0,
		Ratio:            []*big.Int{big.NewInt(1), big.NewInt(1)},
		Amount:           amount,
		Signature: &Signature{
			PubKey:    toPubKey,
			Signature: "3f01d85af2339b1d800ff727c2c9b66af29d6e449cb1fbe25af8643b57fc0f2401a3aa2e1ef236284919fa7edd7423acefab78aa46503f7399a6a06da95a2d05",
		},
		EthereumSignature: &EthSignature{
			Type:      "EthereumSignature",
			Signature: "0xbc74f5046d9e30cc8faa0e4ff39dae207d112aa3e17b93ae12ea20f39928d12548b78986a707aa157345b517e8081a3b25c7223761c419fcf1abc66ef8d9623a1b",
		},
		TimeRange: defaultTimeRange,
	}, order2)

	// expect for invoking of provider.SubmitTxMultiSig() method with properly prepared swap tx and three valid signatures
	providerMock.On("SubmitTxMultiSig", mock.AnythingOfType("*zksync.Swap"),
		mock.AnythingOfType("*zksync.EthSignature"),
		mock.AnythingOfType("*zksync.EthSignature"),
		mock.AnythingOfType("*zksync.EthSignature")).
		Run(func(args mock.Arguments) {
			// assert arguments
			assert.EqualValues(t, &Swap{
				Type:             "Swap",
				SubmitterId:      expAccountId,
				SubmitterAddress: common.HexToAddress(expAddress),
				Orders:           []*Order{order1, order2},
				Amounts:          []*big.Int{amount, amount},
				FeeToken:         0,
				Fee:              totalFee,
				Nonce:            nonce,
				Signature: &Signature{
					PubKey:    expPubKey,
					Signature: "169a3e96a63c983a7c7dc72c0882fb0c563d69803e11ccc9b9a0f0d00855c19dc3ea9726f94b32a1ce5d629ae9132a35e33d63088662e32600bfcd0ea8404e04",
				},
			}, args.Get(0))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0x0a6d3a68c54306e6e4739ab6d7c3144460acaffa31ca6931cdc94478d6aa14443875a986c4c3695219bc3bb14d5e1de9030dccb73deb2150bbf405ce650f43f91b",
			}, args.Get(1))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0xc101d1eecf8f35b74efc9d296f5d236389178f917a33416a15a85795f3b2c88d4e2306edde31106b16f7ffdc008451f57c9334a38bdbda81e608dc01b3c2a6831c",
			}, args.Get(2))
			assert.EqualValues(t, &EthSignature{
				Type:      "EthereumSignature",
				Signature: "0xbc74f5046d9e30cc8faa0e4ff39dae207d112aa3e17b93ae12ea20f39928d12548b78986a707aa157345b517e8081a3b25c7223761c419fcf1abc66ef8d9623a1b",
			}, args.Get(3))
		}).
		Return("txHash", nil).Once()

	txHash, err := wallet.SyncSwap(order1, order2, amount, amount, txFeeDetails.GetTxFee(CreateETH()), accountState.Committed.Nonce)
	require.NoError(t, err)
	assert.EqualValues(t, "txHash", txHash)
}
