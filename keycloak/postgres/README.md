# Bootstrap PostgreSQL DB for use by `keycloak`

## Description

A postgres docker container is used to bootstrap the postgres DB. Normally, `docker compose --env-file ../.env up/down` can be used to start/stop the container. A makefile is provided to simplify the operations.

## Makefile usage

- Start the bootstrap container.
  ```
  $ make [start]
  ```
- Stop the bootstrap container.
  ```
  $ make stop
  ```
- Clean up the PostgreSQL DB.
  ```
  $ make clean
  ```