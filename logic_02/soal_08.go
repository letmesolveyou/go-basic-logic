package logic_02

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal08(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if n-j-i-1 == 0 {
				result[i][j] = start
				start += 2
			}
		}
	}
	return result
}
