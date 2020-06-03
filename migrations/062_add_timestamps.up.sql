ALTER TABLE review_comment_reply ADD deleted_at DATETIME DEFAULT NULL AFTER updated_at;
