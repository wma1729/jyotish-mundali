DIRS = keycloak jyotish/postgres

all:
	@for I in ${DIRS}; do make -C $${I} start COMPOSE_FLAG=--detach; done

stop:
	@for I in ${DIRS}; do make -C $${I} stop; done
