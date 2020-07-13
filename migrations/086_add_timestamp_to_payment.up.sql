ALTER TABLE payment
    ADD owner_deposit_requested_at DATETIME DEFAULT NULL COMMENT 'オーナー入金依頼日時' AFTER total_price;
