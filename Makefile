run: build
	@./bin/app

build:
	@templ generate
	@tailwindcss -i ./view/css/index.css -o ./public/css/output.css
	@go build -o bin/app cmd/main.go
