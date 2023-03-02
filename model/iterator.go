package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	return fmt.Sprintf("%dx%d %s", i.GetWidth()-1, i.GetHeight()-1, i.GetDifficulty())
}

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) GetWidth() Size {
	if i < MinIterator || i > MaxIterator {
		return invalidSize
	}

	if i == 0 {
		// 5 x 5 Normal
		return 5 + 1
	}
	if i == 4 {
		// 5 x 5 Hard
		return 5 + 1
	}
	if i == 10 {
		// 7 x 7 Normal
		return 7 + 1
	}
	if i == 11 {
		// 7 x 7 Hard
		return 7 + 1
	}
	if i == 1 {
		// 10 x 10 Normal
		return 10 + 1
	}
	if i == 5 {
		// 10 x 10 Hard
		return 10 + 1
	}
	if i == 2 {
		// 15 x 15 Normal
		return 15 + 1
	}
	if i == 6 {
		// 15 x 15 Hard
		return 15 + 1
	}
	if i == 3 {
		// 20 x 20 Normal
		return 20 + 1
	}
	if i == 7 {
		// 20 x 20 Hard
		return 20 + 1
	}
	if i == 8 {
		// 25 x 30 Normal
		return 25 + 1
	}
	if i == 9 {
		// 25 x 30 Hard
		return 25 + 1
	}

	return invalidSize
}

func (i Iterator) GetHeight() Size {
	if i < MinIterator || i > MaxIterator {
		return invalidSize
	}

	if i == 0 {
		// 5 x 5 Normal
		return 5 + 1
	}
	if i == 4 {
		// 5 x 5 Hard
		return 5 + 1
	}
	if i == 10 {
		// 7 x 7 Normal
		return 7 + 1
	}
	if i == 11 {
		// 7 x 7 Hard
		return 7 + 1
	}
	if i == 1 {
		// 10 x 10 Normal
		return 10 + 1
	}
	if i == 5 {
		// 10 x 10 Hard
		return 10 + 1
	}
	if i == 2 {
		// 15 x 15 Normal
		return 15 + 1
	}
	if i == 6 {
		// 15 x 15 Hard
		return 15 + 1
	}
	if i == 3 {
		// 20 x 20 Normal
		return 20 + 1
	}
	if i == 7 {
		// 20 x 20 Hard
		return 20 + 1
	}
	if i == 8 {
		// 25 x 30 Normal
		return 30 + 1
	}
	if i == 9 {
		// 25 x 30 Hard
		return 30 + 1
	}

	return invalidSize
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i == 0 {
		// 5x5 Normal
		return medium
	}
	if i == 4 {
		// 5x5 Hard
		return hard
	}
	if i == 10 {
		// 7x7 Normal
		return medium
	}
	if i == 11 {
		// 7x7 Hard
		return hard
	}
	if i == 1 {
		// 10x10 Normal
		return medium
	}
	if i == 5 {
		// 10x10 Hard
		return hard
	}
	if i == 2 {
		// 15x15 Normal
		return medium
	}
	if i == 6 {
		// 15x15 Hard
		return hard
	}
	if i == 3 {
		// 20x20 Normal
		return medium
	}
	if i == 7 {
		// 20x20 Hard
		return hard
	}
	if i == 8 {
		// 25x30 Normal
		return medium
	}
	if i == 9 {
		// 25x30 Hard
		return hard
	}

	return Difficulty((i-1)%3) + 1
}
