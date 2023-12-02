package auth

import (
	dtor "aseprayana-skeleton-go/dto"
	dto "aseprayana-skeleton-go/dto/auth"
	dtol "aseprayana-skeleton-go/dto/limit"
)

func (u *authService) Profile(req dto.ProfileRequest) (*dto.ProfileResponse, error) {
	transaction, err := u.Repo.Profile(req)
	if err != nil {
		return nil, err
	}

	// Fetch limit information
	var limits []dtol.LimitResponse
	for _, limit := range transaction.Limit {
		limits = append(limits, dtol.LimitResponse{
			Limit: limit.Limit,
			Month: limit.Month,
		})
	}

	roles := dtor.RoleUserResponse{
		RoleID: transaction.Role.RoleID,
	}

	return &dto.ProfileResponse{
		FullName:     transaction.FullName,
		Email:        transaction.Email,
		LegalName:    transaction.LegalName,
		Roles:        roles,
		TempatLahir:  transaction.TempatLahir,
		TanggalLahir: transaction.TanggalLahir,
		Gaji:         transaction.Gaji,
		FotoKtp:      transaction.FotoKtp,
		FotoSelfie:   transaction.FotoSelfie,
		Limits:       limits,
	}, nil
}
