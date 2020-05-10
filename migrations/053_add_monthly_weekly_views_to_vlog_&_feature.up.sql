ALTER TABLE vlog
    ADD monthly_views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1ヶ月の閲覧数' AFTER views,
    ADD weekly_views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1週間の閲覧数' AFTER monthly_views;
ALTER TABLE feature
    ADD monthly_views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1ヶ月の閲覧数' AFTER views,
    ADD weekly_views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1週間の閲覧数' AFTER monthly_views;