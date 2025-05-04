package models

type SKL struct {
	IDSKL         uint   `gorm:"column:id_skl;primaryKey" json:"id_skl"`
	IDSiswaFK     uint   `gorm:"column:idsiswa_fk" json:"idsiswa_fk"`
	Kode          string `json:"kode"`
	File          string `json:"file"`
	DownloadCount int    `json:"download_count"`
	Status        string `json:"status"`
	Keterangan    string `json:"keterangan"`

	Siswa Siswa `json:"siswa" gorm:"foreignKey:IDSiswaFK;references:ID"`
}

func (SKL) TableName() string {
	return "skl"
}
