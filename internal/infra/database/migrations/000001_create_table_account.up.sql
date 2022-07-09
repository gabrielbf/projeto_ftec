CREATE TYPE account_type AS ENUM (
    'USER',
    'PARTNER'
);

create table if not exists account (
id bigserial NOT NULL,
email varchar NOT NULL,
password varchar NOT NULL,
status varchar NOT NULL,
type account_type NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT account_primary_key PRIMARY KEY (id)
);
