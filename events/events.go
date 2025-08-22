package events

import (
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/yegres025/app/reminder"
	"regexp"
	"time"
)

type Event struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	StartAt  time.Time `json:"start_at"`
	Priority Priority  `json:"priority"`
	Reminder *reminder.Reminder
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title, dateStr string, priority Priority) (*Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return &Event{}, errors.New("Incorrect date format - " + err.Error())
	}

	validateTitle := IsValidateTitle(title)
	if !validateTitle {
		return &Event{}, errors.New("Incorrect task name")
	}

	err = priority.Validate()
	if err != nil {
		return &Event{}, err
	}

	return &Event{Title: title, StartAt: t, Priority: priority, ID: getNextID(), Reminder: nil}, nil
}

func IsValidateTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-ЯёЁ0-9 ,\\.]{3,30}$"

	matched, err := regexp.MatchString(pattern, title)

	if err != nil {
		return false
	}
	return matched
}

func (e *Event) Update(title, date string, priority Priority) error {
	t, err := dateparse.ParseAny(date)
	if err != nil {
		return errors.New("Incorrect date format")
	}

	e.Title = title
	e.StartAt = t
	e.Priority = Priority(priority)
	return nil
}

func (e Event) Print() {
	fmt.Println(e.Title, e.StartAt)
}

func (e *Event) AddReminder(message, at string, notify func(msg string)) error {
	r, err := reminder.NewReminder(message, at)

	if err != nil {
		return fmt.Errorf("can't added reminder: %w", err)
	}

	e.Reminder = r
	e.Reminder.Start(notify)
	return nil
}

func (e *Event) RemoveReminder() {
	e.Reminder = nil
}
