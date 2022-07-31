package mysql

import (
	"database/sql"

	"github.com/talosrobert/golang-practice/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (sm *SnippetModel) Insert(title, content, expires string) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := sm.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (sm *SnippetModel) Get(id int) (*models.Snippet, error) {

	return nil, nil
}

// This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
