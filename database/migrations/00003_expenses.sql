-- +goose Up
CREATE TABLE IF NOT EXISTS expenses
(
    id                uuid      DEFAULT Uuid_generate_v4() primary key,
    group_id    uuid           NOT NULL,
    user_id     uuid           NOT NULL,
    amount      varchar(255) NOT NULL,
    Title       varchar(255)   NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (group_id) REFERENCES groups (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE TABLE IF NOT EXISTS expense_shares
(
    id           uuid PRIMARY KEY,
    expense_id   uuid           NOT NULL,
    user_id      uuid           NOT NULL,
    share_amount varchar(255) NOT NULL,

    FOREIGN KEY (expense_id) REFERENCES expenses (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose Down
DROP TABLE IF EXISTS expense_shares;
DROP TABLE IF EXISTS expenses;
