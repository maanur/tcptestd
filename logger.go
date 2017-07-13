package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type logger struct {
	wr      io.Writer
	started chan struct{}
	buf     bytes.Buffer
}

var logr logger

func (logr *logger) writer() io.Writer {
	select {
	case <-logr.started:
		return logr.wr
	default:
		return os.Stdout
	}

}

func (logr *logger) runLogger(ctx context.Context) {
	logr.started = make(chan struct{})
	defer func() {
		select {
		case <-logr.started:
		default:
			close(logr.started)
		}
	}()
	var wg sync.WaitGroup
	log.Println("logr starting")
	wg.Add(1)
	go func(ctx context.Context) {
		// Вывод в файл (TODO: создавать новый лог каждый день, архивировать логи.)
		f, err := os.Create(time.Now().Format("0102150405") + ".log")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		//logr.wr = io.MultiWriter(os.Stdout, f, &logr.buf)
		//logr.buf надо переписать под кастомный тип, реализующий io.Writer
		logr.wr = io.MultiWriter(os.Stdout, f)
		close(logr.started)
		wg.Done()
		<-ctx.Done()
	}(ctx)
	wg.Wait()
	log.Println("logr started")
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
