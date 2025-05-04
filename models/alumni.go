package models

// Siswa defines the model for the Siswa (student) table
type Alumni struct {
	ID        uint   `gorm:"column:id_alumni;primaryKey" json:"id_alumni""`
	Nama      string `json:"nama"`
	NISN      string `json:"nisn"`
	Kode      string `json:"kode"`
	File      string `json:"file"`
	StatusSKL string `json:"status_skl"`
}

func (Alumni) TableName() string {
	return "alumni"
}
