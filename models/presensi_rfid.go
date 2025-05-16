package models

// Guru defines the model for the Guru (teacher) table
type PresensiRFID struct {
	ID        uint   `gorm:"column:id_presensi_rfid;primaryKey" json:"id_presensi_rfid"`
	Tanggal   string `json:"tanggal"`
	Waktu     string `json:"waktu"`
	Status    string `json:"status"`
	IDSiswaFK uint   `gorm:"column:idsiswa_fk" json:"idsiswa_fk"`
}

type RekapPresensi struct {
	Tanggal          string `json:"tanggal"`
	IDSiswaFK        uint   `gorm:"column:idsiswa_fk" json:"idsiswa_fk"`
	Masuk            string `json:"masuk"`
	Pulang           string `json:"pulang"`
	JumlahIjinKeluar int    `json:"jumlah_ijin_keluar"`
}

type JumlahPresensi struct {
	IDSiswaFK   uint `gorm:"column:idsiswa_fk" json:"idsiswa_fk"`
	JumlahMasuk int  `json:"jumlah_masuk"`
}

func (PresensiRFID) TableName() string {
	return "presensi_rfid"
}
