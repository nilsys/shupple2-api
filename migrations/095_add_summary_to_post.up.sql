ALTER TABLE post
    ADD summary VARCHAR(255) NOT NULL COMMENT '本文冒頭部分' AFTER title;