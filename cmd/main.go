package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/vshumeyko92/book-service/api"
	"github.com/vshumeyko92/book-service/internal/handler"
	"github.com/vshumeyko92/book-service/internal/repository"
	"github.com/vshumeyko92/book-service/internal/service"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Инициализация репозитория и сервиса книг
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)

	// Создание GRPC-сервера и регистрация обработчика
	grpcServer := grpc.NewServer()
	bookAPI := handler.NewBookAPI(bookService)
	api.RegisterBookServiceServer(grpcServer, bookAPI)

	log.Printf("Starting gRPC server on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
