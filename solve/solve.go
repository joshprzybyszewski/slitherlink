package solve

import (
	"context"
	"time"

	"github.com/joshprzybyszewski/slitherlink/model"
)

func solveWithWorkforce(
	s *state,
	dur time.Duration,
) (model.Solution, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), dur)
	defer cancelFn()

	w := newWorkforce()
	w.start(ctx)
	defer w.stop()
	defer cancelFn()

	return w.solve(ctx, s)
}
