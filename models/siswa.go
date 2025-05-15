package models

// Siswa defines the model for the Siswa (student) table
type Siswa struct {
	ID                  uint    `gorm:"column:id_siswa;primaryKey" json:"id_siswa"`
	NISN                string  `json:"nisn"`
	Nama                string  `json:"nama"`
	IdKelasFK           int     `gorm:"column:idkelas_fk" json:"idkelas_fk"`
	IdJurusanFK         int     `gorm:"column:idjurusan_fk" json:"idjurusan_fk"`
	Agama               string  `gorm:"column:agama" json:"agama"`
	NoIjazahSekolahAsal string  `gorm:"column:no_ijazah_sekolah_asal" json:"no_ijazah_sekolah_asal"`
	NoSkhunSekolahAsal  string  `gorm:"column:no_skhun_sekolah_asal" json:"no_skhun_sekolah_asal"`
	NoUnSekolahAsal     string  `gorm:"column:no_un_sekolah_asal" json:"no_un_sekolah_asal"`
	NoKK                string  `gorm:"column:no_kk" json:"no_kk"`
	NPSNSekolahAsal     string  `gorm:"column:npsn_sekolah_asal" json:"npsn_sekolah_asal"`
	NamaSekolahAsal     string  `gorm:"column:nama_sekolah_asal" json:"nama_sekolah_asal"`
	TempatLahir         string  `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	TanggalLahir        string  `gorm:"column:tanggal_lahir" json:"tanggal_lahir"` // bisa pakai time.Time juga
	BerkebutuhanKhusus  string  `gorm:"column:berkebutuhan_khusus" json:"berkebutuhan_khusus"`
	Alamat              string  `gorm:"column:alamat" json:"alamat"`
	Dusun               string  `gorm:"column:dusun" json:"dusun"`
	RT                  string  `gorm:"column:rt" json:"rt"`
	RW                  string  `gorm:"column:rw" json:"rw"`
	Kelurahan           string  `gorm:"column:kelurahan" json:"kelurahan"`
	Foto                string  `gorm:"column:foto" json:"foto"`
	IdProvinceFK        int     `gorm:"column:idprovince_fk" json:"idprovince_fk"`
	IdCitiesFK          int     `gorm:"column:idcities_fk" json:"idcities_fk"`
	NamaAyah            string  `gorm:"column:nama_ayah" json:"nama_ayah"`
	TempatLahirAyah     string  `gorm:"column:tempat_lahir_ayah" json:"tempat_lahir_ayah"`
	TanggalLahirAyah    string  `gorm:"column:tanggal_lahir_ayah" json:"tanggal_lahir_ayah"`
	PendidikanAyah      string  `gorm:"column:pendidikan_ayah" json:"pendidikan_ayah"`
	PekerjaanAyah       string  `gorm:"column:pekerjaan_ayah" json:"pekerjaan_ayah"`
	PenghasilanAyah     string  `gorm:"column:penghasilan_ayah" json:"penghasilan_ayah"`
	NamaIbu             string  `gorm:"column:nama_ibu" json:"nama_ibu"`
	TempatLahirIbu      string  `gorm:"column:tempat_lahir_ibu" json:"tempat_lahir_ibu"`
	TanggalLahirIbu     string  `gorm:"column:tanggal_lahir_ibu" json:"tanggal_lahir_ibu"`
	PendidikanIbu       string  `gorm:"column:pendidikan_ibu" json:"pendidikan_ibu"`
	PekerjaanIbu        string  `gorm:"column:pekerjaan_ibu" json:"pekerjaan_ibu"`
	PenghasilanIbu      string  `gorm:"column:penghasilan_ibu" json:"penghasilan_ibu"`
	TinggiBadan         int     `gorm:"column:tinggi_badan" json:"tinggi_badan"`
	BeratBadan          int     `gorm:"column:berat_badan" json:"berat_badan"`
	JarakKeSekolah      int     `gorm:"column:jarak_ke_sekolah" json:"jarak_ke_sekolah"`
	WaktuKeSekolah      int     `gorm:"column:waktu_ke_sekolah" json:"waktu_ke_sekolah"`
	JumlahSaudara       int     `gorm:"column:jumlah_saudara" json:"jumlah_saudara"`
	JenisKelamin        string  `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Saldo               float64 `gorm:"column:saldo" json:"saldo"`
	IsAlumni            int     `gorm:"column:is_alumni" json:"is_alumni"`
	IdDepartmentFK      int     `gorm:"column:iddepartment_fk" json:"iddepartment_fk"`
	RFID                string  `gorm:"column:rfid" json:"rfid"`
	TahunLulus          string  `gorm:"column:tahun_lulus" json:"tahun_lulus"`
	Kelas               Kelas   `gorm:"foreignKey:IdKelasFK;references:IdKelas" json:"kelas"`
	Jurusan             Jurusan `gorm:"foreignKey:IdJurusanFK;references:IdJurusan" json:"jurusan"`
}

func (Siswa) TableName() string {
	return "siswa"
}

type Kelas struct {
	IdKelas int    `gorm:"column:id_kelas;primaryKey" json:"id_kelas"`
	Kelas   string `gorm:"column:kelas" json:"kelas"`
}

func (Kelas) TableName() string {
	return "kelas"
}

type Jurusan struct {
	IdJurusan int    `gorm:"column:id_jurusan;primaryKey" json:"id_jurusan"`
	Jurusan   string `gorm:"column:jurusan" json:"jurusan"`
	Singkatan string `gorm:"column:singkatan" json:"singkatan"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}
