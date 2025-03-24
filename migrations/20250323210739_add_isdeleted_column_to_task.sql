-- +goose Up
-- +goose StatementBegin
ALTER TABLE task
ADD isdeleted BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE task
DROP COLUMN isdeleted;
-- +goose StatementEnd
