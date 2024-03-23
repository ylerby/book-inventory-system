package book_inventory_system_repository

import (
	domain "book-inventory-system/internal/domain"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

type Option func(repository *Repository) error

func WithDump(
	adminDumpFilePath,
	authorDumpFilePath,
	bookDumpFilePath,
	genreDumpFilePath,
	instanceDumpFilePath,
	languageDumpFilePath,
	productionDumpFilePath,
	readerDumpFilePath,
	userDumpFilePath string,
) Option {
	return func(r *Repository) error {
		admins, err := adminDump(adminDumpFilePath)
		if err != nil {
			return fmt.Errorf("admins.json dump error: %w", err)
		}

		for _, admin := range admins.Admins {
			r.admins[admin.AdminID] = domain.AdminMapField{}
		}

		authors, err := authorDump(authorDumpFilePath)
		if err != nil {
			return fmt.Errorf("authors.json dump error: %w", err)
		}

		for _, author := range authors.Authors {
			r.author[author.AuthorID] = domain.AuthorMapField{
				Name:         author.Name,
				Surname:      author.Surname,
				Patronymic:   author.Patronymic,
				ProductionID: author.ProductionID,
			}
		}

		books, err := bookDump(bookDumpFilePath)
		if err != nil {
			return fmt.Errorf("books.json dump error: %w", err)
		}

		for _, book := range books.Books {
			r.books[book.BookID] = domain.BookMapField{
				Name:         book.Name,
				AuthorID:     book.AuthorID,
				GenreID:      book.GenreID,
				ProductionID: book.ProductionID,
				LanguageID:   book.LanguageID,
				Description:  book.Description,
			}
		}

		genres, err := genreDump(genreDumpFilePath)
		if err != nil {
			return fmt.Errorf("genres.json dump error: %w", err)
		}

		for _, genre := range genres.Genres {
			r.genres[genre.GenreID] = domain.GenreMapField{
				Name: genre.Name,
			}
		}

		instances, err := instanceDump(instanceDumpFilePath)
		if err != nil {
			return fmt.Errorf("instances.json dump error: %w", err)
		}

		for _, instance := range instances.Instances {
			r.instance[instance.InstanceID] = domain.InstanceMapField{
				BookID: instance.BookID,
				Status: instance.Status,
			}
		}

		languages, err := languageDump(languageDumpFilePath)
		if err != nil {
			return fmt.Errorf("languages.json dump error: %w", err)
		}

		for _, language := range languages.Languages {
			r.language[language.LanguageID] = domain.LanguageMapField{
				Name: language.Name,
			}
		}

		productions, err := productionDump(productionDumpFilePath)
		if err != nil {
			return fmt.Errorf("productions.json dump error: %w", err)
		}

		for _, production := range productions.Productions {
			r.production[production.ProductionID] = domain.ProductionMapField{
				Name: production.Name,
			}
		}

		readers, err := readerDump(readerDumpFilePath)
		if err != nil {
			return fmt.Errorf("readers.json dump error: %w", err)
		}

		for _, reader := range readers.Readers {
			r.reader[reader.ReaderID] = domain.ReaderMapField{
				InstanceID: reader.InstanceID,
			}
		}

		users, err := userDump(userDumpFilePath)
		if err != nil {
			return fmt.Errorf("users.json dump error: %w", err)
		}

		for _, user := range users.Users {
			r.user[user.UserID] = domain.UserMapField{
				Name:         user.Name,
				Password:     user.Password,
				LoginStatus:  user.LoginStatus,
				RegisterDate: user.RegisterDate,
			}
		}

		return nil
	}
}

func adminDump(filepath string) (*domain.Admin, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Admin
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func authorDump(filepath string) (*domain.Author, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Author
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func bookDump(filepath string) (*domain.Book, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Book
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func genreDump(filepath string) (*domain.Genre, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Genre
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func instanceDump(filepath string) (*domain.Instance, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Instance
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func languageDump(filepath string) (*domain.Language, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Language
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func productionDump(filepath string) (*domain.Production, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Production
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func readerDump(filepath string) (*domain.Reader, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.Reader
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}

func userDump(filepath string) (*domain.User, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dump domain.User
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}

	return &dump, nil
}
