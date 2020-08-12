ALTER TABLE post
    ADD beginning VARCHAR(1024) NOT NULL COMMENT '本文冒頭部分' AFTER title;