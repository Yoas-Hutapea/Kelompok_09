package models

type User struct {
	Nama         string `json:"nama"`
	NIK          string `json:"nik"`
	NoTelp       string `json:"no_telp"`
	Alamat       string `json:"alamat"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Usia         int    `json:"usia"`
	JenisKelamin string `json:"jenis_kelamin"`
	Pekerjaan    string `json:"pekerjaan"`
	Agama        string `json:"agama"`
	KK           string `json:"kk"`
	Gambar       string `json:"gambar"`
	Password     string `json:"password"`
}
