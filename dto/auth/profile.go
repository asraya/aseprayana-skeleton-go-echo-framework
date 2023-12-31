package dto

import (
	dtor "aseprayana-skeleton-go/dto"
	dto "aseprayana-skeleton-go/dto/limit"
)

type ProfileRequest struct {
	ID string `param:"id" validate:"required"`
}

type ProfileResponse struct {
	ID           uint                  `json:"id"`
	FullName     string                `json:"full_name"`
	LegalName    string                `json:"legal_name"`
	Email        string                `json:"email"`
	Roles        dtor.RoleUserResponse `json:"role"`
	TempatLahir  string                `json:"tempat_lahir"`
	TanggalLahir string                `json:"tanggal_lahir"`
	Gaji         string                `json:"gaji"`
	FotoKtp      string                `json:"foto_ktp"`
	FotoSelfie   string                `json:"foto_selfie"`
	Limits       []dto.LimitResponse   `json:"limits"`
}
