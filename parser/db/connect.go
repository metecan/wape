package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	// Creating the table if it is not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS titles (keyword varchar NOT NULL, title varchar NOT NULL, epoch int NOT NULL)")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
