package logic_03

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal05(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	mid := n / 2
	for i := 0; i <= mid; i++ {
		for j := 0; j <= i; j++ {
			if i >= j {
				if i%2 == 0 {
					result[i][j] = start
					result[i][n-1-j] = start
					result[n-i-1][j] = start
					result[n-i-1][n-1-j] = start
				} else {
					result[i][i-j] = start
					result[i][j+(n-i-1)] = start
					result[n-1-i][i-j] = start
					result[n-1-i][j+(n-i-1)] = start
				}
			}
			start += 2
		}
	}
	return result
}
