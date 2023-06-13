# Book Service

Простой сервис для работы с книгами и авторами.

## Установка

1. Клонируйте репозиторий:

   ```shell
   git clone https://github.com/vshumeyko92/book-service.git
   
   Перейдите в директорию проекта:
cd book-service
   Соберитк проект с помощью команды make:
make build
   Запустите приложение:
make run

## Использование

GRPC API
Сервис предоставляет следующие методы GRPC API:

GetBooks: получение списка книг
GetAuthors: получение списка авторов
Пример использования GRPC-клиента на Go:

package main

import (
	"context"
	"log"

	"github.com/vshumeyko92/book-service/api"
	"google.golang.org/grpc"
)

func main() {
	// Подключение к серверу GRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Создание клиента GRPC
	client := api.NewBookServiceClient(conn)

	// Выполнение запроса GetBooks
	booksReq := &api.GetBooksRequest{}
	booksResp, err := client.GetBooks(context.Background(), booksReq)
	if err != nil {
		log.Fatalf("failed to get books: %v", err)
	}

	// Обработка ответа
	for _, book := range booksResp.Books {
		log.Printf("Book ID: %d, Title: %s, Author ID: %d", book.Id, book.Title, book.AuthorId)
	}

	// Выполнение запроса GetAuthors
	authorsReq := &api.GetAuthorsRequest{}
	authorsResp, err := client.GetAuthors(context.Background(), authorsReq)
	if err != nil {
		log.Fatalf("failed to get authors: %v", err)
	}

	// Обработка ответа
	for _, author := range authorsResp.Authors {
		log.Printf("Author ID: %d, Name: %s", author.Id, author.Name)
	}
}


## Тестирование

Модульные тесты добавлены для основных компонентов проекта. Чтобы запустить тесты, выполните команду 'make test'
