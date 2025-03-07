package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"wanderloge/database"
	"wanderloge/router"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	database.DBMigrate(DB)

	r := router.SetupRouter()
	r.Run(":8080")
}
