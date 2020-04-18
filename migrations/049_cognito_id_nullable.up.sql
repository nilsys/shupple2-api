ALTER TABLE user
	MODIFY cognito_id varchar(255) COMMENT 'cognitoから返るsub。wordpress側にしか居ないユーザーの場合nullになる',
	DROP INDEX idx_user_cognito_id,
	ADD UNIQUE KEY user_cognito_id(cognito_id);
