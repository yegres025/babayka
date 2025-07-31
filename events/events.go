package events

import (
	"errors"
	"github.com/araddon/dateparse"
	"regexp"
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

func IsValidateTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-ЯёЁ0-9 ,\\.]{3,30}$"

	matched, err := regexp.MatchString(pattern, title)

	if err != nil {
		return false
	}
	return matched
}
