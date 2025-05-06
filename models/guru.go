package models

// Guru defines the model for the Guru (teacher) table
type Guru struct {
	ID   uint   `gorm:"column:id_guru;primaryKey" json:"id_guru"`
	NISN string `json:"nisn"`
	Nama string `json:"nama"`
}

func (Guru) TableName() string {
	return "guru"
}
