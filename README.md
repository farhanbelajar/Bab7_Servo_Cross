Buat Project
```
go mod init [nama_project]
``` 
Download 5 ibrary (module) Golang
```
go get -u "github.com/gin-gonic/gin" 
go get -u "github.com/lib/pq" 
go get -u "github.com/rubenv/sql-migrate" 
go get -u "github.com/joho/godotenv" 
go get -u "github.com/gin-contrib/cors"
```
Fungsi:
Mengatasi error CORS ketika API diakses oleh frontend (Flutter, React, JS).
Biasanya dipakai agar request dari frontend tidak diblokir oleh browser.



Codingan 1_initiate.sql

```sql
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Bab7Servo (
    code INTEGER PRIMARY KEY,
    StatusServo INTEGER
);

-- +migrate StatementEnd
```

Codingan database.go

```package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var DBMigrasi embed.FS
var DBKonesi *sql.DB

func DBMigrate(con *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: DBMigrasi,
		Root:       "sql_migrations",
	}

	n, Errs := migrate.Exec(con, "postgres", migrations, migrate.Up)
	if Errs != nil {
		panic(Errs)
	}

	DBKonesi = con

	fmt.Println("Migrasi Sukses", n, migrations)
}
```
Codingan StatusEntity.go

```
package entities

type Status struct {
	Code        int `json:"id"`
	ServoStatus int `json:"srv_status"`
}
```

Codingan StatusRepo.go

```
package repositories

import (
	"database/sql"
	"farhan_s/entities"
)

func UtkString(db *sql.DB) (err error) {
	sql := "INSERT INTO Bab7Servo(Code, StatusServo) values(1,0)"
	_, err = db.Query(sql)
	return err
}

func LihatStatus(db *sql.DB) (result []entities.Status, err error) {
	sql := "SELECT * FROM Bab7Servo"
	rows, err := db.Query(sql)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Status
		err = rows.Scan(&data.Code, &data.ServoStatus)
		if err != nil {
			return
		}
		result = append(result, data)
	}
	return
}

func UbahStatus(db *sql.DB, status entities.Status) (err error) {
	sql := "UPDATE Bab7Servo SET StatusServo = $1 WHERE Code = 1"
	_, err = db.Exec(sql, status.ServoStatus)
	return
}
```

Codingan StatusController.go

```
package controllers

import (
	"farhan_s/database"
	"farhan_s/entities"
	"farhan_s/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitProj(c *gin.Context) {
	err := repositories.UtkString(database.DBKonesi)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetStatus(c *gin.Context) {
	var result gin.H
	status, err := repositories.LihatStatus(database.DBKonesi)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": status,
		}
	}
	c.JSON(http.StatusOK, result)
}

func UpdateStatus(c *gin.Context) {
	var status entities.Status
	Ganti, _ := strconv.Atoi(c.Param("Ganti"))
	status.ServoStatus = Ganti
	err := repositories.UbahStatus(database.DBKonesi, status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ServoStatus": status.ServoStatus})
}
```

Codingan StatusRouter.go

```
package routers

import (
	"farhan_s/controllers"
	"time"

	"github.com/gin-contrib/cors" // sama ini
	"github.com/gin-gonic/gin"
)

func MulaiServer() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/servo/init-proj", controllers.InitProj)
	router.GET("/servo/status", controllers.GetStatus)
	router.PUT("/servo/update/:Ganti", controllers.UpdateStatus)
	return router
}
```

Codingan main.go
```
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
```
