package event

import (
	"dev-11/internal/apperror"
	"dev-11/internal/delivery/api"
	"dev-11/internal/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	NoRouteError     = apperror.NewAppError(nil, "no route", "no route", "ES-4-4")
	MethodNotAllowed = apperror.NewAppError(nil, "method not allowed", "method not allowed", "ES-4-5")
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) api.Handler {
	return &handler{service: service}
}

func (h handler) NoRoute(w http.ResponseWriter, r *http.Request) error {
	return NoRouteError
}

func (h handler) Register(router *http.ServeMux) {
	router.HandleFunc("*", Log(h.NoRoute))

	router.HandleFunc("/create_event", Log(h.CreateEvent))
	router.HandleFunc("/update_event", Log(h.UpdateEvent))
	router.HandleFunc("/delete_event", Log(h.DeleteEvent))

	router.HandleFunc("/events_for_day", Log(h.EventsForDay))
	router.HandleFunc("/events_for_week", Log(h.EventsForWeek))
	router.HandleFunc("/events_for_month", Log(h.EventsForMonth))

}

func (h handler) CreateEvent(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var dto EventDTO
		err = json.Unmarshal(body, &dto)
		if err != nil {
			return err
		}

		event := dto.toEvent()
		h.service.SetEvent(dto.UserID, event)

		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}

func (h handler) UpdateEvent(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var dto EventDTO
		err = json.Unmarshal(body, &dto)
		if err != nil {
			return err
		}
		event := dto.toEvent()

		h.service.SetEvent(dto.UserID, event)

		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}

func (h handler) DeleteEvent(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var dto EventDTO
		err = json.Unmarshal(body, &dto)
		if err != nil {
			return err
		}
		event := dto.toEvent()

		h.service.SetEvent(dto.UserID, event)

		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}

func (h handler) EventsForDay(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		params := r.URL.Query()
		date := params.Get("date")
		uuid := params.Get("user_id")
		t, err := h.service.IsValid(date)
		if err != nil {
			return err
		}

		n, err := strconv.Atoi(uuid)
		events := h.service.GetEvents(n, t, time.Hour*24)
		data, err := json.Marshal(events)
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}

func (h handler) EventsForWeek(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		params := r.URL.Query()
		date := params.Get("date")
		uuid := params.Get("user_id")
		t, err := h.service.IsValid(date)
		if err != nil {
			return err
		}

		n, err := strconv.Atoi(uuid)
		events := h.service.GetEvents(n, t, time.Hour*24*7)
		data, err := json.Marshal(events)
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}

func (h handler) EventsForMonth(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		params := r.URL.Query()
		date := params.Get("date")
		uuid := params.Get("user_id")
		t, err := h.service.IsValid(date)
		if err != nil {
			return err
		}

		n, err := strconv.Atoi(uuid)
		events := h.service.GetEvents(n, t, time.Hour*24*30)
		data, err := json.Marshal(events)
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	default:
		return MethodNotAllowed
	}
	return nil
}
