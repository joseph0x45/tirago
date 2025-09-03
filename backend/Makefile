.PHONY: up down logs psql reset
include .env

build-dev:
	go build -o server.out .

build-release:
	@echo 'not implemented'

new-migration:
	./migrate.bin create -ext sql -dir migrations -seq $(name)

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
