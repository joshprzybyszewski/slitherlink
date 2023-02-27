package solve

func findGimmes(
	s *state,
) {

	for _, n := range s.nodes {
		if n.Num == 0 {
			s.avoidHor(n.Row, n.Col)
			s.avoidHor(n.Row+1, n.Col)
			s.avoidVer(n.Row, n.Col)
			s.avoidVer(n.Row, n.Col+1)
		}
	}
}
