ALTER TABLE card
    ADD last4 VARCHAR(255) NOT NULL AFTER card_id,
    ADD expired VARCHAR(255) NOT NULL AFTER last4;
