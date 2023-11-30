package handlers

import (
	m "aseprayana-skeleton-go/middlewares"
	s "aseprayana-skeleton-go/services/transaction"

	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type domainHandler struct {
	domainService s.TransactionService
	jwt           m.JWTService
}

func NewTransactionHandler(TransactionS s.TransactionService, jwtS m.JWTService) TransactionHandler {
	return &domainHandler{
		domainService: TransactionS,
		jwt:           jwtS,
	}
}
