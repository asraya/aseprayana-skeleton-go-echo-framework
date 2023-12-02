package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
)

func (b *transactionService) Create(req dto.CreateRequest) (dto.TransactionResponse, error) {

	transaction, err := b.TransactionR.Create(req)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	response := dto.TransactionResponse{
		ID:            transaction.ID,
		NomorKontrak:  transaction.NomorKontrak,
		OTR:           transaction.OTR,
		AdminFee:      transaction.AdminFee,
		JumlahCicilan: transaction.JumlahCicilan,
		JumlahBunga:   transaction.JumlahBunga,
		NamaAsset:     transaction.NamaAsset,
		CreatedBy:     req.CreatedBy,
	}

	return response, nil
}
