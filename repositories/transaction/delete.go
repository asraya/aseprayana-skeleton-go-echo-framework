package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
)

func (b *transactionRepository) Delete(req dto.DeleteRequest) (*dto.TransactionResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return nil, err
	}

	if err := b.DB.Delete(&entity.Transaction{}, req.ID).Error; err != nil {
		return nil, err
	}

	return nil, err
}
