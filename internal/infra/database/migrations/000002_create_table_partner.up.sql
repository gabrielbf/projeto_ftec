create table if not exists partner (
	id bigserial PRIMARY KEY NOT NULL,
	account_id bigserial NOT NULL,
	name varchar NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL DEFAULT NULL,

  CONSTRAINT fk_account
	  FOREIGN KEY(account_id)
		  REFERENCES account(id)
);
