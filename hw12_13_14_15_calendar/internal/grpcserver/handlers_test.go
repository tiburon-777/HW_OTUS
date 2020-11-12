package grpcserver

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/suite"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage"
	oslog "log"
	"testing"
	"time"
)

var conf = config.Config{Server: config.Server{Address: "localhost", Port: "50511"}, GRPC: config.Server{Address: "localhost", Port: "50512"}, Logger: config.Logger{File: "calendar.log", Level: "INFO", MuteStdout: false}, Storage: config.Storage{InMemory: true, SQLHost: "", SQLPort: "", SQLDbase: "", SQLUser: "", SQLPass: ""}}

var storeConf = store.Config(conf.Storage)

var testEvt1 = Event{Title: "Test event 1", Date: buildTimestampWraped("01.01.2001"), Latency: ptypes.DurationProto(time.Hour * 24 * 1), Note: "First gen", UserID: 1111, NotifyTime: ptypes.DurationProto(time.Minute * 10)}
var testEvt2 = Event{Title: "Test event 2", Date: buildTimestampWraped("02.02.2002"), Latency: ptypes.DurationProto(time.Hour * 24 * 2), Note: "Second gen", UserID: 2222, NotifyTime: ptypes.DurationProto(time.Minute * 20)}
var testEvt3 = Event{Title: "Test event 3", Date: buildTimestampWraped("03.03.2003"), Latency: ptypes.DurationProto(time.Hour * 24 * 3), Note: "Third gen", UserID: 3333, NotifyTime: ptypes.DurationProto(time.Minute * 30)}

func TestCalendarSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	srv Service
}

func (suite *TestSuite) SetupTest() {
	log, err := logger.New(conf)
	if err != nil {
		oslog.Fatal("can't init logger")
	}
	suite.srv = Service{App: *app.New(log, store.NewStore(storeConf))}
}

func (s *TestSuite) TestCreateEvent() {
	controlEvent := testEvt1
	s.createEvent(&controlEvent)
	createdEvent := s.getEventByDate(controlEvent.Date)
	s.Equal(&controlEvent, createdEvent[0])
}

func (s *TestSuite) TestUpdateEvent() {
	oldEvent := testEvt1
	s.createEvent(&oldEvent)
	newEvent := testEvt2
	s.updateEvent(oldEvent, newEvent)
	updatedEvent := s.getEventByID(oldEvent.ID)
	s.Equal(newEvent.Title, updatedEvent.Title)
}

func (s *TestSuite) TestDeleteEvent() {
	testEvent := testEvt1
	s.createEvent(&testEvent)
	createdEvent := s.getEventByDate(testEvent.Date)
	s.Equal(&testEvent, createdEvent[0])
	s.deleteEvent(createdEvent[0].ID)
	controlEvent := s.getEventByDate(testEvent.Date)
	s.Equal(0, len(controlEvent))
}

func (s *TestSuite) TestListEvent() {
	testEvent := testEvt1
	s.createEvent(&testEvent)
	eventList := s.listEvents()
	s.Equal(1, len(eventList))
	s.Equal(testEvt1.Title, eventList[0].Title)
}

func (s *TestSuite) TestGetEventByID() {
	testEvent := testEvt1
	s.createEvent(&testEvent)
	eventList := s.getEventByID(testEvent.ID)
	s.Equal(testEvent.Title, eventList.Title)
}

func (s *TestSuite) TestGetEventByDate() {
	testEvent := testEvt1
	s.createEvent(&testEvent)
	eventList := s.getEventByDate(testEvent.Date)
	s.Equal(1, len(eventList))
	s.Equal(testEvent.Title, eventList[0].Title)
}

func (s *TestSuite) createEvent(ev *Event) {
	createdEvent, err := s.srv.Create(nil, eventToCreateRequest(*ev))
	s.NoError(err, "can,t create event")
	s.NotEqual(createdEvent.ID, 0, `message ID may not be a "0"`)
	ev.ID = createdEvent.ID
}

func (s *TestSuite) updateEvent(oldEvent Event, newEvent Event) {
	_, err := s.srv.Update(nil, eventToUpdateRequest(oldEvent, newEvent))
	s.NoError(err, "can,t update event")
}

func (s *TestSuite) deleteEvent(i int64) {
	_, err := s.srv.Delete(nil, &DeleteReq{ID: i})
	s.NoError(err, "can,t delete event")
}

func (s *TestSuite) listEvents() []*Event {
	r, err := s.srv.List(nil, nil)
	s.NoError(err, "can't list events")
	return r.Events

}

func (s *TestSuite) getEventByID(i int64) Event {
	r, err := s.srv.GetByID(nil, idToGetByIDRequest(i))
	s.NoError(err, "can't get event by id")
	s.Equal(1, len(r.Events), "length of slice in responce must be a \"1\"")
	return *r.Events[0]
}

func (s *TestSuite) getEventByDate(d *timestamp.Timestamp) []*Event {
	r, err := s.srv.GetByDate(nil, tstampToGetByDateRequest(d))
	s.NoError(err, "can't get event by date")
	return r.Events
}

func eventToCreateRequest(ev Event) (req *CreateReq) {
	return &CreateReq{Title: ev.Title, Date: ev.Date, Latency: ev.Latency, Note: ev.Note, UserID: ev.UserID, NotifyTime: ev.NotifyTime}
}

func eventToUpdateRequest(oldEvent Event, newEvent Event) (req *UpdateReq) {
	return &UpdateReq{ID: oldEvent.ID, Event: &Event{Title: newEvent.Title, Date: newEvent.Date, Latency: newEvent.Latency, Note: newEvent.Note, UserID: newEvent.UserID, NotifyTime: newEvent.NotifyTime}}
}

func idToGetByIDRequest(i int64) *GetByIDReq {
	return &GetByIDReq{ID: i}
}

func tstampToGetByDateRequest(t *timestamp.Timestamp) *GetByDateReq {
	return &GetByDateReq{Date: t, Range: 0}
}

func buildTimestampWraped(in string) *timestamp.Timestamp {
	tm, _ := time.Parse("02.01.2006", in)
	cdate, _ := ptypes.TimestampProto(tm)
	return cdate
}
