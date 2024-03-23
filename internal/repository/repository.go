package book_inventory_system_repository

import (
	domain "book-inventory-system/internal/domain"
	"errors"
	"fmt"
	"sync"
)

type Repository struct {
	mu         *sync.Mutex
	admins     map[int]domain.AdminMapField
	books      map[int]domain.BookMapField
	genres     map[int]domain.GenreMapField
	language   map[int]domain.LanguageMapField
	author     map[int]domain.AuthorMapField
	production map[int]domain.ProductionMapField
	instance   map[int]domain.InstanceMapField
	user       map[int]domain.UserMapField
	reader     map[int]domain.ReaderMapField
}

func New(opts ...Option) (*Repository, error) {
	errs := make([]error, 0)
	r := new(Repository)
	r.mu = new(sync.Mutex)

	r.admins = make(map[int]domain.AdminMapField)
	r.books = make(map[int]domain.BookMapField)
	r.genres = make(map[int]domain.GenreMapField)
	r.language = make(map[int]domain.LanguageMapField)
	r.author = make(map[int]domain.AuthorMapField)
	r.production = make(map[int]domain.ProductionMapField)
	r.instance = make(map[int]domain.InstanceMapField)
	r.user = make(map[int]domain.UserMapField)
	r.reader = make(map[int]domain.ReaderMapField)

	for _, opt := range opts {
		err := opt(r)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return r, errors.Join(errs...)
}

func (r *Repository) ReturnBook(id int) error {
	r.mu.Lock()

	instance, ok := r.instance[id]
	if !ok {
		return fmt.Errorf("record not found")
	}

	instance.Status = 1

	r.mu.Unlock()

	return nil
}
