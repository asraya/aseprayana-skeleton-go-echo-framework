package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
)

func (b *transactionService) Update(req dto.UpdateRequest) (dto.TransactionResponse, error) {
	transactionBody := dto.UpdateRequest{
		NomorKontrak:  req.NomorKontrak,
		OTR:           req.OTR,
		AdminFee:      req.AdminFee,
		JumlahCicilan: req.JumlahCicilan,
		JumlahBunga:   req.JumlahBunga,
		NamaAsset:     req.NamaAsset,
		UpdatedBy:     req.UpdatedBy,
	}

	Transaction, err := b.TransactionR.Update(req)
	if err != nil {
		return Transaction, err
	}

	response := dto.TransactionResponse{
		NomorKontrak:  transactionBody.NomorKontrak,
		OTR:           transactionBody.OTR,
		AdminFee:      transactionBody.AdminFee,
		JumlahCicilan: transactionBody.JumlahCicilan,
		JumlahBunga:   transactionBody.JumlahBunga,
		NamaAsset:     transactionBody.NamaAsset,
		UpdatedBy:     transactionBody.UpdatedBy,
	}

	return response, nil
}
