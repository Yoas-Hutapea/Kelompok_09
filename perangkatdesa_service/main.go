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

type PerangkatDesa struct {
	Nama    string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/perangkat")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func AddPerangkatDesaHandler(w http.ResponseWriter, r *http.Request) {
	var perangkat PerangkatDesa
	err := json.NewDecoder(r.Body).Decode(&perangkat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform validation
	// ...

	// Add perangkat desa to the database
	query := "INSERT INTO perangkat (nama, jabatan) VALUES (?, ?)"
	_, err = db.Exec(query, perangkat.Nama, perangkat.Jabatan)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perangkat Desa added successfully"))
}

func UpdatePerangkatDesaHandler(w http.ResponseWriter, r *http.Request) {
	// Update perangkat desa logic
	// ...
}

func DeletePerangkatDesaHandler(w http.ResponseWriter, r *http.Request) {
	// Delete perangkat desa logic
	// ...
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/perangkatdesa", AddPerangkatDesaHandler).Methods("POST")
	r.HandleFunc("/perangkatdesa/{id}", UpdatePerangkatDesaHandler).Methods("PUT")
	r.HandleFunc("/perangkatdesa/{id}", DeletePerangkatDesaHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", r))
}
