package gameplate

// InfoLine minimum result structure
type InfoLine struct {
	ScotterPoint  [][]int
	WildPoint     [][]int
	LinePoint     [][]int
	LineSymbolNum [][]int
	// MainSymbol    int
	Score         int64
	JackPartScore int64
	WinRate       int
}

// NewLineInfo Get default init NewLineInfo
func NewLineInfo() InfoLine {
	var result InfoLine
	// result.MainSymbol = EmptyNum
	return result
}

// AddNewPoint add new line point
func (l *InfoLine) AddNewPoint(symbolNum int, point int, option PlateOption) {

	var wildPoint []int
	var scotterPoint []int
	var symbolNums []int

	symbolNums = append(symbolNums, symbolNum)
	if isWild, _ := option.IsWild(symbolNum); isWild {
		wildPoint = append(wildPoint, point)
	} else if isScotter, _ := option.IsScotter(symbolNum); isScotter {
		scotterPoint = append(scotterPoint, point)
	}

	l.WildPoint = append(l.WildPoint, wildPoint)
	l.ScotterPoint = append(l.ScotterPoint, scotterPoint)
	l.LinePoint = append(l.LinePoint, []int{point})
	l.LineSymbolNum = append(l.LineSymbolNum, symbolNums)
}
