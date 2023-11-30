package dto

// Register
type AuthRegisterRequest struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	KTP          string `json:"ktp"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gaji         string `json:"gaji"`
	FotoKtp      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Password     string `json:"password" form:"password" validate:"required,password"`
}

type AuthRegisterResponse struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	KTP          string `json:"ktp"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gaji         string `json:"gaji"`
	FotoKtp      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
	Email        string `json:"email"`
	Password     string `json:"password" form:"password" validate:"required,password"`
}
