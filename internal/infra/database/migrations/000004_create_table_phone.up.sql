create table if not exists phone (
id bigserial NOT NULL,
--reference_uuid uuid  UNIQUE NOT NULL,
contact_id int NOT NULL,
phone_number varchar NOT NULL,
description varchar NOT NULL,
status varchar NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT phone_primary_key PRIMARY KEY (id),
CONSTRAINT phone_fk_contact
FOREIGN KEY(contact_id)
REFERENCES contact(id)
ON DELETE CASCADE
);