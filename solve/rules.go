package solve

import (
	"sort"

	"github.com/joshprzybyszewski/slitherlink/model"
)

var (
	emptyApply = func(*state) {}
)

type path struct {
	model.Coord
	IsHorizontal bool
}

type affectsApply struct {
	affects int

	fn applyFn
}

type rules struct {
	// if "this" row/col changes, then run these other checks
	// [row][col]
	horizontals [maxPinsPerLine][maxPinsPerLine]affectsApply
	verticals   [maxPinsPerLine][maxPinsPerLine]affectsApply

	// unknowns describes the paths that aren't initialized known.
	// They should exist in a sorted manner, where the first one has the most
	// rules associated with it, and so on. This can be used to find "the most
	// interesting space to investigate next."
	unknowns []path
}

func newRules(
	width, height model.Size,
) *rules {
	r := rules{
		unknowns: make([]path, 0, 2*int(width)*int(height-1)),
	}

	for i := range r.horizontals {
		for j := range r.horizontals[i] {
			r.horizontals[i][j].fn = emptyApply
			r.verticals[i][j].fn = emptyApply
		}
	}

	return &r
}

func (r *rules) populateRules(
	s *state,
) {

	for i := range s.nodes {
		if s.nodes[i].Num == 0 {
			// already done in gimmes
			continue
		}
		nr := newNumberRule(
			s.nodes[i].Num,
			s.nodes[i].Row, s.nodes[i].Col,
		)

		r.addHorizontalRule(s.nodes[i].Row, s.nodes[i].Col, &nr)
		r.addHorizontalRule(s.nodes[i].Row+1, s.nodes[i].Col, &nr)
		r.addVerticalRule(s.nodes[i].Row, s.nodes[i].Col, &nr)
		r.addVerticalRule(s.nodes[i].Row, s.nodes[i].Col+1, &nr)

		switch s.nodes[i].Num {
		case 1:
			r.addOneRules(s.nodes[i])
		case 3:
			r.addThreeRules(s.nodes[i])
		}
	}

	var pins [maxPinsPerLine + 2][maxPinsPerLine + 2]rule
	for row := model.Dimension(1); row <= model.Dimension(s.height)+1; row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.width)+1; col++ {
			pins[row][col] = newDefaultRule(row, col)
		}
	}

	for row := model.Dimension(1); row <= model.Dimension(s.height); row++ {
		for col := model.Dimension(1); col < model.Dimension(s.width); col++ {
			r.addHorizontalRule(row, col, &pins[row][col])
			r.addHorizontalRule(row, col, &pins[row][col+1])
		}
	}

	for col := model.Dimension(1); col <= model.Dimension(s.width); col++ {
		for row := model.Dimension(1); row < model.Dimension(s.height); row++ {
			r.addVerticalRule(row, col, &pins[row][col])
			r.addVerticalRule(row, col, &pins[row+1][col])
		}
	}

	// TODO
	// for a corner of a (3) node, if the other two are avoids, then that corner lines.
}

func (r *rules) populateUnknowns(
	s *state,
) {

	for row := model.Dimension(1); row <= model.Dimension(s.height); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.width); col++ {

			if !s.hasHorDefined(row, col) {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: true,
				})
			}

			if !s.hasVerDefined(row, col) {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: false,
				})
			}
		}
	}

	var ni, nj int
	var dri, dci, drj, dcj, tmp int
	iw := int(s.width)
	ih := int(s.height)

	sort.Slice(r.unknowns, func(i, j int) bool {
		if r.unknowns[i].IsHorizontal {
			ni = r.horizontals[r.unknowns[i].Row][r.unknowns[i].Col].affects
		} else {
			ni = r.verticals[r.unknowns[i].Row][r.unknowns[i].Col].affects
		}

		if r.unknowns[j].IsHorizontal {
			nj = r.horizontals[r.unknowns[j].Row][r.unknowns[j].Col].affects
		} else {
			nj = r.verticals[r.unknowns[j].Row][r.unknowns[j].Col].affects
		}

		if ni != nj {
			return ni < nj
		}

		// There are the same number of rules for this segment.
		// Prioritize a segment that is closer to the outer wall
		dri = int(r.unknowns[i].Row) - 1
		if tmp = iw - int(r.unknowns[i].Row); tmp < dri {
			dri = tmp
		}
		dci = int(r.unknowns[i].Col) - 1
		if tmp = ih - int(r.unknowns[i].Col); tmp < dci {
			dci = tmp
		}

		drj = int(r.unknowns[j].Row) - 1
		if tmp = iw - int(r.unknowns[j].Row); tmp < drj {
			drj = tmp
		}
		dcj = int(r.unknowns[j].Col) - 1
		if tmp = ih - int(r.unknowns[j].Col); tmp < dcj {
			dcj = tmp
		}
		ni = dri + dci
		nj = drj + dcj
		if ni != nj {
			return ni < nj
		}

		// They are equally close to the outer wall.
		// Prioritize the one in the top left.
		if r.unknowns[i].Row != r.unknowns[j].Row {
			return r.unknowns[i].Row < r.unknowns[j].Row
		}
		if r.unknowns[i].Col != r.unknowns[j].Col {
			return r.unknowns[i].Col < r.unknowns[j].Col
		}

		// Check horizontal first.
		return r.unknowns[i].IsHorizontal
	})
}

