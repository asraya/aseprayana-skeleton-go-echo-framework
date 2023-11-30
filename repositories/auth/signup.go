package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"
	"aseprayana-skeleton-go/util"
)

func (u *authRepository) Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {
	user := entity.User{
		ID:           util.GenerateRandomString(),
		FullName:     req.FullName,
		KTP:          req.KTP,
		LegalName:    req.LegalName,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		Gaji:         req.Gaji,
		Email:        req.Email,
		Password:     req.Password,
	}

	if err := u.DB.Save(&user).First(&user).Error; err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	response := dto.AuthRegisterResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		LegalName:    user.LegalName,
		KTP:          user.KTP,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir,
		Gaji:         user.Gaji,
		Email:        user.Email,
		Password:     user.Password,
	}

	return response, nil
}
