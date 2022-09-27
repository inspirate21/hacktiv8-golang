package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "postgres"
	dbname = "hacktiv8-golang"
)

var DB *sql.DB

func ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetPostgres() *sql.DB {
	return DB
}
