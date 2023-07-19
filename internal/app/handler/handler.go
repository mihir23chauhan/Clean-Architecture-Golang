package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mihirtunnel/cleanArchitecture/internal/app/book"
	"github.com/mihirtunnel/cleanArchitecture/internal/domain/model"
)

type BookHandler struct {
	bookController *book.BookController
}

func NewBookHandler(bookController *book.BookController) *BookHandler {
	return &BookHandler{bookController: bookController}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookController.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.bookController.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book by ID"})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newBook, err := h.bookController.AddBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book.ID = bookID

	updatedBook, err := h.bookController.UpdateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	if updatedBook == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = h.bookController.DeleteBook(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
