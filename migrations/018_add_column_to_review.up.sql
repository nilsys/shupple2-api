ALTER TABLE review
	ADD travel_date DATE NOT NULL COMMENT '旅行日' AFTER views,
	ADD accompanying TINYINT NOT NULL COMMENT '同行者の種類' AFTER travel_date;
