package log

import (
	"fmt"
	"os"
	"time"
)

type Log struct {
	LogClient
	Server LogServer
}

func NewLog(client LogClient, sever LogServer) *Log {
	if client == nil {
		client = &DefaultClient{}
	}
	if sever == nil {
		create, err := os.Create("./taichi.log")
		if err != nil {
			return nil
		}
		sever = &DefaultServer{file: create}
	}
	return &Log{LogClient: client, Server: sever}
}

type DefaultClient struct {
}

func (receiver *DefaultClient) INFO(M ...any) string {
	var s = fmt.Sprintf("[INFO] %s", time.Now().Format("2006-01-02 15:04:05"))
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
func (receiver *DefaultClient) ERROR(M ...any) string {
	var s = fmt.Sprintf("[ERROR] %s", time.Now().Format("2006-01-02 15:04:05"))
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

type DefaultServer struct {
	file *os.File
}

func (d *DefaultServer) Out(string2 string) {
	println(string2)
	_, err := d.file.WriteString(string2)
	if err != nil {
		println(err.Error())
	}

}

func (receiver *Log) INFO(M ...any) {
	log := receiver.LogClient.INFO(M...)
	receiver.Server.Out(log)

}
func (receiver *Log) ERROR(M ...any) {
	log := receiver.LogClient.ERROR(M...)
	receiver.Server.Out(log)
}
