package generator

import (
	"fmt"
	"time"
)

type prefixer struct {
	NoLog    bool
	NoTime   bool
	Pagename string
}

func (bp *prefixer) prefix(logLevel string) string {
	if bp.NoLog {
		return ""
	}

	if bp.NoTime {
		return fmt.Sprintf("[%s] %s: ", bp.Pagename, logLevel)
	}

	now := time.Now().Format(time.RFC822)
	return fmt.Sprintf("[%s][%s] %s: ", now, bp.Pagename, logLevel)
}
