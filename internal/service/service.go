package book_inventory_system_service

import (
	config "book-inventory-system/internal/config"
	logger "book-inventory-system/pkg/logger"
)

type repository interface {
	ReturnBook(id int) error
}

type Service struct {
	r   repository
	l   logger.Logger
	cfg *config.Config
}

func New(
	r repository,
	l logger.Logger,
	cfg *config.Config,
) *Service {
	return &Service{
		r:   r,
		l:   l,
		cfg: cfg,
	}
}

func (s *Service) ReturnBook(id int) error {
	err := s.r.ReturnBook(id)
	if err != nil {
		return err
	}

	return nil
}
