package tcptestd

import "time"

// Shedule - массив значений времени
type Shedule struct {
	Time      []time.Time
	Proximity time.Duration
}
