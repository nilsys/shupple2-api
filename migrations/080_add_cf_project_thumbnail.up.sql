CREATE TABLE IF NOT EXISTS cf_project_snapshot_thumbnail (
  cf_project_snapshot_id  BIGINT UNSIGNED NOT NULL,
  priority                INT UNSIGNED NOT NULL COMMENT '何番目か(1開始)',
  thumbnail               VARCHAR(255) NOT NULL,
  created_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT cf_project_thumbnail_cf_project_snapshot_id FOREIGN KEY(cf_project_snapshot_id) REFERENCES cf_project_snapshot(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE cf_project_snapshot DROP thumbnail;
