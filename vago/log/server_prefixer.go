package log

import (
	"fmt"
	"time"
)

type serverPrefixer struct {
}

func (bp *serverPrefixer) prefix(logLevel string) string {
	now := time.Now().Format(time.RFC822)
	return fmt.Sprintf("[%s] %s: ", now, logLevel)
}
