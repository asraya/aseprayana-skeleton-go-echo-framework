package handlers

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	res "aseprayana-skeleton-go/util/response"

	"github.com/labstack/echo/v4"
)

// Delete
// @Summary Delete transaction
// @Description Delete transaction
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
// @Router /api/v1/transaction/{id} [delete]
func (b *domainHandler) Delete(c echo.Context) error {
	var req dto.DeleteRequest
	deletedBy, ok := c.Get("DeletedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.DeletedBy = deletedBy

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = idUint

	data, err := b.domainService.Delete(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
