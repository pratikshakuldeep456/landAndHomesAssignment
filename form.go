package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact int64  `json:"contact"`
}

func identifyHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "postgresql://username:password@localhost:5432/database_name?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Contact WHERE name =$1 OR  email = $2 OR contact = $3", requestBody.Email, requestBody.Contact)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func main() {

	http.HandleFunc("/identify", identifyHandler)

	println("hi0")
	log.Fatal(http.ListenAndServe(":8080", nil))

	println("hi10")
}

//create a form with inputs linke name, email and contact number with validation. insert data into database using MySql.(use any preferred technology
