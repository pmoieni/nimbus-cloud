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
