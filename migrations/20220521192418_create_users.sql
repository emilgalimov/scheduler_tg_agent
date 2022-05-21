-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id      BIGINT,
    chat_id BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
