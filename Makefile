
run:
	go run main.go

build:
	go build -o bin/addigy main.go

build-move:
	go build -o /usr/local/bin/addigy main.go && chmod +x /usr/local/bin/addigy
