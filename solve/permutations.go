package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/slitherlink/model"
)

type applyFn func(*state)

type permutationsFactorySubstate struct {
	perms applyFn

	// known is the last known row/col with a value
	known model.Dimension
	// numCrossings is the number of lines currently known to be in this row/col
	numCrossings int
}

const (
	permutationsFactoryNumVals = 8
)

type permutationsFactory struct {
	vals    [permutationsFactoryNumVals]applyFn
	numVals uint16
}

func newPermutationsFactory() permutationsFactory {
	return permutationsFactory{}
}

func (pf *permutationsFactory) save(
	a applyFn,
) {
	pf.vals[pf.numVals] = a
	pf.numVals++
}

func (pf *permutationsFactory) hasRoomForNumEmpty(
	numEmpty int,
) bool {
	if numEmpty == 0 {
		return false
	}
	numPerms := 1 << (numEmpty - 1)
	rem := len(pf.vals) - int(pf.numVals)
	return numPerms <= rem
}

func (pf *permutationsFactory) populate(
	cur *state,
) {
	numEmptyInCol, col := pf.getBestNextStartingCol(cur)
	if col == 0 || !pf.hasRoomForNumEmpty(numEmptyInCol) {
		pf.populateBestRow(cur)
		pf.populateSimple(cur)
		return
	}

	numEmptyInRow, row := pf.getBestNextStartingRow(cur)
	if row == 0 || !pf.hasRoomForNumEmpty(numEmptyInRow) {
		pf.populateBestColumn(cur)
		pf.populateSimple(cur)
		return
	}

	if numEmptyInCol < numEmptyInRow {
		pf.populateBestColumn(cur)
	} else {
		pf.populateBestRow(cur)
	}

	pf.populateSimple(cur)
}

