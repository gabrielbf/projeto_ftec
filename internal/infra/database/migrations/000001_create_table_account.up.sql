create table if not exists account (
	id bigserial NOT NULL,
	--reference_uuid uuid  UNIQUE NOT NULL,
	email varchar NOT NULL,
	password varchar NOT NULL,
	status varchar NOT NULL,
	type varchar NOT NULL
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT primary_key PRIMARY KEY (id)
);