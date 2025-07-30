package main

import (
	"fmt"
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
	e, err = events.NewEvent("Сон", "2024-07-15 23:30")

	if err != nil {
		println(err.Error())
		return
	}

	calendar.AddEvent("event2", e)
	e, err = events.NewEvent("Обед", "2024-07-15 13:30")

	if err != nil {
		println(err.Error())
		return
	}

	calendar.AddEvent("event3", e)
	calendar.ShowEvents(calendar.EventsMap)
	fmt.Scanln()
}
