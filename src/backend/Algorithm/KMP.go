package Algorithm

func sufPref(input string) []int {
	i, j, m := 1, 0, len(input)

	bx := make([]int, m)

	for i < m {
		if input[j] == input[i] {
			bx[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = bx[j-1]
		} else {
			bx[i] = 0
			i++
		}
	}

	return bx
}

func KMP(input string, text string) int {
	n, m, i, j := len(text), len(input), 0, 0
	b := sufPref(input)

	for i < n {
		if input[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = b[j-1]
		} else {
			i++
		}
	}

	return -1
}
