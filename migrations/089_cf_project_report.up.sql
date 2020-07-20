ALTER TABLE post
  ADD cf_project_id BIGINT(20) UNSIGNED DEFAULT NULL COMMENT '記事をプロジェクトレポートにする時、どのプロジェクトに紐付けるか' AFTER toc,
  ADD FOREIGN KEY post_cf_project_id(cf_project_id) REFERENCES cf_project(id);
