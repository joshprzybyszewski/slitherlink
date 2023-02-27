package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	return fmt.Sprintf("%dx%d %s", i.GetSize(), i.GetSize(), i.GetDifficulty())
}

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) GetSize() Size {
	if i < MinIterator || i > MaxIterator {
		return invalidSize
	}

	if i == 0 {
		return six
	}
	if i <= 3 {
		return eight
	}
	if i <= 6 {
		return ten
	}
	if i <= 9 {
		return fifteen
	}
	if i <= 12 {
		return twenty
	}
	if i == 13 {
		return daily
	}
	if i == 14 {
		return weekly
	}
	if i == 15 {
		return monthly
	}
	if i <= 18 {
		return twentyfive
	}

	return invalidSize
}

func (i Iterator) GetHeight() Size {
	if i == 0 {
		// 5x5 Normal
		return 30
	}
	if i == 4 {
		// 5x5 Hard
		return 30
	}
	if i == 10 {
		// 7x7 Normal
		return 30
	}
	if i == 11 {
		// 7x7 Hard
		return 30
	}
	if i == 1 {
		// 10x10 Normal
		return 30
	}
	if i == 5 {
		// 10x10 Hard
		return 30
	}
	if i == 2 {
		// 15x15 Normal
		return 30
	}
	if i == 6 {
		// 15x15 Hard
		return 30
	}
	if i == 3 {
		// 20x20 Normal
		return 30
	}
	if i == 7 {
		// 20x20 Hard
		return 30
	}
	if i == 8 {
		// 25x30 Normal
		return 30
	}
	if i == 9 {
		// 25x30 Hard
		return 30
	}

	return invalidSize
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i == 0 {
		return easy
	}
	if i >= 13 && i <= 15 {
		return hard
	}
	if i > 15 {
		return Difficulty(i - 15)
	}

	return Difficulty((i-1)%3) + 1
}
