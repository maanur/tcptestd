package main

import (
	"io"
	"log"

	"fmt"
	"strings"

	"github.com/maanur/testtcpmail/tcpmprobe"
)

/*
func init() {
	actions["tcpmonitor"] = func() Action {
		return new(tcpmonitor)
	}
}
*/

type tcpmonitor struct {
	server []struct {
		addr string
		err  error
	}
}

func (tcpm *tcpmonitor) Run(output io.Writer) error {
	iserr := false
	logger := log.New(output, "[tcpmonitor] ", log.Lshortfile|log.LstdFlags)
	for _, server := range tcpm.server {
		server.err = tcpmprobe.MonRun(server.addr, logger)
		if server.err != nil {
			logger.Println(server.addr, server.err)
			iserr = true
		}
	}
	if iserr {
		return fmt.Errorf("Errors found")
	}
	return nil
}

func (*tcpmonitor) Name() string {
	return "tcpmonitor"
}

func (tcpm *tcpmonitor) GetParam(str string) (string, error) {
	var ds []string
	sstr := strings.Split(str, " ")
	for _, s := range sstr {
		if strings.Contains(s, ":") {
			ds = append(ds, s)
		}
	}
	if len(ds) == 0 {
		return "", fmt.Errorf("No correct address found")
	}
	for _, dss := range ds {
		tcpm.server = append(tcpm.server, struct {
			addr string
			err  error
		}{dss, nil})
	}
	return strings.Join(ds, string('\n')), nil
}
