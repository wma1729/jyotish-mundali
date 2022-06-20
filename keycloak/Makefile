KEYCLOAK_BASE ?= $(shell pwd)
STEP_CA_BASE  ?= $(shell dirname `pwd`)/step-ca

include $(KEYCLOAK_BASE)/.env

HOST_CERT      = $(STEP_CA_BASE)/data/$(KC_HOSTNAME).crt
HOST_KEY       = $(STEP_CA_BASE)/data/$(KC_HOSTNAME).key

.PHONY: keystoredir $(HOST_CERT)

all: keystoredir keystore/jm-keystore.jks

keystoredir:
	mkdir -p $(KEYCLOAK_BASE)/keystore

keystore/jm-keystore.jks: keystore/certs.p12
	keytool -importkeystore \
        -deststorepass $(KC_KS_PWD) -destkeystore $@ \
        -srcstorepass $(KC_KS_PWD) -srckeystore $^ \
		-srcstoretype PKCS12

keystore/certs.p12: keystore/root.crt $(HOST_CERT)
	openssl pkcs12 -export \
		-in $(HOST_CERT) -inkey $(HOST_KEY) -CAfile keystore/root.crt \
		-name $(KC_HOSTNAME) -password pass:$(KC_KS_PWD) -out $@

keystore/root.crt:
	cat /etc/ssl/certs/ca-certificates.crt $(STEP_CA_BASE)/data/certs/root_ca.crt > $@

$(HOST_CERT):
	export STEP_CA_BASE=$(STEP_CA_BASE) && \
	$(STEP_CA_BASE)/gen-cert.sh $(KC_HOSTNAME)

start:
	sudo docker compose --env-file $(KEYCLOAK_BASE)/.env up

stop:
	sudo docker compose --env-file $(KEYCLOAK_BASE)/.env down

clean:
	/bin/rm -rf $(KEYCLOAK_BASE)/keystore
