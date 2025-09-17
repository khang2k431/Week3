APP_NAME = realtime-chat

run:
	go run main.go

build:
	go build -o $(APP_NAME) main.go

tidy:
	go mod tidy

clean:
	rm -f $(APP_NAME)

test:
	go test ./...
