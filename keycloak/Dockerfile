FROM quay.io/keycloak/keycloak:latest as builder

ENV KC_FEATURES=token-exchange
ENV KC_DB=postgres
RUN /opt/keycloak/bin/kc.sh build

FROM quay.io/keycloak/keycloak:latest
COPY --from=builder /opt/keycloak/ /opt/keycloak/
RUN mkdir -p /opt/keycloak/keystore
WORKDIR /opt/keycloak

ENTRYPOINT [ "/opt/keycloak/bin/kc.sh", "start" ]
