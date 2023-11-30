package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"
)

func (u *authRepository) Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {
	user := entity.User{
		ID:           req.ID,
		FullName:     req.FullName,
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
		FullName:     user.FullName,
		LegalName:    user.LegalName,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir,
		Gaji:         user.Gaji,
		Email:        user.Email,
		Password:     user.Password,
	}

	return response, nil
}
