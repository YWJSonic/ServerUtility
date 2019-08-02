package gameplate

// EmptyNum is mean not this symbol
const EmptyNum = -9999

// InfoLine243 minimum result structure
type InfoLine243 struct {
	// plate info
	ScotterPoint   [][]int `json:"ScotterPoint"`
	WildPoint      [][]int `json:"WildPoint"`
	WinSymbolPoint [][]int `json:"WinSymbolPoint"`
	WinSymbolNum   [][]int `json:"WinSymbolNum"`
	LineNum        [][]int `json:"LineNum"`
	LinePoint      [][]int `json:"LinePoint"`

	// pay info
	Score          int64 `json:"Score"`
	JackPotScore   int64 `json:"JackPotScore"`
	SpecialWinRate int64 `json:"SpecialWinRate"`
	LineWinRate    int   `json:"LineWinRate"`

	// // bound game info
	// RespinCount   int `json:"RespinCount,omitempty"`
	// FreeGameCount int `json:"FreeGameCount,omitempty"`

	// // bonud game flag
	// IsLink int `json:"IsLink"`
	// IsRespin      int `json:"IsRespin"`      // Respin Game some scroll respin
	// IsFreeGame    int `json:"IsFreeGame"`    // Free Game free spin
	// IsScotterGame int `json:"IsScotterGame"` // Scotter Game spcial game
}

// WildCount ...
func (I *InfoLine243) WildCount() int {
	var wildCount int

	for _, wildPointArray := range I.WildPoint {
		wildCount += len(wildPointArray)
	}

	return wildCount
}

// ScotterCount ...
func (I *InfoLine243) ScotterCount() int {
	var scotterCount int

	for _, scrtterPointArray := range I.ScotterPoint {
		scotterCount += len(scrtterPointArray)
	}

	return scotterCount
}

// AddNewPoint add new line point
func (I *InfoLine243) AddNewPoint(symbolNum int, point int, option PlateOption) {

	var wildPoint []int
	var scotterPoint []int
	var symbolNums []int

	symbolNums = append(symbolNums, symbolNum)
	if isWild, _ := option.IsWild(symbolNum); isWild {
		wildPoint = append(wildPoint, point)
	} else if isScotter, _ := option.IsScotter(symbolNum); isScotter {
		scotterPoint = append(scotterPoint, point)
	}

	I.WildPoint = append(I.WildPoint, wildPoint)
	I.ScotterPoint = append(I.ScotterPoint, scotterPoint)
	I.WinSymbolPoint = append(I.WinSymbolPoint, []int{point})
	I.WinSymbolNum = append(I.WinSymbolNum, symbolNums)
}

// AddNewWinSymol ...
func (I *InfoLine243) AddNewWinSymol(symbolNums []int, linePoint []int, option PlateOption) {

	var wildPoint = make([]int, 0)
	var scotterPoint = make([]int, 0)

	for Index, rowSymbolNum := range symbolNums {
		if isWild, _ := option.IsWild(rowSymbolNum); isWild {
			wildPoint = append(wildPoint, linePoint[Index])
		} else if isScotter, _ := option.IsScotter(rowSymbolNum); isScotter {
			scotterPoint = append(scotterPoint, linePoint[Index])
		}
	}

	I.WildPoint = append(I.WildPoint, wildPoint)
	I.ScotterPoint = append(I.ScotterPoint, scotterPoint)
	I.WinSymbolPoint = append(I.WinSymbolPoint, linePoint)
	I.WinSymbolNum = append(I.WinSymbolNum, symbolNums)
}

// NewInfoLine243 Get default init NewLineInfo
func NewInfoLine243() InfoLine243 {
	var result InfoLine243
	return result
}
