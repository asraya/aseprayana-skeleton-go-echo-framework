package auth

import (
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

	return &dto.ProfileResponse{
		ID:           transaction.ID,
		FullName:     transaction.FullName,
		Email:        transaction.Email,
		LegalName:    transaction.LegalName,
		TempatLahir:  transaction.TempatLahir,
		TanggalLahir: transaction.TanggalLahir,
		Gaji:         transaction.Gaji,
		Password:     transaction.Password,
		FotoKtp:      transaction.FotoKtp,
		FotoSelfie:   transaction.FotoSelfie,
		Limits:       limits,
	}, nil
}
