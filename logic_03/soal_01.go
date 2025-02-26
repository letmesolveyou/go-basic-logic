package logic_03

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal01(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				if i%2 == 0 {
					result[i][j] = start
					start += 2
				} else {
					result[i][i-j] = start
					start += 2
				}
			}
		}
	}
	return result
}
