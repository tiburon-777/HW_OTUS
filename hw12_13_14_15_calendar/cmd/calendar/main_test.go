package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var _ = func() bool {
	testing.Init()
	return true
}()

const testPortBase = 3000

var testEvent01 = public.CreateReq{
	Title:      "Test event 01",
	Date:       time2pbtimestamp(time.Now().Add(30 * time.Second)),
	Latency:    dur2pbduration(24 * time.Hour),
	Note:       "Note of test event 01",
	NotifyTime: dur2pbduration(5 * time.Minute),
	UserID:     1111,
}

var testEvent02 = public.CreateReq{
	Title:      "Test event 02",
	Date:       time2pbtimestamp(time.Now().Add(60 * time.Second)),
	Latency:    dur2pbduration(2 * 24 * time.Hour),
	Note:       "Note of test event 02",
	NotifyTime: dur2pbduration(5 * time.Minute),
	UserID:     2222,
}

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	go func(ctx context.Context) {
		main()
	}(ctx)
	time.Sleep(1 * time.Second)

	c := m.Run()

	cancel()
	os.Exit(c)
}

func TestPublicGRPCEndpoint(t *testing.T) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	publicAPI, err := public.NewClient(ctx, "localhost", "50051")
	require.NoError(t, err)

	wg.Add(5)
	// Реализовать тесты логики приложения:
	t.Run("test public GRPC.Create and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))

		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.NoError(t, err)
		require.Equal(t, 1, len(resp2.Events))
		require.Equal(t, testEvent01.Title, resp2.Events[0].Title)
		require.Equal(t, testEvent01.UserID, resp2.Events[0].UserID)
		require.Equal(t, testEvent01.Date.Seconds, resp2.Events[0].Date.Seconds)
		require.Equal(t, testEvent01.Note, resp2.Events[0].Note)
	})

	t.Run("test public GRPC.Create, GRPC.Update and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))
		_, err = publicAPI.Update(ctx, &public.UpdateReq{ID: resp1.ID, Event: &public.Event{ID: resp1.ID, Title: testEvent02.Title, Date: testEvent02.Date, Latency: testEvent02.Latency, Note: testEvent02.Note, UserID: testEvent02.UserID, NotifyTime: testEvent02.NotifyTime}})
		require.NoError(t, err)
		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.NoError(t, err)
		require.Equal(t, 1, len(resp2.Events))
		require.Equal(t, testEvent02.Title, resp2.Events[0].Title)
		require.Equal(t, testEvent02.UserID, resp2.Events[0].UserID)
		require.Equal(t, testEvent02.Date.Seconds, resp2.Events[0].Date.Seconds)
		require.Equal(t, testEvent02.Note, resp2.Events[0].Note)
	})

	t.Run("test public GRPC.Create, GRPC.Delete and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))
		_, err = publicAPI.Delete(ctx, &public.DeleteReq{ID: resp1.ID})
		require.NoError(t, err)
		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.Error(t, err)
		require.Nil(t, resp2)
	})

	t.Run("test public GRPC.Create and GRPC.List", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		resp2, err := publicAPI.Create(ctx, &testEvent02)
		require.NoError(t, err)
		require.NotEqual(t, resp1.ID, resp2.ID)

		list, err := publicAPI.List(ctx, &empty.Empty{})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(list.Events), 2)
		var e1, e2 bool
		for _, v := range list.Events {
			if v.ID == resp1.ID {
				e1 = true
			}
			if v.ID == resp2.ID {
				e2 = true
			}
		}
		require.True(t, e1)
		require.True(t, e2)
	})

	t.Run("test public GRPC.Create and GRPC.GetByDate", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		list, err := publicAPI.GetByDate(ctx, &public.GetByDateReq{Date: testEvent01.Date, Range: public.QueryRange_DAY})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(list.Events), 2)
		var e1 bool
		for _, v := range list.Events {
			if v.ID == resp1.ID {
				e1 = true
			}
		}
		require.True(t, e1)
	})

	wg.Wait()
}

func TestPublicAPIEndpoint(t *testing.T) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	publicAPI, err := public.NewClient(ctx, "localhost", "50051")
	require.NoError(t, err)

	wg.Add(5)
	// Реализовать тесты логики приложения:
	t.Run("test public GRPC.Create and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))

		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.NoError(t, err)
		require.Equal(t, 1, len(resp2.Events))
		require.Equal(t, testEvent01.Title, resp2.Events[0].Title)
		require.Equal(t, testEvent01.UserID, resp2.Events[0].UserID)
		require.Equal(t, testEvent01.Date.Seconds, resp2.Events[0].Date.Seconds)
		require.Equal(t, testEvent01.Note, resp2.Events[0].Note)
	})

	t.Run("test public GRPC.Create, GRPC.Update and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))
		_, err = publicAPI.Update(ctx, &public.UpdateReq{ID: resp1.ID, Event: &public.Event{ID: resp1.ID, Title: testEvent02.Title, Date: testEvent02.Date, Latency: testEvent02.Latency, Note: testEvent02.Note, UserID: testEvent02.UserID, NotifyTime: testEvent02.NotifyTime}})
		require.NoError(t, err)
		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.NoError(t, err)
		require.Equal(t, 1, len(resp2.Events))
		require.Equal(t, testEvent02.Title, resp2.Events[0].Title)
		require.Equal(t, testEvent02.UserID, resp2.Events[0].UserID)
		require.Equal(t, testEvent02.Date.Seconds, resp2.Events[0].Date.Seconds)
		require.Equal(t, testEvent02.Note, resp2.Events[0].Note)
	})

	t.Run("test public GRPC.Create, GRPC.Delete and GRPC.GetById", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		require.Greater(t, resp1.ID, int64(0))
		_, err = publicAPI.Delete(ctx, &public.DeleteReq{ID: resp1.ID})
		require.NoError(t, err)
		resp2, err := publicAPI.GetByID(ctx, &public.GetByIDReq{ID: resp1.ID})
		require.Error(t, err)
		require.Nil(t, resp2)
	})

	t.Run("test public GRPC.Create and GRPC.List", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		resp2, err := publicAPI.Create(ctx, &testEvent02)
		require.NoError(t, err)
		require.NotEqual(t, resp1.ID, resp2.ID)

		list, err := publicAPI.List(ctx, &empty.Empty{})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(list.Events), 2)
		var e1, e2 bool
		for _, v := range list.Events {
			if v.ID == resp1.ID {
				e1 = true
			}
			if v.ID == resp2.ID {
				e2 = true
			}
		}
		require.True(t, e1)
		require.True(t, e2)
	})

	t.Run("test public GRPC.Create and GRPC.GetByDate", func(t *testing.T) {
		defer wg.Done()
		resp1, err := publicAPI.Create(ctx, &testEvent01)
		require.NoError(t, err)
		list, err := publicAPI.GetByDate(ctx, &public.GetByDateReq{Date: testEvent01.Date, Range: public.QueryRange_DAY})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(list.Events), 2)
		var e1 bool
		for _, v := range list.Events {
			if v.ID == resp1.ID {
				e1 = true
			}
		}
		require.True(t, e1)
	})

	wg.Wait()
}

func time2pbtimestamp(t time.Time) *timestamp.Timestamp {
	r, err := ptypes.TimestampProto(t)
	if err != nil {
		log.Fatalf("cant convert Time to Timestamp: %s", err.Error())
	}
	return r
}

func dur2pbduration(t time.Duration) *duration.Duration {
	return ptypes.DurationProto(t)
}
