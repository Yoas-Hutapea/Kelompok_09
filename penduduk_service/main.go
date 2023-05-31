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

type User struct {
	Nama     string `json:"nama"`
	NIK      string `json:"nik"`
	NoTelp   string `json:"no_telp"`
	Alamat   string `json:"alamat"`
	Password string `json:"password"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/users")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform validation
	// ...

	// Add user to the database
	query := "INSERT INTO users (nama, nik, no_telp, alamat, password) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, user.Nama, user.NIK, user.NoTelp, user.Alamat, user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User added successfully"))
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Update user logic
	// ...
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Delete user logic
	// ...
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", AddUserHandler).Methods("POST")
	r.HandleFunc("/users/{nik}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{nik}", DeleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}
