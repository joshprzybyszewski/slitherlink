package adapter

import (
	"os"
	"time"
)

func init() {
	dep := os.Getenv(`DEPLOYED`)
	if dep != `` {
		maxTimeout = 50 * time.Millisecond
	}
}
