package storage

import (
	"errors"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	error_migration = errors.New("failed make migrations to DB")
)

func (s *Storage) makeMigration() error {
	log.Println("Making migration.")
	migrationDir := "migrations"

	err := goose.Up(s.db, migrationDir)
	if err != nil {
		return error_migration
	}
	return nil
}
