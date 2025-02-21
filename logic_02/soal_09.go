package logic_02

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal09(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || j+i == n-1 {
				result[i][j] = start
				start += 2
			}
		}
	}
	return result
}
