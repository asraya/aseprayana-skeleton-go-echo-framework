package handlers

import (
	dto "aseprayana-skeleton-go/dto/transaction"

	"github.com/labstack/echo/v4"
)

func (b *domainHandler) GetAll(c echo.Context) error {
	var paginationDTO dto.GetAllRequest
	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	users, err := b.domainService.GetAll(paginationDTO)
	if err != nil {
		return c.JSON(500, "Error fetching users")
	}

	return c.JSON(200, users)
}
