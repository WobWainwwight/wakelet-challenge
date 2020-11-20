package nasa

import (
	"encoding/json"
	"net/http"
)

func GetEvents() ([]Event, error) {
	nasaEndpoint := "https://eonet.sci.gsfc.nasa.gov/api/v3/events"

	resp, err := http.Get(nasaEndpoint)

	if err != nil {
		return nil, err
	}

	var eventsWrapper EventsWrapper

	decodingErr := json.NewDecoder(resp.Body).Decode(&eventsWrapper)

	if decodingErr != nil {
		return nil, decodingErr
	}

	return eventsWrapper.Events, nil
}
