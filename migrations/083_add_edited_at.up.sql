ALTER TABLE comic ADD edited_at DATETIME DEFAULT NULL;
UPDATE comic SET edited_at = updated_at;
ALTER TABLE comic MODIFY edited_at DATETIME NOT NULL;

ALTER TABLE feature ADD edited_at DATETIME DEFAULT NULL;
UPDATE feature SET edited_at = updated_at;
ALTER TABLE feature MODIFY edited_at DATETIME NOT NULL;

ALTER TABLE post ADD edited_at DATETIME DEFAULT NULL;
UPDATE post SET edited_at = updated_at;
ALTER TABLE post MODIFY edited_at DATETIME NOT NULL;

ALTER TABLE tourist_spot ADD edited_at DATETIME DEFAULT NULL;
UPDATE tourist_spot SET edited_at = updated_at;
ALTER TABLE tourist_spot MODIFY edited_at DATETIME NOT NULL;

ALTER TABLE vlog ADD edited_at DATETIME DEFAULT NULL;
UPDATE vlog SET edited_at = updated_at;
ALTER TABLE vlog MODIFY edited_at DATETIME NOT NULL;
