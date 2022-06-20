POSTGRES_BASE ?= $(shell pwd)
KEYCLOAK_BASE ?= $(shell dirname $(POSTGRES_BASE))

.PHONY: start stop clean

all: start

datadir:
	mkdir -p $(POSTGRES_BASE)/db

start: datadir
	sudo docker compose --env-file $(KEYCLOAK_BASE)/.env up


stop:
	sudo docker compose --env-file $(KEYCLOAK_BASE)/.env down


clean:
	sudo /bin/rm -rf $(POSTGRES_BASE)/db
