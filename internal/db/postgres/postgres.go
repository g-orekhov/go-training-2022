package postgres

import (
	"os"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func OpenDB() (db.Session, error) {
	//TODO: recieve hardcoded params from function caller
	var settings = postgresql.ConnectionURL{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DBNAME"),
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
