package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
)

func (b *transactionRepository) GetById(req dto.GetByIdRequest) (*dto.TransactionResponse, error) {
	tr := entity.Transaction{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.TransactionResponse{
		NomorKontrak: tr.NomorKontrak,
	}

	return response, nil
}
