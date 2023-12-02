package entity

type Role struct {
	ID         uint     `gorm:"primary_key:auto_increment" json:"id"`
	Role       string   `gorm:"role" json:"role"`
	RoleUserID RoleUser `json:"role_user" gorm:"foreignKey:RoleID"`
}
