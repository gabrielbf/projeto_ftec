create table if not exists post (
id bigserial NOT NULL,
--reference_uuid uuid  UNIQUE NOT NULL,
description varchar NOT NULL,
pet_id bigserial NOT NULL,
status varchar NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT post_primary_key PRIMARY KEY (id),
CONSTRAINT post_fk_pet
FOREIGN KEY(pet_id)
REFERENCES pet(id)
ON DELETE SET NULL
);