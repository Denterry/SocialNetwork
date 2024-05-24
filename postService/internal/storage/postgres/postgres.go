package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Denterry/SocialNetwork/postService/internal/config"
	_ "github.com/lib/pq"
)

// Initialization of PostgreSQL Database
func InitDb(dbConfig *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbConfig.Storage.Host, dbConfig.Storage.User, dbConfig.Storage.Password, dbConfig.Storage.Name, dbConfig.Storage.Port)

	Db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Database!")

	return Db, nil
}
