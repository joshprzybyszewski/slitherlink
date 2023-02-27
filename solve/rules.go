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
	size model.Size,
) *rules {
	r := rules{
		unknowns: make([]path, 0, 2*int(size)*int(size-1)),
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
		nr := newNumberRule(
			s.nodes[i].Num,
			s.nodes[i].Row, s.nodes[i].Col,
		)

		r.addHorizontalRule(s.nodes[i].Row, s.nodes[i].Col, &nr)
		r.addHorizontalRule(s.nodes[i].Row+1, s.nodes[i].Col, &nr)
		r.addVerticalRule(s.nodes[i].Row, s.nodes[i].Col, &nr)
		r.addVerticalRule(s.nodes[i].Row, s.nodes[i].Col+1, &nr)
	}

	var pins [maxPinsPerLine][maxPinsPerLine]rule
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
			pins[row][col] = newDefaultRule(row, col)
		}
	}

	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col < model.Dimension(s.size); col++ {
			r.addHorizontalRule(row, col, &pins[row][col])
			r.addHorizontalRule(row, col, &pins[row][col+1])
		}
	}

	for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
		for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
			r.addVerticalRule(row, col, &pins[row][col])
			r.addVerticalRule(row, col, &pins[row+1][col])
		}
	}
}

func (r *rules) populateUnknowns(
	s *state,
) {

	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {

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
	is := int(s.size)

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
		if tmp = is - int(r.unknowns[i].Row); tmp < dri {
			dri = tmp
		}
		dci = int(r.unknowns[i].Col) - 1
		if tmp = is - int(r.unknowns[i].Col); tmp < dci {
			dci = tmp
		}

		drj = int(r.unknowns[j].Row) - 1
		if tmp = is - int(r.unknowns[j].Row); tmp < drj {
			drj = tmp
		}
		dcj = int(r.unknowns[j].Col) - 1
		if tmp = is - int(r.unknowns[j].Col); tmp < dcj {
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
	r.verticals[row][col].affects += rule.affects
	prev := r.verticals[row][col].fn
	r.verticals[row][col].fn = func(s *state) {
		rule.check(s)
		if !s.hasInvalid {
			prev(s)
		}
	}
}
