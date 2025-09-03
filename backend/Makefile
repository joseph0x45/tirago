.PHONY: up down logs psql reset migrate-up migrate-down migrate-new
include .env

MIGRATE_DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE_PATH=./migrations
MIGRATE_BIN=./migrate.bin

migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATE_PATH) -database $(MIGRATE_DB_URL) up

migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATE_PATH) -database $(MIGRATE_DB_URL) down 1

migrate-new:
	@test -n "$(name)" || (echo ">> name= required"; exit 1)
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATE_PATH) -seq $(name)

build-dev:
	go build -o server.out .

run:
	./server.out

build-release:
	@echo 'not implemented'

up:
	docker run -d \
		--name $(POSTGRES_CONTAINER) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(POSTGRES_PORT):5432 \
		postgres

down:
	docker stop $(POSTGRES_CONTAINER) || true
	docker rm $(POSTGRES_CONTAINER) || true

logs:
	docker logs -f $(POSTGRES_CONTAINER)

psql:
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

reset: down
	$(MAKE) up
