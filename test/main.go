package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maanur/tcptestd/sheduler"
)

func main() {
	var wg sync.WaitGroup
	var test tester
	shed := sheduler.TestShedule.Mkshed(test)
	/*for _, ev := range shed.Events {
		fmt.Println(ev.Time.String())
	}*/
	fmt.Println("StdOut flag...")
	go func() {
		for {
			ev := <-shed.Sheduler
			ev.Tester.Run()
		}
	}()
	wg.Add(1)
	go func() {
		shed.Run()
		defer wg.Done()
	}()
	wg.Wait()
}

type tester struct{}

func (t tester) Run() {
	fmt.Println(time.Now().String() + " - tested")
}