func (pf *permutationsFactory) populateBestColumn(
	cur *state,
) {
	if pf.numVals > 0 {
		return
	}

	numEmpty, col := pf.getBestNextStartingCol(cur)
	if col == 0 || !pf.hasRoomForNumEmpty(numEmpty) {
		return
	}

	pf.buildColumn(
		cur,
		col,
		permutationsFactorySubstate{
			known:        0,
			numCrossings: getNumLinesInCol(cur, col),
			perms: func(s *state) {
				if s.hasInvalid {
					return
				}
				if getNextEmptyRow(s, col, 0) != 0 {
					fmt.Printf("Didn't fill the whole column\nColumn: %d\n%s\n", col, s)
					panic(`didn't fill the whole col?`)
				}
				if getNumLinesInCol(s, col)%2 != 0 {
					fmt.Printf("Wrong Number of Lines\nColumn: %d\n%s\n", col, s)
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)

}

func (pf *permutationsFactory) buildColumn(
	s *state,
	col model.Dimension,
	cur permutationsFactorySubstate,
) {
	row := getNextEmptyRow(s, cur.known+1, col)
	if row == 0 {
		// there wasn't an empty column found.
		if cur.numCrossings%2 == 0 {
			pf.save(cur.perms)
		}
		return
	}

	a := permutationsFactorySubstate{
		known:        row,
		numCrossings: cur.numCrossings,
		perms: func(s *state) {
			s.avoidHor(row, col)
			cur.perms(s)
		},
	}
	l := permutationsFactorySubstate{
		known:        row,
		numCrossings: cur.numCrossings + 1,
		perms: func(s *state) {
			s.lineHor(row, col)
			cur.perms(s)
		},
	}

	if a.numCrossings >= int(s.height)/2 {
		pf.buildColumn(s, col, a)
		pf.buildColumn(s, col, l)
	} else {
		pf.buildColumn(s, col, l)
		pf.buildColumn(s, col, a)
	}
}

func (pf *permutationsFactory) populateBestRow(
	cur *state,
) {
	if pf.numVals > 0 {
		return
	}

	numEmpty, row := pf.getBestNextStartingRow(cur)
	if row == 0 || !pf.hasRoomForNumEmpty(numEmpty) {
		// couldn't find a good starting row?
		return
	}

	pf.buildRow(
		cur,
		row,
		permutationsFactorySubstate{
			known:        0,
			numCrossings: getNumLinesInRow(cur, row),
			perms: func(s *state) {
				if s.hasInvalid {
					return
				}
				if getNextEmptyCol(s, row, 0) != 0 {
					panic(`didn't fill the whole row?`)
				}
				if getNumLinesInRow(s, row)%2 != 0 {
					fmt.Printf("Wrong Number of Lines\nRow: %d\n%s\n", row, s)
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)
}

func (pf *permutationsFactory) buildRow(
	s *state,
	row model.Dimension,
	cur permutationsFactorySubstate,
) {

	col := getNextEmptyCol(s, row, cur.known+1)
	if col == 0 {
		if cur.numCrossings%2 == 0 {
			pf.save(cur.perms)
		}
		return
	}

	a := permutationsFactorySubstate{
		known:        col,
		numCrossings: cur.numCrossings,
		perms: func(s *state) {
			s.avoidVer(row, col)
			cur.perms(s)
		},
	}

	l := permutationsFactorySubstate{
		known:        col,
		numCrossings: cur.numCrossings + 1,
		perms: func(s *state) {
			s.lineVer(row, col)
			cur.perms(s)
		},
	}

	if a.numCrossings >= int(s.width)/2 {
		pf.buildRow(s, row, a)
		pf.buildRow(s, row, l)
	} else {
		pf.buildRow(s, row, l)
		pf.buildRow(s, row, a)
	}
}

func (pf *permutationsFactory) populateSimple(
	s *state,
) {
	pf.populateNextNode(s)

	if pf.numVals > 0 {
		return
	}

	c, isHor, ok := s.getMostInterestingPath()
	if !ok {
		return
	}

	constantDim := c.Row
	travelDim := c.Col
	if isHor {
		constantDim = c.Col
		travelDim = c.Row
	}

	pf.buildSimple(
		s,
		isHor,
		constantDim,
		permutationsFactorySubstate{
			known:        travelDim,
			numCrossings: 0,
			perms: func(s *state) {
			},
		},
	)
}

func (pf *permutationsFactory) populateNextNode(
	s *state,
) {

	if pf.numVals > 0 {
		return
	}

	// solve threes first, then ones, then twos.
	var unsolved [3]model.Node
	for _, n := range s.nodes {
		if !isNodeSolved(s, n) {
			if n.Num == 3 {
				unsolved[0] = n
				break
			}
			if unsolved[n.Num].Row == 0 {
				unsolved[n.Num] = n
			}
		}
	}
	node := unsolved[0]
	if node.Row == 0 {
		node = unsolved[1]
		if node.Row == 0 {
			node = unsolved[2]
			if node.Row == 0 {
				// did not find an unsolved node
				return
			}
		}
	}

	if node.Num == 3 {
		// avoid left
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
		// avoid top
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
		// avoid right
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// avoid down
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
	} else if node.Num == 1 {
		// line left
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// line top
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// line right
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
		// line down
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
	} else if node.Num == 2 {
		// line left+top
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// line left+right
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
		// line left+down
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.lineVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// line top+right
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.avoidHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
		// line top+down
		pf.save(func(s *state) {
			s.lineHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.avoidVer(node.Row, node.Col+1)
		})
		// line right+down
		pf.save(func(s *state) {
			s.avoidHor(node.Row, node.Col)
			s.avoidVer(node.Row, node.Col)
			s.lineHor(node.Row+1, node.Col)
			s.lineVer(node.Row, node.Col+1)
		})
	}
}

func isNodeSolved(
	s *state,
	n model.Node,
) bool {
	return s.hasHorDefined(n.Row, n.Col) &&
		s.hasVerDefined(n.Row, n.Col) &&
		s.hasHorDefined(n.Row+1, n.Col) &&
		s.hasVerDefined(n.Row, n.Col+1)
}

func (pf *permutationsFactory) buildSimple(
	s *state,
	travelCol bool,
	constantDim model.Dimension,
	cur permutationsFactorySubstate,
) {

	var travelDim model.Dimension
	if travelCol {
		travelDim = getNextEmptyRow(s, cur.known+1, constantDim)
	} else {
		travelDim = getNextEmptyCol(s, constantDim, cur.known+1)
	}

	ap := func(s *state) {
		if travelCol {
			s.avoidHor(cur.known, constantDim)
		} else {
			s.avoidVer(constantDim, cur.known)
		}
		cur.perms(s)
	}

	lp := func(s *state) {
		if travelCol {
			s.lineHor(cur.known, constantDim)
		} else {
			s.lineVer(constantDim, cur.known)
		}
		cur.perms(s)
	}

	if travelDim == 0 || cur.numCrossings >= 4 {
		pf.save(ap)
		pf.save(lp)
		return
	}

	a := permutationsFactorySubstate{
		known:        travelDim,
		numCrossings: cur.numCrossings + 1,
		perms:        ap,
	}

	l := permutationsFactorySubstate{
		known:        travelDim,
		numCrossings: cur.numCrossings + 1,
		perms:        lp,
	}

	pf.buildSimple(s, travelCol, constantDim, a)
	pf.buildSimple(s, travelCol, constantDim, l)
}

func (pf *permutationsFactory) getBestNextStartingRow(
	s *state,
) (int, model.Dimension) {
	var rowByNumEmpty [maxPinsPerLine]model.Dimension
	var ne int

	for row := model.Dimension(1); row < model.Dimension(s.height); row++ {
		ne = int(s.width) - int(s.crossings.rows[row]) - int(s.crossings.rowsAvoid[row])
		rowByNumEmpty[ne] = row
	}

	return pf.chooseStart(rowByNumEmpty)
}

func (pf *permutationsFactory) getBestNextStartingCol(
	s *state,
) (int, model.Dimension) {
	var colByNumEmpty [maxPinsPerLine]model.Dimension
	var ne int

	for col := model.Dimension(1); col < model.Dimension(s.width); col++ {
		ne = int(s.height) - int(s.crossings.cols[col]) - int(s.crossings.colsAvoid[col])
		colByNumEmpty[ne] = col
	}

	return pf.chooseStart(colByNumEmpty)
}

func (pf *permutationsFactory) chooseStart(
	byNumEmpty [maxPinsPerLine]model.Dimension,
) (int, model.Dimension) {

	for numEmpty := 2; numEmpty < len(byNumEmpty); numEmpty++ {
		if byNumEmpty[numEmpty] > 0 {
			return numEmpty, byNumEmpty[numEmpty]
		}
	}

	// it's unlikely that all rows are filled...
	return 0, 0
}

func getNextEmptyRow(
	s *state,
	row, col model.Dimension,
) model.Dimension {
	for ; row <= model.Dimension(s.height); row++ {
		if !s.hasHorDefined(row, col) {
			return row
		}
	}
	return 0
}

func getNumLinesInCol(
	s *state,
	col model.Dimension,
) int {
	return int(s.crossings.cols[col])
}

func getNextEmptyCol(
	s *state,
	row, col model.Dimension,
) model.Dimension {
	for ; col <= model.Dimension(s.width); col++ {
		if !s.hasVerDefined(row, col) {
			return col
		}
	}
	return 0
}

func getNumLinesInRow(
	s *state,
	row model.Dimension,
) int {
	return int(s.crossings.rows[row])
}
