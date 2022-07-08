create table if not exists location (
	id bigserial NOT NULL,
	--reference_uuid uuid  UNIQUE NOT NULL,
	country varchar NOT NULL,
	state varchar NOT NULL,
	city varchar NOT NULL,
	neighborhood varchar NOT NULL,
	street varchar NOT NULL,
	number varchar NOT NULL,
	complement varchar NOT NULL,
	cep varchar NOT NULL,
	coordinates varchar NOT NULL,
	description varchar NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT primary_key PRIMARY KEY (id)
);
