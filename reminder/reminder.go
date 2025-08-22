package reminder

import (
	"errors"
	"fmt"
	"time"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	timer   *time.Timer
}

var ErrEmptyTitle = errors.New("message is empty")
var ErrDateFormat = errors.New("incorrect date")

func NewReminder(message, at string) (*Reminder, error) {
	if len(message) <= 0 {
		return &Reminder{}, fmt.Errorf("can't create reminder: %w", ErrEmptyTitle)
	}

	d, err := time.ParseDuration(at)
	if err != nil {
		return &Reminder{}, fmt.Errorf("can't create reminder: %w", ErrDateFormat)
	}
	t := time.Now().Add(d)

	return &Reminder{
		Message: message,
		At:      t,
		Sent:    false,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Reminder!", r.Message)
	r.Sent = true
}

func (r *Reminder) Start(notify func(msg string)) {
	delay := r.At.Sub(time.Now())

	if delay < 0 {
		go func() {
			notify(r.Message)
			r.Sent = true
		}()
	}
	r.timer = time.AfterFunc(delay, func() {
		notify(r.Message)
		r.Sent = true
	})
}

func (r *Reminder) Stop() {
	if r.timer != nil {
		r.timer.Stop()
		r.timer = nil
	}
}
