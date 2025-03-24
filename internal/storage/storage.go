package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/badprinter/task_manager/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	// error_connection = errors.New("failed create connection to DB")
	error_ping = errors.New("failed ping to DB")
)

// Основная структура для работы с базой данной.
// Она нужна для того чтобы подлючаться к БД и CRUD
type Storage struct {
	db *sql.DB
}

func NewStorageAndConnect() (*Storage, error) {
	db_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	db, err := sql.Open("pgx", db_string)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, error_ping
	}

	storage := &Storage{
		db,
	}
	err = storage.makeMigration()
	if err != nil {
		return nil, err
	}
	return storage, nil

}

func (s *Storage) Close() {
	s.db.Close()
}
