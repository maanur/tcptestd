package tcptestd

import (
	"time"
)

// ShAction представляет собой запланированный на Time запуск тестера Tester
type Event struct {
	Time   time.Time
	Tester *Tester
}

// Shedule - массив значений времени
type Shedule struct {
	Events    []Event
	Proximity time.Duration
}

var Sheduler chan Event

// Run проверяет время с интервалом в sh.Proximity и выполняет запуск
func (sh *Shedule) Run() {
	t := time.NewTicker(sh.Proximity)
	for {
		now := <-t.C
		for _, event := range sh.Events {
			if now.Round(sh.Proximity).Equal(evemt.Time.Round(sh.Proximity)) {
				Sheduler <- event.Tester.Run()
			}
		}
	}
}
