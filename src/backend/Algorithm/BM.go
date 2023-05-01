package Algorithm

import (
	"math"
)

func buildLast(input string) []int {
	bl := make([]int, 128)
	for i := 0; i < 128; i++ {
		bl[i] = -1
	}
	for i := 0; i < len(input); i++ {
		bl[input[i]] = i
	}

	return bl
}

func BM(input string, text string) (int, float64) {

	bl := buildLast(input)
	simLr := HammingDistance(input, text)
	n, m := len(text), len(input)
	i := m - 1
	j := m - 1

	if i > n-1 {
		return -1, simLr
	}

	for {
		if input[j] == text[i] {
			if j == 0 {
				return i, 100
			} else {
				i--
				j--
			}
		} else {
			lo := bl[text[i]]
			i += m - int(math.Min(float64(j), float64(lo+1)))
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}

	return -1, simLr
}
