-- +goose Up
CREATE TABLE IF NOT EXISTS groups
(
    id         uuid PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    owner_id   uuid         NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (owner_id) REFERENCES users (id)
);
CREATE INDEX IF NOT EXISTS name_idx ON groups (name);
CREATE INDEX IF NOT EXISTS owner_id_idx ON groups (owner_id);

CREATE TABLE IF NOT EXISTS group_members
(
    id       uuid PRIMARY KEY,
    group_id  uuid NOT NULL,
    user_id   uuid NOT NULL,
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (group_id) REFERENCES groups (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose Down
DROP TABLE IF EXISTS group_members;
DROP INDEX IF EXISTS name_idx;
DROP INDEX IF EXISTS owner_id_idx;
DROP TABLE IF EXISTS groups;
