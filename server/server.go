package main

import (
	"encoding/json"
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
	// Get hold of query string
	orderBy := r.URL.Query().Get("order-by")

	if orderBy != "time" && orderBy != "title" {
		http.Error(w, "Order by time or title", http.StatusBadRequest)
		return
	}

	events, err := eventsrepository.GetEvents(orderBy)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)

}
