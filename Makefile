build:
		go build -o bin/todoapi

run:  build
		./bin/todoapi

test:
		go test -v ./...
