package solve

import "github.com/joshprzybyszewski/slitherlink/model"

func findGimmes(
	s *state,
) {

	for _, n := range s.zeros {
		s.avoidHor(n.Row, n.Col)
		s.avoidHor(n.Row+1, n.Col)
		s.avoidVer(n.Row, n.Col)
		s.avoidVer(n.Row, n.Col+1)
	}

	// gimmes: two (3) nodes in a row => the between-er is a line
	claimNearbyThrees(s)
}

func claimNearbyThrees(
	s *state,
) {
	var nodes [maxPinsPerLine]model.DimensionBit

	var c model.DimensionBit

	for _, n := range s.nodes {
		if n.Num != 3 {
			continue
		}
		c = n.Col.Bit()
		nodes[n.Row] |= c
		if (nodes[n.Row-1])&c != 0 {
			s.lineHor(n.Row, n.Col)
		}

		if (nodes[n.Row])&(c>>1) != 0 {
			s.lineVer(n.Row, n.Col)
		}
	}

}
