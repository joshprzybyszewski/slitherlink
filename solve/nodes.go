package solve

import (
	"fmt"
	"time"

	"github.com/joshprzybyszewski/slitherlink/model"
)

func FromNodes(
	width, height model.Size,
	ns []model.Node,
) (model.Solution, error) {
	return FromNodesWithTimeout(
		width, height,
		ns,
		maxAttemptDuration,
	)
}

func FromNodesWithTimeout(
	width, height model.Size,
	ns []model.Node,
	dur time.Duration,
) (model.Solution, error) {

	s := newState(width, height, ns)

	ss := settle(&s)
	if ss == solved {
		return s.toSolution(), nil
	} else if ss == invalid {
		fmt.Printf("%s\n", &s)
		panic(`bad initialization`)
	}

	return solveWithWorkforce(
		&s,
		dur,
	)
}
