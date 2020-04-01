ALTER TABLE vlog
  ADD editor_id BIGINT UNSIGNED NOT NULL COMMENT '編集者' AFTER user_id,
  DROP COLUMN editor_name,
  ADD CONSTRAINT vlog_editor_id FOREIGN KEY(editor_id) REFERENCES user(id)
