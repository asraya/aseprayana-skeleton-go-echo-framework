package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	m "aseprayana-skeleton-go/middlewares"
	repositories "aseprayana-skeleton-go/repositories/transaction"
)

type TransactionService interface {
	GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.TransactionResponse, error)
	Create(req dto.CreateRequest) (dto.TransactionResponse, error)
	Update(req dto.UpdateRequest) (dto.TransactionResponse, error)
	Delete(req dto.DeleteRequest) (dto.TransactionResponse, error)
}

type transactionService struct {
	TransactionR repositories.TransactionRepository
	jwt          m.JWTService
}

func NewTransactionService(TransactionR repositories.TransactionRepository, jwtS m.JWTService) TransactionService {
	return &transactionService{
		TransactionR: TransactionR,
		jwt:          jwtS,
	}
}
