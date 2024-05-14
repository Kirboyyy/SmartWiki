package sqlite

import (
	"database/sql"

	"github.com/kirboyyy/smartwiki/internal/domain"
)

type PageRepository struct {
	DB *sql.DB
}

func (r *PageRepository) GetPage(id int) (*domain.Page, error) {
	row := r.DB.QueryRow("SELECT id, title, content, created_at FROM pages WHERE id = ?", id)
	var page domain.Page
	err := row.Scan(&page.ID, &page.Title, &page.Content, &page.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *PageRepository) CreatePage(page *domain.Page) error {
	stmt, err := r.DB.Prepare("INSERT INTO pages(title, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(page.Title, page.Content)
	return err
}

func (r *PageRepository) UpdatePage(page *domain.Page) (*domain.Page, error) {
	stmt, err := r.DB.Prepare("UPDATE pages SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(page.Title, page.Content, page.ID)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *PageRepository) DeletePageById(id int) error {
	stmt, err := r.DB.Prepare("DELETE pages WHERE id = ?")
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
