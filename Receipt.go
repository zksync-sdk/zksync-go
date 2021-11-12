package zksync

import (
	"github.com/pkg/errors"
	"time"
)

type TransactionReceiptProcessor struct {
	provider      Provider
	checkInterval time.Duration
	timeout       time.Duration
}

func NewTransactionReceiptProcessor(provider Provider) *TransactionReceiptProcessor {
	return &TransactionReceiptProcessor{
		provider:      provider,
		checkInterval: time.Second,
		timeout:       5 * time.Minute,
	}
}

func NewTransactionReceiptProcessorDurations(provider Provider,
	checkInterval time.Duration, timeout time.Duration) *TransactionReceiptProcessor {
	return &TransactionReceiptProcessor{
		provider:      provider,
		checkInterval: checkInterval,
		timeout:       timeout,
	}
}

func (rp *TransactionReceiptProcessor) WaitForTransaction(txHash string, status TransactionStatus) (*TransactionDetails, error) {
	checkTransactionDetails := func(td *TransactionDetails, err error) (*TransactionDetails, error) {
		if err != nil {
			return nil, errors.Wrap(err, "failed to get transaction details by txHash")
		}
		if td.Executed {
			switch status {
			case TransactionStatusSent:
				if td.Block != nil {
					return td, nil
				}
			case TransactionStatusCommitted:
				if td.Block != nil && td.Block.Committed {
					return td, nil
				}
			case TransactionStatusVerified:
				if td.Block != nil && td.Block.Verified {
					return td, nil
				}
			}
		}
		return nil, nil
	}
	td, err := checkTransactionDetails(rp.provider.GetTransactionDetails(txHash))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction details by txHash")
	} else if td != nil {
		return td, nil
	}
	for {
		select {
		case <-time.After(rp.checkInterval):
			td, err = checkTransactionDetails(rp.provider.GetTransactionDetails(txHash))
			if err != nil {
				return nil, errors.Wrap(err, "failed to get transaction details by txHash")
			} else if td != nil {
				return td, nil
			}
		case <-time.After(rp.timeout):
			return nil, errors.New("Transaction was not generated, waiting timeout reached")
		}
	}
}
