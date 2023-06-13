# Компилятор Go
GO := go

# Название исполняемого файла
BINARY := book-service

# Флаги для сборки
BUILD_FLAGS := -v

.PHONY: all build run stop clean

all: build

# Сборка приложения
build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY)

# Запуск приложения
run:
	./$(BINARY)

# Остановка приложения
stop:
	killall $(BINARY) || true

# Очистка собранных файлов
clean:
	rm -f $(BINARY)
