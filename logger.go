package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

type logstr struct {
	String string
}

type logger struct {
	wr      io.Writer
	started bool
}

var logr logger

var logbuf = []logstr{logstr{String: "one"}, logstr{String: "two"}}

func (logr *logger) writer() io.Writer {
	if logr.started {
		return logr.wr
	}
	return os.Stdout
}

func (logr *logger) runLogger(ctx context.Context) {
	f, err := os.Create(time.Now().Format("0102150405") + ".log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logr.wr = io.MultiWriter(os.Stdout, f)
	logr.started = true
	<-ctx.Done()
	log.Println("runLogger ended")
}

/*
var logch chan logstr

type logger struct {
	writer io.Writer
	reader io.Reader
}

func newLogger() *logger {
	l := new(logger)
	l.reader, l.writer = io.Pipe()
	return l
}

func (l *logger) runChan() {
	for {
		str := make([]byte, 10)
		_, err := l.reader.Read(str)
		if err != nil || err != io.EOF {
			log.Fatal(err)
		}
		fmt.Println("gotcha")
		logch <- logstr{String: string(str)}
	}
}

func (l *logger) Writer() io.Writer {
	return io.MultiWriter(os.Stdout, l.writer)
}

func (l *logger) runBuf() {
	for {
		select {
		case str := <-logch:
			logbuf = append(logbuf, str)
		default:
		}
		if len(logbuf) > 20 {
			logbuf = logbuf[1:]
		}
	}
}
*/
//TODO: сохранять логи в файлики по датам и добавлять в архив (tar?)
