ALTER TABLE cf_return_gift
  DROP thumbnail,
  DROP sort_order,
  CHANGE latest_cf_return_gift_snapshot_id latest_snapshot_id BIGINT UNSIGNED DEFAULT NULL COMMENT '最新のcf_return_gift_snapshot.id';

ALTER TABLE cf_return_gift_snapshot
  MODIFY delivery_date VARCHAR(255) NOT NULL COMMENT 'お届け時期' AFTER full_amount,
  ADD sort_order INT NOT NULL COMMENT '並び順' AFTER deadline;
