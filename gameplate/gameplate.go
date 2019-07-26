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

// ResultMap243 game result base info
func ResultMap243(scrollIndex, plate, gameresult interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	result["plateindex"] = scrollIndex
	result["plate"] = plate
	result["gameresult"] = gameresult
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
