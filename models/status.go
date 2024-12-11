package models

type Status struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IDLayanan   uint   `json:"id_layanan"`
	Sequence    uint   `json:"sequence"`
}

func (Status) TableName() string {
	return "status"
}