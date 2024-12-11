package models

type RegistrationDetail struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	IDPendaftaran uint `json:"id_pendaftaran" gorm:"primaryKey;autoIncrement:false"`
	IDPersyaratan uint `json:"id_persyaratan" gorm:"primaryKey;autoIncrement:false"`
	IDLayanan     uint `json:"id_layanan" gorm:"primaryKey;autoIncrement:false"`
}

func (RegistrationDetail) TableName() string {
	return "detail_pendaftarans"
}
