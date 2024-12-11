package models

import "time"

type Users struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	IDKecamatan      uint      `json:"id_kecamatan" binding:"required"`
	IDKelurahan      uint      `json:"id_kelurahan" binding:"required"`
	Name             string    `json:"name" binding:"required"`
	Email            string    `json:"email" gorm:"unique" binding:"required,email"`
	EmailVerifiedAt  *time.Time `json:"email_verified_at"`
	Password         string    `json:"-" binding:"required,min=6"`
	NIK              string    `json:"nik" binding:"required"`
	Phone            string    `json:"phone" binding:"required"`
	Avatar           string    `json:"avatar"`
	Hit              int       `json:"hit"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
func (Users) TableName() string {
    return "users"
}
