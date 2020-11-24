package nasa

import (
	"fmt"
	"testing"
	"time"
)

func TestMapGeometriesToEvents(t *testing.T) {
	firstGeometryEvent := GeometryEvent{
		Date: time.Now(),
		Type: "first",
	}

	secondGeometryEvent := GeometryEvent{
		Date: time.Now(),
		Type: "second",
	}

	event := Event{
		ID:      "id1",
		EventID: "eid1",
		Title:   "title",
		Geometry: []GeometryEvent{
			firstGeometryEvent,
			secondGeometryEvent,
		},
	}

	events := MapGeometriesToEvents(&event)

	if len(events) != 2 {
		t.Errorf("Wanted 2 events got %d", len(events))

	}

	if events[0].Time != firstGeometryEvent.Date || events[1].Time != secondGeometryEvent.Date {
		t.Error("Didn't map the geometry's date to event's time")
	}
	fmt.Print(events)
}
