package main

import (
	"github.com/yegres025/app/calendar"
	"github.com/yegres025/app/events"
)

func main() {
	e, err := events.NewEvent("Встреча", "2024-07-15 09:30")

	if err != nil {
		println(err.Error())
		return
	}

	calendar.AddEvent("event1", e)
	calendar.AddEvent("event2", e)
	calendar.AddEvent("event3", e)

	calendar.ShowEvents(calendar.EventsMap)
	calendar.RemoveEvent(calendar.EventsMap, "event2")
	calendar.ShowEvents(calendar.EventsMap)
	calendar.ChangeEvent(calendar.EventsMap, "event1", "Попыхтеть", "2024-08-15 12:30")
	calendar.ShowEvents(calendar.EventsMap)
}
