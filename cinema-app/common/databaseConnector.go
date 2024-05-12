package databaseConnector

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	orm "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() {
	db := createDatabaseConnection()
	applyMigrations(db)
	connectORM(db)
}

func createDatabaseConnection() *sql.DB {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "cinema_db")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

func applyMigrations(db *sql.DB) {
	fmt.Println("Connected! Applying migrations...")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Migrations files found")
	}

	m.Up()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	} else {
		fmt.Println("Migrations applied succesfully!")
	}
}

func connectORM(db *sql.DB) {
	gormDB, err := gorm.Open(orm.New(orm.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect GORM")
	}

	DB = gormDB
}

func GetORMConnection() *gorm.DB {
	return DB
}
