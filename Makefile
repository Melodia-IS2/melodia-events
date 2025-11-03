# Variables
DOCKER_COMPOSE=docker compose
NETWORK_NAME = melodia-network

.PHONY: up-app up-infra up-test up-dev down logs deps lint test-coverage kill ensure-network

ensure-network:
	@echo "🔍 Verificando red $(NETWORK_NAME)..."
	@if ! docker network inspect $(NETWORK_NAME) >/dev/null 2>&1; then \
		echo "🌐 Red $(NETWORK_NAME) no existe, creando..."; \
		docker network create --driver bridge $(NETWORK_NAME); \
	else \
		echo "✅ Red $(NETWORK_NAME) ya existe"; \
	fi

up-app: ensure-network
	@echo "🔹 Levantando APP + DB + Minio..."
	$(DOCKER_COMPOSE) --profile app up -d

up-infra: ensure-network
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
