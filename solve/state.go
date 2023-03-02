package solve

import (
	"fmt"
	"strings"

	"github.com/joshprzybyszewski/slitherlink/model"
)

const (
	all64Bits model.DimensionBit = 0xFFFFFFFFFFFFFFFF
)

type state struct {
	rules ruleCheckCollector
	nodes []model.Node
	zeros []model.Node

	width, height model.Size
	hasInvalid    bool

	paths pathCollector

	crossings crossings

	// [row]colBitMask
	horizontalLines  [maxPinsPerLine]model.DimensionBit
	horizontalAvoids [maxPinsPerLine]model.DimensionBit

	// [col]rowBitMask
	verticalLines  [maxPinsPerLine]model.DimensionBit
	verticalAvoids [maxPinsPerLine]model.DimensionBit
}

func newState(
	width, height model.Size,
	ns []model.Node,
) state {
	r := newRules(width, height)
	rcc := newRuleCheckCollector(r)

	s := state{
		nodes:     make([]model.Node, 0, len(ns)),
		zeros:     make([]model.Node, 0, len(ns)),
		width:     width,
		height:    height,
		crossings: newCrossings(width, height),
		rules:     rcc,
	}

	// offset all of the input nodes by positive one
	for i := range ns {
		n := ns[i]
		n.Row++
		n.Col++
		if n.Num == 0 {
			s.zeros = append(s.zeros, n)
		} else {
			s.nodes = append(s.nodes, n)
		}
	}
	s.paths = newPathCollector(s.nodes)

	findGimmes(&s)

	r.populateRules(&s)

	s.initialize()

	r.populateUnknowns(&s)

	return s
}

func (s *state) initialize() {
	s.verticalAvoids[0] = all64Bits
	avoidVer := model.DimensionBit(1 | (1 << s.height))
	for i := 1; i <= int(s.width); i++ {
		s.verticalAvoids[i] |= avoidVer
	}
	s.verticalAvoids[s.width+1] = all64Bits

	s.horizontalAvoids[0] = all64Bits
	avoidHor := model.DimensionBit(1 | (1 << s.width))
	for i := 1; i <= int(s.height); i++ {
		s.horizontalAvoids[i] |= avoidHor
	}
	s.horizontalAvoids[s.height+1] = all64Bits

	if checkEntireRuleset(s) == invalid {
		fmt.Printf("Invalid State:\n%s\n", s)
		panic(`state initialization is not valid?`)
	}
}

func (s *state) toSolution() model.Solution {
	sol := model.Solution{
		Height: s.height,
		Width:  s.width,
	}

	// each line needs to be shifted by one.
	for i := 0; i < int(s.height); i++ {
		sol.Horizontals[i] = (s.horizontalLines[i+1]) >> 1
	}

	for i := 0; i < int(s.width); i++ {
		sol.Verticals[i] = (s.verticalLines[i+1]) >> 1
	}

	return sol
}

func (s *state) getMostInterestingPath() (model.Coord, bool, bool) {
	c, isHor, ok := s.paths.getInteresting(s)
	if ok {
		return c, isHor, true
	}

	for _, pp := range s.rules.rules.unknowns {
		if pp.IsHorizontal {
			if !s.hasHorDefined(pp.Row, pp.Col) {
				return pp.Coord, pp.IsHorizontal, true
			}
		} else {
			if !s.hasVerDefined(pp.Row, pp.Col) {
				return pp.Coord, pp.IsHorizontal, true
			}
		}
	}
	// there are no more interesting paths left. Likely this means that there's
	// an error in the state and we need to abort.
	return model.Coord{}, false, false
}

func (s *state) hasHorDefined(r, c model.Dimension) bool {
	return (s.horizontalLines[r]|s.horizontalAvoids[r])&c.Bit() != 0
}

func (s *state) horAt(r, c model.Dimension) (bool, bool) {
	return s.horLineAt(r, c), s.horAvoidAt(r, c)
}

func (s *state) horLineAt(r, c model.Dimension) bool {
	return s.horizontalLines[r]&c.Bit() != 0
}

func (s *state) horAvoidAt(r, c model.Dimension) bool {
	return s.horizontalAvoids[r]&c.Bit() != 0
}

