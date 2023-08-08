package log

type Logger interface {
	Info(string, ...any)
	Warning(string, ...any)
	Error(error, ...any)
}
