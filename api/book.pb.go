// Package api содержит сгенерированный код для сервиса книг.
// Не изменяйте этот файл вручную. Он будет сгенерирован автоматически.
package api

import (
	context "context"
	fmt "fmt"
)

// BookServiceClient является клиентом для взаимодействия с сервисом книг.
type BookServiceClient interface {
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error)
	GetAuthors(ctx context.Context, in *GetAuthorsRequest, opts ...grpc.CallOption) (*GetAuthorsResponse, error)
}

// BookServiceServer является сервером для обработки запросов к сервису книг.
type BookServiceServer interface {
	GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error)
	GetAuthors(context.Context, *GetAuthorsRequest) (*GetAuthorsResponse, error)
}

// RegisterBookServiceServer регистрирует сервер книг в указанном сервере GRPC.
func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBooks(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookService/GetAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAuthors(ctx, req.(*GetAuthorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetBooksRequest представляет запрос на получение списка книг.
type GetBooksRequest struct {
}

// GetBooksResponse представляет ответ со списком книг.
type GetBooksResponse struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

// GetAuthorsRequest представляет запрос на получение списка авторов.
type GetAuthorsRequest struct {
}

// GetAuthorsResponse представляет ответ со списком авторов.
type GetAuthorsResponse struct {
	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

// Book представляет книгу.
type Book struct {
	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AuthorId int64    `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Author   *Author  `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
}

// Author представляет автора книги.
type Author struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func init() {
	fmt.Println("Initializing book.pb.go")
}

var (
	// Дескриптор сервиса книг
	_ServiceDesc = grpc.ServiceDesc{
		ServiceName: "api.BookService",
		HandlerType: (*BookServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetBooks",
				Handler:    _BookService_GetBooks_Handler,
			},
			{
				MethodName: "GetAuthors",
				Handler:    _BookService_GetAuthors_Handler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "book.proto",
	}

	// Дескриптор сервиса книг для регистрации
	_BookService_serviceDesc = grpc.ServiceDesc{
		ServiceName: "api.BookService",
		HandlerType: (*BookServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetBooks",
				Handler:    _BookService_GetBooks_Handler,
			},
			{
				MethodName: "GetAuthors",
				Handler:    _BookService_GetAuthors_Handler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "book.proto",
	}
)

