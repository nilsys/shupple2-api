ALTER TABLE cf_return_gift_snapshot
    ADD is_cancelable TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'キャンセル可能か' AFTER full_amount,
    ADD deadline      DATETIME DEFAULT NULL COMMENT '有効期限(宿泊券の場合のみ入る)' AFTER is_cancelable;
