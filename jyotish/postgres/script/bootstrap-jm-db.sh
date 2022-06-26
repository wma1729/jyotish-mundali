#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" --dbname "${POSTGRES_DB}" <<-EOSQL
	CREATE USER ${JYOTISH_USER} PASSWORD '${JYOTISH_PASSWORD}';
	CREATE DATABASE ${JYOTISH_DB} WITH ENCODING 'UTF8';
	GRANT ALL PRIVILEGES ON DATABASE ${JYOTISH_DB} TO ${JYOTISH_USER};

	CREATE TYPE language AS ENUM ('english', 'hindi');

	CREATE TABLE users (
		email text PRIMARY KEY,
		name text NOT NULL,
		lang language DEFAULT 'english',
		description text DEFAULT '',
		astrologer boolean DEFAULT false
	);
EOSQL
