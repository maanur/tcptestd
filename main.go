package main

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
)

// Operator executes appropriate command with arguments, passed to it as single string
type Operator interface {
	CallName() string
	Operate(io.Writer, string) func(*gin.Context)
}

// Integrator runs (while context is not expired) on background and provides appropriate gin.HandlerFunc for returning results to user.
type Integrator interface {
	CallName() string
	Run(*context.Context, io.Writer)
	HandlerFunc(*gin.Context)
}

var queries = make(chan context.Context)
var run = make(chan func())

func main() {
	fmt.Println(len(receivers))
	var wg sync.WaitGroup
	wg.Add(1)
	go func(w io.Writer) {
		defer wg.Done()
		web(w)
	}(logger())
	wg.Wait()
}
