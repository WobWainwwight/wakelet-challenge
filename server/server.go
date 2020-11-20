package main

import (
	"fmt"
	"log"
	"net/http"

	"wakelet-challenge/nasa"
)

func main() {
	// Add events to DynamoDb
	events, err := nasa.GetEvents()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(events)
	}

	http.HandleFunc("/events", HandleEvents)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HandleEvents")
}
