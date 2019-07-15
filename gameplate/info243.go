package gameplate

// EmptyNum is mean not this symbol
const EmptyNum = -9999

// InfoLine243 minimum result structure
type InfoLine243 struct {
	ScotterPoint  [][]int
	WildPoint     [][]int
	LinePoint     [][]int
	LineSymbolNum [][]int
	MainSymbol    int
	Score         int64
	JackPartScore int64
	WinRate       int
}

// NewInfoLine243 Get default init NewLineInfo
func NewInfoLine243() InfoLine243 {
	var result InfoLine243
	result.MainSymbol = EmptyNum
	return result
}

// AddNewPoint add new line point
func (l *InfoLine243) AddNewPoint(col []int, symbol []int, option PlateOption) {

	var wildPoint []int
	var scotterPoint []int
	var symbolNums []int
	var symbolNum int

	for _, point := range col {
		symbolNum = symbol[point]
		symbolNums = append(symbolNums, symbolNum)
		if isWild, wildNum := option.IsWild(symbolNum); isWild {
			wildPoint = append(wildPoint, point)
			if l.MainSymbol == EmptyNum {
				l.MainSymbol = wildNum
			}
		} else if isScotter, scotterNum := option.IsScotter(symbolNum); isScotter {
			scotterPoint = append(scotterPoint, point)
			if l.MainSymbol == EmptyNum {
				l.MainSymbol = scotterNum
			}
		} else {
			l.MainSymbol = symbol[point]
		}
	}

	l.WildPoint = append(l.WildPoint, wildPoint)
	l.ScotterPoint = append(l.ScotterPoint, scotterPoint)
	l.LinePoint = append(l.LinePoint, col)
	l.LineSymbolNum = append(l.LineSymbolNum, symbolNums)
}

// WinLineCount get symbol win count
func (l *InfoLine243) WinLineCount() int {
	return len(l.LinePoint)
}

// Checkout Count line score
func (l *InfoLine243) Checkout(betMoney int64, winRate int) {
	l.WinRate = winRate
	l.Score = betMoney * int64(winRate)
}
