ALTER TABLE cf_return_gift
  ADD latest_cf_return_gift_snapshot_id BIGINT UNSIGNED DEFAULT NULL COMMENT '最新のcf_return_gift_snapshot.id',
  ADD CONSTRAINT cf_return_gift_latest_cf_return_gift_snapshot_id FOREIGN KEY(latest_cf_return_gift_snapshot_id) REFERENCES cf_return_gift_snapshot(id);
