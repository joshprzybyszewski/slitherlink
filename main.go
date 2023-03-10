package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/joshprzybyszewski/puzzler/compete"
	pmodel "github.com/joshprzybyszewski/puzzler/model"
	"github.com/joshprzybyszewski/puzzler/results"
	"github.com/joshprzybyszewski/puzzler/trial"
	"github.com/joshprzybyszewski/slitherlink/adapter"
	"github.com/joshprzybyszewski/slitherlink/model"
	"github.com/joshprzybyszewski/slitherlink/profile"
)

var (
	updateResults = flag.Bool("results", false, "if set, then it will print the custom benchmark results")

	puzzID = flag.String("puzzID", "", "if set, then this will run a specific puzzle")

	iterStart     = flag.Int("start", int(model.MinIterator), "if set, this will override the iterators starting value")
	iterFinish    = flag.Int("finish", int(model.MaxIterator), "if set, this will override the iterators final value")
	numIterations = flag.Int("numIterations", 1, "set this value to run through the puzzles many times")

	fetchNewPuzzles = flag.Bool("refresh", true, "if set, then it will fetch new puzzles")

	shouldProfile = flag.Bool("profile", false, "if set, will produce a profile output")
)

func main() {
	flag.Parse()

	if *updateResults {
		results.Generate(
			adapter.NewSolver(
				model.MinIterator,
				model.MaxIterator,
				10*time.Second,
			),
		)
		return
	}

	if *shouldProfile {
		defer profile.Start()()
	}

	if *puzzID != `` {
		_ = trial.Run(
			adapter.NewTargetedSolver(
				model.Iterator(*iterStart),
				pmodel.GameID(*puzzID),
				15*time.Second,
			),
		)
		return
	}

	for i := 0; i < *numIterations; i++ {
		err := compete.Run(
			adapter.NewSolver(
				model.Iterator(*iterStart),
				model.Iterator(*iterFinish),
				10*time.Second,
			),
		)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
	}
}
