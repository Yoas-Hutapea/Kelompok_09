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

type LoginRequest struct {
	NIK      string `json:"nik"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform login logic
	// ...

	// Check if the user exists in the database
	query := "SELECT COUNT(*) FROM users WHERE nik = ? AND password = ?"
	var count int
	err = db.QueryRow(query, request.NIK, request.Password).Scan(&count)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var response LoginResponse
	if count > 0 {
		response.Success = true
		response.Message = "Login successful"
	} else {
		response.Success = false
		response.Message = "Invalid credentials"
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
