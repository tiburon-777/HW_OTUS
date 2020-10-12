package grpcserver

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

var conf = config.Config{Server: config.Server{Address: "localhost", Port: "50511"}, Grpc: config.Server{Address: "localhost", Port: "50512"}, Logger: config.Logger{File: "calendar.log", Level: "INFO", MuteStdout: false}, Storage: config.Storage{InMemory: true, SQLHost: "", SQLPort: "", SQLDbase: "", SQLUser: "", SQLPass: ""}}

var storeConf = store.Config(conf.Storage)

func TestService(t *testing.T) {
	log, err := logger.New(conf)
	require.NoError(t, err, "can't create logger")
	srv := Service{App: *app.New(log, store.NewStore(storeConf))}
	tm, _ := time.Parse("02.01.2006", "01.02.2003")
	cdate, err := ptypes.TimestampProto(tm)
	require.NoError(t, err, "can't convert time to proto")
	// Create
	step1, err := srv.Create(nil, &CreateReq{Title: "Test event", Date: cdate, Latency: ptypes.DurationProto(time.Hour * 24), Note: "First gen", UserID: 1111, NotifyTime: ptypes.DurationProto(time.Minute * 5)})
	require.NoError(t, err, "can,t create event")
	require.NotEqual(t, step1.ID, 0, "message ID may not be a \"0\"")
	// GetByID
	step2, err := srv.GetByID(nil, &GetByIDReq{ID: step1.ID})
	require.NoError(t, err, "can't get event by id")
	require.Equal(t, 1, len(step2.Events), "length of slice in responce must be a \"1\"")
	require.Equal(t, &Event{ID: 1, Title: "Test event", Date: cdate, Latency: ptypes.DurationProto(time.Hour * 24), Note: "First gen", UserID: 1111, NotifyTime: ptypes.DurationProto(time.Minute * 5)}, step2.Events[0], "request contains invalid data")
	// Update
	_, err = srv.Update(nil, &UpdateReq{ID: step1.ID, Event: &Event{Title: "Updated event", Date: cdate, Latency: ptypes.DurationProto(time.Hour * 48), Note: "Updated gen", UserID: 2222, NotifyTime: ptypes.DurationProto(time.Minute * 10)}})
	require.NoError(t, err, "can't update event")
	// List
	step3, err := srv.List(nil, nil)
	require.NoError(t, err, "problem with list after Update")
	require.Equal(t, 1, len(step3.Events), "length of slice in responce must be a \"1\"")
	require.Equal(t, &Event{ID: 1, Title: "Updated event", Date: cdate, Latency: ptypes.DurationProto(time.Hour * 48), Note: "Updated gen", UserID: 2222, NotifyTime: ptypes.DurationProto(time.Minute * 10)}, step3.Events[0], "request contains invalid data")
	// GetByDate
	step4, err := srv.GetByDate(nil, &GetByDateReq{Date: cdate, Range: 0})
	require.NoError(t, err, "problem with GetByDate")
	require.Equal(t, 1, len(step4.Events), "length of slice in responce must be a \"1\"")
	require.Equal(t, &Event{ID: 1, Title: "Updated event", Date: cdate, Latency: ptypes.DurationProto(time.Hour * 48), Note: "Updated gen", UserID: 2222, NotifyTime: ptypes.DurationProto(time.Minute * 10)}, step4.Events[0], "request contains invalid data")
	// Delete
	_, err = srv.Delete(nil, &DeleteReq{ID: step4.Events[0].ID})
	require.NoError(t, err, "problem with Delete")
	// List
	step5, err := srv.List(nil, nil)
	require.NoError(t, err, "problem with list after Delete")
	require.Equal(t, 0, len(step5.Events), "length of slice in responce must be a \"0\"")
	// GetByID
	_, err = srv.GetByID(nil, &GetByIDReq{ID: step4.Events[0].ID})
	require.Error(t, err, status.Errorf(codes.NotFound, "event not found"))
}
