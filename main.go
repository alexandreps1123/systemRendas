package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal()
	} else {
		fmt.Println()
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	ROUTER := mux.NewRouter()

	DB, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		fmt.Printf("Cannot connect to database ", dbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", dbDriver)
	}

	defer DB.Close()
}
