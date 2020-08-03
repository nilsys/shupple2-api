ALTER TABLE cf_project
    ADD facebook_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'facebookシェア数',
    ADD twitter_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'twitterシェア数';