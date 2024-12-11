package models

import "time"

type Village struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    IDKecamatan uint     `json:"id_kecamatan"`
    Name       string    `json:"name"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
