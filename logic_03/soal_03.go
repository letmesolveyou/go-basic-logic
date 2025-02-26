package logic_03

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal03(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				if i == 0 {
					result[i][j-i] = start
					start += 3
				} else if i%2 == 0 {
					result[i][j-i] = start
					start += 2
					if j == n-1 {
						start += 1
					}
				} else {
					result[i][n-j-1] = start
					start += 3
					if j == n-1 {
						start -= 1
					}
				}
			}
		}
	}
	result[0][0] = 2
	return result
}
