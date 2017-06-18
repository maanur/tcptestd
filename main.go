package main

import (
	"fmt"
	"sync"
)

var receivers []func()

func main() {
	fmt.Println(len(receivers))
	var wg sync.WaitGroup
	for _, receiver := range receivers {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(receiver)
	}
	wg.Wait()
}
