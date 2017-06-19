package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var receivers []func(io.Writer)

var run = make(chan func())

func main() {
	fmt.Println(len(receivers))
	var wg sync.WaitGroup
	for _, receiver := range receivers {
		wg.Add(1)
		go func(f func(io.Writer)) {
			defer wg.Done()
			f(os.Stdout)
		}(receiver)
	}
	wg.Wait()
	go func() {
		for {
			r := <-run
			r()
		}
	}()
}
