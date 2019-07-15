package gameplate

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ResultMapLink game result base info
func ResultMapLink(scrollIndex, plate, line interface{}, scores int64, islink bool) map[string]interface{} {
	result := make(map[string]interface{})

	result["line"] = line
	result["plateindex"] = scrollIndex
	result["plate"] = plate
	result["scores"] = scores
	if islink {
		result["islink"] = 1
	} else {
		result["islink"] = 0
	}
	return result
}

// ResultMap game result base info
func ResultMap(scrollIndex, plate interface{}, scores int64, islink bool) map[string]interface{} {
	result := make(map[string]interface{})

	result["plateindex"] = scrollIndex
	result["plate"] = plate
	result["scores"] = scores
	if islink {
		result["islink"] = 1
	} else {
		result["islink"] = 0
	}
	return result
}

// NewPlate 1D plate
func NewPlate(plateSize []int, scroll [][]int) ([]int, []int) {
	var ScrollIndex []int
	var plate []int
	var index int

	for i := range plateSize {
		index = rand.Intn(len(scroll[i]))
		plate = append(plate, scroll[i][index])
		ScrollIndex = append(ScrollIndex, index)
	}

	return ScrollIndex, plate
}

// NewPlate2D plate
func NewPlate2D(plateSize []int, scroll [][]int) ([][]int, [][]int) {
	var index, plateIndex int
	ScrollIndex := make([][]int, len(plateSize))
	plate := make([][]int, len(plateSize))

	for scrollIndex, value := range plateSize {
		index = rand.Intn(len(scroll[scrollIndex]))

		for j := 0; j < value; j++ {
			plateIndex = (index + j) % len(scroll[scrollIndex])
			ScrollIndex[scrollIndex] = append(ScrollIndex[scrollIndex], plateIndex)
			plate[scrollIndex] = append(plate[scrollIndex], scroll[scrollIndex][plateIndex])
		}
	}

	return ScrollIndex, plate
}

// Line243ResultArray ...
func Line243ResultArray(plate [][]int, lineMap [][]int, option PlateOption) []InfoLine243 {
	var result []InfoLine243
	var resultLine InfoLine243

	for _, line := range lineMap {
		resultLine = GetLine243Point(plate, line, option)
		if len(resultLine.LinePoint) > option.LineMiniCount {
			result = append(result, resultLine)
		}
	}

	return result
}

// GetLine243Point return line point(index) number
func GetLine243Point(plate [][]int, line []int, option PlateOption) InfoLine243 {
	var pointSymbol int
	var isPointWild bool
	var isMainWild bool

	result := NewInfoLine243()
	mainSymbol := plate[0][line[0]]

	for i, point := range line {
		pointSymbol = plate[i][point]
		isPointWild, _ = option.IsWild(pointSymbol)

		if !isPointWild {
			isMainWild, _ = option.IsWild(mainSymbol)
			if isMainWild {
				mainSymbol = pointSymbol
			} else if mainSymbol != pointSymbol {
				return result
			}
		}
		result.AddNewPoint([]int{point}, plate[i], option)
	}

	return result
}

// PlateToLinePlate ...
func PlateToLinePlate(plate [][]int, lineMap [][]int) [][]int {
	var plateLineMap [][]int
	var plateline []int

	for _, linePoint := range lineMap {
		plateline = []int{}
		for lineIndex, point := range linePoint {
			plateline = append(plateline, plate[lineIndex][point])
		}
		plateLineMap = append(plateLineMap, plateline)
	}

	return plateLineMap
}

// CutSymbolLink get line link array
func CutSymbolLink(symbolLine []int, option PlateOption) []int {
	var newSymbolLine []int
	mainSymbol := symbolLine[0]

	for _, symbol := range symbolLine {
		if isWild, _ := option.IsWild(symbol); isWild {

		} else if isWild, _ := option.IsWild(mainSymbol); isWild {
			mainSymbol = symbol
		} else if symbol != mainSymbol {
			break
		}

		newSymbolLine = append(newSymbolLine, symbol)
	}

	return newSymbolLine
}
