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
