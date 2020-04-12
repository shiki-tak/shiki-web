GOCMD=go
GOBUILD=$(GOCMD) build

BINARY_HOME=build
BINARY_NAME=shiki-web

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -rf ./${BINARY_HOME}
	mkdir ./build

build: clean
	$(GOBUILD) -o ${BINARY_HOME}/$(BINARY_NAME) -v

install: go.sum
	$(GOCMD) install

clean-docker: clean-docker
	docker kill  $(docker ps -a -q  --filter ancestor="shikitak/shiki-web:latest") || true
	docker rm -f $(docker ps -a -q  --filter ancestor="shikitak/shiki-web:latest") || true
	docker rmi -f shikitak/shiki-web:latest || true

build-docker:
		docker build -t shikitak/shiki-web:latest .

run-docker: build-docker
		docker run -p 1323:1323 shikitak/shiki-web:latest
