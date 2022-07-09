create table if not exists pet (
id bigserial NOT NULL,
--reference_uuid uuid  UNIQUE NOT NULL,
description varchar NOT NULL,
species varchar NOT NULL,
partner_id int,
status varchar NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT pet_primary_key PRIMARY KEY (id),
CONSTRAINT pet_fk_partner
FOREIGN KEY(partner_id)
REFERENCES partner(id)
ON DELETE CASCADE
);