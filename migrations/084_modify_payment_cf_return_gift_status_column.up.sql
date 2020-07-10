ALTER TABLE payment_cf_return_gift
    DROP is_canceled,
    DROP is_owner_confirmed,
    ADD gift_type_other_status TINYINT NOT NULL DEFAULT 0 COMMENT 'その他の場合のステータス(その他以外の場合は1以上)' AFTER amount,
    ADD gift_type_reserved_ticket_status TINYINT NOT NULL DEFAULT 0 COMMENT '宿泊券の場合のステータス(宿泊券の場合は1以上)' AFTER gift_type_other_status;
