package zksync

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
