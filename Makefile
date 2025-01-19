all:
	@go generate && go build

clean:
	@rm -f goes
	@rm -f *_string.go

run:
	@go run .

test:
	@go test ./...

build:
	@go build -o goes .
