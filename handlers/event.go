package handlers

import (
	"net/http"

	"github.com/leeola/hob"
	"github.com/leeola/hob/middleware"
	"github.com/pressly/chi"
)

func EventHandler(events hob.Events, actions hob.Actions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := middleware.GetLog(r)

		eventName := chi.URLParam(r, "event")
		actionName, ok := events[eventName]
		if !ok {
			log.Info("unknown event requested", "eventName", eventName)
			http.Error(w, http.StatusText(404), 404)
			return
		}

		action, ok := actions[actionName]
		if !ok {
			log.Info("unknown action configured", "actionName", actionName)
			http.Error(w, http.StatusText(404), 404)
			return
		}

		action.Act(w, r)
	}
}
