package main

import (
	"context"
	"fmt"
	"io"
	"sync"
)

var receivers []func(io.Writer)
var queries = make(chan context.Context)
var run = make(chan func())

func main() {
	fmt.Println(len(receivers))
	var wg sync.WaitGroup
	l := newLogger()
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.runBuf()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.runChan()
	}()
	for _, receiver := range receivers {
		wg.Add(1)
		go func(f func(io.Writer)) {
			defer wg.Done()
			f(l.Writer()) //Вывод
		}(receiver)
	}
	wg.Wait()
}
