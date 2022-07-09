ALTER TABLE email DROP CONSTRAINT email_primary_key;
ALTER TABLE email DROP CONSTRAINT email_fk_contact;
DROP TABLE email;