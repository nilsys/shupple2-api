ALTER TABLE payment_cf_return_gift
    ADD inquiry_code VARCHAR(255) NOT NULL COMMENT 'お問い合わせ番号' AFTER amount;