//go:build lambda

package solve

const (
	numWorkers = 6
)

var (
	maxAttemptDuration = 50 * time.Millisecond
)
