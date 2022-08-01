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
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := sm.DB.QueryRow(query, id)

	record := &models.Snippet{}

	err := row.Scan(&record.ID, &record.Title, &record.Content, &record.Created, &record.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return record, nil
}

// This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := sm.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		record := &models.Snippet{}
		err := rows.Scan(&record.ID, &record.Title, &record.Content, &record.Created, &record.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snippets, nil
}
