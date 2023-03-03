package solve

import (
	"github.com/joshprzybyszewski/slitherlink/model"
)

type settledState uint8

const (
	invalid       settledState = 1
	solved        settledState = 2
	validUnsolved settledState = 3
	unexpected    settledState = 4
)

// returns true if the state is still valid
func settle(
	s *state,
) settledState {
	if s.hasInvalid {
		return invalid
	}

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	if s.rules.runAllChecks(s) == invalid {
		return invalid
	}

	if s.hasInvalid {
		return invalid
	}

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	return validUnsolved
}

func settleCycle(
	s *state,
) settledState {

	// Please only call this if the state has a cycle
	if !s.paths.hasCycle {
		return unexpected
	}

	if s.paths.cycleSeen != s.paths.cycleTarget {
		// there's a cycle, but it doesn't include all of the nodes.
		return invalid
	}

	// we found a state that includes a cycle with all of the nodes.
	// avoid all of the remaining spots in the state, and see if it's
	// still valid: this eliminates the bad state of having tertiary paths set.

	avoidAllUnknowns(s)

	if !hasValidCrossings(s) {
		return invalid
	}

	if checkEntireRuleset(s) != validUnsolved {
		return invalid
	}

	if !s.paths.hasCycle {
		return unexpected
	}

	// re-validate our assumptions after checking all the rules
	if s.hasInvalid || s.paths.cycleSeen != s.paths.cycleTarget {
		return invalid
	}

	return solved
}

func hasValidCrossings(s *state) bool {
	bit := model.DimensionBit(1 << 1)
	for i := 1; i <= int(s.height); i++ {
		if !hasValidCrossingsForHorizontalBit(s, bit) {
			return false
		}
		if i <= int(s.width) &&
			!hasValidCrossingsForVerticalBit(s, bit) {
			return false
		}
		bit <<= 1
	}

	return true
}

func hasValidCrossingsForVerticalBit(
	s *state,
	bit model.DimensionBit,
) bool {
	var l uint8
	for i := 1; i <= int(s.height); i++ {
		if s.horizontalLines[i]&bit == bit {
			l++
		} else if s.horizontalAvoids[i]&bit != bit {
			return true
		}
	}
	return l%2 == 0
}

func hasValidCrossingsForHorizontalBit(
	s *state,
	bit model.DimensionBit,
) bool {
	var l uint8
	for i := 1; i <= int(s.width); i++ {
		if s.verticalLines[i]&bit == bit {
			l++
		} else if s.verticalAvoids[i]&bit != bit {
			return true
		}
	}
	return l%2 == 0
}

func avoidAllUnknowns(
	s *state,
) {
	var col model.Dimension
	for row := model.Dimension(1); row <= model.Dimension(s.height); row++ {
		for col = model.Dimension(1); col <= model.Dimension(s.width); col++ {
			if !s.horLineAt(row, col) {
				s.avoidHor(row, col)
			}
			if !s.verLineAt(row, col) {
				s.avoidVer(row, col)
			}
		}
	}
}

func checkEntireRuleset(s *state) settledState {
	maxC := model.Dimension(s.width).Bit()
	var c model.DimensionBit

	for r := model.Dimension(0); r <= model.Dimension(s.height); r++ {
		s.rules.checkHorizontal(r, 0)
		for c = 1; c <= maxC; c <<= 1 {
			s.rules.checkHorizontal(r, c)
		}
	}

	maxR := model.Dimension(s.height).Bit()
	var r model.DimensionBit
	for c := model.Dimension(0); c <= model.Dimension(s.width); c++ {
		s.rules.checkVertical(0, c)
		for r = 1; r <= maxR; r <<= 1 {
			s.rules.checkVertical(r, c)
		}
	}

	return s.rules.runAllChecks(s)
}
