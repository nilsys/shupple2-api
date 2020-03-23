ALTER TABLE review
    ADD comment_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'レビューにひもづくコメント数' AFTER accompanying,
    ADD FULLTEXT KEY(body) WITH PARSER NGRAM,
    ADD CONSTRAINT review_user_id FOREIGN KEY(user_id) REFERENCES user(id),
    ADD CONSTRAINT review_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id)
