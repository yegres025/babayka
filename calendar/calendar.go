package calendar

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/yegres025/app/events"
)

var EventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	validate := events.IsValidateTitle(e.Title)

	if !validate {
		fmt.Println("Некорректное имя задачи -", e.Title)
		return
	}

	EventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
}

func ShowEvents(eventsMap map[string]events.Event) {
	for _, v := range eventsMap {
		fmt.Println(v.Title, "-", v.StartAt.Format("02.01.2006 15:04"))
	}
}

func RemoveEvent(eventsMap map[string]events.Event, key string) {
	_, ok := eventsMap[key]

	if ok {
		fmt.Println("Задача", key, "удалена")
		delete(eventsMap, key)
		return
	} else {
		fmt.Println("Нет такой задачи")
	}
}

func ChangeEvent(eventsMap map[string]events.Event, key string, newTitle string, newDate string) {
	event, ok := eventsMap[key]
	validate := events.IsValidateTitle(newTitle)
	t, dateErr := dateparse.ParseAny(newDate)

	if !ok {
		fmt.Println("Задача", key, "не найдена")
		return
	}

	if !validate {
		fmt.Println("Некорректное имя задача - ", newTitle)
		return
	}

	if dateErr != nil {
		fmt.Println("Некорректный формат даты")
		return
	}

	event.Title = newTitle
	event.StartAt = t
	eventsMap[key] = event
	fmt.Println("Задача обновлена:", newTitle, "-", newDate)
}
