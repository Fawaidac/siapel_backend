package models

type RegistrationStatusDetail struct {
	IDPendaftaran uint `json:"id_pendaftaran" gorm:"primaryKey;autoIncrement:false"`
	IDStatus      uint `json:"id_status" gorm:"primaryKey;autoIncrement:false"`
}

func (RegistrationStatusDetail) TableName() string {
	return "detail_pendaftaran_status"
}
