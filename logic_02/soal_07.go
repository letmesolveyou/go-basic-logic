package logic_02

import (
	go_print_slice "github.com/letmesolveyou/go-print-slice"
)

func Soal07(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				result[i][j] = 2*start + 1
				start++
			}
		}
	}
	return result
}
