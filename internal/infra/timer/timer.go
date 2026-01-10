package timer

import (	
	"time"
)

type Timer struct {
	Duration time.Duration
	TimeStart time.Time
	TimeEnd time.Time
}

func NewTimer() *Timer {
	return &Timer{
		TimeStart: time.Now(),
	}
}

func (t *Timer) TimeClose()  {
	t.TimeEnd = time.Now()
	t.Duration = t.TimeEnd.Sub(t.TimeStart).Truncate(time.Second)
}


