package solve

import "github.com/joshprzybyszewski/slitherlink/model"

func new1UpperLeftRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check1UpperLeft

	return r
}

func (r *rule) check1UpperLeft(
	s *state,
) {
	lh, ah := s.horAt(r.row, r.col-1)
	lv, av := s.verAt(r.row-1, r.col)
	if ah && av {
		s.avoidVer(r.row, r.col)
		s.avoidHor(r.row, r.col)
		return
	}
	if lh && av || lv && ah {
		s.avoidVer(r.row, r.col+1)
		s.avoidHor(r.row+1, r.col)
		return
	}
}

func new1UpperRightRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check1UpperRight

	return r
}

func (r *rule) check1UpperRight(
	s *state,
) {
	lh, ah := s.horAt(r.row, r.col)
	lv, av := s.verAt(r.row-1, r.col)
	if ah && av {
		s.avoidVer(r.row, r.col)
		s.avoidHor(r.row, r.col-1)
		return
	}
	if lh && av || lv && ah {
		s.avoidVer(r.row, r.col-1)
		s.avoidHor(r.row+1, r.col-1)
		return
	}
}

func new1LowerRightRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check1LowerRight

	return r
}

func (r *rule) check1LowerRight(
	s *state,
) {
	lh, ah := s.horAt(r.row, r.col)
	lv, av := s.verAt(r.row, r.col)
	if ah && av {
		s.avoidVer(r.row-1, r.col)
		s.avoidHor(r.row, r.col-1)
		return
	}
	if lh && av || lv && ah {
		s.avoidVer(r.row-1, r.col-1)
		s.avoidHor(r.row-1, r.col-1)
		return
	}
}

func new1LowerLeftRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check1LowerLeft

	return r
}

func (r *rule) check1LowerLeft(
	s *state,
) {
	lh, ah := s.horAt(r.row, r.col-1)
	lv, av := s.verAt(r.row, r.col)
	if ah && av {
		s.avoidVer(r.row-1, r.col)
		s.avoidHor(r.row, r.col)
		return
	}
	if lh && av || lv && ah {
		s.avoidVer(r.row-1, r.col+1)
		s.avoidHor(r.row-1, r.col)
		return
	}
}
