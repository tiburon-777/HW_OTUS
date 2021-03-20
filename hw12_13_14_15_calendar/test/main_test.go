package main

import (
	"github.com/stretchr/testify/require"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/test/client"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/test/misc"
	"log"
	"sync"
	"testing"
	"time"
)

func TestPublicAPIEndpoints(t *testing.T) {
	cli := []client.Interface{
		// client.GRPCAPI{Ctx: context.Background(), Host: "localhost", Port: "50051", Name: "GRPC API"},
		client.HTTPAPI{BaseURL: "http://localhost:50052", Name: "HTTP REST API"},
	}
	wg := sync.WaitGroup{}
	wg.Add(len(cli) * 5)

	for _, c := range cli {
		t.Run("test "+c.GetName()+" for Create, GetById and Delete", func(t *testing.T) {
			var ids []int64
			defer func() {
				wg.Done()
				clean(c, &ids)
			}()

			resp1, err := c.Create(&misc.TestEvent01)
			require.NoError(t, err)
			require.Greater(t, resp1.ID, int64(0))
			ids = append(ids, resp1.ID)

			resp2, err := c.GetByID(&public.GetByIDReq{ID: resp1.ID})
			require.NoError(t, err)
			require.Equal(t, 1, len(resp2.Events))
			require.Equal(t, misc.TestEvent01.Title, resp2.Events[0].Title)
			require.Equal(t, misc.TestEvent01.UserID, resp2.Events[0].UserID)
			require.Equal(t, misc.TestEvent01.Note, resp2.Events[0].Note)
		})

		t.Run("test "+c.GetName()+" for Create, Update, GetById and Delete", func(t *testing.T) {
			var ids []int64
			defer func() {
				wg.Done()
				clean(c, &ids)
			}()

			resp1, err := c.Create(&misc.TestEvent01)
			require.NoError(t, err)
			require.Greater(t, resp1.ID, int64(0))
			ids = append(ids, resp1.ID)

			err = c.Update(&public.UpdateReq{ID: resp1.ID, Event: &public.Event{ID: resp1.ID, Title: misc.TestEvent02.Title, Date: misc.TestEvent02.Date, Latency: misc.TestEvent02.Latency, Note: misc.TestEvent02.Note, UserID: misc.TestEvent02.UserID, NotifyTime: misc.TestEvent02.NotifyTime}})
			require.NoError(t, err)

			resp2, err := c.GetByID(&public.GetByIDReq{ID: resp1.ID})
			require.NoError(t, err)
			require.Equal(t, 1, len(resp2.Events))
			require.Equal(t, misc.TestEvent02.Title, resp2.Events[0].Title)
			require.Equal(t, misc.TestEvent02.UserID, resp2.Events[0].UserID)
			require.Equal(t, misc.TestEvent02.Note, resp2.Events[0].Note)
		})

		t.Run("test "+c.GetName()+" for Create, Delete and GetById", func(t *testing.T) {
			defer wg.Done()

			resp1, err := c.Create(&misc.TestEvent01)
			require.NoError(t, err)
			require.Greater(t, resp1.ID, int64(0))

			err = c.Delete(&public.DeleteReq{ID: resp1.ID})
			require.NoError(t, err)

			resp2, err := c.GetByID(&public.GetByIDReq{ID: resp1.ID})
			require.Error(t, err)
			require.Nil(t, resp2)
		})

		t.Run("test "+c.GetName()+" for Create, List and Delete", func(t *testing.T) {
			var ids []int64
			defer func() {
				wg.Done()
				clean(c, &ids)
			}()

			resp1, err := c.Create(&misc.TestEvent01)
			require.NoError(t, err)
			require.Greater(t, resp1.ID, int64(0))
			ids = append(ids, resp1.ID)

			resp2, err := c.Create(&misc.TestEvent02)
			require.NoError(t, err)
			require.Greater(t, resp2.ID, int64(0))
			ids = append(ids, resp2.ID)

			resp3, err := c.List()
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(resp3.Events), 2)
			var e1, e2 bool
			for _, v := range resp3.Events {
				if v.ID == resp1.ID {
					e1 = true
				}
				if v.ID == resp2.ID {
					e2 = true
				}
			}
			require.True(t, e1)
			require.True(t, e2)

			err = c.Delete(&public.DeleteReq{ID: resp1.ID})
			require.NoError(t, err)

			err = c.Delete(&public.DeleteReq{ID: resp2.ID})
			require.NoError(t, err)
		})

		t.Run("test "+c.GetName()+" for Create, GetByDate and Delete", func(t *testing.T) {
			var ids []int64
			defer func() {
				wg.Done()
				clean(c, &ids)
			}()

			startDate, err := time.Parse("2006-01-02T15:04:00", "3100-01-01T12:00:00")
			require.NoError(t, err)
			for i := time.Hour; i < (60 * 24 * time.Hour); i = i + 48*time.Hour {
				resp1, err := c.Create(&public.CreateReq{
					Title:      "Test event 02",
					Date:       misc.Time2pbtimestamp(startDate.Add(i)),
					Latency:    misc.Dur2pbduration(2 * 24 * time.Hour),
					Note:       "Note of test event 02",
					NotifyTime: misc.Dur2pbduration(5 * time.Minute),
					UserID:     2222,
				})
				require.NoError(t, err)
				require.Greater(t, resp1.ID, int64(0))
				ids = append(ids, resp1.ID)
			}

			resp2, err := c.GetByDate(&public.GetByDateReq{Date: misc.Time2pbtimestamp(startDate), Range: public.QueryRange_DAY})
			require.NoError(t, err)
			require.Equal(t, 1, len(resp2.Events))

			resp3, err := c.GetByDate(&public.GetByDateReq{Date: misc.Time2pbtimestamp(startDate), Range: public.QueryRange_WEEK})
			require.NoError(t, err)
			require.Equal(t, 4, len(resp3.Events))

			resp4, err := c.GetByDate(&public.GetByDateReq{Date: misc.Time2pbtimestamp(startDate), Range: public.QueryRange_MONTH})
			require.NoError(t, err)
			require.Equal(t, 16, len(resp4.Events))
		})
	}
}

func clean(c client.Interface, ids *[]int64) {
	for _, m := range *ids {
		err := c.Delete(&public.DeleteReq{ID: m})
		if err != nil {
			log.Println("error when try to clean DB: ", err.Error())
		}
	}
}
