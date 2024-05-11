package main

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "cinema_db")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	fmt.Println(os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"))

	var err error
	db, err := sql.Open("postgres", psqlInfo)	
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
