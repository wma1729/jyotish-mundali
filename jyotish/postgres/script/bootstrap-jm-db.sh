#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" --dbname "${POSTGRES_DB}" <<-EOSQL
	CREATE USER ${JYOTISH_USER} PASSWORD '${JYOTISH_PASSWORD}';
	CREATE DATABASE ${JYOTISH_DB} WITH ENCODING 'UTF8';
	GRANT ALL PRIVILEGES ON DATABASE ${JYOTISH_DB} TO ${JYOTISH_USER};
EOSQL

psql -v ON_ERROR_STOP=1 --username "${JYOTISH_USER}" --dbname "${JYOTISH_DB}" <<-EOSQL
	CREATE OR REPLACE FUNCTION trigger_update_timestamp() RETURNS TRIGGER AS \$\$
	BEGIN
		NEW.updated_at = NOW();
		RETURN NEW;
	END;
	\$\$ LANGUAGE plpgsql;

	CREATE TYPE language AS ENUM ('en', 'hi');

	CREATE TABLE users (
		email text PRIMARY KEY,
		name text NOT NULL,
		description text NOT NULL DEFAULT '',
		lang language DEFAULT 'en',
		astrologer boolean DEFAULT false,
		public boolean DEFAULT true,
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER update_timestamp
		BEFORE UPDATE ON users FOR EACH ROW
		EXECUTE PROCEDURE trigger_update_timestamp();

	CREATE TABLE profiles (
		id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		email text REFERENCES users(email),
		name text NOT NULL,
		dob timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		city text NOT NULL DEFAULT '',
		state text NOT NULL DEFAULT '',
		country text NOT NULL DEFAULT '',
		details jsonb NOT NULL DEFAULT '{}'::jsonb,
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER update_timestamp
		BEFORE UPDATE ON profiles FOR EACH ROW
		EXECUTE PROCEDURE trigger_update_timestamp();

EOSQL
