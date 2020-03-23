ALTER TABLE tourist_spot ADD rate DOUBLE DEFAULT NULL COMMENT '自社レーティング' AFTER search_inn_url;
ALTER TABLE tourist_spot ADD vendor_rate DOUBLE DEFAULT NULL COMMENT '外部レーティング' AFTER rate ;
