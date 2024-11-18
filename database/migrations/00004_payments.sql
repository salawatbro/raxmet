-- +goose Up
CREATE TABLE IF NOT EXISTS payments
(
    id         uuid      DEFAULT Uuid_generate_v4() primary key,
    paid_by    uuid         NOT NULL,
    paid_to    uuid         NOT NULL,
    amount     varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (paid_by) REFERENCES users (id),
    FOREIGN KEY (paid_to) REFERENCES users (id)
);
-- +goose Down
DROP TABLE IF EXISTS payments;