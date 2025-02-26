package logic_03

import go_print_slice "github.com/letmesolveyou/go-print-slice"

func Soal06(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1
	mid := n / 2

	var validIdx bool
	for i := 0; i < mid+1; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				validIdx = isValididx(i, j, n)
				if validIdx {
					result[i][j] = start
					result[n-1-i][j] = start
					start += 2
				}
			} else {
				validIdx = isValididx(i, j, n)
				if validIdx {
					result[i][n-j-1] = start
					result[n-1-i][n-j-1] = start
					start += 2
				}
			}
		}
	}
	return result
}

func isValididx(i int, j int, n int) bool {
	// (0,0) -> true
	// (0, 1) -> true
	// (1, 1) -> true
	// (1, 0) -> false
	// (1 , 1) -> true
	//...
	// (1, 7) -> true
	// (1, 8) -> false
	if i > j || j > n-1-i {
		return false
	} else {
		return true
	}
}
