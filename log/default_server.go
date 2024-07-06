package log

import (
	"fmt"
	"os"
)

type DefaultServer struct {
	file *os.File
	path string
	p    chan string
}

func (d *DefaultServer) Start() {
	go func() {
		defer func() {
			println("已退出")
		}()
		for {
			v, ok := <-d.p
			if ok {
				fmt.Printf(v)
				_, err := d.file.WriteString(v)
				if err != nil {
					println("log write error", err.Error())
				}
			} else {

				return
			}
		}
	}()

}
func (d *DefaultServer) Out(msg string) {
	d.p <- msg

}

func (d *DefaultServer) FilePath(path string) {
	d.path = path
}

func (d *DefaultServer) Close() {
	close(d.p)
}
