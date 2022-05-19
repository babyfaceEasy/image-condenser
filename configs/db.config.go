package configs

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() (*sql.DB, error) {
	db_username := os.Getenv("DB_USERNAME")
	db_connection := os.Getenv("DB_CONNECTION")
	db_port := os.Getenv("DB_PORT")
	db_host := os.Getenv("DB_HOST")
	db_database := os.Getenv("DB_DATABASE")
	db_password := os.Getenv("DB_PASSWORD")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_password, db_host, db_port, db_database)

	db, err := sql.Open(db_connection, connString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
