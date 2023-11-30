package dto

type LimitResponse struct {
	Limit float64 `gorm:"limit" json:"limit"`
	Month uint    `gorm:"month" json:"month"`
}
