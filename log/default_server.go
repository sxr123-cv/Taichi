package log

import "os"

type DefaultServer struct {
	file *os.File
	path string
}

func (d *DefaultServer) Out(string2 string) {
	_, err := d.file.WriteString(string2)
	if err != nil {
		println(err.Error())
	}
}
func (d *DefaultServer) FilePath(path string) {
	d.path = path
}
