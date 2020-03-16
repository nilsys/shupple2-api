ALTER TABLE hashtag
	ADD post_count INT NOT NULL DEFAULT 0 AFTER name,
	ADD review_count INT NOT NULL DEFAULT 0 AFTER post_count,
	ADD score INT NOT NULL DEFAULT 0 AFTER review_count;
