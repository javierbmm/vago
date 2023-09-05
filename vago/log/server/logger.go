package server

import (
	"fmt"
	"log"
	"os"
)

type ServerLogger struct {
	prefixer prefixer
	logger   *log.Logger
}

const (
	INFO    = "INFO"
	ERROR   = "ERROR"
	WARNING = "WARNING"
)

func (bl *ServerLogger) Init() ServerLogger {
	bl.logger = log.New(os.Stdout, "", 0)
	bl.prefixer = prefixer{}

	return *bl
}

func (bl *ServerLogger) Info(text string, args ...any) {
	bl.setPrefix(INFO)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}

func (bl *ServerLogger) Error(text error, args ...any) {
	bl.setPrefix(ERROR)
	if args == nil {
		bl.logger.Fatal(text)
	} else {
		bl.logger.Fatal(fmt.Sprintf(text.Error(), args...))
	}
}

func (bl *ServerLogger) Warning(text string, args ...any) {
	bl.setPrefix(WARNING)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}

func (bl *ServerLogger) setPrefix(level string) {
	bl.logger.SetPrefix(bl.prefixer.prefix(level))
}

func (bl *ServerLogger) Log(pagename string, requester string) {
	msg := fmt.Sprintf("Sending [%s] to [%s]", pagename, requester)
	bl.setPrefix(INFO)
	bl.logger.Println(msg)
}
