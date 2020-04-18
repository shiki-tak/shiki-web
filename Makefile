GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BINARY_HOME=build
BINARY_NAME=shiki-web

golangci-lint:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint: golangci-lint
	$(GOCMD) mod verify
	golangci-lint run

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -rf ./${BINARY_HOME}
	mkdir ./build

build: clean lint test
	$(GOBUILD) -o ${BINARY_HOME}/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

install: go.sum
	$(GOCMD) install

clean-docker:
	docker kill  $(docker ps -a -q  --filter ancestor="shikitak/shiki-web:latest") || true
	docker rm -f $(docker ps -a -q  --filter ancestor="shikitak/shiki-web:latest") || true
	docker rmi -f shikitak/shiki-web:latest || true

build-docker: clean-docker
	docker build -t shikitak/shiki-web:latest .

run-server: build-docker
	docker run -p 1323:1323 shikitak/shiki-web:latest

run-mongodb:
	docker-compose -f ./app/db/mongo/docker-compose.yml up -d

run-mysql:
	rm -rf ./app/db/mysql/data
	docker-compose -f ./app/db/mysql/docker-compose.yml up -d
