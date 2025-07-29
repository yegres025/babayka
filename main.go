package main

import (
	"fmt"
	"github.com/yegres025/app/calendar"
	"github.com/yegres025/app/events"
	"time"
)

func main() {
	e := events.Event{
		Title:   "встреча",
		StartAt: time.Now(),
	}

	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлен")
	fmt.Scanln()
}
