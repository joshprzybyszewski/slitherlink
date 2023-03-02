package solve

import "github.com/joshprzybyszewski/slitherlink/model"

func newNumberRule(
	n int,
	row, col model.Dimension,
) rule {
	r := rule{
		row: row,
		col: col,
	}
	switch n {
	case 1:
		r.affects = 3
		r.check = r.checkOne
	case 2:
		r.affects = 2
		r.check = r.checkTwo
	case 3:
		r.affects = 3
		r.check = r.checkThree
	default:
		panic(`dev error`)
	}
	return r
}

func (r *rule) checkOne(
	s *state,
) {
	var l, ah, av, ah1, av1 bool
	l, ah = s.horAt(r.row, r.col)
	if l {
		s.avoidVer(r.row, r.col)
		s.avoidHor(r.row+1, r.col)
		s.avoidVer(r.row, r.col+1)
		return
	}
	l, av = s.verAt(r.row, r.col)
	if l {
		s.avoidHor(r.row, r.col)
		s.avoidHor(r.row+1, r.col)
		s.avoidVer(r.row, r.col+1)
		return
	}
	l, ah1 = s.horAt(r.row+1, r.col)
	if l {
		s.avoidVer(r.row, r.col)
		s.avoidHor(r.row, r.col)
		s.avoidVer(r.row, r.col+1)
		return
	}
	l, av1 = s.verAt(r.row, r.col+1)
	if l {
		s.avoidHor(r.row, r.col)
		s.avoidHor(r.row+1, r.col)
		s.avoidVer(r.row, r.col)
		return
	}

	if ah && av && ah1 && av1 {
		r.setInvalid(s)
		return
	}

	if !ah {
		if av && ah1 && av1 {
			s.lineHor(r.row, r.col)
		}
		return
	}
	if !av {
		if ah1 && av1 {
			s.lineVer(r.row, r.col)
		}
		return
	}
	if !ah1 {
		if av1 {
			s.lineHor(r.row+1, r.col)
		}
		return
	}
	s.lineVer(r.row, r.col+1)
}

func (r *rule) checkTwo(
	s *state,
) {
	nl, na := uint8(0), uint8(0)
	l, a := s.horAt(r.row, r.col)
	if l {
		nl++
	} else if a {
		na++
	}
	l, a = s.verAt(r.row, r.col)
	if l {
		nl++
	} else if a {
		na++
	}
	l, a = s.horAt(r.row+1, r.col)
	if l {
		nl++
	} else if a {
		na++
	}
	l, a = s.verAt(r.row, r.col+1)
	if l {
		nl++
	} else if a {
		na++
	}

	if nl+na == 4 {
		// all set
		return
	}
	if nl != 2 && na != 2 {
		// we don't know enough to complete the square
		return
	}
	if !s.hasHorDefined(r.row, r.col) {
		if nl == 2 {
			s.avoidHor(r.row, r.col)
		} else {
			s.lineHor(r.row, r.col)
		}
	}
	if !s.hasVerDefined(r.row, r.col) {
		if nl == 2 {
			s.avoidVer(r.row, r.col)
		} else {
			s.lineVer(r.row, r.col)
		}
	}
	if !s.hasHorDefined(r.row+1, r.col) {
		if nl == 2 {
			s.avoidHor(r.row+1, r.col)
		} else {
			s.lineHor(r.row+1, r.col)
		}
	}
	if !s.hasVerDefined(r.row, r.col+1) {
		if nl == 2 {
			s.avoidVer(r.row, r.col+1)
		} else {
			s.lineVer(r.row, r.col+1)
		}
	}
}

func (r *rule) checkThree(
	s *state,
) {
	var a, lh, lv, lh1, lv1 bool
	lh, a = s.horAt(r.row, r.col)
	if a {
		s.lineVer(r.row, r.col)
		s.lineHor(r.row+1, r.col)
		s.lineVer(r.row, r.col+1)
		return
	}
	lv, a = s.verAt(r.row, r.col)
	if a {
		s.lineHor(r.row, r.col)
		s.lineHor(r.row+1, r.col)
		s.lineVer(r.row, r.col+1)
		return
	}
	lh1, a = s.horAt(r.row+1, r.col)
	if a {
		s.lineVer(r.row, r.col)
		s.lineHor(r.row, r.col)
		s.lineVer(r.row, r.col+1)
		return
	}
	lv1, a = s.verAt(r.row, r.col+1)
	if a {
		s.lineVer(r.row, r.col)
		s.lineHor(r.row, r.col)
		s.lineHor(r.row+1, r.col)
		return
	}

	if lh && lv && lh1 && lv1 {
		r.setInvalid(s)
		return
	}

	if !lh {
		if lv && lh1 && lv1 {
			s.avoidHor(r.row, r.col)
		}
		return
	}
	if !lv {
		if lh1 && lv1 {
			s.avoidVer(r.row, r.col)
		}
		return
	}
	if !lh1 {
		if lv1 {
			s.avoidHor(r.row+1, r.col)
		}
		return
	}
	s.avoidVer(r.row, r.col+1)
}
