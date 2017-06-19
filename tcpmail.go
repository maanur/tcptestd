package main

import (
	"io"
	"log"

	"github.com/maanur/testtcpmail/tcpmprobe"
)

func tcpmonitor(addr string, output io.Writer) {
	logger := log.New(output, "[tcpmonitor] ", log.Lshortfile|log.LstdFlags)
	err := tcpmprobe.MonRun(addr, logger)
	if err != nil {
		logger.Println(err)
	}
}
