ALTER TABLE cf_project_snapshot DROP achieved_price;
ALTER TABLE cf_project ADD achieved_price INT UNSIGNED DEFAULT 0 COMMENT '達成金額' AFTER favorite_count;
