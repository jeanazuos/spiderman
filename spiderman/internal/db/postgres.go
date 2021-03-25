package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func CreatePool() (*pgxpool.Pool, error) {

	connectionString := "postgres://esaizziy:nIBPRvEuqVsdw6DQywfu_4L22nKuRVps@tuffi.db.elephantsql.com:5432/esaizziy"
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func Init() {
	pool, err := CreatePool()
	if err != nil {
		log.Fatal("Unable to create connection pool", "error", err)
		os.Exit(1)
	}
	db = pool
}

func GetDb() *pgxpool.Pool {
	if db == nil {
		CreatePool()
	}

	return db
}
