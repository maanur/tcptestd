package main

import (
	"io"
)

var actions map[string]func() Action

// Action реализует интерфейс для действий, которые могут вызывать боты.
// Name() нужно чтобы понимать, что вызвали
// GetParam() вычленяет параметры из строки, вызвавшей срабатываение (это очень криво)
// Run() запускает действие с выводом результата в io.Writer; предполагается, что параметры уже получены
type Action interface {
	Name() string
	GetParam(string) (string, error)
	Run(io.Writer) error
}
