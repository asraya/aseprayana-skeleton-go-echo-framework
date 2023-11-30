package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"
)

func (u *authRepository) Signin(req dto.AuthLoginRequest) (*entity.User, error) {
	var existingUser entity.User
	err := u.DB.Preload("Limit").Where("email = ?", req.Email).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
