BINARY_NAME=posts
export CONFIG_PATH=./config.yml

build: dep
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BINARY_NAME) main.go

dep:
	go mod download

run: build
	./bin/$(BINARY_NAME)

docker-run:
	docker-compose up --build