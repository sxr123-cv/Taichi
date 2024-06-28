package log

type LogType interface {
	GetLog() string
}

type LogClient interface {
	INFO(...any) string
	WARN(...any) string
	DEBUG(...any) string
	ERROR(...any) string
}

type LogServer interface {
	Out(log string)
	FilePath(path string)
	Close()
}
