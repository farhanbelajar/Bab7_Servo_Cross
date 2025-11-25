package main

import (
	"database/sql"
	"farhan_s/database"
	"farhan_s/routers"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "aan080203"
	dbName   = "Database_Bab7han"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	var PORT = ":50001"

	psqlInfo := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbName,
	)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Gagal Masuk Database", psqlInfo)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.MulaiServer().Run(PORT)
	fmt.Println("DB Terhubung")
}
