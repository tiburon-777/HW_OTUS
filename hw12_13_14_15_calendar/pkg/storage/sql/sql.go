package sql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
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
		return fmt.Errorf("can't connect to SQL server: %w", err)
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
		return -1, fmt.Errorf("can't create event in SQL DB: %w", err)
	}
	idint64, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("can't get LastInsertId from SQL DB: %w", err)
	}
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
	if err != nil {
		return fmt.Errorf("can't update event in SQL DB: %w", err)
	}
	return nil
}

func (s *Storage) Delete(id event.ID) error {
	_, err := s.db.Exec(
		`DELETE from events where id=$1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("can't delete event from SQL DB: %w", err)
	}
	return nil
}

func (s *Storage) List() (map[event.ID]event.Event, error) {
	res := make(map[event.ID]event.Event)
	results, err := s.db.Query(`SELECT (id,title,date,latency,note,userID,notifyTime) from events ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("can't get list of events from SQL DB: %w", err)
	}
	defer results.Close()
	for results.Next() {
		var id event.ID
		var evt event.Event
		var dateRaw string
		err = results.Scan(&id, &evt.Title, &dateRaw, &evt.Latency, &evt.Note, &evt.UserID, &evt.NotifyTime)
		if err != nil {
			return nil, fmt.Errorf("can't parce list of events getted from SQL DB: %w", err)
		}
		evt.Date, err = time.Parse(dateTimeLayout, dateRaw)
		if err != nil {
			return nil, fmt.Errorf("can't parce Date getted from SQL DB: %w", err)
		}
		res[id] = evt
	}
	if results.Err() != nil {
		return nil, fmt.Errorf("something happens while we try to parce lines getted from SQL DB: %w", results.Err())
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
		return nil, fmt.Errorf("can't get list of events from SQL DB: %w", err)
	}
	defer results.Close()
	for results.Next() {
		var id event.ID
		var evt event.Event
		var dateRaw string
		err = results.Scan(&id, &evt.Title, &dateRaw, &evt.Latency, &evt.Note, &evt.UserID, &evt.NotifyTime)
		if err != nil {
			return nil, fmt.Errorf("can't parce list of events getted from SQL DB: %w", err)
		}
		evt.Date, err = time.Parse(dateTimeLayout, dateRaw)
		if err != nil {
			return nil, fmt.Errorf("can't parce Date getted from SQL DB: %w", err)
		}
		res[id] = evt
	}
	if results.Err() != nil {
		return nil, fmt.Errorf("something happens while we try to parce lines getted from SQL DB: %w", results.Err())
	}
	return res, nil
}

func (s *Storage) GetNotifications() (map[event.ID]event.Event, error) {
	res := make(map[event.ID]event.Event)
	results, err := s.db.Query(
		`SELECT (id,title,date,latency,note,userID,notifyTime)
				from events
				where 	($1>date-notifyTime) AND (NOT notified)
				ORDER BY id`,
		time.Now())
	if err != nil {
		return nil, fmt.Errorf("can't get list of events from SQL DB: %w", err)
	}
	defer results.Close()
	for results.Next() {
		var id event.ID
		var evt event.Event
		var dateRaw string
		err = results.Scan(&id, &evt.Title, &dateRaw, &evt.Latency, &evt.Note, &evt.UserID, &evt.NotifyTime)
		if err != nil {
			return nil, fmt.Errorf("can't parce list of events getted from SQL DB: %w", err)
		}
		evt.Date, err = time.Parse(dateTimeLayout, dateRaw)
		if err != nil {
			return nil, fmt.Errorf("can't parce Date getted from SQL DB: %w", err)
		}
		res[id] = evt
	}
	if results.Err() != nil {
		return nil, fmt.Errorf("something happens while we try to parce lines getted from SQL DB: %w", results.Err())
	}
	return res, nil
}

func (s *Storage) SetNotified(id event.ID) error {
	_, err := s.db.Exec(
		`UPDATE events set 
		notified=true 
		where id=$1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("can't set event as notified in SQL DB: %w", err)
	}
	return nil
}

func (s *Storage) PurgeOldEvents(days int64) (int64, error) {
	r, err := s.db.Exec(
		`DELETE from events where date<$1`,
		time.Now().Add(-time.Duration(days*24*60*60)),
	)
	if err != nil {
		return 0, fmt.Errorf("can't delete old events from SQL DB: %w", err)
	}
	l, err := r.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("can't run RowAffected on SQL DB query: %w", err)
	}
	return l, nil
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
