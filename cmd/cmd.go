package cmd

import (
	"errors"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/yegres025/app/calendar"
	"github.com/yegres025/app/events"
	"github.com/yegres025/app/logger"
	"github.com/yegres025/app/reminder"
	"os"
	"strings"
	"sync"
)

type Cmd struct {
	calendar *calendar.Calendar
	log      []string
	mutexLog sync.Mutex
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	cmd := strings.ToLower(parts[0])
	c.saveLog(strings.Join(parts, " "))
	switch cmd {
	case ADD:
		if len(parts) < 4 {
			fmt.Println("Format: add \"название события\" \"дата и время\" \"приоритет\"")
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println("Add error:", err)
			c.saveLog(err.Error())
		} else {
			fmt.Println("Event:", e.Title, "added")
			c.saveLog("Event:" + e.Title + "added")
		}
	case EXIT:
		defer logger.CloseFile()
		err = c.calendar.Save()
		c.calendar.Close()
		os.Exit(0)
	case LIST:
		events, err := c.calendar.ShowEvents()
		if err != nil {
			fmt.Println(err)
			c.saveLog(err.Error())
			logger.Error("error in ShowEvents " + err.Error())
		}

		for _, e := range events {
			line := fmt.Sprintf("%s - %s", e.Title, e.Priority)
			fmt.Println(line)
			c.saveLog(line)
			logger.Info("show event from ShowEvents: " + line)
		}
		logger.System("the function is called: ShowEvents")
	case REMOVE:
		title := parts[1]

		notification, err := c.calendar.RemoveEvent(title)
		c.saveLog(notification)
		if err != nil {
			fmt.Println(err.Error())
			c.saveLog(err.Error())
		}

	case UPDATE:
		prevTitle := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])
		err := c.calendar.EditEvent(prevTitle, title, date, priority)
		line := fmt.Sprintf("%s - %s - %s - %s- %s", "Updated", prevTitle, title, date, priority)
		logger.System("the function is called: EditEvent")
		logger.Info("update event " + line)
		c.saveLog(line)
		fmt.Println(line)
		if err != nil {
			fmt.Println(err.Error())
			logger.Error("error in EditEvent" + err.Error())
			c.saveLog(err.Error())
		}

	case ADD_REMINDER:
		title := parts[1]
		message := parts[2]
		at := parts[3]

		r, err := c.calendar.SetEventReminder(title, message, at)
		if errors.Is(err, reminder.ErrDateFormat) {
			fmt.Println("Can't set reminder with incorrect date")
			logger.Error("error in SetEventReminder: Can't set reminder with incorrect date")
		}

		if errors.Is(err, reminder.ErrEmptyTitle) {
			fmt.Println("Can't set reminder witn empty message")
			logger.Error("error in SetEventReminder: Can't set reminder witn empty message")
		}
		c.saveLog(r)
		logger.System("the function is called: SetEventReminder")
		logger.Info("added reminder" + r)
	case REMOVE_REMINDER:
		title := parts[1]

		m, err := c.calendar.CancelEventReminder(title)
		c.saveLog(m)
		if err != nil {
			fmt.Println(err.Error())
			c.saveLog(err.Error())
		}
	case SHOW_LOG:
		c.showLogList()

	default:
		fmt.Println("Unknown input team")
		fmt.Println("Enter ‘help’ for a list of commands")
	}

	fmt.Println(">>", input)
	c.saveLog(input)
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestion := []prompt.Suggest{
		{Text: ADD, Description: "Добавить событие"},
		{Text: LIST, Description: "Показать все события"},
		{Text: REMOVE, Description: "Удалить событие"},
		{Text: UPDATE, Description: "Изменить событие"},
		{Text: ADD_REMINDER, Description: "Добавить напоминание на событие"},
		{Text: REMOVE_REMINDER, Description: "Удалить напоминание на событии\n title"},
		{Text: HELP, Description: "Показать справку"},
		{Text: EXIT, Description: "Выйти из программы"},
		{Text: SHOW_LOG, Description: "Показать логи"},
	}

	return prompt.FilterHasPrefix(suggestion, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	err := c.calendar.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)

	go func() {
		for msg := range c.calendar.Notification {
			fmt.Println(msg)
			c.saveLog(msg)
		}
	}()
	p.Run()
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
	}
}

func (c *Cmd) saveLog(value string) {
	c.mutexLog.Lock()
	defer c.mutexLog.Unlock()
	c.log = append(c.log, value)
}

func (c *Cmd) showLogList() {
	c.mutexLog.Lock()
	defer c.mutexLog.Unlock()
	if len(c.log) == 0 {
		fmt.Println("Logs list is empty")
		return
	}

	for _, log := range c.log {
		fmt.Println(log)
	}
}
