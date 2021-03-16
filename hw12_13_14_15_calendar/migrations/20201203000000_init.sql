-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
	id serial NOT NULL,
	title varchar(255) NOT NULL,
	date timestamp NOT NULL,
	latency int8 NOT NULL,
	note text NULL,
	userID int8 NOT NULL,
	notifyTime int8 NULL,
	notified bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
