package handlers

import (
	"climbing-journal/data"
	"log"
	"net/http"
)

type Entries struct {
	l *log.Logger
}

func NewEntries(l *log.Logger) *Entries {
	return &Entries{l}
}

func (e *Entries) GetEntries(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle GET entries")

	lp := data.GetEntries()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert JSON", http.StatusInternalServerError)
	}
}
