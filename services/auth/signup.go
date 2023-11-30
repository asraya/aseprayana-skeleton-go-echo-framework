package auth

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/util"
)

func (u *authService) Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {

	encryp := util.EncryptPassword(&req)
	if encryp != nil {
		return dto.AuthRegisterResponse{}, encryp
	}

	user := dto.AuthRegisterRequest{
		ID:           req.ID,
		FullName:     req.FullName,
		LegalName:    req.LegalName,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		Gaji:         req.Gaji,
		Email:        req.Email,
		Password:     req.Password,
	}

	createdUser, err := u.Repo.Signup(user)
	if err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	response := dto.AuthRegisterResponse{
		FullName:     createdUser.FullName,
		LegalName:    createdUser.LegalName,
		TempatLahir:  createdUser.TempatLahir,
		TanggalLahir: createdUser.TanggalLahir,
		Gaji:         createdUser.Gaji,
		Email:        createdUser.Email,
		Password:     createdUser.Password,
	}

	return response, nil
}
