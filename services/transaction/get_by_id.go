package services

import dto "aseprayana-skeleton-go/dto/transaction"

func (b *transactionService) GetById(req dto.GetByIdRequest) (*dto.TransactionResponse, error) {
	transaction, err := b.TransactionR.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
