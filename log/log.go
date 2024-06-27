package log

import (
	"fmt"
	"os"
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
		sever = &DefaultServer{file: create, T: make(chan string)}
		sever.Print()
	}
	return &Log{LogClient: client, Server: sever}
}

type DefaultClient struct {
}

func (receiver *DefaultClient) INFO(M ...any) string {
	s := fmt.Sprintf("%s")
	for _, m := range M {
		v, ok := m.(LogType)
		if ok {
			s += fmt.Sprintf("[INFO]%s\n", v.GetLog())
		} else {
			s += fmt.Sprintf("[INFO]%+v\n", m)
		}
	}
	return s
}
func (receiver *DefaultClient) ERROR(M ...any) string {
	var s string
	for _, m := range M {
		v, ok := m.(LogType)
		if ok {
			s += fmt.Sprintf("[ERROR]%s\n", v.GetLog())
		} else {
			s += fmt.Sprintf("[ERROR]%+v\n", m)
		}
	}
	return s
}

type DefaultServer struct {
	file *os.File
	T    chan string
}

func (d *DefaultServer) Out(string2 string) {
	d.T <- string2
}
func (d *DefaultServer) Print() {
	go func() {
		defer close(d.T)
		for {
			v, ok := <-d.T
			if ok {
				fmt.Println(v)
				_, err := d.file.WriteString(v)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}()
}

func (receiver *Log) INFO(M ...any) {
	log := receiver.LogClient.INFO(M...)
	receiver.Server.Out(log)

}
func (receiver *Log) ERROR(M ...any) {
	log := receiver.LogClient.ERROR(M...)
	receiver.Server.Out(log)
}
