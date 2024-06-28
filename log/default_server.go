package log

import "os"

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
