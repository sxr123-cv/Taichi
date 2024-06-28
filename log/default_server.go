package log

import (
	"fmt"
	"os"
)

type DefaultServer struct {
	file *os.File
	path string
}

func (d *DefaultServer) Out(msg string) {
	fmt.Printf(msg)
	_, err := d.file.WriteString(msg)
	if err != nil {
		println("log write error")
	}
}
func (d *DefaultServer) FilePath(path string) {
	d.path = path
}
