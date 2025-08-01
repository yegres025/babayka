package main

import (
	"fmt"
	"github.com/yegres025/app/calendar"
)

func main() {

	e1, err1 := calendar.AddEvent("Поспать", "2024-07-15 09:30", "low")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
		return
	}

	e2, err2 := calendar.AddEvent("Поесть", "2024-07-15 22:10", "medium")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
		return
	}

	_, err3 := calendar.AddEvent("Погулять", "2024-07-15 19:30", "high")
	if err3 != nil {
		fmt.Println("Ошибка:", err3)
		return
	}
	calendar.ShowEvents()
	calendar.RemoveEvent(e1.ID)
	calendar.ShowEvents()
	calendar.ChangeEvent(e2.ID, "Пошабить", "2025-07-15 19:30", "high")
	calendar.ShowEvents()

}
