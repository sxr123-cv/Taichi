package log

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

type Log struct {
	LogClient
	Server LogServer
}

func (receiver *Log) INFO(M ...any) {
	log := receiver.LogClient.INFO(M...)
	receiver.Server.Out(log)

}
func (receiver *Log) ERROR(M ...any) {
	log := receiver.LogClient.ERROR(M...)
	receiver.Server.Out(log)
}

func (receiver *Log) WARN(M ...any) {
	log := receiver.LogClient.WARN(M...)
	receiver.Server.Out(log)

}
func (receiver *Log) DEBUG(M ...any) {
	log := receiver.LogClient.DEBUG(M...)
	receiver.Server.Out(log)
}

func (receiver *Log) SetFilePath(filePath string) *Log {
	receiver.Server.FilePath(filePath)
	return receiver
}

func NewLog(client LogClient, sever LogServer) *Log {
	if client == nil {
		client = &DefaultClient{}
	}
	if sever == nil {
		sever = &DefaultServer{}
	}
	return &Log{LogClient: client, Server: sever}
}
