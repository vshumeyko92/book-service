package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vshumeyko92/book-service/internal/model"
	"github.com/vshumeyko92/book-service/internal/repository"
)

func TestBookService_GetBooks(t *testing.T) {
	// Создание мока репозитория книг
	mockBookRepo := &repository.MockBookRepository{
		GetBooksFunc: func() ([]*model.Book, error) {
			// Возвращаем тестовые данные книг
			books := []*model.Book{
				{ID: 1, Title: "Book 1", AuthorID: 1},
				{ID: 2, Title: "Book 2", AuthorID: 2},
			}
			return books, nil
		},
	}

	// Создание экземпляра сервиса книг с использованием мока репозитория
	bookService := NewBookService(mockBookRepo)

	// Вызов метода GetBooks сервиса книг
	books, err := bookService.GetBooks()

	// Проверка результатов
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, "Book 1", books[0].Title)
	assert.Equal(t, "Book 2", books[1].Title)
}

func TestBookService_GetAuthors(t *testing.T) {
	// Создание мока репозитория авторов
	mockAuthorRepo := &repository.MockAuthorRepository{
		GetAuthorsFunc: func() ([]*model.Author, error) {
			// Возвращаем тестовые данные авторов
			authors := []*model.Author{
				{ID: 1, Name: "Author 1"},
				{ID: 2, Name: "Author 2"},
			}
			return authors, nil
		},
	}

	// Создание экземпляра сервиса книг с использованием мока репозитория авторов
	bookService := NewBookService(nil, mockAuthorRepo)

	// Вызов метода GetAuthors сервиса книг
	authors, err := bookService.GetAuthors()

	// Проверка результатов
	assert.NoError(t, err)
	assert.Len(t, authors, 2)
	assert.Equal(t, "Author 1", authors[0].Name)
	assert.Equal(t, "Author 2", authors[1].Name)
}
