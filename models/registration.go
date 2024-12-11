package models

import "time"

type Registration struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID      uint    `json:"user_id"`
    OperatorID    uint    `json:"operator_id"`
    Tiket      string    `json:"slug" gorm:"unique"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (Registration) TableName() string {
    return "pendaftarans"
}