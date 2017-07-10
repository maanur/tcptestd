package main

import (
	"context"
	"io"
	"log"

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

var backRunList []Integrator

func main() {
	work, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	logr.runLogger(work) //Запуск обработчика лога
	log := log.New(logr.writer(), "[main] ", log.Lshortfile|log.LstdFlags)
	for _, backRun := range backRunList {
		go backRun.Run(work, logr.writer())
		log.Println("Started " + backRun.CallName())
	}
	<-work.Done()
}
