package server

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Server struct {
	pgDB *sql.DB
	// dbManager lib.DBManager
}

func NewServer(config *Config) (*Server, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	dbConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "vikas.rao", "root", "go-fiber-api")
	pgDB, err := sql.Open("postgres", dbConnection)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &Server{
		pgDB: pgDB,
	}, nil
}
