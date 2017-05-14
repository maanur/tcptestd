package sheduler

import "time"

type simpleshed struct{}

// TestShedule создает простое расписание для теста
var TestShedule simpleshed

func (simpleshed) Mkshed(tester Tester) *Shedule {
	sh := new(Shedule)
	sh.Proximity = 10 * time.Second
	var t Event
	t.Tester = tester
	t.Time = time.Now().Round(sh.Proximity).Add(sh.Proximity)
	sh.Events = append(sh.Events, t)
	for i := 1; i <= 10; i++ {
		t.Time = t.Time.Add(sh.Proximity)
		sh.Events = append(sh.Events, t)
	}
	return sh
}
