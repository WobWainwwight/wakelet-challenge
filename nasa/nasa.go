package nasa

import (
	"encoding/json"
	"net/http"
)

func GetEvents() ([]Event, error) {
	nasaEndpoint := "https://eonet.sci.gsfc.nasa.gov/api/v3/events"

	resp, err := http.Get(nasaEndpoint)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var eventsWrapper EventsWrapper

	jsonErr := json.NewDecoder(resp.Body).Decode(&eventsWrapper)

	// I want the every nasa event to have id = nasa_event
	// and every EventId == nasa event id (EONET_123)
	if len(eventsWrapper.Events) == 0 {
		return nil, jsonErr
	}

	for i := range eventsWrapper.Events {
		eventsWrapper.Events[i].ID = "nasa_event"
	}

	return eventsWrapper.Events, nil
}
