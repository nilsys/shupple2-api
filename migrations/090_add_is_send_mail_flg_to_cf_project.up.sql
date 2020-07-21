ALTER TABLE cf_project
    ADD is_sent_achievement_email TINYINT(1) NOT NULL DEFAULT 0 COMMENT '達成メール送信済か' AFTER achieved_price;
