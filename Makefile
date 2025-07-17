run: build
	@./bin/app

templ: 
	@templ generate --watch --proxy=http://localhost:8080 --cmd="make"

tailwind:
	@tailwindcss -i ./view/css/index.css -o ./public/css/output.css --watch

reload:
	@tailwindcss -i ./view/css/index.css -o ./public/css/output.css
	@templ generate 

build:
	@templ generate
	@tailwindcss -i ./view/css/index.css -o ./public/css/output.css
	@go build -o bin/app cmd/main.go

install:
	@go get ./...
	@go mod tidy
