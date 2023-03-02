package solve

func findGimmes(
	s *state,
) {

	for _, n := range s.zeros {
		s.avoidHor(n.Row, n.Col)
		s.avoidHor(n.Row+1, n.Col)
		s.avoidVer(n.Row, n.Col)
		s.avoidVer(n.Row, n.Col+1)
	}

	// TODO
	// gimmes: three (3) nodes in a row => the between-ers are lines
}
