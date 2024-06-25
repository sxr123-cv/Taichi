package log

import (
	"fmt"
	"time"
)

type Log struct {
	Level       string
	StdoutLevel int
	Msg         string
	Get         func(a string)
}

func (l Log) Out(msg string) {
	fmt.Printf("[%s][%s]:%s\n", time.Now().Format("2006-01-02 15:04:05"), l.Level, msg)
}

type Children struct {
	Log
}

func (l Children) Out(msg string) {
	fmt.Printf("[%s][%s]:%s\n", l.Level, time.Now().Format("2006-01-02 15:04:05"), msg)
}

var L Children

func INFO(msg string) {
	L.Level = "INFO"
	L.Out(msg)
}

func ERROR(msg string) {
	L.Level = "ERROR"
	L.Out(msg)
}

func DEBUG(msg string) {
	L.Level = "DEBUG"
	L.Out(msg)
}

func WARN(msg string) {
	L.Level = "WARN"
	L.Out(msg)
}
