create table if not exists partner (
id bigserial NOT NULL,
--reference_uuid uuid  UNIQUE NOT NULL,
account_id bigserial NOT NULL,
kind varchar NOT NULL,
name varchar NOT NULL,
location_id bigserial NOT NULL,
contact_id bigserial NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT partner_primary_key PRIMARY KEY (id),
CONSTRAINT partner_fk_contact
FOREIGN KEY(contact_id)
REFERENCES contact(id)
ON DELETE cascade,
CONSTRAINT partner_fk_location
FOREIGN KEY(location_id)
REFERENCES location(id)
ON DELETE cascade,
CONSTRAINT partner_fk_account
FOREIGN KEY(account_id)
REFERENCES account(id)
ON DELETE CASCADE
);