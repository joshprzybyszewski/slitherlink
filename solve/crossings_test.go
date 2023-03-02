package solve

import (
	"testing"

	"github.com/joshprzybyszewski/slitherlink/model"
	"github.com/stretchr/testify/assert"
)

func TestCrossings(t *testing.T) {
	// This is puzz 7,272,378
	s := newState(
		model.MinIterator.GetWidth(),
		model.MinIterator.GetHeight(),
		puzz8x8_3368748_Nodes,
	)

	expString := `0 X 1 X 2 X 3 X 4 X 5 X 6 X 7 X 
X X X X X X X X X X X X X X X X 
1 X * - * X * X * - * - * X     
X X | 3 |   X   |   X   | X X X 
2 X * X * - * X * X * - * X     
X X |   X 2 | 3 |   | 3 X X X X 
3 X * - * X * - * X * - * X     
X X X 2 |   X 1 X 0 X 3 | X X X 
4 X * X * - * X * X * - * X     
X X X 0 X   | 2 X   | 3 X X X X 
5 X * X * X * - * X * - * X     
X X X   X   X   |   X   | X X X 
6 X * X * X * X * - * - * X     
X X X X X X X X X X X X X X X X 
7 X   X   X   X   X   X   X   X 
X X   X   X   X   X   X   X X X 
`

	assert.Equal(t, expString, s.String())
}

var (
	puzz8x8_3368748_Nodes = []model.Node{{
		Coord: model.Coord{
			Row: 0,
			Col: 0,
		},
		Num: 3,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 1,
		},
		Num: 2,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 2,
		},
		Num: 3,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 4,
		},
		Num: 3,
	}, {
		Coord: model.Coord{
			Row: 2,
			Col: 0,
		},
		Num: 2,
	}, {
		Coord: model.Coord{
			Row: 2,
			Col: 2,
		},
		Num: 1,
	}, {
		Coord: model.Coord{
			Row: 2,
			Col: 3,
		},
		Num: 0,
	}, {
		Coord: model.Coord{
			Row: 2,
			Col: 4,
		},
		Num: 3,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 0,
		},
		Num: 0,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 2,
		},
		Num: 2,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 4,
		},
		Num: 3,
	}}
)
