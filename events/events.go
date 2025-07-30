package events

import (
	"errors"
	"github.com/araddon/dateparse"
	"time"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат даты")
	}

	return Event{Title: title, StartAt: t}, nil
}
