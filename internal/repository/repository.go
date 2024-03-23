package book_inventory_system_repository

import (
	domain "book-inventory-system/internal/domain"
	"errors"
	"fmt"
	"sync"
)

const (
	inUse int = iota
	inLibrary
	outOfUser
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
	defer r.mu.Unlock()

	instance, ok := r.instance[id]
	if !ok {
		return fmt.Errorf("instance not found")
	}

	switch instance.Status {
	case inUse:
		instance.Status = inLibrary
		r.instance[id] = instance
		return nil
	default:
		return fmt.Errorf("instance already in library")
	}
}

func (r *Repository) TakeBook(id int) (*domain.BookMapField, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	instance, ok := r.instance[id]
	if !ok {
		return nil, fmt.Errorf("instance not found")
	}

	switch instance.Status {
	case inLibrary:
		for bookID := range r.books {
			if instance.BookID == bookID {
				book := r.books[bookID]
				instance.Status = inUse
				r.instance[id] = instance
				return &book, nil
			}
		}
	default:
		return nil, fmt.Errorf("you can`t take an instance")
	}

	return nil, nil
}

func (r *Repository) UpdateLoginStatus(id int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.user[id]
	if !ok {
		return fmt.Errorf("user not found")
	}

	user.LoginStatus = status
	r.user[id] = user

	return nil
}

//todo: black list(map[int]struct{})

func (r *Repository) BanUser(userID, adminID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.admins[adminID]
	if !ok {
		return fmt.Errorf("invalid admin id")
	}

	_, ok = r.admins[adminID]
	if !ok {
		return fmt.Errorf("permission denied")
	}

	_, ok = r.user[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}

	delete(r.user, userID)
	return nil
}

func (r *Repository) UpdateInstanceStatus(instanceID, status int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	instance, ok := r.instance[instanceID]
	if !ok {
		return fmt.Errorf("instance not found")
	}

	instance.Status = status
	r.instance[instanceID] = instance

	return nil
}

func (r *Repository) CheckAvailability(instanceID int) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	instance, ok := r.instance[instanceID]
	if !ok {
		return false, fmt.Errorf("instance not found")
	}

	switch instance.Status {
	case inLibrary:
		return true, nil
	default:
		return false, nil
	}
}

func (r *Repository) CountPublishedBooks(authorID int) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.author[authorID]
	if !ok {
		return 0, fmt.Errorf("author not found")
	}

	var counter int

	for _, book := range r.books {
		if book.AuthorID == authorID {
			counter++
		}
	}

	return counter, nil
}

func (r *Repository) CheckBorrowBooks(readerID int) ([]domain.BookMapField, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reader, ok := r.reader[readerID]
	if !ok {
		return nil, fmt.Errorf("reader not found")
	}

	books := make([]domain.BookMapField, 0)

	for _, instance := range reader.InstanceID {
		for id := range r.instance {
			if instance == id {
				bookID := r.instance[id].BookID
				book := r.books[bookID]
				books = append(books, book)
			}
		}
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("empty books list")
	}

	return books, nil
}
