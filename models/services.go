package models

import "time"

type Service struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Code      string    `json:"code"`
    Sect      string    `json:"sect"`
    Name      string    `json:"name"`
    Detail    string    `json:"detail"`
    Slug      string    `json:"slug"`
    Image     string    `json:"image"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (Service) TableName() string {
    return "layanans"
}