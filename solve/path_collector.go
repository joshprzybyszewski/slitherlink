package solve

import (
	"github.com/joshprzybyszewski/slitherlink/model"
)

type pair struct {
	a model.Coord
	b model.Coord

	nodePaths int
}

func newEmptyPair() pair {
	return pair{}
}

func (p *pair) isEmpty() bool {
	return p.a.Col == 0
}

func (p *pair) isHorizontallyClose() bool {
	return p.a.Row == p.b.Row &&
		(p.a.Col == p.b.Col+1 || p.a.Col == p.b.Col-1)
}

func (p *pair) isVerticallyClose() bool {
	return p.a.Col == p.b.Col &&
		(p.a.Row == p.b.Row+1 || p.a.Row == p.b.Row-1)
}

type pathCollector struct {
	pairs [maxPinsPerLine][maxPinsPerLine]pair

	nodes [maxPinsPerLine]model.DimensionBit

	hasCycle bool

	cycleSeen   int
	cycleTarget int
}

func newPathCollector(
	nodes []model.Node,
) pathCollector {
	pc := pathCollector{}

	for _, n := range nodes {
		pc.nodes[n.Row] |= n.Col.Bit()
		pc.cycleTarget += n.Num
	}

	return pc
}

func (pc *pathCollector) getInteresting(
	s *state,
) (model.Coord, bool, bool) {
	var c model.Coord

	maxRow := model.Dimension(s.height)
	maxCol := model.Dimension(s.width)

	for c.Row = model.Dimension(1); c.Row <= maxRow; c.Row++ {
		for c.Col = model.Dimension(1); c.Col <= maxCol; c.Col++ {
			if !pc.pairs[c.Row][c.Col].isEmpty() {
				if !pc.pairs[c.Row+1][c.Col].isEmpty() {
					if !s.hasVerDefined(c.Row, c.Col) {
						return c, false, true
					}
				}
				if !pc.pairs[c.Row][c.Col+1].isEmpty() {
					if !s.hasHorDefined(c.Row, c.Col) {
						return c, true, true
					}
				}
			}
		}
	}

	return c, false, false
}

func (pc *pathCollector) numNodes(
	a model.Coord,
	b model.Coord,
) int {
	n := 0

	c1 := a
	var c2 model.Coord

	if b.Row < c1.Row {
		c1 = b
		c2 = c1
		c2.Col--
	} else if b.Col < c1.Col {
		c1 = b
		c2 = c1
		c2.Row--
	} else if a.Col == b.Col {
		c2 = c1
		c2.Col--
	} else if a.Row == b.Row {
		c2 = c1
		c2.Row--
	}

	if pc.nodes[c1.Row]&(c1.Col.Bit()) != 0 {
		n++
	}
	if pc.nodes[c2.Row]&(c2.Col.Bit()) != 0 {
		n++
	}
	return n
}

func (pc *pathCollector) addHorizontal(
	row, col model.Dimension,
	s *state,
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row,
		Col: col + 1,
	}

	pc.add(
		mya, myb,
		s,
	)
}

func (pc *pathCollector) addVertical(
	row, col model.Dimension,
	s *state,
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row + 1,
		Col: col,
	}
	pc.add(
		mya, myb,
		s,
	)
}

func (pc *pathCollector) add(
	mya, myb model.Coord,
	s *state,
) {
	l := pc.pairs[mya.Row][mya.Col]
	r := pc.pairs[myb.Row][myb.Col]

	if l.isEmpty() && r.isEmpty() {
		p := pair{
			a: mya,
			b: myb,
		}

		p.nodePaths += pc.numNodes(mya, myb)
		// fmt.Printf("empty: %d / %d\n%s\n", p.nodePaths, pc.cycleTarget, s)

		pc.pairs[mya.Row][mya.Col] = p
		pc.pairs[myb.Row][myb.Col] = p
		return
	}

	if !l.isEmpty() && !r.isEmpty() {
		pc.pairs[mya.Row][mya.Col] = newEmptyPair()
		pc.pairs[myb.Row][myb.Col] = newEmptyPair()

		if l == r {
			if pc.hasCycle {
				// a second cycle? this is bad news.
				pc.cycleSeen = -1
				return
			}
			pc.hasCycle = true
			pc.cycleSeen = l.nodePaths
			return
		}

		p := l
		p.nodePaths += r.nodePaths
		p.nodePaths += pc.numNodes(mya, myb)
		// fmt.Printf("both: %d / %d\n%s\n", p.nodePaths, pc.cycleTarget, s)
		if p.a == mya {
			if r.a == myb {
				p.a = r.b
			} else {
				p.a = r.a
			}
		} else {
			if r.a == myb {
				p.b = r.b
			} else {
				p.b = r.a
			}
		}
		pc.pairs[p.a.Row][p.a.Col] = p
		pc.pairs[p.b.Row][p.b.Col] = p
		pc.checkNewPair(
			p,
			s,
		)
		return
	}

	if !l.isEmpty() {
		p := l
		p.nodePaths += pc.numNodes(mya, myb)
		// fmt.Printf("l: %d / %d\n%s\n", p.nodePaths, pc.cycleTarget, s)
		pc.pairs[mya.Row][mya.Col] = newEmptyPair()
		if p.a == mya {
			p.a = myb
		} else {
			p.b = myb
		}

		pc.pairs[p.a.Row][p.a.Col] = p
		pc.pairs[p.b.Row][p.b.Col] = p
		pc.checkNewPair(
			p,
			s,
		)
		return
	}

	p := r
	p.nodePaths += pc.numNodes(mya, myb)
	// fmt.Printf("r: %d / %d\n%s\n", p.nodePaths, pc.cycleTarget, s)
	pc.pairs[myb.Row][myb.Col] = newEmptyPair()
	if p.a == myb {
		p.a = mya
	} else {
		p.b = mya
	}

	pc.pairs[p.a.Row][p.a.Col] = p
	pc.pairs[p.b.Row][p.b.Col] = p

	pc.checkNewPair(
		p,
		s,
	)
}

func (pc *pathCollector) checkNewPair(
	p pair,
	s *state,
) {
	if pc.hasCycle || s.hasInvalid {
		return
	}
	if p.isEmpty() {
		return
	}

	h := p.isHorizontallyClose()
	if !h && !p.isVerticallyClose() {
		return
	}

	r := p.a.Row
	if p.b.Row < r {
		r = p.b.Row
	}
	c := p.a.Col
	if p.b.Col < c {
		c = p.b.Col
	}

	// only need to check the state when we're about to write a line.
	// re-writing an avoid is no problem.
	// fmt.Printf("%d / %d\n", p.nodePaths, pc.cycleTarget)
	if p.nodePaths >= pc.cycleTarget-3 {
		cpy := *s
		if h {
			if !cpy.horAvoidAt(r, c) {
				cpy.lineHor(r, c)
			}
		} else {
			if !cpy.verAvoidAt(r, c) {
				cpy.lineVer(r, c)
			}
		}
		ss := settle(&cpy)
		if ss == solved || ss == validUnsolved {
			settle(s)
			return
		}
	}

	if h {
		s.avoidHor(r, c)
	} else {
		s.avoidVer(r, c)
	}
}
