create table if not exists account (
	id bigserial PRIMARY KEY NOT NULL,
	email varchar NOT NULL,
	password varchar NOT NULL,
	status varchar NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL
);
