package Algorithm

import (
	"fmt"
	"strconv"
)
import "math"

func Calculate(querry string) string {
	// Basis
	if isANumber(querry) {
		return querry
	} else if IsQuerryAtom(querry) {
		if isKurung(querry, 1, len(querry)-2) {
			return convertToString(calculateTwoNumber(querry[1 : len(querry)-1]))
		} else {
			return convertToString(calculateTwoNumber(querry))
		}
	} else {
		solveAtom, start, end := GetQuerry(querry)
		solveAtom = Calculate(solveAtom)
		var solveGlobal string
		fmt.Println("Awal")
		fmt.Println(querry)
		fmt.Println(solveAtom)
		if isKurung(querry, start, end) {
			fmt.Println("Kurung")
			solveGlobal = querry[:start-1] + solveAtom + querry[end+2:]
		} else {
			fmt.Println("Bukan")
			solveGlobal = querry[:start] + solveAtom + querry[end+1:]
		}
		fmt.Println(solveGlobal)
		return Calculate(solveGlobal)
	}
}

func calculateTwoNumber(querryAtom string) float64 {
	operand1, operator, operand2 := getOperandOperator(querryAtom)
	var retVal float64
	if operator == "^" {
		retVal = math.Pow(float64(operand1), float64(operand2))
	} else if operator == "*" {
		retVal = operand1 * operand2
	} else if operator == "/" {
		retVal = operand1 / operand2
	} else if operator == "+" {
		retVal = operand1 + operand2
	} else if operator == "-" {
		retVal = operand1 - operand2
	}
	return retVal
}

func getOperandOperator(querryAtom string) (float64, string, float64) {
	var operand2Turn bool = false
	var tempOperand1 string
	var operator string
	var tempOperand2 string
	var currentChar string
	var firstChar bool = true
	for _, character := range querryAtom {
		currentChar = string(character)
		if firstChar {
			tempOperand1 += currentChar
			firstChar = false
		} else if operand2Turn {
			tempOperand2 += currentChar
		} else if isOperator(currentChar) {
			operand2Turn = true
			operator = currentChar
		} else {
			tempOperand1 += currentChar
		}
	}
	operand1 := convertToFloat(tempOperand1)
	operand2 := convertToFloat(tempOperand2)

	return operand1, operator, operand2
}

func convertToFloat(operand string) float64 {
	if string(operand[0]) == "-" {
		tempRetVal := string(operand[1:])
		retVal, _ := strconv.ParseFloat(tempRetVal, 64)
		return -1 * retVal
	} else {
		retVal, _ := strconv.ParseFloat(operand, 64)
		return retVal
	}
}

func isOperator(symbol string) bool {
	validOperator := map[string]bool{}
	for _, v := range []string{"+", "-", "*", "/", "^"} {
		validOperator[v] = true
	}

	return validOperator[symbol]
}

func isANumber(symbol string) bool {
	validNumber := map[string]bool{}
	for _, v := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."} {
		validNumber[v] = true
	}
	if len(symbol) == 0 {
		return false
	} else {
		for i := 1; i < len(symbol); i++ {
			currentChar := string(symbol[i])
			if i == 0 && !(currentChar == "-" || validNumber[currentChar]) {
				return false
			} else if !validNumber[currentChar] {
				return false
			}
		}
	}
	return true
}

