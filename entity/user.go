package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"primary_key" json:"id"`
	KTP          string         `gorm:"ktp" json:"ktp"`
	FullName     string         `gorm:"type:varchar(255)" json:"full_name"`
	LegalName    string         `gorm:"type:varchar(255)" json:"legal_name"`
	Email        string         `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	TempatLahir  string         `gorm:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir string         `gorm:"tanggal_lahir" json:"tanggal_lahir"`
	Gaji         string         `gorm:"gaji" json:"gaji"`
	FotoKtp      string         `gorm:"foto_ktp" json:"foto_ktp"`
	FotoSelfie   string         `gorm:"foto_selfie" json:"foto_selfie"`
	Password     string         `gorm:"password" json:"password"`
	Limit        []Limit        `json:"limit" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
