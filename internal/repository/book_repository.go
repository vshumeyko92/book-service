package repository

import (
	"errors"

	"github.com/vshumeyko92/book-service/internal/model"
)

// BookRepository представляет репозиторий для работы с данными книг.
type BookRepository struct {
	books []*model.Book
}

// NewBookRepository создает новый экземпляр BookRepository.
func NewBookRepository() *BookRepository {
	// Здесь можно добавить логику для инициализации репозитория,
	// например, загрузку данных из базы данных.
	return &BookRepository{
		books: make([]*model.Book, 0),
	}
}

// GetBooks возвращает список всех книг.
func (r *BookRepository) GetBooks() ([]*model.Book, error) {
	return r.books, nil
}

// GetAuthors возвращает список всех авторов.
func (r *BookRepository) GetAuthors() ([]*model.Author, error) {
	authorsMap := make(map[int64]*model.Author)
	for _, book := range r.books {
		authorsMap[book.AuthorID] = &model.Author{
			ID:   book.AuthorID,
			Name: book.AuthorName,
		}
	}

	authors := make([]*model.Author, 0, len(authorsMap))
	for _, author := range authorsMap {
		authors = append(authors, author)
	}

	return authors, nil
}

// AddBook добавляет новую книгу.
func (r *BookRepository) AddBook(book *model.Book) error {
	r.books = append(r.books, book)
	return nil
}

// GetBookByID возвращает книгу по ее идентификатору.
func (r *BookRepository) GetBookByID(id int64) (*model.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

// GetBooksByAuthorID возвращает список книг, принадлежащих определенному автору.
func (r *BookRepository) GetBooksByAuthorID(authorID int64) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	for _, book := range r.books {
		if book.AuthorID == authorID {
			books = append(books, book)
	
