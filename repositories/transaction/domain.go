package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetById(req dto.GetByIdRequest) (*dto.TransactionResponse, error)
	GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error)
	Create(req dto.CreateRequest) (*dto.TransactionResponse, error)
	Update(req dto.UpdateRequest) (dto.TransactionResponse, error)
	Delete(req dto.DeleteRequest) (*dto.TransactionResponse, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: DB,
	}
}
