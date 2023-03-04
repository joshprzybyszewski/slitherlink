//go:build lambda

package solve

import "time"

const (
	numWorkers = 6
)

var (
	maxAttemptDuration = 50 * time.Millisecond
)
