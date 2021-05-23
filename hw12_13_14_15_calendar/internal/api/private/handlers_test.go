package private

import (
	"github.com/stretchr/testify/suite"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
	oslog "log"
	"testing"
	"time"
)

var conf = config.Calendar{
	HTTP:    config.Server{Address: "localhost", Port: "50511"},
	GRPC:    config.Server{Address: "localhost", Port: "50512"},
	API:     config.Server{Address: "localhost", Port: "50513"},
	Logger:  config.Logger{File: "calendar.log", Level: "INFO", MuteStdout: false},
	Storage: config.Storage{InMemory: true, SQLHost: "", SQLPort: "", SQLDbase: "", SQLUser: "", SQLPass: ""}}

var storeConf = store.Config(conf.Storage)

var testEvt = event.Event{Title: "Test event 1", Date: time.Now(), Latency: time.Hour * 24, Note: "First gen", UserID: 1111, NotifyTime: 0, Notified: false}

func TestCalendarSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	srv Service
}

func (suite *TestSuite) SetupTest() {
	log, err := logger.New(logger.Config(conf.Logger))
	if err != nil {
		oslog.Fatal("can't init logger")
	}
	suite.srv = Service{App: *calendar.New(log, store.NewStore(storeConf))}
}

func (s *TestSuite) TestGetNotifications() {
	controlEvent := testEvt
	s.createEvent(controlEvent)

	returnedEvent := s.getNotifications()
	s.Equal(controlEvent.Title, returnedEvent.Title)
	s.Equal(controlEvent.UserID, returnedEvent.UserID)
}

func (s *TestSuite) TestSetNotified() {
	controlEvent := testEvt
	id := s.createEvent(controlEvent)

	s.setNotified(id)
	returnedEvent := s.getEventByID(id)
	s.Equal(true, returnedEvent.Notified)
}

func (s *TestSuite) createEvent(ev event.Event) event.ID {
	createdEvent, err := s.srv.App.Storage.Create(ev)
	s.NoError(err, "can,t create event")
	s.NotEqual(createdEvent, 0, `message ID may not be a "0"`)
	return createdEvent
}

func (s *TestSuite) getNotifications() Event {
	evs, err := s.srv.GetNotifications(nil, nil)
	s.NoError(err, "can't get events")
	s.Equal(1, len(evs.Events), "length of slice in responce must be a \"1\"")
	return *evs.Events[0]
}

func (s *TestSuite) setNotified(id event.ID) {
	_, err := s.srv.SetNotified(nil, &SetReq{ID: int64(id)})
	s.NoError(err, "can't mark event as notified")
}

func (s *TestSuite) getEventByID(id event.ID) event.Event {
	r, ok := s.srv.App.Storage.GetByID(id)
	s.Equal(true, ok, "can't get event by id")
	return r
}
