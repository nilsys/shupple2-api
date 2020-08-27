ALTER TABLE notice
    ADD endpoint VARCHAR(255) NOT NULL COMMENT '通知からの遷移先エンドポイント' AFTER action_target_id;
