package data

import (
	"encoding/json"
	"io"
)

type Entry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Text  string `json:"text"`
}

func (e *Entry) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(e)
}

type Entries []*Entry

func (e *Entries) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(e)
}

func GetEntries() Entries {
	return entryList
}

var entryList = []*Entry{
	{
		ID:    1,
		Title: "Climbing day 1",
		Type:  "Climbing",
		Text:  "Bouldering in the gym",
	},
	{
		ID:    2,
		Title: "Climbing day 2",
		Type:  "Climbing",
		Text:  "Bouldering in the gym",
	},
}
