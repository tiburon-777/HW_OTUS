package sqlstorage

import (
	"database/sql"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

type Config struct {
	User  string
	Pass  string
	Host  string
	Port  string
	Dbase string
}

type Storage struct {
	db *sql.DB
}

func New(conf Config) *Storage {
	return &Storage{}
}

func (s *Storage) Connect(config Config) error {
	var err error
	s.db, err = sql.Open("mysql", config.User+":"+config.Pass+"@tcp("+config.Host+":"+config.Port+")/"+config.Dbase)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Create(event event.Event) (int64, error) {
	res, err := s.db.Exec(
		`INSERT INTO events (title, date) VALUES ($1, $2)`,
		event.Title,
		event.Date,
	)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (s *Storage) Update(id int64, event event.Event) error {
	_, err := s.db.Exec(
		`UPDATE events set title=$1, date=$2 where id=$3`,
		event.Title,
		event.Date,
		id,
	)
	return err
}

func (s *Storage) Delete(id int64) error {
	_, err := s.db.Exec(
		`DELETE from events where id=$1`,
		id,
	)
	return err
}

func (s *Storage) List() (map[int64]event.Event, error) {
	res := make(map[int64]event.Event)
	results, err := s.db.Query(
		`SELECT (id,title,date) from events ORDER BY id`)
	if err != nil || results.Err() != nil {
		return res, err
	}
	defer results.Close()
	for results.Next() {
		var tmp struct {
			id    int64
			title string
			date  string
		}
		err = results.Scan(&tmp.id, &tmp.title, &tmp.date)
		if err != nil {
			return res, err
		}
		res[tmp.id] = event.Event{Title: tmp.title, Date: tmp.date}
	}
	return res, nil
}

func (s *Storage) GetByID(id int64) (event.Event, bool) {
	var res event.Event
	err := s.db.QueryRow(
		`SELECT (title,date) from events where id=$1`, id).Scan(res.Title, res.Date)
	if err != nil {
		return res, false
	}
	return res, true
}
