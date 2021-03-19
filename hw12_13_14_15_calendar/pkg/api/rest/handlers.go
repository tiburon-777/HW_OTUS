package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FromRESTCreate(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		bodyReq := public.CreateReq{}
		bodyIn, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			err503("can't body of the HTTP API request", err, calendar.Logger, r)
		}
		err = json.Unmarshal(bodyIn, &bodyReq)
		if err != nil {
			err503("can't unmarshal data from HTTP API request", err, calendar.Logger, r)
		}
		evt, err := createReq2Event(bodyReq)
		if err != nil {
			err503("can't convert types", err, calendar.Logger, r)
		}
		id, err := calendar.Storage.Create(evt)
		if err != nil {
			err503("can't create event through HTTP API", err, calendar.Logger, r)
		}
		bodyOut,err := json.Marshal(&public.CreateRsp{ID: int64(id)})
		r.WriteHeader(201)
		r.Write(bodyOut)
	}
}

func FromRESTUpdate(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		paramID, err := strconv.Atoi(mux.Vars(req)["ID"])
		if err != nil {
			err503("can't get request parameter", err, calendar.Logger, r)
		}
		bodyReq := public.Event{}
		bodyIn, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			err503("can't body of the HTTP API request", err, calendar.Logger, r)
		}
		err = json.Unmarshal(bodyIn, &bodyReq)
		if err != nil {
			err503("can't unmarshal data from HTTP API request", err, calendar.Logger, r)
		}
		evt, err := pubEvent2Event(bodyReq)
		if err != nil {
			err503("can't convert types", err, calendar.Logger, r)
		}
		err = calendar.Storage.Update(event.ID(paramID), evt)
		if err != nil {
			err503("can't update event through HTTP API", err, calendar.Logger, r)
		}
		r.WriteHeader(200)
	}
}

func FromRESTDelete(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		paramID, err := strconv.Atoi(mux.Vars(req)["ID"])
		if err != nil {
			err503("can't get request parameter", err, calendar.Logger, r)
		}
		err = calendar.Storage.Delete(event.ID(paramID))
		if err != nil {
			err503("can't create event from HTTP API", err, calendar.Logger, r)
		}
		r.WriteHeader(200)
	}
}

func FromRESTList(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		evs, err := calendar.Storage.List()
		if err != nil {
			err503("can't list events through HTTP API", err, calendar.Logger, r)
		}
		events, err := events2pubEvents(evs)
		if err != nil {
			err503("can't convert types", err, calendar.Logger, r)
		}
		bodyOut,err := json.Marshal(&events)
		r.WriteHeader(200)
		r.Write(bodyOut)
	}
}

func FromRESTGetByID(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		paramID, err := strconv.Atoi(mux.Vars(req)["ID"])
		if err != nil {
			err503("can't get request parameter", err, calendar.Logger, r)
		}
		ev, _ := calendar.Storage.GetByID(event.ID(paramID))
		event, err := event2pubEvent(ev)
		if err != nil {
			err503("can't convert types", err, calendar.Logger, r)
		}
		bodyOut,err := json.Marshal(&event)
		r.WriteHeader(200)
		r.Write(bodyOut)


		log.Println(paramID)
		r.WriteHeader(555)
	}
}

func FromRESTGetByDate(calendar *calendar.App) http.HandlerFunc {
	return func(r http.ResponseWriter, req *http.Request) {
		paramRange := mux.Vars(req)["Range"]
		paramDate, err := time.Parse("2006-01-02 15:04:05", mux.Vars(req)["Date"])
		if err != nil {
			err503("can't parse date from request parameter", err, calendar.Logger, r)
		}
		evs, err := calendar.Storage.GetByDate(paramDate, paramRange)
		if err != nil {
			err503("can't list events through HTTP API", err, calendar.Logger, r)
		}
		events, err := events2pubEvents(evs)
		if err != nil {
			err503("can't convert types", err, calendar.Logger, r)
		}
		bodyOut,err := json.Marshal(&events)
		r.WriteHeader(200)
		r.Write(bodyOut)
	}
}

func err503(s string, err error, l logger.Interface, r http.ResponseWriter) {
	l.Errorf(s,": ", err.Error())
	r.WriteHeader(503)
}