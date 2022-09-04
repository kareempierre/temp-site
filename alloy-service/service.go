package alloy

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var Pool *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	var (
		dbUser    = os.Getenv("DB_USER")
		dbPwd     = os.Getenv("DB_PASS")
		dbTCPHost = os.Getenv("INSTANCE_HOST")
		dbPort    = os.Getenv("DB_PORT")
		dbName    = os.Getenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}

	Pool = dbPool
}
