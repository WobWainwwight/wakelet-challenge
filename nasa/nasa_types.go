package nasa

type EventsWrapper struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Events      []Event `json:"events"`
}

type Event struct {
	ID          string `dynamodbav:"id"`
	EventID     string `json:"id" dynamodbav:"event_id" `
	Title       string `json:"title" dynamodbav:"title"`
	Time        string `json:"time" dynamodbav:"time"`
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
	Geometry []GeometryEvent `json:"geometry"`
}

type GeometryEvent struct {
	MagnitudeValue float64   `json:"magnitudeValue"`
	MagnitudeUnit  string    `json:"magnitudeUnit"`
	Date           string    `json:"date"`
	Type           string    `json:"type"`
	Coordinates    []float64 `json:"coordinates"`
}
