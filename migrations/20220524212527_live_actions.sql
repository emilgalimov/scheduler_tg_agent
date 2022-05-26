-- +goose Up
-- +goose StatementBegin
CREATE TABLE active_live_actions
(
    chat_id BIGINT PRIMARY KEY,
    name    VARCHAR,
    state   VARCHAR,
    data    json NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE active_live_actions;
-- +goose StatementEnd
