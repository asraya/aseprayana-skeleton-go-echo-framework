package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
)

func (b *transactionRepository) Delete(req dto.DeleteRequest) (dto.TransactionResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Transaction{}).Error; err != nil {
		return dto.TransactionResponse{}, err
	}

	response := dto.TransactionResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
