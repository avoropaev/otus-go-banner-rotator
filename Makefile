include .env

BIN_FILENAME := otus_go_banner_rotator
BIN := "./bin/$(BIN_FILENAME)"
DOCKER_COMPOSE_PATH="./deployments/docker-compose.yaml"

build:
	docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) build
	go build -v -o $(BIN) .

up:
	docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) up -d
	while ! docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) exec --user postgres db psql -c "select 'db ready!'" > /dev/null; do sleep 1; done;
	while ! curl -f -s http://localhost:15672 > /dev/null; do sleep 1; done;

down:
	docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) down --remove-orphans

install-migrator:
	(which goose > /dev/null) || go install github.com/pressly/goose/v3/cmd/goose@latest

migrate: install-migrator
	goose --dir ./migrations postgres ${POSTGRES_URI} up

run: build up migrate
	mkdir -p logs
	nohup $(BIN) serve-http --config ./config/banner_rotator_config.yaml 0<&- &> ./logs/banner_rotator.log &

stop:
	killall $(BIN_FILENAME)
	docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) down --remove-orphans

test:
	go test -race -v -count 100 ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

install-protoc:
ifeq (, $(shell which protoc))
ifeq ($(shell uname -s),Darwin)
	brew install protobuf
else
	apt install -y protobuf-compiler
endif
endif
	(which protoc-gen-go > /dev/null) || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	(which protoc-gen-go-grpc > /dev/null) || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate: install-protoc
	protoc \
	--proto_path=./api/ --go_out=./internal/server/pb --go-grpc_out=./internal/server/pb \
	--grpc-gateway_out ./internal/server/pb --grpc-gateway_opt logtostderr=true,paths=import,generate_unbound_methods=true \
	api/*.proto
	go generate ./...

integration-tests: up
	go test -v ./... --tags integration; \
	e=$$?; \
	docker-compose --env-file .env -f $(DOCKER_COMPOSE_PATH) down --remove-orphans; \
	exit $$e

.PHONY: build up down migrate run stop test lint generate
