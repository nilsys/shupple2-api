ALTER TABLE tourist_spot
  MODIFY lat DOUBLE COMMENT '緯度',
  MODIFY lng DOUBLE COMMENT '経度',
  MODIFY search_inn_url VARCHAR(1024) DEFAULT NULL COMMENT '関連宿を探すためのstaywayURL';
