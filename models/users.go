package models

type Users struct {
	ID                 uint   `gorm:"column:id;primaryKey" json:"id"`
	Username           string `gorm:"column:username" json:"username"`
	Password           string `gorm:"column:password" json:"-"`
	FirstName          string `gorm:"column:first_name" json:"first_name"`
	LastName           string `gorm:"column:last_name" json:"last_name"`
	Company            string `gorm:"column:company" json:"company"`
	Phone              string `gorm:"column:phone" json:"phone"`
	Foto               string `gorm:"column:foto" json:"foto"`
	AnggotaID          uint   `gorm:"column:anggota_id" json:"anggota_id"`
	Table              string `gorm:"column:table" json:"table"`
	IsWalas            bool   `gorm:"column:is_walas" json:"is_walas"`
	IsAbsen            bool   `gorm:"column:is_absen" json:"is_absen"`
	C                  bool   `gorm:"column:c" json:"c"`
	R                  bool   `gorm:"column:r" json:"r"`
	U                  bool   `gorm:"column:u" json:"u"`
	D                  bool   `gorm:"column:d" json:"d"`
	DeviceID           string `gorm:"column:device_id" json:"device_id"`
	PembimbingPrakerin string `gorm:"column:pembimbing_prakerin" json:"pembimbing_prakerin"`

	Groups []Group `gorm:"many2many:users_groups;joinForeignKey:UserID;joinReferences:GroupID" json:"groups"`
}

type Group struct {
	ID   uint   `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
