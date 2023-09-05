package server

import (
	"fmt"
	"time"
)

type prefixer struct {
}

func (bp *prefixer) prefix(logLevel string) string {
	now := time.Now().Format(time.RFC822)
	return fmt.Sprintf("[%s] %s: ", now, logLevel)
}
