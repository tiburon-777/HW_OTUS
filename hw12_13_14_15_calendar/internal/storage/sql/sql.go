package sql

import (
	"database/sql"
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

const dateTimeLayout = "2006-01-02 15:04:00 -0700"

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

func (s *Storage) Create(ev event.Event) (event.ID, error) {
	res, err := s.db.Exec(
		`INSERT INTO events 
		(title, date, latency, note, userID, notifyTime) VALUES 
		($1, $2, $3, $4, $5, $6)`,
		ev.Title,
		ev.Date.Format(dateTimeLayout),
		ev.Latency,
		ev.Note,
		ev.UserID,
		ev.NotifyTime,
	)
	if err != nil {
		return -1, err
	}
	idint64, err := res.LastInsertId()
	return event.ID(idint64), err
}

func (s *Storage) Update(id event.ID, event event.Event) error {
	_, err := s.db.Exec(
		`UPDATE events set 
		title=$1, date=$2, latency=$3, note=$4, userID=$5, notifyTime=$6
		where id=$7`,
		event.Title,
		event.Date.Format(dateTimeLayout),
		event.Latency,
		event.Note,
		event.UserID,
		event.NotifyTime,
		id,
	)
	return err
}

func (s *Storage) Delete(id event.ID) error {
	_, err := s.db.Exec(
		`DELETE from events where id=$1`,
		id,
	)
	return err
}

func (s *Storage) List() (map[event.ID]event.Event, error) {
	res := make(map[event.ID]event.Event)
	results, err := s.db.Query(`SELECT (id,title,date,latency,note,userID,notifyTime) from events ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var id event.ID
		var evt event.Event
		var dateRaw string
		err = results.Scan(&id, &evt.Title, &dateRaw, &evt.Latency, &evt.Note, &evt.UserID, &evt.NotifyTime)
		if err != nil {
			return nil, err
		}
		evt.Date, err = time.Parse(dateTimeLayout, dateRaw)
		if err != nil {
			return nil, err
		}
		res[id] = evt
	}
	if results.Err() != nil {
		return nil, results.Err()
	}
	return res, nil
}

func (s *Storage) GetByID(id event.ID) (event.Event, bool) {
	var res event.Event
	var dateRaw string
	err := s.db.QueryRow(
		`SELECT (id,title,date,latency,note,userID,notifyTime) from events where id=$1`, id).Scan(res.Title, dateRaw, res.Latency, res.Note, res.UserID, res.NotifyTime)
	if err != nil {
		return res, false
	}
	dateParced, err := time.Parse(dateTimeLayout, dateRaw)
	if err != nil {
		return res, false
	}
	res.Date = dateParced
	return res, true
}

func (s *Storage) GetByDate(startDate time.Time, rng string) (map[event.ID]event.Event, error) {
	res := make(map[event.ID]event.Event)
	results, err := s.db.Query(
		`SELECT (id,title,date,latency,note,userID,notifyTime)
				from events
				where 	(date>$1 AND date<$2) OR
						(date+latency>$1 AND date+latency<$2) OR
						(date<$1 AND date+latency>$2)
				ORDER BY id`,
		startDate,
		getEndDate(startDate, rng))
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var id event.ID
		var evt event.Event
		var dateRaw string
		err = results.Scan(&id, &evt.Title, &dateRaw, &evt.Latency, &evt.Note, &evt.UserID, &evt.NotifyTime)
		if err != nil {
			return nil, err
		}
		evt.Date, err = time.Parse(dateTimeLayout, dateRaw)
		if err != nil {
			return nil, err
		}
		res[id] = evt
	}
	if results.Err() != nil {
		return nil, results.Err()
	}
	return res, nil
}

func getEndDate(startDate time.Time, rng string) time.Time {
	switch rng {
	case "DAY":
		return startDate.AddDate(0, 0, 1)
	case "WEEK":
		return startDate.AddDate(0, 0, 7)
	case "MONTH":
		return startDate.AddDate(0, 1, 0)
	default:
		return startDate
	}
}
