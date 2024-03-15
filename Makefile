start:
	go run github.com/a-h/templ/cmd/templ@latest generate --watch --cmd="go run cmd/main.go"
generate:
	go run github.com/a-h/templ/cmd/templ@latest generate

build: generate
	go build -ldflags="-s -w" -o bin/satontdev cmd/main.go