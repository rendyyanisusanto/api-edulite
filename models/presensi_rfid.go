package models

// Guru defines the model for the Guru (teacher) table
type PresensiRFID struct {
	ID        uint   `gorm:"column:id_presensi_rfid;primaryKey" json:"id_presensi_rfid"`
	Tanggal   string `json:"tanggal"`
	Waktu     string `json:"waktu"`
	Status    string `json:"status"`
	IDSiswaFK uint   `gorm:"column:idsiswa_fk" json:"idsiswa_fk"`
}

func (PresensiRFID) TableName() string {
	return "presensi_rfid"
}
