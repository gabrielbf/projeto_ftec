create table if not exists _user (
id bigserial NOT NULL,
--reference_uuid uuid  UNIQUE NOT NULL,
first_name varchar NOT NULL,
last_name varchar NOT NULL,
account_id bigserial NOT NULL,
contact_id bigserial NOT NULL,
created_at timestamptz NOT NULL,
updated_at timestamptz NOT NULL,
deleted_at timestamptz NULL DEFAULT NULL,

CONSTRAINT user_primary_key PRIMARY KEY (id),
CONSTRAINT user_fk_contact
FOREIGN KEY(contact_id)
REFERENCES contact(id)
ON DELETE cascade,
CONSTRAINT user_fk_account
FOREIGN KEY(account_id)
REFERENCES account(id)
ON DELETE cascade
);
