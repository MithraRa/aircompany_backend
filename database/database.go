package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Host := os.Getenv("db_host")
	Port := os.Getenv("db_port")
	Password := os.Getenv("db_pass")
	User := os.Getenv("db_user")
	DBName := os.Getenv("db_name")
	SSLMode := os.Getenv("db_sslmode")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Host, Port, User, Password, DBName, SSLMode,
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Something went wrong with Open()", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Something went wrong with Ping()", err)
	}
	log.Println("Successfully connected to Postgres")
}

func GetDB() *sql.DB {
	return db
}
