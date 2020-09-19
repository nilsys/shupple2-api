ALTER TABLE user
    ADD associate_id VARCHAR(255) DEFAULT NULL COMMENT '業務委託でプロモーションを依頼してる場合の識別子、リターンギフト購入アフィリンク等で使用する想定' AFTER device_token;

ALTER TABLE payment
    ADD associate_user_id BIGINT UNSIGNED DEFAULT NULL COMMENT '購入経路が業務委託でプロモーションを依頼したuserからの場合,そのuser.id' AFTER remark,
    ADD CONSTRAINT payment_associate_user_id FOREIGN KEY(associate_user_id) REFERENCES user(id);
