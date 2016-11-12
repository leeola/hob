package handlers

import "net/http"

func LongpollHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(501), 501)
	}
}