func (r *rules) addHorizontalRule(
	row, col model.Dimension,
	rule *rule,
) {
	if rule == nil || rule.check == nil {
		panic(`ahh`)
	}

	r.horizontals[row][col].affects += rule.affects
	prev := r.horizontals[row][col].fn
	r.horizontals[row][col].fn = func(s *state) {
		rule.check(s)
		if !s.hasInvalid {
			prev(s)
		}
	}
}

func (r *rules) addVerticalRule(
	row, col model.Dimension,
	rule *rule,
) {
	if rule == nil || rule.check == nil {
		panic(`ahh`)
	}

	r.verticals[row][col].affects += rule.affects
	prev := r.verticals[row][col].fn
	r.verticals[row][col].fn = func(s *state) {
		rule.check(s)
		if !s.hasInvalid {
			prev(s)
		}
	}
}

func (r *rules) addOneRules(
	n model.Node,
) {
	if n.Num != 1 {
		return
	}

	ul := new1UpperLeftRule(n.Row, n.Col)
	r.addHorizontalRule(n.Row, n.Col-1, &ul)
	r.addHorizontalRule(n.Row, n.Col, &ul)
	r.addVerticalRule(n.Row-1, n.Col, &ul)
	r.addVerticalRule(n.Row, n.Col, &ul)

	ur := new1UpperRightRule(n.Row, n.Col+1)
	r.addHorizontalRule(n.Row, n.Col, &ur)
	r.addHorizontalRule(n.Row, n.Col+1, &ur)
	r.addVerticalRule(n.Row-1, n.Col+1, &ur)
	r.addVerticalRule(n.Row, n.Col+1, &ur)

	lr := new1LowerRightRule(n.Row+1, n.Col+1)
	r.addHorizontalRule(n.Row+1, n.Col, &lr)
	r.addHorizontalRule(n.Row+1, n.Col+1, &lr)
	r.addVerticalRule(n.Row, n.Col+1, &lr)
	r.addVerticalRule(n.Row+1, n.Col+1, &lr)

	ll := new1LowerLeftRule(n.Row+1, n.Col)
	r.addHorizontalRule(n.Row+1, n.Col-1, &ll)
	r.addHorizontalRule(n.Row+1, n.Col, &ll)
	r.addVerticalRule(n.Row, n.Col, &ll)
	r.addVerticalRule(n.Row+1, n.Col+1, &ll)
}

func (r *rules) addThreeRules(
	n model.Node,
) {
	if n.Num != 3 {
		return
	}

	ul := new3UpperLeftRule(n.Row, n.Col)
	r.addHorizontalRule(n.Row, n.Col-1, &ul)
	r.addHorizontalRule(n.Row, n.Col, &ul)
	r.addVerticalRule(n.Row-1, n.Col, &ul)
	r.addVerticalRule(n.Row, n.Col, &ul)

	ur := new3UpperRightRule(n.Row, n.Col+1)
	r.addHorizontalRule(n.Row, n.Col, &ur)
	r.addHorizontalRule(n.Row, n.Col+1, &ur)
	r.addVerticalRule(n.Row-1, n.Col+1, &ur)
	r.addVerticalRule(n.Row, n.Col+1, &ur)

	lr := new3LowerRightRule(n.Row+1, n.Col+1)
	r.addHorizontalRule(n.Row+1, n.Col, &lr)
	r.addHorizontalRule(n.Row+1, n.Col+1, &lr)
	r.addVerticalRule(n.Row, n.Col+1, &lr)
	r.addVerticalRule(n.Row+1, n.Col+1, &lr)

	ll := new3LowerLeftRule(n.Row+1, n.Col)
	r.addHorizontalRule(n.Row+1, n.Col-1, &ll)
	r.addHorizontalRule(n.Row+1, n.Col, &ll)
	r.addVerticalRule(n.Row, n.Col, &ll)
	r.addVerticalRule(n.Row+1, n.Col+1, &ll)
}
