package models

type Requirement struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    uint   `json:"name"`
	Data    uint   `json:"data"`
	IsValid string `json:"is_valid"`
}

func (Requirement) TableName() string {
	return "persyaratans"
}