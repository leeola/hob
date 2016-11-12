package routes

import (
	"net/http"

	"github.com/leeola/hob"
	"github.com/leeola/hob/handlers"
	"github.com/pressly/chi"
)

func ApiDefault(e hob.Events, a hob.Actions) http.Handler {
	return ApiVersion0(e, a)
}

func ApiVersion0(e hob.Events, a hob.Actions) http.Handler {
	r := chi.NewRouter()
	r.Get("/events/:event", handlers.EventHandler(e, a))
	r.Get("/actions/:action/longpoll", handlers.LongpollHandler())
	return r
}
