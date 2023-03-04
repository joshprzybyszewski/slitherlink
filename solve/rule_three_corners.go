package solve

import "github.com/joshprzybyszewski/slitherlink/model"

func new3UpperLeftRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check3UpperLeft

	return r
}

func (r *rule) check3UpperLeft(
	s *state,
) {
	l, ah := s.horAt(r.row, r.col-1)
	if l {
		s.avoidVer(r.row-1, r.col)

		s.lineVer(r.row, r.col+1)
		s.lineHor(r.row+1, r.col)

		return
	}
	l, av := s.verAt(r.row-1, r.col)
	if l {
		s.avoidHor(r.row, r.col-1)

		s.lineVer(r.row, r.col+1)
		s.lineHor(r.row+1, r.col)

		return
	}
	if !ah || !av {
		return
	}
	s.lineVer(r.row, r.col)
	s.lineHor(r.row, r.col)
}

func new3UpperRightRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check3UpperRight

	return r
}

func (r *rule) check3UpperRight(
	s *state,
) {
	l, ah := s.horAt(r.row, r.col)
	if l {
		s.avoidVer(r.row-1, r.col)

		s.lineVer(r.row, r.col-1)
		s.lineHor(r.row+1, r.col-1)

		return
	}
	l, av := s.verAt(r.row-1, r.col)
	if l {
		s.avoidHor(r.row, r.col)

		s.lineVer(r.row, r.col-1)
		s.lineHor(r.row+1, r.col-1)

		return
	}
	if !ah || !av {
		return
	}
	s.lineVer(r.row, r.col)
	s.lineHor(r.row, r.col-1)
}

func new3LowerRightRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check3LowerRight

	return r
}

func (r *rule) check3LowerRight(
	s *state,
) {
	l, ah := s.horAt(r.row, r.col)
	if l {
		s.avoidVer(r.row, r.col)

		s.lineVer(r.row-1, r.col-1)
		s.lineHor(r.row-1, r.col-1)

		return
	}
	l, av := s.verAt(r.row, r.col)
	if l {
		s.avoidHor(r.row, r.col)

		s.lineVer(r.row-1, r.col-1)
		s.lineHor(r.row-1, r.col-1)

		return
	}
	if !ah || !av {
		return
	}
	s.lineVer(r.row-1, r.col)
	s.lineHor(r.row, r.col-1)
}

func new3LowerLeftRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row:     row,
		col:     col,
		affects: 3,
	}
	r.check = r.check3LowerLeft

	return r
}

func (r *rule) check3LowerLeft(
	s *state,
) {
	l, ah := s.horAt(r.row, r.col-1)
	if l {
		s.avoidVer(r.row, r.col)

		s.lineVer(r.row-1, r.col+1)
		s.lineHor(r.row-1, r.col)

		return
	}
	l, av := s.verAt(r.row, r.col)
	if l {
		s.avoidHor(r.row, r.col-1)

		s.lineVer(r.row-1, r.col+1)
		s.lineHor(r.row-1, r.col)

		return
	}
	if !ah || !av {
		return
	}
	s.lineVer(r.row-1, r.col)
	s.lineHor(r.row, r.col)
}
