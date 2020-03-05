ALTER TABLE category ADD metasearch_area_id BIGINT COMMENT '宿泊検索側のarea_id'  AFTER type;
ALTER TABLE category ADD metasearch_sub_area_id BIGINT COMMENT '宿泊検索側のsub_area_id'  AFTER metasearch_area_id;
ALTER TABLE category ADD metasearch_sub_sub_area_id BIGINT COMMENT '宿泊検索側のsub_sub_area_id'  AFTER metasearch_sub_area_id;
ALTER TABLE category ADD metasearch_inn_type_id BIGINT COMMENT '宿泊検索側の宿タイプid'  AFTER metasearch_sub_sub_area_id;
ALTER TABLE category ADD metasearch_discerning_condition_id BIGINT COMMENT '宿泊検索側のこだわり条件'  AFTER metasearch_inn_type_id;
