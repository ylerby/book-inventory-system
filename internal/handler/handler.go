package book_inventory_system_handler

import (
	domain "book-inventory-system/internal/domain"
	logger "book-inventory-system/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
)

type service interface {
	ReturnBook(id int) error
	TakeBook(id int) (*domain.BookMapField, error)
	UpdateLoginStatus(id int, status string) error
	BanUser(userID, adminID int) error
	UpdateInstanceStatus(instanceID, status int) error
	CheckAvailability(instanceID int) (bool, error)
	CountPublishedBooks(authorID int) (int, error)
	CheckBorrowBooks(readerID int) ([]domain.BookMapField, error)
}

type Handler struct {
	l logger.Logger
	s service
}

func New(l logger.Logger, s service) *Handler {
	return &Handler{
		l: l,
		s: s,
	}
}

func (h *Handler) InitRoutes(address string, ch chan error) {
	router := gin.Default()
	router.POST("/", h.main)
	router.GET("/return_book", h.returnBook)
	router.GET("/take_book", h.takeBook)
	router.GET("/update_login_status", h.updateLoginStatus)
	router.GET("/ban_user", h.banUser)
	router.GET("/update_instance_status", h.updateInstanceStatus)
	router.GET("/check_availability", h.checkAvailability)
	router.GET("/check_borrow_books", h.checkBorrowBooks)

	err := router.Run(address)
	if err != nil {
		ch <- err
		return
	}
}

func (h *Handler) main(ctx *gin.Context) {
	_, err := ctx.Writer.Write([]byte("main page"))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) returnBook(ctx *gin.Context) {
	bookID := ctx.Query("book_id")

	intBookID, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	err = h.s.ReturnBook(intBookID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write([]byte("the book has been returned"))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) takeBook(ctx *gin.Context) {
	bookID := ctx.Query("book_id")

	intBookID, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	book, err := h.s.TakeBook(intBookID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write(response)
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) updateLoginStatus(ctx *gin.Context) {
	loginStatus := ctx.Query("login_status")
	userID := ctx.Query("user_id")

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	err = h.s.UpdateLoginStatus(intUserID, loginStatus)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write([]byte("login status has been updated"))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) banUser(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	adminID := ctx.Query("admin_id")

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	intAdminID, err := strconv.Atoi(adminID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	err = h.s.BanUser(intUserID, intAdminID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write([]byte("user has been banned"))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) updateInstanceStatus(ctx *gin.Context) {
	instanceStatus := ctx.Query("instance_status")
	instanceID := ctx.Query("instance_id")

	intInstanceStatus, err := strconv.Atoi(instanceStatus)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	intInstanceID, err := strconv.Atoi(instanceID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	err = h.s.UpdateInstanceStatus(intInstanceID, intInstanceStatus)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write([]byte("instance status has been updated"))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) checkAvailability(ctx *gin.Context) {
	instanceID := ctx.Query("instance_id")

	intInstanceID, err := strconv.Atoi(instanceID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	isAvailable, err := h.s.CheckAvailability(intInstanceID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write([]byte(fmt.Sprintf("is available: %t", isAvailable)))
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}

func (h *Handler) checkBorrowBooks(ctx *gin.Context) {
	readerID := ctx.Query("reader_id")

	intReaderID, err := strconv.Atoi(readerID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte("internal server error"))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	books, err := h.s.CheckBorrowBooks(intReaderID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		if err != nil {
			h.l.Errorf("response error: %v", err)
			return
		}

		return
	}

	ctx.Status(http.StatusOK)
	_, err = ctx.Writer.Write(response)
	if err != nil {
		h.l.Errorf("response error: %v", err)
		return
	}
}
