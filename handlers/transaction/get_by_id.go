package handlers

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	res "aseprayana-skeleton-go/util/response"

	"github.com/labstack/echo/v4"
)

// Get By ID
// @Summary Get By ID transaction
// @Description Get By ID transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Authorization Bearer"
// @Param id path int true "id path"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/transaction/{id} [get]
func (b *domainHandler) GetById(c echo.Context) error {
	var req dto.GetByIdRequest

	// idStr := c.Param("id")

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = idUint

	transaction, err := b.domainService.GetById(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}
