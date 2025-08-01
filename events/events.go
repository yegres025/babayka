package events

import (
	"errors"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"regexp"
	"time"
)

type Event struct {
	ID       string
	Title    string
	StartAt  time.Time
	Priority string
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title, dateStr, priority string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат даты")
	}

	validateTitle := IsValidateTitle(title)

	if !validateTitle {
		return Event{}, errors.New("Некорректное имя задачи")
	}

	return Event{Title: title, StartAt: t, ID: getNextID(), Priority: priority}, nil
}

func IsValidateTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-ЯёЁ0-9 ,\\.]{3,30}$"

	matched, err := regexp.MatchString(pattern, title)

	if err != nil {
		return false
	}
	return matched
}
