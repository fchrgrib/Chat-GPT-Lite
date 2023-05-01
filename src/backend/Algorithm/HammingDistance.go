package Algorithm

import "math"

func HammingDistance(input, text string) float64 {

	min, max := int(math.Min(float64(len(input)), float64(len(text)))), int(math.Max(float64(len(input)), float64(len(text))))
	distance := 0

	for i := 0; i < min; i++ {
		if input[i] != text[i] {
			distance++
		}
	}

	if len(input) != len(text) {
		distance += max - min
	}
	return 100 - float64(distance)/float64(max)*100
}
