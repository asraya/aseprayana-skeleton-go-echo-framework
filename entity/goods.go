package entity

type Goods struct {
	ID    uint    `gorm:"primary_key:auto_increment" json:"id"`
	Price float64 `gorm:"price" json:"price"`
	Goods string  `gorm:"goods" json:"goods"`
	Tenor int     `gorm:"tenor" json:"tenor"`
}
