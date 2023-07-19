package database

import (
	"database/sql"
	"log"

	"github.com/mihirtunnel/cleanArchitecture/internal/domain/model"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) GetAllBooks() ([]model.Book, error) {
	query := "SELECT * FROM books"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Failed to fetch books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear)
		if err != nil {
			log.Printf("Failed to scan book row: %v", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) GetBookByID(id int) (*model.Book, error) {
	query := "SELECT * FROM books WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var book model.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Book not found
		}
		log.Printf("Failed to fetch book by ID: %v", err)
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) AddBook(book *model.Book) (*model.Book, error) {
	query := "INSERT INTO books (title, author, publicationYear) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, book.Title, book.Author, book.PublicationYear)
	if err != nil {
		log.Printf("Failed to insert book: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)

	return book, nil
}

func (r *BookRepository) UpdateBook(book *model.Book) (*model.Book, error) {
	originalBook, err := r.GetBookByID(book.ID)
	if err != nil {
		log.Printf("Failed to update book: %v", err)
		return nil, err
	}
	if originalBook == nil {
		return originalBook, err
	}

	if book.Author == "" {
		book.Author = originalBook.Author
	}
	if book.PublicationYear == 0 {
		book.PublicationYear = originalBook.PublicationYear
	}
	if book.Title == "" {
		book.Title = originalBook.Title
	}

	query := "UPDATE books SET title = ?, author = ?, publicationYear = ? WHERE id = ?"

	_, err = r.DB.Exec(query, book.Title, book.Author, book.PublicationYear, book.ID)

	if err != nil {
		log.Printf("Failed to update book: %v", err)
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete book: %v", err)
		return err
	}

	return nil
}
