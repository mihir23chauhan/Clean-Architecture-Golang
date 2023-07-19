package book

import (
	"github.com/mihirtunnel/cleanArchitecture/internal/domain/model"
	"github.com/mihirtunnel/cleanArchitecture/internal/domain/usecase"
)

type BookController struct {
	repo usecase.BookUseCase
}

func NewBookController(repo usecase.BookUseCase) *BookController {
	return &BookController{repo: repo}
}

func (uc *BookController) GetAllBooks() ([]model.Book, error) {
	return uc.repo.GetAllBooks()
	
}

func (uc *BookController) GetBookByID(id int) (*model.Book, error) {
	return uc.repo.GetBookByID(id)
}

func (uc *BookController) AddBook(book *model.Book) (*model.Book, error) {
	
	return uc.repo.AddBook(book)
}

func (uc *BookController) UpdateBook(book *model.Book) (*model.Book, error) {
	return uc.repo.UpdateBook(book)
}

func (uc *BookController) DeleteBook(id int) error {
	return uc.repo.DeleteBook(id)
}
