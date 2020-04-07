package gameplate

import (
	"math/rand"
	"time"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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

// ResultMapLine game result base info
func ResultMapLine(scrollIndex, plate, gameresult interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	result["plateindex"] = scrollIndex
	result["plate"] = plate
	result["gameresult"] = gameresult
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

// NewPlateOneByOne2D ....
func NewPlateOneByOne2D(plateSize []int, scroll [][]int) ([][]int, [][]int) {
	var index int
	ScrollIndex := make([][]int, len(plateSize))
	plate := make([][]int, len(plateSize))

	for scrollIndex, value := range plateSize {

		for j := 0; j < value; j++ {
			index = rand.Intn(len(scroll[scrollIndex]))
			ScrollIndex[scrollIndex] = append(ScrollIndex[scrollIndex], index)
			plate[scrollIndex] = append(plate[scrollIndex], scroll[scrollIndex][index])
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

// PlateSymbolCollect collect symbol in plate
func PlateSymbolCollect(targetSymbolNum int, plate [][]int, option PlateOption, funcOption map[string]interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	var symbolPointCollation = make([][]int, 0)
	var symbolNumCollation = make([][]int, 0)
	var targetSymbolCount, symbolCollectCount int

	var isIncludeWild = foundation.InterfaceToBool(funcOption["isincludewild"])
	var isSeachAllPlate = foundation.InterfaceToBool(funcOption["isseachallplate"])
	var IsScotterTarget, _ = option.IsScotter(targetSymbolNum)

	for _, colArray := range plate {
		var rowPointArray []int
		var rowSymbolArray []int
		for rowIndex, rowSymbol := range colArray {

			if targetSymbolNum == rowSymbol {
				targetSymbolCount++
				symbolCollectCount++
				rowSymbolArray = append(rowSymbolArray, rowSymbol)
				rowPointArray = append(rowPointArray, rowIndex)
			} else if !IsScotterTarget {
				IsWild, _ := option.IsWild(rowSymbol)

				if isIncludeWild && IsWild {
					symbolCollectCount++
					rowSymbolArray = append(rowSymbolArray, rowSymbol)
					rowPointArray = append(rowPointArray, rowIndex)
				}
			}
		}

		if len(rowPointArray) <= 0 && !isSeachAllPlate {
			break
		}
		symbolNumCollation = append(symbolNumCollation, rowSymbolArray)
		symbolPointCollation = append(symbolPointCollation, rowPointArray)
	}

	result["targetsymbolcount"] = targetSymbolCount
	result["symbolcollectcount"] = symbolCollectCount
	result["symbolnumcollation"] = symbolNumCollation
	result["symbolpointcollation"] = symbolPointCollation
	return result

}

// PlateSymbolCount collect target symbol
func PlateSymbolCount(targetSymbol int, plate [][]int) int {
	var count int

	for _, cols := range plate {
		for _, rowSymnol := range cols {
			if rowSymnol == targetSymbol {
				count++
			}
		}
	}
	return count
}

// PlateSymbolPoint collect target symbol point
func PlateSymbolPoint(targetSymbol int, plate [][]int) (int, [][]int) {
	var count int
	var point = make([][]int, len(plate))

	for colIndex, cols := range plate {
		for _, rowSymnol := range cols {
			if rowSymnol == targetSymbol {
				count++
				point[colIndex] = append(point[colIndex], colIndex)
			}
		}
	}
	return count, point
}

// CutSymbolLink get line link array
func CutSymbolLink(symbolLine []int, option PlateOption) []int {
	var newSymbolLine []int
	mainSymbol := symbolLine[0]

	for _, symbol := range symbolLine {
		if isScotter, _ := option.IsScotter(symbol); isScotter {
			if isMainScotter, _ := option.IsScotter(mainSymbol); !isMainScotter {
				break
			}
		} else if isWild, _ := option.IsWild(symbol); isWild {
			if isMainScotter, _ := option.IsScotter(mainSymbol); isMainScotter {
				break
			}
		} else if isWild, _ := option.IsWild(mainSymbol); isWild {
			mainSymbol = symbol
		} else if symbol != mainSymbol {
			break
		}

		newSymbolLine = append(newSymbolLine, symbol)
	}

	return newSymbolLine
}

// LineMulitResult return line mulit win if was
func LineMulitResult(line []int, option PlateOption) [][]int {
	var result [][]int
	winLineMap := make(map[int][]int)
	mainSymbol := line[0]

	if isWild, _ := option.IsWild(mainSymbol); !isWild {
		return [][]int{line}
	}

	for _, value := range line {

		if mainSymbol == value { // is wild and same wild
			winLineMap[value] = append(winLineMap[value], value)
		} else if isWild, _ := option.IsWild(value); isWild { // is wild but different wild
			break
		} else { // is not wild and not equal mainsymbol
			break
		}

	}

	result = append(result, line)
	for _, winLine := range winLineMap {
		result = append(result, winLine)
	}

	return result
}
