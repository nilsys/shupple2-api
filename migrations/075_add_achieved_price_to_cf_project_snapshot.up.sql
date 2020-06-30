ALTER TABLE cf_project_snapshot ADD achieved_price INT UNSIGNED DEFAULT 0 COMMENT '達成金額' AFTER goal_price;
