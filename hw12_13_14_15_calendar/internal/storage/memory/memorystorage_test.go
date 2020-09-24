package memorystorage

import (
	"github.com/stretchr/testify/require"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	s := New()

	t.Run("Empty storage", func(t *testing.T) {
		require.Equal(t,0, len(s.Events))
	})

	id,err:=s.Create(event.Event{ Title:"event1",Date:"11.11.1111" })

	t.Run("Create events", func(t *testing.T) {
		require.NoError(t,err)
		require.Equal(t,1, len(s.Events))
		require.Equal(t,event.Event{ Title:"event1",Date:"11.11.1111" }, s.Events[id])
	})

	t.Run("Update event", func(t *testing.T) {
		err:=s.Update(id,event.Event{ Title:"event1_modifyed",Date:"22.11.22222" })
		require.NoError(t,err)
		require.Equal(t,1, len(s.Events))
		require.Equal(t,event.Event{ Title:"event1_modifyed",Date:"22.11.22222" }, s.Events[id])
	})

	t.Run("List event", func(t *testing.T) {
		res,err:=s.List()
		require.NoError(t,err)
		require.Equal(t,1, len(res))
		require.Equal(t,event.Event{ Title:"event1_modifyed",Date:"22.11.22222" }, res[id])
	})

	t.Run("Get event by ID", func(t *testing.T) {
		res,ok := s.GetByID(id)
		require.Equal(t,ok,true)
		require.Equal(t,event.Event{ Title:"event1_modifyed",Date:"22.11.22222" }, res)
	})

	t.Run("Get event by fake ID", func(t *testing.T) {
		res,ok := s.GetByID(53663)
		require.Equal(t,ok,false)
		require.Equal(t,event.Event{}, res)
	})

	t.Run("Delete event", func(t *testing.T) {
		err := s.Delete(id)
		require.NoError(t,err)
		require.Equal(t,0, len(s.Events))
	})

}