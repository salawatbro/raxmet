-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id                uuid                  DEFAULT Uuid_generate_v4() primary key,
    name              varchar(255) NOT NULL,
    email             varchar(255) NOT NULL,
    password          varchar(255) NOT NULL,
    profile_pic       TEXT,
    is_email_verified BOOLEAN               DEFAULT FALSE,
    role              varchar(255)          DEFAULT 'user' NOT NULL,
    created_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users;
