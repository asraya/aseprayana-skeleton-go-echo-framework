package handlers

import (
	dto "aseprayana-skeleton-go/dto/auth"
	res "aseprayana-skeleton-go/util/response"

	"github.com/labstack/echo/v4"
)

// Login
// @Security BearerAuth
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.AuthLoginRequest true "request body"
// @Success 200 {object} dto.AuthLoginResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/auth/login [post]
func (u *domainHandler) Signin(c echo.Context) error {
	var req dto.AuthLoginRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	data, err := u.serviceAuth.Signin(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)

}
