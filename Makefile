GOCMD=go

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

install: go.sum
	$(GOCMD) install