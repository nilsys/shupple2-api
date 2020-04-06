ALTER TABLE vlog 
  DROP user_sns,
  DROP FOREIGN KEY vlog_editor_id,
  DROP editor_id;

CREATE TABLE IF NOT EXISTS vlog_editor (
  vlog_id    BIGINT UNSIGNED NOT NULL,
  user_id    BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(vlog_id, user_id),
  CONSTRAINT vlog_editor_vlog_id FOREIGN KEY(vlog_id) REFERENCES vlog(id),
  CONSTRAINT vlog_editor_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'vlogの共同編集者';
