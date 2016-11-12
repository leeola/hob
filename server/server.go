package server

import (
	"net/http"

	"github.com/inconshreveable/log15"
	"github.com/leeola/hob"
	"github.com/leeola/hob/middleware"
	"github.com/leeola/hob/routes"
	"github.com/pressly/chi"
)

func ListenAndServe(bind string, c hob.Config) {
	r := chi.NewRouter()

	r.Use(middleware.Logging(log15.New()))

	v0Routes := routes.ApiDefault(c.Events, c.Actions)
	r.Mount("/", v0Routes)
	r.Mount("/v0", v0Routes)

	http.ListenAndServe(bind, r)
}
