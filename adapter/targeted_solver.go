package adapter

import (
	"time"

	"github.com/joshprzybyszewski/puzzler/model"
	slm "github.com/joshprzybyszewski/slitherlink/model"
)

type targetedSolver struct {
	iter model.Iterator
	id   model.GameID

	timeout time.Duration
}

func NewTargetedSolver(
	iter slm.Iterator,
	id model.GameID,
	timeout time.Duration,
) targetedSolver {
	if iter < slm.MinIterator {
		panic(`unexpected`)
	}
	if iter > slm.MinIterator {
		panic(`unexpected`)
	}
	if timeout < 0 {
		timeout = 0
	} else if timeout > maxTimeout {
		timeout = maxTimeout
	}

	return targetedSolver{
		iter:    model.Iterator(iter),
		id:      id,
		timeout: timeout,
	}
}

func (s targetedSolver) Iterator() model.Iterator {
	return s.iter
}

func (s targetedSolver) IteratorString() model.IteratorString {
	return model.IteratorString(slm.Iterator(s.iter).String())
}

func (s targetedSolver) GameID() model.GameID {
	return s.id
}

func (s targetedSolver) Solve(g *model.Game) error {
	return solveGame(g, s.timeout)
}

func (s targetedSolver) Pretty(g model.Game) string {
	return `TODO`
}
