package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

var logbufbyte []byte
var logbuf = func(b []byte) *bufio.ReadWriter {
	buf := bytes.NewBuffer(b)
	return bufio.NewReadWriter(bufio.NewReader(buf), bufio.NewWriterSize(buf, 3000))
}(logbufbyte)

//TODO: сохранять логи в файлики по датам и добавлять в архив (tar?)
//TODO: держать буфер с последним логом для вывода на WEB-страничку
func logger() io.Writer {
	return os.Stdout
}
