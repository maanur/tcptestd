package tcptestd

import "time"

type simpleshed struct{}

func (simpleshed) Mkshed(tester Tester) *Shedule {
	sh := new(Shedule)
	sh.Proximity = 10 * time.Second
	var t ShAction
	t.Tester = tester
	t.Time = time.Now().Round(sh.Proximity).Add(sh.Proximity)
	sh.Items = append(sh.Items, t)
	for i := 1; i <= 10; i++ {
		t.Time = t.Time.Add(sh.Proximity)
		sh.Items = append(sh.Items, t)
	}
	return sh
}
