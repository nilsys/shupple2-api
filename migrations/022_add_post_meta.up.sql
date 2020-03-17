ALTER TABLE post
  ADD is_sticky TINYINT NOT NULL DEFAULT 0 COMMENT '新着順で上に固定して表示するかどうか' AFTER toc,
  ADD INDEX idx_post_sticky_updated_at(is_sticky, updated_at),
  ADD hide_ads TINYINT NOT NULL DEFAULT 0 COMMENT '広告を非表示にするかどうか' AFTER is_sticky,
  ADD seo_title VARCHAR(1024) NOT NULL COMMENT 'SEO用のtitle' AFTER views,
  ADD seo_description VARCHAR(1024) NOT NULL COMMENT 'SEO用のdescription' AFTER seo_title
