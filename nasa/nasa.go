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

	if len(eventsWrapper.Events) == 0 {
		return nil, jsonErr
	}

	return eventsWrapper.Events, nil
}
