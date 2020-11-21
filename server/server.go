package main

import (
	"fmt"
	"log"
	"net/http"

	"wakelet-challenge/events-repository"
	"wakelet-challenge/nasa"
)

func main() {
	// Get events
	events, err := nasa.GetEvents()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(events)

	// Add to db
	eventsrepository.CreateMany(events)

	http.HandleFunc("/events", HandleEvents)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HandleEvents")
}
