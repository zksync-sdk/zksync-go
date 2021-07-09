package zksync

type TransactionType string

func (t TransactionType) getType() interface{} {
	switch t {
	case TransactionTypeChangePubKeyOnchain, TransactionTypeChangePubKeyECDSA, TransactionTypeChangePubKeyCREATE2:
		// custom object instead of string
		return TransactionTypeChangePubKey{ChangePubKey: string(t)}
	default:
		return string(t)
	}
}

type ZksTransaction interface {
	getType() string
}

type SignedTransaction struct {
	transaction       ZksTransaction
	ethereumSignature *EthSignature
}

func (tx *SignedTransaction) getTransaction() ZksTransaction {
	return tx.transaction
}
