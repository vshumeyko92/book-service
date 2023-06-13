package service

import (
	"github.com/vshumeyko92/book-service/internal/model"
	"github.com/vshumeyko92/book-service/internal/repository"
)

// BookService представляет сервис для работы с данными книг.
type BookService struct {
	bookRepository repository.BookRepository
}

// NewBookService создает новый экземпляр BookService.
func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

// GetBooks возвращает список всех книг.
func (s *BookService) GetBooks() ([]*model.Book, error) {
	return s.bookRepository.GetBooks()
}

// GetAuthors возвращает список всех авторов.
func (s *BookService) GetAuthors() ([]*model.Author, error) {
	return s.bookRepository.GetAuthors()
}