func (s *state) avoidHor(r, c model.Dimension) {
	b := c.Bit()
	if s.hasInvalid || s.horizontalAvoids[r]&b == b {
		// already avoided
		return
	}
	s.horizontalAvoids[r] |= b
	if s.horizontalLines[r]&s.horizontalAvoids[r] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.rules.checkHorizontal(r, b)
	s.crossings.avoidHor(c, s)
}

func (s *state) lineHor(r, c model.Dimension) {
	b := c.Bit()
	if s.hasInvalid || s.horizontalLines[r]&b == b {
		// already a line
		return
	}
	s.horizontalLines[r] |= b
	if s.horizontalLines[r]&s.horizontalAvoids[r] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.rules.checkHorizontal(r, b)
	s.crossings.lineHor(c, s)
	s.paths.addHorizontal(r, c, s)
}

func (s *state) hasVerDefined(r, c model.Dimension) bool {
	return (s.verticalLines[c]|s.verticalAvoids[c])&r.Bit() != 0
}

func (s *state) verAt(r, c model.Dimension) (bool, bool) {
	return s.verLineAt(r, c), s.verAvoidAt(r, c)
}

func (s *state) verLineAt(r, c model.Dimension) bool {
	return s.verticalLines[c]&r.Bit() != 0
}

func (s *state) verAvoidAt(r, c model.Dimension) bool {
	return s.verticalAvoids[c]&r.Bit() != 0
}

func (s *state) avoidVer(r, c model.Dimension) {
	b := r.Bit()
	if s.hasInvalid || s.verticalAvoids[c]&b == b {
		// already avoided
		return
	}
	s.verticalAvoids[c] |= b
	if s.verticalLines[c]&s.verticalAvoids[c] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.rules.checkVertical(b, c)
	s.crossings.avoidVer(r, s)
}

func (s *state) lineVer(r, c model.Dimension) {
	b := r.Bit()
	if s.hasInvalid || s.verticalLines[c]&b == b {
		// already avoided
		return
	}
	s.verticalLines[c] |= b
	if s.verticalLines[c]&s.verticalAvoids[c] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.rules.checkVertical(b, c)
	s.crossings.lineVer(r, s)
	s.paths.addVertical(r, c, s)
}

const (
	confusedSpace       byte = '@'
	horizontalLineSpace byte = '-'
	verticalLineSpace   byte = '|'
	avoidSpace          byte = 'X'
)

func (s *state) String() string {
	var sb strings.Builder

	var isLine, isAvoid bool

	for r := 0; r <= int(s.height+1); r++ {
		for c := 0; c <= int(s.width+1); c++ {
			sb.WriteByte(s.getDot(model.Dimension(r), model.Dimension(c)))
			sb.WriteByte(' ')
			isLine, isAvoid = s.horAt(model.Dimension(r), model.Dimension(c))
			if isLine && isAvoid {
				sb.WriteByte(confusedSpace)
			} else if isLine {
				sb.WriteByte(horizontalLineSpace)
			} else if isAvoid {
				sb.WriteByte(avoidSpace)
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := 0; c <= int(s.width+1); c++ {
			isLine, isAvoid = s.verAt(model.Dimension(r), model.Dimension(c))
			if isLine && isAvoid {
				sb.WriteByte(confusedSpace)
			} else if isLine {
				sb.WriteByte(verticalLineSpace)
			} else if isAvoid {
				sb.WriteByte(avoidSpace)
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
			sb.WriteByte(s.getNode(model.Dimension(r), model.Dimension(c)))
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (s *state) getNode(r, c model.Dimension) byte {
	for _, n := range s.nodes {
		if n.Row != r || n.Col != c {
			continue
		}
		return '0' + byte(n.Num)
	}
	for _, n := range s.zeros {
		if n.Row != r || n.Col != c {
			continue
		}
		return '0' + byte(n.Num)
	}
	// if r == 0 || c == 0 || r >= model.Dimension(s.height) || c >= model.Dimension(s.width) {
	// 	return 'X'
	// }
	return ' '
}

func (s *state) getDot(r, c model.Dimension) byte {
	if r == 0 {
		return '0' + byte(c%10)
	}
	if c == 0 {
		return '0' + byte(r%10)
	}
	if r > model.Dimension(s.height) || c > model.Dimension(s.width) {
		return ' '
	}
	return '*'
}
