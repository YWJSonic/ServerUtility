package gameplate

// PlateOption plate symbol index option
type PlateOption struct {
	Wild          []int         // symbol index
	Scotter       []int         // symbol index
	Space         []int         // symbol index
	Group         map[int][]int // groupID symbol index
	LineMiniCount int
}

// EmptyNum this number don't use, is use compare does symbol set.
func (p *PlateOption) EmptyNum() int {
	return EmptyNum
}

// IsWild wild compare
func (p *PlateOption) IsWild(item int) (bool, int) {
	for _, wild := range p.Wild {
		if wild == item {
			return true, wild
		}
	}
	return false, EmptyNum
}

// IsScotter Scotter compare
func (p *PlateOption) IsScotter(item int) (bool, int) {
	for _, scotter := range p.Scotter {
		if scotter == item {
			return true, scotter
		}
	}
	return false, EmptyNum
}

// IsSpace Space compare
func (p *PlateOption) IsSpace(item int) (bool, int) {
	for _, space := range p.Space {
		if space == item {
			return true, space
		}
	}
	return false, EmptyNum
}
