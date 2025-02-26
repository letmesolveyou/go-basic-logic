package logic_03

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal04(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				if i%2 == 0 {
					result[i][n-j-1] = start
					start += 2
				} else {
					result[i][j+(n-i-1)] = start
					start += 2
				}
			}
		}
	}
	return result
}
