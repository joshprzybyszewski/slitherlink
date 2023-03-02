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

func (s *Solution) ToAnswer() string {
	if s == nil || s.Width == 0 || s.Height == 0 {
		return ``
	}

	numRows := int(s.Height) - 1
	numCols := int(s.Width) - 1
	var rows, cols strings.Builder
	rows.Grow(numRows * (numCols - 1))
	cols.Grow((numRows - 1) * numCols)

	for row := Dimension(0); row <= Dimension(numRows); row++ {
		for col := Dimension(0); col < Dimension(numCols); col++ {
			if s.Horizontals[row]&col.Bit() != 0 {
				_ = rows.WriteByte('y')
			} else {
				_ = rows.WriteByte('n')
			}
		}
	}

	for row := Dimension(0); row < Dimension(numRows); row++ {
		for col := Dimension(0); col <= Dimension(numCols); col++ {
			if s.Verticals[col]&row.Bit() != 0 {
				_ = cols.WriteByte('y')
			} else {
				_ = cols.WriteByte('n')
			}
		}
	}

	return rows.String() + cols.String()
}
