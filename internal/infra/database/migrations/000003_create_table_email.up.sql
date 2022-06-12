create table if not exists email (
	id bigserial NOT NULL,
	--reference_uuid uuid  UNIQUE NOT NULL,
	contact_id int NOT NULL,
	email varchar NOT NULL,
	description varchar NOT NULL,
	status varchar NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT primary_key PRIMARY KEY (id)
  CONSTRAINT fk_contact
	FOREIGN KEY(contact_id)
		REFERENCES contact(id)
		ON DELETE CASCADE
);
