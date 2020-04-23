ALTER TABLE area_category
    ADD sort_order BIGINT UNSIGNED DEFAULT NULL COMMENT 'sort順' AFTER metasearch_sub_sub_area_id;
