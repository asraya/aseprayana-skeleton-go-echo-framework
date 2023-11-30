package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            string         `gorm:"primary_key" json:"id"`
	NomorKontrak  string         `gorm:"type:varchar(255)" json:"nomor_kontrak" form:"nomor_kontrak" validate:"required,nomor_kontrak"`
	OTR           string         `gorm:"type:varchar(255)" json:"otr" form:"otr" validate:"required,otr"`
	AdminFee      string         `gorm:"type:varchar(255)" json:"admin_fee"`
	JumlahCicilan string         `gorm:"type:varchar(255)" json:"jumlah_cicilan"`
	JumlahBunga   string         `gorm:"type:varchar(255)" json:"jumlah_bunga"`
	NamaAsset     string         `gorm:"type:varchar(255)" json:"nama_asset"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
