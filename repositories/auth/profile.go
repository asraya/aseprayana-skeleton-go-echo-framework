package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"
)

func (u *authRepository) Profile(req dto.ProfileRequest) (*entity.User, error) {
	var existingUser entity.User
	err := u.DB.Preload("Limit").Preload("Role").Where("id = ?", req.ID).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
