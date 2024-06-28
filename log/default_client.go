package log

import (
	"fmt"
	"time"
)

type DefaultClient struct {
}

func (d *DefaultClient) INFO(M ...any) string {
	return d.Print(INFO, M...)
}
func (d *DefaultClient) ERROR(M ...any) string {
	return d.Print(ERROR, M...)
}
func (d *DefaultClient) DEBUG(M ...any) string {
	return d.Print(DEBUG, M...)
}
func (d *DefaultClient) WARN(M ...any) string {
	return d.Print(WARNING, M...)
}
func (d *DefaultClient) Print(level int, M ...any) string {
	var levelStr = "INFO"
	switch level {
	case INFO:
		levelStr = "INFO"
	case DEBUG:
		levelStr = "DEBUG"
	case WARNING:
		levelStr = "WARNING"
	case ERROR:
		levelStr = "ERROR"
	}

	var s = fmt.Sprintf("%s [%s] ", levelStr, time.Now().Format("2006-01-02 15:04:05"))
	for _, m := range M {
		v, ok := m.(LogType)
		if ok {
			s += fmt.Sprintf("%s\n", v.GetLog())
		} else {
			s += fmt.Sprintf("%+v\n", m)
		}
	}
	return s
}
