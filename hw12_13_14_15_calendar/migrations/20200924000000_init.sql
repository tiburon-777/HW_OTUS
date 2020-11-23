-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
	id int(16) NOT NULL AUTO_INCREMENT,
	title varchar(255) NOT NULL,
	date datetime NOT NULL,
	latency int(16) NOT NULL,
	note text,
	userID int(16),
	notifyTime int(16),
	notified bool
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
