package nasa

import (
	"time"
)

type EventsWrapper struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Events      []Event `json:"events"`
}

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Closed      string `json:"closed"`
	Categories  []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"categories"`
	Sources []struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"sources"`
	Geometry []struct {
		MagnitudeValue interface{} `json:"magnitudeValue"`
		MagnitudeUnit  interface{} `json:"magnitudeUnit"`
		Date           time.Time   `json:"date"`
		Type           string      `json:"type"`
		Coordinates    []float64   `json:"coordinates"`
	} `json:"geometry"`
}
