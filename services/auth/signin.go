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

	token, err := u.jwt.GenerateToken(user.ID, user.FullName)
	return &dto.AuthLoginResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Token:    token,
	}, nil
}
