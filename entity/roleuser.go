package entity

type RoleUser struct {
	ID     uint   `gorm:"primary_key:auto_increment" json:"id"`
	UserID string `gorm:"type:varchar(36);index" json:"user_id"`
	RoleID uint   `gorm:"type:varchar(36);index" json:"role_id"`
}
