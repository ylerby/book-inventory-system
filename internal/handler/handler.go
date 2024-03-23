package book_inventory_system_handler

import (
	logger "book-inventory-system/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type service interface {
	ReturnBook(id int) error
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
		_, err = ctx.Writer.Write([]byte("internal server error"))
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
