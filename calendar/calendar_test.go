package calendar

import (
	"github.com/yegres025/app/events"
	"testing"
)

func TestAdd(t *testing.T) {
	c := NewCalendar(nil)
	event, err := c.AddEvent("Memes", "7/7/2025", events.PriorityLow)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if event == nil {
		t.Errorf("expected event, got nil")
	}

	if event.Title != "Memes" {
		t.Errorf("expected title Memes, got %v", event.Title)
	}

}

func TestSearch(t *testing.T) {
	c := NewCalendar(nil)

	added, err := c.AddEvent("Memes", "7/7/2025", events.PriorityLow)

	if err != nil {
		t.Fatalf("AddEvent error %v", err)
	}

	id, err := c.Search("Memes")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if id == "" || added.ID != id {
		t.Errorf("error search id %v", err)
	}
}

func TestShowMap(t *testing.T) {
	c := NewCalendar(nil)

	_, err := c.AddEvent("Memes", "7/7/2025", events.PriorityLow)

	if err != nil {
		t.Fatalf("AddEvent error %v", err)
	}

	eventMap, err := c.ShowEvents()

	if eventMap == nil {
		t.Errorf("expected map, but got nil")
	}

	if err != nil {
		t.Errorf("unexpected error, but %v", err)
	}
}
