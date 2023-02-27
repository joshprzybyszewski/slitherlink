package solve

import "github.com/joshprzybyszewski/slitherlink/model"

func newNumberRule(
	n int,
	row, col model.Dimension,
) rule {
	r := rule{
		affects: 2,
		row:     row,
		col:     col,
	}
	switch n {
	case 1:
		r.check = r.checkOne
	case 2:
		r.check = r.checkTwo
	case 3:
		r.check = r.checkThree
	default:
		panic(`dev error`)
	}
	return r
}

func (r *rule) checkOne(
	s *state,
) {

}

func (r *rule) checkTwo(
	s *state,
) {

}

func (r *rule) checkThree(
	s *state,
) {

}
