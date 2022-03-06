package event

import (
	"dev-11/internal/apperror"
	"errors"
	"log"
	"net/http"
	"time"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Log(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appError *apperror.AppError
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		t1 := time.Now()
		err := h(w, r)
		t2 := time.Now()

		if err != nil {
			if errors.As(err, &appError) {
				if errors.Is(err, NoRouteError) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(NoRouteError.Marshal())

					log.Printf("Method: %s, Path: %s%s, Status:%d, %s", r.Method, r.Host, r.RequestURI, http.StatusNotFound, t2.Sub(t1).String())
					return
				} else if errors.Is(err, MethodNotAllowed) {
					w.WriteHeader(http.StatusMethodNotAllowed)
					w.Write(MethodNotAllowed.Marshal())

					log.Printf("Method: %s, Path: %s%s, Status:%d, %s", r.Method, r.Host, r.RequestURI, http.StatusMethodNotAllowed, t2.Sub(t1).String())
					return
				}
			}

			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write(apperror.NewSystemError(err).Marshal())

			log.Printf("Method: %s, Path: %s%s, Status:%d, %s", r.Method, r.Host, r.RequestURI, http.StatusServiceUnavailable, t2.Sub(t1).String())
			return
		}
		w.WriteHeader(http.StatusOK)

		log.Printf("Method: %s, Path: %s%s, Status:%d, %s", r.Method, r.Host, r.RequestURI, http.StatusOK, t2.Sub(t1).String())
	}
}
