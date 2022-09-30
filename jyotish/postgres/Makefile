POSTGRES_BASE ?= $(shell pwd)

.PHONY: start stop clean

all: start

datadir:
	mkdir -p $(POSTGRES_BASE)/db

start: datadir
	sudo docker compose --env-file $(POSTGRES_BASE)/.env up $(COMPOSE_FLAG)


stop:
	sudo docker compose --env-file $(POSTGRES_BASE)/.env down


clean:
	sudo /bin/rm -rf $(POSTGRES_BASE)/db
