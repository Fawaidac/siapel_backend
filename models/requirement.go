package models

type Requirement struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string   `json:"name"`
	Data    string   `json:"data"`
	IsValid string `json:"is_valid"`
}

func (Requirement) TableName() string {
	return "persyaratans"
}