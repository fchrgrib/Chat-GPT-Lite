package Algorithm

func SufPref(input string) []int {
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
