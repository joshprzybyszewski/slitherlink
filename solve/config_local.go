//go:build !lambda

package solve

import "time"

const (
	// numWorkers = 1
	numWorkers = 4
)

var (
	maxAttemptDuration = 5 * time.Second
)