func GetQuerry(querry string) (string, int, int) {
	startIndex := 0
	endIndex := len(querry) - 1
	startIndexKurung := -1
	endIndexKurung := -1
	startIndexPangkat := -1
	endIndexPangkat := -1
	startIndexKali := -1
	endIndexKali := -1
	startIndexBagi := -1
	endIndexBagi := -1
	startIndexTambah := -1
	endIndexTambah := -1
	startIndexKurang := -1
	endIndexKurang := -1

	prevNumberStartIndex := 0

	for i, x := range querry {
		currentChar := string(x)
		if currentChar == "(" {
			if startIndexKurung == -1 {
				startIndexKurung = i + 1
			}
		} else if currentChar == ")" {
			endIndexKurung = i - 1
		} else if currentChar == "^" && startIndexPangkat == -1 {
			startIndexPangkat = prevNumberStartIndex
			endIndexPangkat = getPostNumberEndIndex(querry, i)
		} else if currentChar == "*" && startIndexKali == -1 {
			startIndexKali = prevNumberStartIndex
			endIndexKali = getPostNumberEndIndex(querry, i)
		} else if currentChar == "/" && startIndexBagi == -1 {
			startIndexBagi = prevNumberStartIndex
			endIndexBagi = getPostNumberEndIndex(querry, i)
		} else if currentChar == "+" && startIndexTambah == -1 {
			startIndexTambah = prevNumberStartIndex
			endIndexTambah = getPostNumberEndIndex(querry, i)
		} else if currentChar == "-" && startIndexKurang == -1 && i != 0 && !isOperator(string(querry[i-1])) {
			startIndexKurang = prevNumberStartIndex
			endIndexKurang = getPostNumberEndIndex(querry, i)
		}
		if isOperator(currentChar) && i != 0 && !isOperator(string(querry[i-1])) {
			prevNumberStartIndex = i + 1

		}
	}

	if startIndexKurung != -1 {
		return querry[startIndexKurung : endIndexKurung+1], startIndexKurung, endIndexKurung
	}
	if startIndexPangkat != -1 {
		return querry[startIndexPangkat : endIndexPangkat+1], startIndexPangkat, endIndexPangkat
	}
	if startIndexKali != -1 && startIndexBagi != -1 {
		if startIndexKali < startIndexBagi {
			return querry[startIndexKali : endIndexKali+1], startIndexKali, endIndexKali
		} else {
			return querry[startIndexBagi : endIndexBagi+1], startIndexBagi, endIndexBagi
		}
	}
	if startIndexKali != -1 {
		return querry[startIndexKali : endIndexKali+1], startIndexKali, endIndexKali
	}
	if startIndexBagi != -1 {
		return querry[startIndexBagi : endIndexBagi+1], startIndexBagi, endIndexBagi
	}
	if startIndexTambah != -1 && startIndexKurang != -1 {
		if startIndexTambah < startIndexKurang {
			return querry[startIndexTambah : endIndexTambah+1], startIndexTambah, endIndexTambah
		} else {
			return querry[startIndexKurang : endIndexKurang+1], startIndexKurang, endIndexKurang
		}
	}
	if startIndexTambah != -1 {
		return querry[startIndexTambah : endIndexTambah+1], startIndexTambah, endIndexTambah
	}
	if startIndexKurang != -1 {
		return querry[startIndexKurang : endIndexKurang+1], startIndexKurang, endIndexKurang
	}
	return querry[startIndex : endIndex+1], startIndex, endIndex
}

func IsQuerryAtom(querry string) bool {
	firstChar := true
	operatorCount := 0
	prevOperator := false
	for _, character := range querry {
		currentChar := string(character)
		if operatorCount > 1 {
			return false
		} else if firstChar {
			firstChar = false
		} else if currentChar == "-" {
			if !prevOperator {
				operatorCount++
				prevOperator = true
			} else {
				prevOperator = false
			}
		} else if isOperator(currentChar) {
			operatorCount++
			prevOperator = true
		} else {
			prevOperator = false
		}
	}
	if operatorCount == 0 {
		return false
	}
	return true
}

func convertToString(number float64) string {
	retVal := strconv.FormatFloat(number, 'f', -1, 64)
	return retVal
}

func getPostNumberEndIndex(querry string, startIndex int) int {
	var i int
	negFound := false
	for i = startIndex + 1; i < len(querry); i++ {
		currentChar := string(querry[i])
		if currentChar == "-" && !negFound {
			negFound = true
		} else if (currentChar == "-" && negFound) || (currentChar != "-" && isOperator(currentChar)) {
			return i - 1
		}
	}
	return i - 1
}

func isKurung(querry string, start int, end int) bool {
	if start == 0 || end == len(querry)-1 {
		return false
	} else if string(querry[start-1]) == "(" && string(querry[end+1]) == ")" {
		return true
	} else {
		return false
	}
}
