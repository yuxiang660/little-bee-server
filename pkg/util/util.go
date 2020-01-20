package util

import (
	"fmt"
	"os"
	"time"
)

var (
	pid = os.Getpid()
)

// NewTraceID generates a trace id which has pid and time.
func NewTraceID() string {
	return fmt.Sprintf(
		"pid-%d [start at %s]",
		pid,
		time.Now().Format(time.Stamp),
	)
}