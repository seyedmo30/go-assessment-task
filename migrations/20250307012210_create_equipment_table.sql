-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS equipment
(
    id                 BIGSERIAL PRIMARY KEY,
    name               CHARACTER VARYING(252) NOT NULL,
    stock              BIGINT                 NOT NULL DEFAULT 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS equipment;
-- +goose StatementEnd
