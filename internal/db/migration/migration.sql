-- +migrate Up
CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email text NOT NULL,
    password text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (email)
);

-- +migrate UP
CREATE TABLE files (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name text NOT NULL,
    object_name text NOT NULL,
    owner BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(object_name)
);

-- +migrate UP
CREATE TABLE user_file_relation (
    user_id BIGINT NOT NULL,
    file_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (file_id) REFERENCES files(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(user_id, file_id)
);
