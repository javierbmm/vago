package log

import (
	"fmt"
	"time"
)

type generatorPrefixer struct {
	NoLog    bool
	NoTime   bool
	Pagename string
}

func (bp *generatorPrefixer) prefix(logLevel string) string {
	if bp.NoLog {
		return ""
	}

	if bp.NoTime {
		return fmt.Sprintf("[%s] %s: ", bp.Pagename, logLevel)
	}

	now := time.Now().Format(time.RFC822)
	return fmt.Sprintf("[%s][%s] %s: ", now, bp.Pagename, logLevel)
}
