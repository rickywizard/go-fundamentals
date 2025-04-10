package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := os.Getenv("DB_CONN")

	DB, err = sql.Open("mysql", dbConnection)

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Println("Connected to database")

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		UNIQUE (email)
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Fatal("Error creating users tables", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		log.Fatal("Error creating events tables", err)
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		log.Fatal("Error creating registrations tables", err)
	}
}
