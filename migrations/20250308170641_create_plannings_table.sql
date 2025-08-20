-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS plannings
(
    id        BIGSERIAL PRIMARY KEY,
    equipment BIGINT NOT NULL,
    quantity  BIGINT NOT NULL,
    start_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    end_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plannings;
-- +goose StatementEnd
