package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logstr struct {
	String string
}

var logch chan logstr
var logbuf []logstr

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

//TODO: сохранять логи в файлики по датам и добавлять в архив (tar?)
