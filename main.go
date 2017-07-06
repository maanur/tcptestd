package main

import (
	"context"
	"io"

	"time"

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
	Run(context.Context, io.Writer)
	HandlerFunc(*gin.Context)
}

var queries = make(chan context.Context)
var run = make(chan func())

func main() {
	mainctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	go logr.runLogger(mainctx)
	go web.Run(mainctx, logr.writer())
	for {
		select {
		case <-mainctx.Done():
			break
		default:
		}
	}
}
