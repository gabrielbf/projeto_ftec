create table if not exists _user (
	id bigserial PRIMARY KEY NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	account_id bigserial NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

	CONSTRAINT fk_account
		FOREIGN KEY(account_id)
			REFERENCES account(id)
);
