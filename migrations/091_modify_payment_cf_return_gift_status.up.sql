ALTER TABLE payment_cf_return_gift
    MODIFY COLUMN gift_type_other_status TINYINT NULL,
    MODIFY COLUMN gift_type_reserved_ticket_status TINYINT NULL;
