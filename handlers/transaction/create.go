package handlers

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	res "aseprayana-skeleton-go/util/response"

	"github.com/labstack/echo/v4"
)

// Create
// @Summary Create transaction
// @Description Create transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Authorization Bearer"
// @Param request body dto.CreateRequest true "request body"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/transaction [post]
func (b *domainHandler) Create(c echo.Context) error {
	var req dto.CreateRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	data, err := b.domainService.Create(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)

}
