package gameplate

// InfoLine minimum result structure
type InfoLine struct {
	// plate info
	ScotterPoint    [][]int `json:"ScotterPoint"`
	WildPoint       [][]int `json:"WildPoint"`
	LineSymbolPoint [][]int `json:"LineSymbolPoint"`
	LineSymbolNum   [][]int `json:"LineSymbolNum"`

	// pay info
	Score          int64 `json:"Score"`
	JackPotScore   int64 `json:"JackPotScore"`
	SpecialWinRate int64 `json:"SpecialWinRate"`
	LineWinRate    int   `json:"LineWinRate"`

	// // bound game info
	// RespinCount   int `json:"RespinCount,omitempty"`
	// FreeGameCount int `json:"FreeGameCount,omitempty"`

	// // bonud game flag
	// IsRespin      int `json:"IsRespin"`      // Respin Game some scroll respin
	// IsFreeGame    int `json:"IsFreeGame"`    // Free Game free spin
	// IsScotterGame int `json:"IsScotterGame"` // Scotter Game spcial game
}

// WildCount ...
func (I *InfoLine) WildCount() int {
	var wildCount int

	for _, wildPointArray := range I.WildPoint {
		wildCount += len(wildPointArray)
	}

	return wildCount
}

// ScotterCount ...
func (I *InfoLine) ScotterCount() int {
	var scotterCount int

	for _, scrtterPointArray := range I.ScotterPoint {
		scotterCount += len(scrtterPointArray)
	}

	return scotterCount
}

// AddEmptyPoint add new empty point
func (I *InfoLine) AddEmptyPoint() {

	var wildPoint = make([]int, 0)
	var scotterPoint = make([]int, 0)
	var symbolNums = make([]int, 0)

	// symbolNums = append(symbolNums, symbolNum)
	// if isWild, _ := option.IsWild(symbolNum); isWild {
	// 	wildPoint = append(wildPoint, point)
	// } else if isScotter, _ := option.IsScotter(symbolNum); isScotter {
	// 	scotterPoint = append(scotterPoint, point)
	// }

	I.WildPoint = append(I.WildPoint, wildPoint)
	I.ScotterPoint = append(I.ScotterPoint, scotterPoint)
	I.LineSymbolPoint = append(I.LineSymbolPoint, []int{})
	I.LineSymbolNum = append(I.LineSymbolNum, symbolNums)
}

// AddNewPoint add new line point
func (I *InfoLine) AddNewPoint(symbolNum int, point int, option PlateOption) {

	var wildPoint = make([]int, 0)
	var scotterPoint = make([]int, 0)
	var symbolNums = make([]int, 0)

	symbolNums = append(symbolNums, symbolNum)
	if isWild, _ := option.IsWild(symbolNum); isWild {
		wildPoint = append(wildPoint, point)
	} else if isScotter, _ := option.IsScotter(symbolNum); isScotter {
		scotterPoint = append(scotterPoint, point)
	}

	I.WildPoint = append(I.WildPoint, wildPoint)
	I.ScotterPoint = append(I.ScotterPoint, scotterPoint)
	I.LineSymbolPoint = append(I.LineSymbolPoint, []int{point})
	I.LineSymbolNum = append(I.LineSymbolNum, symbolNums)
}

// AddNewLine ...
func (I *InfoLine) AddNewLine(symbolNums []int, linePoint []int, option PlateOption) {

	var wildPoint = make([]int, 0)
	var scotterPoint = make([]int, 0)

	// symbolNums = append(symbolNums, symbolNum)
	for Index, rowSymbolNum := range symbolNums {
		if isWild, _ := option.IsWild(rowSymbolNum); isWild {
			wildPoint = append(wildPoint, linePoint[Index])
		} else if isScotter, _ := option.IsScotter(rowSymbolNum); isScotter {
			scotterPoint = append(scotterPoint, linePoint[Index])
		}
	}

	I.WildPoint = append(I.WildPoint, wildPoint)
	I.ScotterPoint = append(I.ScotterPoint, scotterPoint)
	I.LineSymbolPoint = append(I.LineSymbolPoint, linePoint)
	I.LineSymbolNum = append(I.LineSymbolNum, symbolNums)
}

// NewInfoLine Get default init NewLineInfo
func NewInfoLine() InfoLine {
	var result InfoLine
	return result
}
