package models

// Siswa defines the model for the Siswa (student) table
type Siswa struct {
	ID   uint   `gorm:"column:id_siswa;primaryKey" json:"id_siswa"`
	NISN string `json:"nisn"`
	Nama string `json:"nama"`
}

func (Siswa) TableName() string {
	return "siswa"
}
