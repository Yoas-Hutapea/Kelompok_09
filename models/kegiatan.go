package models

type Kegiatan struct {
	Judul        string `json:"judul"`
	Tempat       string `json:"tempat"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalAkhir string `json:"tanggal_akhir"`
	Deskripsi    string `json:"deskripsi"`
}
