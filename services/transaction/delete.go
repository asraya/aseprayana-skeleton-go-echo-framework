package services

import dto "aseprayana-skeleton-go/dto/transaction"

func (b *transactionService) Delete(req dto.DeleteRequest) (*dto.TransactionResponse, error) {
	tr := dto.DeleteRequest{
		ID: req.ID,
	}
	_, err := b.TransactionR.Delete(tr)
	if err != nil {
		return nil, err
	}

	return nil, err
}
