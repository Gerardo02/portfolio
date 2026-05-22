-- +goose Up
CREATE TABLE users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    oauth_id TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL DEFAULT ''
);

CREATE TABLE tasks (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	description TEXT NOT NULL DEFAULT '',
	done BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE tasks;
DROP TABLE users;
