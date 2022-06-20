# Keycloak identity provider

A customized image (see [Dockerfile](./Dockerfile)) of keycloak is used with inspiration from [Running keycloak in a container](https://www.keycloak.org/server/containers).

## Prepare java keystore

A java keystore is created for keycloak to use. Use the makefile to prepare it.
```
$ make
```

## Start Keycloak server with backend PostgreSQL

- Directly using docker compose.
  ```
  sudo docker compose --env-file .env up
  ```
- Using make file.
  ```
  $ make start
  ```