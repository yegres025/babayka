package calendar

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/yegres025/app/events"
)

var calendarEvents = make(map[string]events.Event)

func AddEvent(title, date, priority string) (events.Event, error) {
	e, err := events.NewEvent(title, date, priority)

	if err != nil {
		return events.Event{}, err
	}

	calendarEvents[e.ID] = e
	return e, err
}

func ShowEvents() {
	for _, v := range calendarEvents {
		fmt.Println(v.Title, "-", v.StartAt.Format("02.01.2006 15:04"), "(", v.Priority, ")")
	}
}

func RemoveEvent(key string) {
	_, ok := calendarEvents[key]

	if ok {
		fmt.Println("Задача", calendarEvents[key].Title, "удалена")
		delete(calendarEvents, key)
		return
	} else {
		fmt.Println("Нет такой задачи")
	}
}

func ChangeEvent(key, newTitle, newDate, newPriority string) {
	event, ok := calendarEvents[key]
	if !ok {
		fmt.Println("Задача", key, "не найдена")
		return
	}

	t, dateErr := dateparse.ParseAny(newDate)
	if dateErr != nil {
		fmt.Println("Некорректный формат даты")
		return
	}

	validateTitle := events.IsValidateTitle(newTitle)
	if !validateTitle {
		fmt.Println("Некорректное имя задачи -", newTitle)
		return
	}

	validatePriority := events.IsValidateTitle(newPriority)
	if !validatePriority {
		fmt.Println("Некорректное значение приоритета -", newPriority)
	}

	event.Title = newTitle
	event.StartAt = t
	event.Priority = newPriority
	calendarEvents[key] = event
	fmt.Println("Задача обновлена:", newTitle, "-", newDate)
}
