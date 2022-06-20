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

Files                   | Description
------------------------|---------------------------------
prepare-keystore.sh     | Prepares the Java keystore that can be consumed by keycloak for TLS.
keystore                | Location where Java keystore is saved. It is then bind mounted into the keycloak container.
Dockerfile              | Docker file for customized keycloak.
docker-compose.yml      | Build the customized keycloak if required. Both postgres and keycloak containers are started from here.
.env                    | A plethora of environment variables.