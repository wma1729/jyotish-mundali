# keycloak identity provider

## Components

Files                   | Description
------------------------|---------------------------------
prepare-keystore.sh     | Prepares the Java keystore that can be consumed by keycloak for TLS.
keystore                | Location where Java keystore is saved. It is then bind mounted into the keycloak container.
Dockerfile              | Docker file for customized keycloak.
docker-compose.yml      | Build the customized keycloak if required. Both postgres and keycloak containers are started from here.
.env                    | A plethora of environment variables.