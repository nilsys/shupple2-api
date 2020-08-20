ALTER TABLE user
    ADD is_non_login TINYINT(1) DEFAULT 0 COMMENT '非ログインユーザーかどうか' AFTER header_uuid;
