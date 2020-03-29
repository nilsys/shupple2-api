CREATE TABLE IF NOT EXISTS notice
(
    id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '通知ID',
    user_id            BIGINT UNSIGNED NOT NULL COMMENT '通知を受け取るユーザID',
    triggered_user_id  BIGINT UNSIGNED NOT NULL COMMENT '通知を引き起こすアクションをしたユーザID',
    action_type        TINYINT         NOT NULL COMMENT 'アクションの種別',
    action_target_type TINYINT         NOT NULL COMMENT 'アクション対象',
    action_target_id   BIGINT UNSIGNED NOT NULL COMMENT 'アクション対象のID',
    is_read            TINYINT(1)      NOT NULL COMMENT '既読フラグ',
    created_at         DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT notice_user_id FOREIGN KEY (user_id) REFERENCES user (id),
    CONSTRAINT notice_triggered_user_id FOREIGN KEY (user_id) REFERENCES user (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
    COMMENT = 'ユーザーがハッシュタグをフォローしたコマンドを保存するテーブル';