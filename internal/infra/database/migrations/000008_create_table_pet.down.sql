ALTER TABLE pet DROP CONSTRAINT pet_primary_key;
ALTER TABLE pet DROP CONSTRAINT pet_fk_partner;
DROP TABLE pet;