CREATE TABLE IF NOT EXISTS user_tourist_spot (
  user_id         BIGINT UNSIGNED NOT NULL COMMENT 'user.id',
  tourist_spot_id BIGINT UNSIGNED NOT NULL COMMENT 'tourist_spot.id',
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, tourist_spot_id),
  CONSTRAINT user_tourist_spot_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_tourist_spot_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーと関係する（オーナー等）tourist_spot';