package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn   *gorm.DB
	user     = os.Getenv("MYCAP_DB_USER")
	password = os.Getenv("MYCAP_DB_PASSWORD")
	host     = os.Getenv("MYCAP_DB_HOST")
	db       = os.Getenv("MYCAP_DB_NAME")
	port     = os.Getenv("MYCAP_DB_PORT")
)

// Connect creates a connection to database
func Connect() (err error) {
	port, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, port)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
