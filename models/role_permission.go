package models

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey;autoIncrement:false"`
	PermissionID uint `gorm:"primaryKey;autoIncrement:false"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
