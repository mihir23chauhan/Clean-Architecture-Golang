package usecase

import "github.com/mihirtunnel/cleanArchitecture/internal/domain/model"

type BookUseCase interface {
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id int) (*model.Book, error)
	AddBook(book *model.Book) (*model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	DeleteBook(id int) error
}
