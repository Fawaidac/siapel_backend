package utils

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Hit  int
}

func GenerateOperator(db *gorm.DB, roleName string) (uint, error) {
	if roleName == "" {
		return 0, errors.New("role name is required")
	}

	var user User

	err := db.Table("users").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.name = ?", roleName).
		Where("users.status = ?", "aktif").
		Order("users.hit ASC").
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("no active user found with the specified role")
		}
		return 0, err
	}

	err = db.Table("users").Where("id = ?", user.ID).Update("hit", gorm.Expr("hit + 1")).Error
	if err != nil {
		return 0, errors.New("failed to update user hit count")
	}

	return user.ID, nil
}