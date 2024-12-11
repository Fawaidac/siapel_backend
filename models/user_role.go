package models

type UserRole struct {
	RoleID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID uint `gorm:"primaryKey;autoIncrement:false"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
