ALTER TABLE _user DROP CONSTRAINT user_primary_key;
ALTER TABLE _user DROP CONSTRAINT user_fk_contact;
ALTER TABLE _user DROP CONSTRAINT user_fk_account;
DROP TABLE _user;