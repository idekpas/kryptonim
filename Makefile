NAME=kryptonim-app
VERSION=1.1.1

.PHONY: build
build:
	@go build -o $(NAME)

.PHONY: run
run: build
	@./$(NAME) -e dev

.PHONY: clean
clean:
	@rm -f $(NAME)

.PHONY: deps
deps:
	@go mod download

.PHONY: test
test:
	@go test -v ./...
