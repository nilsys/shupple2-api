ALTER TABLE user
    ADD uid VARCHAR(255) NOT NULL COMMENT 'ユーザーが決める一意のID' AFTER id,
    ADD url  VARCHAR(255) DEFAULT NULL COMMENT 'ユーザーが自由に設定するURL',
    ADD facebook_url  VARCHAR(255) DEFAULT NULL COMMENT 'FacebookURL',
    ADD instagram_url  VARCHAR(255) DEFAULT NULL COMMENT 'InstagramURL',
    ADD twitter_url  VARCHAR(255) DEFAULT NULL COMMENT 'TwitterURL',
    ADD living_area  VARCHAR(255) DEFAULT NULL COMMENT '居住エリア';

UPDATE user SET uid = concat("userdummy", id);

ALTER TABLE user ADD CONSTRAINT UNIQUE INDEX user_uid(uid);
