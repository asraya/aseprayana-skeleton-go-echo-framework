package entity

type Limit struct {
	ID     uint    `gorm:"primary_key:auto_increment" json:"id"`
	UserID uint    `gorm:"user_id" json:"user_id"`
	Limit  float64 `gorm:"limit" json:"limit"`
	Month  uint    `gorm:"month" json:"month"`
}
