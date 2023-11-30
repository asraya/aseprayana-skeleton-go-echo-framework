package auth

import (
	dto "aseprayana-skeleton-go/dto/auth"

	"aseprayana-skeleton-go/util"
	"errors"
)

func (u *authService) Signin(req dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	user, err := u.Repo.Signin(req)
	if err != nil || util.VerifyPassword(user.Password, req.Password) != nil {
		return nil, errors.New("Password verification failed")
	}
	// Fetch limit information
	// var limits []dtol.LimitResponse
	// for _, limit := range user.Limit {
	// 	limits = append(limits, dtol.LimitResponse{
	// 		Limit: limit.Limit,
	// 		Month: limit.Month,
	// 	})
	// }
	token, err := u.jwt.GenerateToken(user.ID, user.FullName)
	return &dto.AuthLoginResponse{
		ID:       util.GenerateRandomString(),
		FullName: user.FullName,
		Email:    user.Email,
		Token:    token,
		// Limits:   limits,
	}, nil
}
