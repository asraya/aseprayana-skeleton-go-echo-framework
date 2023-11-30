package repositories

import (
	"aseprayana-skeleton-go/entity"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock interface {
	GetTransactionsRepository() ([]*entity.Transaction, error)
	GetTransactionRepository(id string) (*entity.Transaction, error)
	CreateRepository(Transaction entity.Transaction) (*entity.Transaction, error)
	UpdateRepository(id string, TransactionBody entity.Transaction) (*entity.Transaction, error)
	DeleteRepository(id string) error
}

type ItransactionRepositoryMock struct {
	Mock mock.Mock
}

func NewTransactionRepositoryMock(mock mock.Mock) TransactionRepositoryMock {
	return &ItransactionRepositoryMock{
		Mock: mock,
	}
}

func (b *ItransactionRepositoryMock) GetTransactionsRepository() ([]*entity.Transaction, error) {
	args := b.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transactions := args.Get(0).([]*entity.Transaction)

	return transactions, nil
}

func (b *ItransactionRepositoryMock) GetTransactionRepository(id string) (*entity.Transaction, error) {
	args := b.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transaction := args.Get(0).(entity.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) CreateRepository(transactionData entity.Transaction) (*entity.Transaction, error) {
	args := u.Mock.Called(transactionData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transaction := args.Get(0).(entity.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) UpdateRepository(id string, transactionData entity.Transaction) (*entity.Transaction, error) {
	args := u.Mock.Called(id, transactionData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transaction := args.Get(0).(entity.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
