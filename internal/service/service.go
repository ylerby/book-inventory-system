package book_inventory_system_service

import (
	config "book-inventory-system/internal/config"
	domain "book-inventory-system/internal/domain"
	logger "book-inventory-system/pkg/logger"
)

type repository interface {
	ReturnBook(id int) error
	TakeBook(id int) (*domain.BookMapField, error)
	UpdateLoginStatus(id int, status string) error
	BanUser(userID, adminID int) error
	UpdateInstanceStatus(instanceID, status int) error
	CheckAvailability(instanceID int) (bool, error)
	CountPublishedBooks(authorID int) (int, error)
	CheckBorrowBooks(readerID int) ([]domain.BookMapField, error)
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

func (s *Service) TakeBook(id int) (*domain.BookMapField, error) {
	book, err := s.r.TakeBook(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) UpdateLoginStatus(id int, status string) error {
	err := s.r.UpdateLoginStatus(id, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) BanUser(userID, adminID int) error {
	err := s.r.BanUser(userID, adminID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateInstanceStatus(instanceID, status int) error {
	err := s.r.UpdateInstanceStatus(instanceID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckAvailability(instanceID int) (bool, error) {
	isAvailable, err := s.r.CheckAvailability(instanceID)
	if err != nil {
		return false, err
	}

	return isAvailable, nil
}

func (s *Service) CountPublishedBooks(authorID int) (int, error) {
	count, err := s.r.CountPublishedBooks(authorID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Service) CheckBorrowBooks(readerID int) ([]domain.BookMapField, error) {
	books, err := s.r.CheckBorrowBooks(readerID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
