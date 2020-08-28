ALTER TABLE user
    ADD device_token VARCHAR(255) DEFAULT NULL COMMENT 'fcmデバイストークン' AFTER is_non_login;
