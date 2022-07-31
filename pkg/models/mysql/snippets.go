package mysql

import (
	"database/sql"

	"github.com/talosrobert/golang-practice/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (sm *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (sm *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
