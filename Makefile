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