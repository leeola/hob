package hob

import "net/http"

type Action interface {
	Act(w http.ResponseWriter, r *http.Request)
}
