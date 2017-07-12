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
	var wg sync.WaitGroup
	var writers []io.Writer
	defer func() {
		select {
		case <-logr.started:
		default:
			close(logr.started)
		}
	}()
	log.Println("logr starting")
	wg.Add(2)
	go func(ctx context.Context) {
		// Вывод в файл (TODO: создавать новый лог каждый день, архивировать логи.)
		f, err := os.Create(time.Now().Format("0102150405") + ".log")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		writers = append(writers, f)
		wg.Done()
		<-ctx.Done()
	}(ctx)
	go func(ctx context.Context) {
		// Вывод в буфер для отображения в вебе

		//logr.wr = io.MultiWriter(logr.wr, lbuf)
		writers = append(writers, &logr.buf)
		wg.Done()
		<-ctx.Done()
	}(ctx)
	go func(ctx context.Context) {
		logr.wr = io.MultiWriter(append(writers, os.Stdout)...)
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
