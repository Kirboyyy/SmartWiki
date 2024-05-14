package domain

import (
	"html/template"
	"time"
)

type Page struct {
	ID        int           `form:"id"`
	Title     string        `form:"title"`
	Content   template.HTML `form:"content"`
	CreatedAt time.Time
}

type PageRepository interface {
	GetPage(id int) (*Page, error)
	CreatePage(page *Page) error
	UpdatePage(page *Page) (*Page, error)
	DeletePageById(id int) error
}
