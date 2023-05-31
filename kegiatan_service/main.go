package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Kelompok_09"
)

type Kegiatan struct {
	Judul        string `json:"judul"`
	Tempat       string `json:"tempat"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalAkhir string `json:"tanggal_akhir"`
	Deskripsi    string `json:"deskripsi"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/kegiatan")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func AddKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	var kegiatan Kegiatan
	err := json.NewDecoder(r.Body).Decode(&kegiatan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform validation
	// ...

	// Add kegiatan to the database
	query := "INSERT INTO kegiatan (judul, tempat, tanggal_mulai, tanggal_akhir, deskripsi) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, kegiatan.Judul, kegiatan.Tempat, kegiatan.TanggalMulai, kegiatan.TanggalAkhir, kegiatan.Deskripsi)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kegiatan added successfully"))
}

func UpdateKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	// Update kegiatan logic
	// ...
}

func DeleteKegiatanHandler(w http.ResponseWriter, r *http.Request) {
	// Delete kegiatan logic
	// ...
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/kegiatan", AddKegiatanHandler).Methods("POST")
	r.HandleFunc("/kegiatan/{id}", UpdateKegiatanHandler).Methods("PUT")
	r.HandleFunc("/kegiatan/{id}", DeleteKegiatanHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8083", r))
}
