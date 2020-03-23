ALTER TABLE review
    ADD views BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '閲覧数' AFTER favorite_count,
	ADD travel_date DATE NOT NULL COMMENT '旅行日' AFTER views,
	ADD accompanying TINYINT NOT NULL COMMENT '同行者の種類' AFTER travel_date;
