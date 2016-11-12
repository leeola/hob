package hob

import (
	"errors"
	"net/http"
)

type Events map[string]string
type Actions map[string]Action

type Linker struct {
	events  Events
	actions Actions
}

func NewLinker(e Events, a Actions) *Linker {
	return &Linker{
		events:  e,
		actions: a,
	}
}

func (l *Linker) Trigger(e string, w http.ResponseWriter, r *http.Request) error {
	return errors.New("not impl")
}
