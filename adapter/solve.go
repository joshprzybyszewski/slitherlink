package adapter

import (
	"strings"
	"time"

	"github.com/joshprzybyszewski/puzzler/model"
	slm "github.com/joshprzybyszewski/slitherlink/model"
	"github.com/joshprzybyszewski/slitherlink/solve"
)

func solveGame(
	g *model.Game,
	timeout time.Duration,
) error {
	iter := slm.Iterator(g.Iterator)
	ns := convertTask(iter, g.Task)

	sol, err := solve.FromNodes(
		iter.GetWidth(),
		iter.GetHeight(),
		ns,
		timeout,
	)
	if err != nil {
		return err
	}

	g.Answer = convertAnswer(&sol)
	return nil
}

func convertTask(
	iter slm.Iterator,
	task model.Task,
) []slm.Node {
	var r, c slm.Dimension
	maxC := slm.Dimension(iter.GetWidth() - 1)
	output := make([]slm.Node, 0, len(task)/2)

	for _, b := range task {
		if b >= '0' && b <= '9' {
			output = append(output, slm.Node{
				Coord: slm.Coord{
					Row: r,
					Col: c,
				},
				Num: int(b - '0'),
			})
		} else {
			c += slm.Dimension(b - 'a')
		}

		c++

		if c >= maxC {
			r += (c / maxC)
			c %= maxC
		}
	}

	return output
}

func convertAnswer(
	sol *slm.Solution,
) model.Answer {
	if sol == nil || sol.Width == 0 || sol.Height == 0 {
		return ``
	}

	numRows := int(sol.Height) - 1
	numCols := int(sol.Width) - 1
	var rows, cols strings.Builder
	rows.Grow(numRows * (numCols - 1))
	cols.Grow((numRows - 1) * numCols)

	for row := slm.Dimension(0); row <= slm.Dimension(numRows); row++ {
		for col := slm.Dimension(0); col < slm.Dimension(numCols); col++ {
			if sol.Horizontals[row]&col.Bit() != 0 {
				_ = rows.WriteByte('y')
			} else {
				_ = rows.WriteByte('n')
			}
		}
	}

	for row := slm.Dimension(0); row < slm.Dimension(numRows); row++ {
		for col := slm.Dimension(0); col <= slm.Dimension(numCols); col++ {
			if sol.Verticals[col]&row.Bit() != 0 {
				_ = cols.WriteByte('y')
			} else {
				_ = cols.WriteByte('n')
			}
		}
	}

	return model.Answer(rows.String() + cols.String())
}
