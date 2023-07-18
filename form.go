package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	ID      int    `json:"id"`
	Link    string `json:"link"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.POST("/api/add_contact", func(c *gin.Context) {
		var contact Contact
		if err := c.ShouldBindJSON(&contact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Perform validation on contact.Name, contact.Email, and contact.Contact if needed
		// Insert the data into the database
		_, err := db.Exec("INSERT INTO contacts (link, name, email, contact) VALUES (?, ?, ?, ?)",
			contact.Link, contact.Name, contact.Email, contact.Contact)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
	})

	// Run the server
	router.Run(":8080")
}


//create a api with inputs link name, email and contact number with validation. insert data into database using MySql
