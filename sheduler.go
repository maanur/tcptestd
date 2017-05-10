package tcptestd

import (
	"time"
)

// Sheduler - некий объект, задающий расписание.
type Sheduler interface {
	Mkshed() Shedule
}

type simpleshed struct{}

func (simpleshed) Mkshed() (sh Shedule) {
	sh.Proximity = time.Minute
	t := time.Now().Round(sh.Proximity).Add(sh.Proximity)
	sh.Time = append(sh.Time, t)
	for i := 1; i <= 10; i++ {
		t = t.Add(sh.Proximity)
		sh.Time = append(sh.Time, t)
	}
	return
}
