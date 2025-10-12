# Variables
DOCKER_COMPOSE=docker compose

.PHONY: up-app up-infra up-test up-dev down logs deps lint test-coverage kill

up-app:
	@echo "🔹 Levantando APP + DB + Minio..."
	$(DOCKER_COMPOSE) --profile app up -d

up-infra:
	@echo "🔹 Levantando solo DB + Minio..."
	$(DOCKER_COMPOSE) --profile infra up -d

up-test:
	@echo "🔹 Levantando entorno de TEST (DB-test + Minio-test)..."
	$(DOCKER_COMPOSE) --env-file ./.env.test --profile test up -d

up-dev:
	@echo "🔹 Levantando entorno DEV (hot-reload)..."
	$(DOCKER_COMPOSE) --profile dev up -d

down:
	@echo "🔹 Bajando todos los servicios..."
	$(DOCKER_COMPOSE) down -v

logs:
	@echo "🔹 Mostrando logs de todos los servicios..."
	$(DOCKER_COMPOSE) logs -f

deps:
	go mod tidy
	go mod verify

lint:
	go vet ./...
	staticcheck ./...
	go fmt ./...

swagger:
	swag init -g cmd/service/main.go -o docs --parseDependency --parseInternal

kill:
	docker stop $$(docker ps -aq) && docker rm $$(docker ps -aq) && docker rmi -f $$(docker images -q) && docker volume rm $$(docker volume ls -q) && docker network prune -f


TEST_PATH ?= ./...

COVERAGE_FILE ?= coverage.out

test-coverage:
	go test -coverpkg=./... -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@open coverage.html
