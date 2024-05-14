package service

import "github.com/kirboyyy/smartwiki/internal/domain"

type PageService struct {
	PageRepo domain.PageRepository
}

func (ps *PageService) GetPage(id int) (*domain.Page, error) {
	return ps.PageRepo.GetPage(id)
}

func (ps *PageService) CreatePage(page *domain.Page) error {
	return ps.PageRepo.CreatePage(page)
}

func (ps *PageService) UpdatePage(page *domain.Page) (*domain.Page, error) {
	return ps.PageRepo.UpdatePage(page)
}

func (ps *PageService) DeletePageById(id int) error {
	return ps.PageRepo.DeletePageById(id)
}
