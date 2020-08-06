ALTER TABLE cf_project
    ADD is_sent_new_post_email BIGINT NULL COMMENT '最新の報告(post)の通知メールが送信済か' AFTER is_sent_achievement_email,
    ADD latest_post_id BIGINT NULL COMMENT '最新の報告(post)id' AFTER is_sent_new_post_email;