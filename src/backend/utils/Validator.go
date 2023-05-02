package utils

import (
	"backend/algorithm"
	"unicode"
)

func IsCalculate(querry string) bool {
	for _, x := range querry {
		currentChar := string(x)
		if !(unicode.IsDigit(x) || algorithm.IsOperator(currentChar) || currentChar == "(" || currentChar == ")") {
			return false
		}
	}
	return true
}
