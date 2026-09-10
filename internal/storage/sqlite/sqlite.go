package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/borntow1n/pastebin-mini/internal/lib/random"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS pastebin(
			id INTEGER PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			text TEXT NOT NULL);
			CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
		`)
	defer stmt.Close()
	if err != nil {
		fmt.Errorf("%s:%w", op, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Errorf("%s:%w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveTEXT(textToSave string) (int64, error) {
	const op = "storage.sqlite.SaveTEXT"

	stmt, err := s.db.Prepare("INSERT INTO pastebin(text, alias) VALUES(?, ?)")
	defer stmt.Close()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(textToSave, random.NewRandomString(random.NewRandomNum(5, 9)))
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
	return id, err
}
