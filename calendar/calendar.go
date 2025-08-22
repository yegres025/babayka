package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yegres025/babayka/events"
	"github.com/yegres025/babayka/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

type Duplicate struct {
	ID    string
	Title string
}

func NewCalendar(s storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: map[string]*events.Event{},
		storage:        s,
		Notification:   make(chan string),
	}
}

func (c *Calendar) AddEvent(title, date string, priority events.Priority) (*events.Event, error) {
	e, err := events.NewEvent(title, date, priority)

	if err != nil {
		return &events.Event{}, err
	}

	c.calendarEvents[e.ID] = e
	return e, err
}

func (c Calendar) ShowEvents() (map[string]events.Event, error) {
	if len(c.calendarEvents) == 0 {
		return nil, errors.New("Calendar is empty")
	}

	eventsMap := make(map[string]events.Event)

	if len(c.calendarEvents) == 0 {

		return nil, errors.New("Calendat is empty")
	}

	for id, v := range c.calendarEvents {
		eventsMap[id] = *v
	}
	return eventsMap, nil
}

func (c *Calendar) RemoveEvent(title string) (string, error) {
	for id, event := range c.calendarEvents {
		if event.Title == title {
			delete(c.calendarEvents, id)
			return "Task " + c.calendarEvents[id].Title + " deleted", nil
		}
	}
	return "", errors.New("There is no such task")
}

func (c *Calendar) EditEvent(prevTitle, newTitle, newDate string, newPriority events.Priority) error {
	isExisting := false
	var targetEvent *events.Event

	for _, event := range c.calendarEvents {
		if event.Title == prevTitle {
			isExisting = true
			targetEvent = event
			break
		}

	}
	if isExisting {
		validateTitle := events.IsValidateTitle(newTitle)
		if !validateTitle {
			return errors.New("Incorrect task name")
		}
		err := targetEvent.Update(newTitle, newDate, newPriority)
		return err
	}

	return nil
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)

	if err != nil {
		return errors.New("Serialization error: " + err.Error())
	}

	err = c.storage.Save(data)
	return err
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.calendarEvents)
	return err
}

func (c *Calendar) SetEventReminder(title, message, at string) (string, error) {
	id, err := c.Search(title)
	if err != nil {
		return "", err
	}

	e, exist := c.calendarEvents[id]
	if !exist {
		return "", errors.New("Task not found")
	}

	err = e.AddReminder(message, at, c.Notify)

	if err != nil {
		return "", fmt.Errorf("set reminder failed: %w", err)
	}
	return "Add reminder:" + message + at, nil
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) Close() {
	if c.Notification != nil {
		close(c.Notification)
		c.Notification = nil
	}
}

func (c *Calendar) CancelEventReminder(title string) (string, error) {
	id, err := c.Search(title)

	if err != nil {
		return "", err
	}
	e, exist := c.calendarEvents[id]

	if !exist {
		return "", errors.New("Task not found")
	}

	if e.Reminder != nil {
		e.Reminder.Stop()
		e.RemoveReminder()
	}
	return "Reminder deleted", nil
}

func (c *Calendar) DuplicateChecker(title string) (bool, string) {
	counter := 0
	for _, t := range c.calendarEvents {
		if title == t.Title {
			counter += 1
		}

		if counter > 1 {
			return true, title
		}

	}
	return false, ""
}

func (c *Calendar) ShowDuplicates(title string) []Duplicate {
	var duplicates []Duplicate

	for _, e := range c.calendarEvents {
		if title == e.Title {
			duplicates = append(duplicates, Duplicate{Title: title, ID: e.ID})
		}
	}
	return duplicates
}

func (c *Calendar) Search(title string) (string, error) {
	for _, e := range c.calendarEvents {
		if title == e.Title {
			return e.ID, nil
		}
	}
	return "", errors.New("Event not found")
}
