package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

// Initialization of PostgreSQL Database
func InitDb() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file, please recheck!")
		log.Fatal(err)
	}

	// Read .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Set up postgres sql to open it
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	Db, err = sql.Open("postgres", psqlSetup)
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Database!")

	// Set up tables in database
	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS account (
						user_id serial PRIMARY KEY,
						username VARCHAR(50) UNIQUE NOT NULL,
						password VARCHAR(50) NOT NULL,
						name VARCHAR(50),
						surname VARCHAR(50),
						birthday DATE,
						email VARCHAR(255),
						phone VARCHAR(20))`)
	if err != nil {
		fmt.Println("Couldn't create an account table, recheck you request!")
		log.Fatal(err)
	}
}