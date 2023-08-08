package loggers

import (
	"fmt"
	"log"
	"os"
	"time"
)

type BuilderLogger struct {
	pagename string
	logger   *log.Logger
}

const (
	INFO    = "INFO"
	ERROR   = "ERROR"
	WARNING = "WARNING"
)

func (bl *BuilderLogger) Init(pagename string) BuilderLogger {
	bl.pagename = pagename
	bl.logger = log.New(os.Stdout, "", 0)

	return *bl
}

func (bl *BuilderLogger) prefix(prefixLevel string) *BuilderLogger {
	now := time.Now().Format(time.RFC822)
	prefix := fmt.Sprintf("%s [%s][%s]: ", prefixLevel, now, bl.pagename)
	bl.logger.SetPrefix(prefix)

	return bl
}

func (bl *BuilderLogger) Info(text string, args ...any) {
	bl.prefix(INFO)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}

func (bl *BuilderLogger) Error(text error, args ...any) {
	bl.prefix(ERROR)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text.Error(), args...))
	}
}

func (bl *BuilderLogger) Warning(text string, args ...any) {
	bl.prefix(WARNING)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}
