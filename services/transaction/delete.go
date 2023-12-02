package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
)

func (b *transactionService) Delete(req dto.DeleteRequest) (dto.TransactionResponse, error) {
	transactionBody := dto.DeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.TransactionR.Delete(req)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	response := dto.TransactionResponse{
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
