package main

import (
	"fmt"
	"github.com/yegres025/app/calendar"
	"github.com/yegres025/app/cmd"
	"github.com/yegres025/app/storage"
)

func main() {
	s := storage.NewJsonStorage("calendar")
	c := calendar.NewCalendar(s)
	cli := cmd.NewCmd(c)
	cli.Run()

	err := c.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

}
