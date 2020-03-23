ALTER TABLE category
    ADD metasearch_discerning_condition_id BIGINT COMMENT '宿泊検索側のこだわり条件'  AFTER metasearch_inn_type_id,
    ADD FULLTEXT KEY(name) WITH PARSER NGRAM
