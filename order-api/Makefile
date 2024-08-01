.PHONY: generate
generate:
	@echo "==> Go Generating"
	@go generate ./...
	@echo "Running swag init"
	@swag init -d app/gateways/http -g router.go -o docs/swagger

.PHONY: tests
tests:
	@echo "==> Running tests"
	@go test -v ./...

.PHONY: run
run:
	@echo "==> Running server"
	@go run main.go

.PHONY: deps
deps:
	@echo "==> Setup: Tidying modules"
	go mod tidy
	@echo "==> Setup: Getting dependencies"
	go mod download

.PHONY: test-coverage
test-coverage:
	@echo "==> Test: Running Coverage Tests"
	gotest ./... -v -cover

.PHONY: build
build:
	@echo "==> Building binary"
	go build -o bin/user-api main.go

.PHONY: build-docker
build-docker:
	@echo "==> Running docker"
	docker compose up --build

.PHONY: start-docker
start-docker:
	@echo "==> Starting docker"
	docker compose up -d

.PHONY: stop-docker
stop-docker:
	@echo "==> Stopping docker"
	docker compose down