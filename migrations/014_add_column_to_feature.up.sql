ALTER TABLE feature ADD views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '閲覧数' AFTER body;
ALTER TABLE feature ADD facebook_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'facebookシェア数' AFTER views;
ALTER TABLE feature ADD twitter_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'twitterシェア数' AFTER facebook_count;
