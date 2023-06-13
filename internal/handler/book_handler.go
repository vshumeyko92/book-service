package handler

import (
	"context"

	"github.com/vshumeyko92/book-service/api"
	"github.com/vshumeyko92/book-service/internal/service"
)

// BookAPI представляет обработчик GRPC-запросов для сервиса книг.
type BookAPI struct {
	bookService service.BookService
}

// NewBookAPI создает новый экземпляр BookAPI.
func NewBookAPI(bookService service.BookService) *BookAPI {
	return &BookAPI{
		bookService: bookService,
	}
}

// GetBooks обрабатывает запрос на получение списка книг.
func (b *BookAPI) GetBooks(ctx context.Context, req *api.GetBooksRequest) (*api.GetBooksResponse, error) {
	books, err := b.bookService.GetBooks()
	if err != nil {
		return nil, err
	}

	// Преобразование списка книг в формат GRPC-ответа
	var bookResponses []*api.Book
	for _, book := range books {
		bookResponses = append(bookResponses, &api.Book{
			Id:       book.ID,
			Title:    book.Title,
			AuthorId: book.AuthorID,
		})
	}

	return &api.GetBooksResponse{
		Books: bookResponses,
	}, nil
}

// GetAuthors обрабатывает запрос на получение списка авторов.
func (b *BookAPI) GetAuthors(ctx context.Context, req *api.GetAuthorsRequest) (*api.GetAuthorsResponse, error) {
	authors, err := b.bookService.GetAuthors()
	if err != nil {
		return nil, err
	}

	// Преобразование списка авторов в формат GRPC-ответа
	var authorResponses []*api.Author
	for _, author := range authors {
		authorResponses = append(authorResponses, &api.Author{
			Id:   author.ID,
			Name: author.Name,
		})
	}

	return &api.GetAuthorsResponse{
		Authors: authorResponses,
	}, nil
}
