package model

import "strings"

type Solution struct {
	Height Size
	Width  Size

	Horizontals [MaxPinsPerLine]DimensionBit
	Verticals   [MaxPinsPerLine]DimensionBit
}

func (s *Solution) String() string {
	var sb strings.Builder

	for r := Dimension(0); r < Dimension(s.Height); r++ {
		for c := Dimension(0); c < Dimension(s.Width); c++ {
			sb.WriteByte('*')
			sb.WriteByte(' ')
			if s.Horizontals[r]&c.Bit() != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := Dimension(0); c < Dimension(s.Width); c++ {
			if s.Verticals[c]&r.Bit() != 0 {
				sb.WriteByte('|')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
			sb.WriteByte(' ')
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (s *Solution) Pretty(
	nodes []Node,
) string {
	var sb strings.Builder

	for r := Dimension(0); r < Dimension(s.Height); r++ {
		for c := Dimension(0); c < Dimension(s.Width); c++ {
			sb.WriteByte('*')
			sb.WriteByte(' ')
			if s.Horizontals[r]&c.Bit() != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := Dimension(0); c < Dimension(s.Width); c++ {
			if s.Verticals[c]&r.Bit() != 0 {
				sb.WriteByte('|')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
			sb.WriteByte(getNodeChar(nodes, r, c))
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func getNodeChar(
	nodes []Node,
	r, c Dimension,
) byte {
	for _, n := range nodes {
		if n.Row == r && n.Col == c {
			return '0' + byte(n.Num)
		}
	}
	return ' '
}
