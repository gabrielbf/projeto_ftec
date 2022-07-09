ALTER TABLE phone DROP CONSTRAINT phone_primary_key;
ALTER TABLE phone DROP CONSTRAINT phone_fk_contact;
DROP TABLE phone;