package fetch

import (
	"fmt"

	"github.com/joshprzybyszewski/slitherlink/model"
)

type input struct {
	ID    string
	param string
	task  string

	iter model.Iterator
}

func (i input) String() string {
	return fmt.Sprintf("Puzzle %s (Iter: %d, Size: %d x %d, Difficulty: %s)",
		i.ID,
		i.iter,
		i.iter.GetWidth(),
		i.iter.GetHeight(),
		i.iter.GetDifficulty(),
	)
}

func (i input) Task() string {
	return i.task
}

func (i input) ToNodes() []model.Node {
	var r, c model.Dimension
	maxC := model.Dimension(i.iter.GetWidth() - 1)
	output := make([]model.Node, 0, len(i.task)/2)

	for _, b := range i.task {
		if b >= '0' && b <= '9' {
			output = append(output, model.Node{
				Coord: model.Coord{
					Row: r,
					Col: c,
				},
				Num: int(b - '0'),
			})
		} else {
			c += model.Dimension(b - 'a')
		}

		c++

		if c >= maxC {
			r += (c / maxC)
			c %= maxC
		}
	}

	return output
}
