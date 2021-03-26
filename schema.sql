CREATE TABLE tweets
(
	id              varchar   NOT NULL PRIMARY KEY,
	text            varchar   NOT NULL,
	author_id       varchar   NOT NULL,
	conversation_id varchar,
	created_at      timestamp NOT NULL,
	query           varchar   NOT NULL
);

CREATE TABLE users
(
	id              varchar   NOT NULL PRIMARY KEY,
	username        varchar   NOT NULL,
	name            varchar   NOT NULL,
	created_at      timestamp NOT NULL,
	followers_count integer   NOT NULL,
	following_count integer   NOT NULL,
	tweet_count     integer   NOT NULL
);
