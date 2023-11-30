package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
)

// func (b *transactionRepository) UpdateRepository(id string, TransactionBody entity.Transaction) (*entity.Transaction, error) {
// 	Transaction, err := b.GetById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = b.DB.Where("ID = ?", id).Updates(entity.Transaction{
// 		NamaAsset: TransactionBody.NamaAsset,
// 	}).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	Transaction.NamaAsset = TransactionBody.NamaAsset

// 	return Transaction, nil
// }

func (b *transactionRepository) Update(req dto.UpdateRequest) (dto.TransactionResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.Transaction{
		NomorKontrak:  req.NomorKontrak,
		OTR:           req.OTR,
		AdminFee:      req.AdminFee,
		JumlahCicilan: req.JumlahCicilan,
		JumlahBunga:   req.JumlahBunga,
		NamaAsset:     req.NamaAsset,
	}

	transaction, err := b.GetById(tr)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.Transaction{
		NomorKontrak:  request.NomorKontrak,
		OTR:           request.OTR,
		AdminFee:      request.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		JumlahBunga:   request.JumlahBunga,
		NamaAsset:     request.NamaAsset,
	}).Error
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	response := dto.TransactionResponse{
		NomorKontrak:  request.NomorKontrak,
		OTR:           request.OTR,
		AdminFee:      transaction.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		JumlahBunga:   transaction.JumlahBunga,
		NamaAsset:     request.NamaAsset,
	}

	return response, nil
}
