ALTER TABLE payment
    ADD remark VARCHAR(1024) NOT NULL COMMENT '備考' AFTER commission_price;
