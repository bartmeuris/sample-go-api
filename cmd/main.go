package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bartmeuris/sample-go-api/api"
	"github.com/bartmeuris/sample-go-api/models"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
	//swagger "github.com/emicklei/go-restful-swagger12"
)

func openDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./testdb.db")
	if err != nil {
		return nil
	}
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\n", 0))
	return db
}

func main() {
	db := openDb()
	defer db.Close()
	if err := models.Migrate(db); err != nil {
		log.Panicf("Error migrating database: %s", err)
	}
	wsContainer, err := api.RegisterAll(db)
	if err != nil {
		log.Panicf("Error creating API: %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", wsContainer)

	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swagger-ui/"))))

	server := &http.Server{Addr: "localhost:8080", Handler: mux}
	log.Printf("Waiting for connections...")
	log.Fatal(server.ListenAndServe())
}
