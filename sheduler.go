package tcptestd

import (
	"sync"
	"time"
)

// Event представляет собой событие: в момент Time запустить тестер Tester
type Event struct {
	Time   time.Time
	Tester Tester
}

// Shedule - массив значений времени
type Shedule struct {
	Events    []Event
	Proximity time.Duration
	Sheduler  chan *Event
}

// Run проверяет время с интервалом в sh.Proximity и передает событие в канал Sheduler
func (sh *Shedule) Run() {
	var wg sync.WaitGroup
	t := time.NewTicker(sh.Proximity)
	wg.Add(1)
	sh.Sheduler = make(chan *Event)
	go func() {
		for {
			now := <-t.C
			for _, event := range sh.Events {
				if now.Round(sh.Proximity).Equal(event.Time.Round(sh.Proximity)) {
					sh.Sheduler <- &event
				}
			}
			if now.After(sh.Events[len(sh.Events)-1].Time.Round(sh.Proximity)) {
				t.Stop()
				defer wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
