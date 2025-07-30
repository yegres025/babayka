package calendar

import (
	"fmt"
	"github.com/yegres025/app/events"
)

var EventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	EventsMap[key] = e
	fmt.Println("Событие добавлено", e)
}

func ShowEvents(eventsMap map[string]events.Event) {
	for _, v := range eventsMap {
		fmt.Println(v.Title, "-", v.StartAt.Format("02.01.2006 15:04"))
	}
}
