package log

type LogType interface {
	GetLog() string
}

type LogClient interface {
	INFO(...any) string
	ERROR(...any) string
}

type LogServer interface {
	Out(log string)
	Print()
}
