package models

type Permission struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Roles []Role `gorm:"many2many:role_permissions;"`
}

func (Permission) TableName() string {
	return "permissions"
}