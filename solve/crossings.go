package solve

import (
	"github.com/joshprzybyszewski/slitherlink/model"
)

type crossings struct {
	// [col]numLines/AvoidsInThatCol
	cols      [maxPinsPerLine]model.Dimension
	colsAvoid [maxPinsPerLine]model.Dimension

	// [row]numLines/AvoidsInThatRow
	rows      [maxPinsPerLine]model.Dimension
	rowsAvoid [maxPinsPerLine]model.Dimension

	// This is the "target" for an _almost_ empty row/col
	targetRow model.Dimension
	targetCol model.Dimension
}

func newCrossings(
	width, height model.Size,
) crossings {
	return crossings{
		targetRow: model.Dimension(width) - 1,
		targetCol: model.Dimension(height) - 1,
	}
}

func (c *crossings) lineHor(
	col model.Dimension,
	s *state,
) {
	if col == 0 {
		return
	}
	c.cols[col]++
	if c.colsAvoid[col]+c.cols[col] == c.targetRow {
		c.completeCol(col, s)
	}

}

func (c *crossings) avoidHor(
	col model.Dimension,
	s *state,
) {
	if col == 0 {
		return
	}
	c.colsAvoid[col]++
	if c.colsAvoid[col]+c.cols[col] == c.targetRow {
		c.completeCol(col, s)
	}
}

func (c *crossings) completeCol(
	col model.Dimension,
	s *state,
) {
	row := getEmptyCrossingInColumn(s, col)
	if c.cols[col]%2 == 0 {
		s.avoidHor(row, col)
	} else {
		s.lineHor(row, col)
	}
}

func (c *crossings) lineVer(
	row model.Dimension,
	s *state,
) {
	if row == 0 {
		return
	}
	c.rows[row]++
	if c.rowsAvoid[row]+c.rows[row] == c.targetCol {
		c.completeRow(row, s)
	}
}

func (c *crossings) avoidVer(
	row model.Dimension,
	s *state,
) {
	if row == 0 {
		return
	}
	c.rowsAvoid[row]++
	if c.rowsAvoid[row]+c.rows[row] == c.targetCol {
		c.completeRow(row, s)
	}
}

func (c *crossings) completeRow(
	row model.Dimension,
	s *state,
) {
	col := getEmptyCrossingInRow(s, row)
	if c.rows[row]%2 == 0 {
		s.avoidVer(row, col)
	} else {
		s.lineVer(row, col)
	}
}

func getEmptyCrossingInColumn(
	s *state,
	col model.Dimension,
) model.Dimension {
	for row := model.Dimension(1); row <= model.Dimension(s.height); row++ {
		if !s.hasHorDefined(row, col) {
			return row
		}
	}
	return 0
}

func getEmptyCrossingInRow(
	s *state,
	row model.Dimension,
) model.Dimension {
	for col := model.Dimension(1); col <= model.Dimension(s.width); col++ {
		if !s.hasVerDefined(row, col) {
			return col
		}
	}
	return 0
}
