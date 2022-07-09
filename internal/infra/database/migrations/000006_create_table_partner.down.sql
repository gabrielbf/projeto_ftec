ALTER TABLE partner DROP CONSTRAINT partner_primary_key;
ALTER TABLE partner DROP CONSTRAINT partner_fk_contact;
ALTER TABLE partner DROP CONSTRAINT partner_fk_location;
ALTER TABLE partner DROP CONSTRAINT partner_fk_account;
DROP TABLE partner;
