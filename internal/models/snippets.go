package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expire  time.Time
}


type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO SNIPPETS(title, content,created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {

		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT * FROM SNIPPETS
			WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)
	s := &Snippet{}
	err := row.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expire)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
			WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	row, err := m.DB.Query(stmt)

	snippet := []*Snippet{}
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &Snippet{}
		err = row.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expire)
		if err != nil {
			return nil, err
		}

		snippet = append(snippet, s)
	}
	if row.Err() != nil {
		return nil, err
	}
	return snippet, nil
}
