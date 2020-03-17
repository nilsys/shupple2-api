ALTER TABLE category
  ADD slug VARCHAR(255) NOT NULL COMMENT 'wordpressのslugを入れる' AFTER name,
  ADD INDEX idx_category_slug(slug);
