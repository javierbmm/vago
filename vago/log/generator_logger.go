package log

import (
	"fmt"
	"log"
	"os"
)

type GeneratorLogger struct {
	pagename string
	logger   *log.Logger
	prefixer generatorPrefixer
}

func (bl *GeneratorLogger) Init(pagename string, noLog bool, noTime bool) GeneratorLogger {
	bl.pagename = pagename
	bl.logger = log.New(os.Stdout, "", 0)
	bl.prefixer = generatorPrefixer{
		NoLog:    noLog,
		NoTime:   noTime,
		Pagename: pagename,
	}

	return *bl
}

func (bl *GeneratorLogger) Info(text string, args ...any) {
	bl.setPrefix(INFO)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}

func (bl *GeneratorLogger) Error(text error, args ...any) {
	bl.setPrefix(ERROR)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text.Error(), args...))
	}
}

func (bl *GeneratorLogger) Warning(text string, args ...any) {
	bl.setPrefix(WARNING)
	if args == nil {
		bl.logger.Println(text)
	} else {
		bl.logger.Println(fmt.Sprintf(text, args...))
	}
}

func (bl *GeneratorLogger) setPrefix(level string) {
	bl.logger.SetPrefix(bl.prefixer.prefix(level))
}
