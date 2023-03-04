//go:build lambda

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joshprzybyszewski/slitherlink/fetch"
	"github.com/joshprzybyszewski/slitherlink/model"
	"github.com/joshprzybyszewski/slitherlink/solve"
)

func main() {
	iterStr := os.Getenv(`ITER`)
	if iterStr != `` {
		val, err := strconv.Atoi(iterStr)
		if err != nil {
			panic(err)
		}
		err = compete(model.Iterator(val))
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
		fmt.Printf("Success!\n")
		os.Exit(0)
	}

	var err error

	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		err = compete(iter)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
	}

	fmt.Printf("Success!\n")
	os.Exit(0)
}

func compete(iter model.Iterator) error {

	fmt.Printf("Starting %s\n\t%s\n\n\n", iter, time.Now())
	input, err := fetch.GetNewPuzzle(iter)
	if err != nil {
		return err
	}

	ns := input.ToNodes()

	t0 := time.Now()
	sol, err := solve.FromNodes(
		iter.GetWidth(),
		iter.GetHeight(),
		ns,
	)
	defer func(dur time.Duration) {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Solution:\n%s\n", sol.Pretty(ns))
		fmt.Printf("Duration: %s\n\n\n", dur)
	}(time.Since(t0))

	if err != nil {
		return err
	}

	return fetch.Submit(
		&input,
		&sol,
	)
}
