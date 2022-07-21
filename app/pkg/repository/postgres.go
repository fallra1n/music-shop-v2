package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	cTables := NewCrateTables(db)
	if err := cTables.CreateAllTables(); err != nil {
		logrus.Fatalf("error init tables: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
