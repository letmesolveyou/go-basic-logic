package main

import (
	"github.com/letmesolveyou/go-basic-logic/logic_02"
	go_print_slice "github.com/letmesolveyou/go-print-slice"
)

func main() {
	matrix := logic_02.Soal09(9)
	go_print_slice.PrintSlice2D(matrix)

	//slice := []int{1, 2, 3}
	//go_print_slice.PrintSlice(slice)
	//
	//fmt.Println()
	//fmt.Println()
	//
	//slice2 := [][]int{{1, 2, 3}, {1, 2, 3}, {7, 8, 9}}
	//go_print_slice.PrintSlice2D(slice2)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 1
	//n := 9
	//slice3 := make([][]int, n)
	//
	//for i := 0; i < n; i++ {
	//	slice3[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		slice3[i][j] = 1 + 2*j
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice3)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 2
	//slice4 := make([][]int, n)
	//
	//for i := 0; i < n; i++ {
	//	slice4[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		slice4[i][j] = 2 + 2*j
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice4)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 3
	//slice5 := make([][]int, n)
	//num5 := 1
	//for i := 0; i < n; i++ {
	//	slice5[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		slice5[i][j] = num5
	//		num5 += 2
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice5)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 4
	//slice6 := make([][]int, n)
	//num6 := 1
	//
	//for i := 0; i < n; i++ {
	//	slice6[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		slice6[i][j] = num6
	//		num6 += 3
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice6)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 5
	//slice7 := make([][]int, n)
	//num7 := 1
	//
	//for i := 0; i < n; i++ {
	//	slice7[i] = make([]int, n)
	//	if i%2 == 0 {
	//		for j := 0; j < n; j++ {
	//			slice7[i][j] = num7
	//			num7 += 2
	//		}
	//	} else {
	//		for j := n - 1; j >= 0; j-- {
	//			slice7[i][j] = num7
	//			num7 += 2
	//		}
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice7)
	//
	//fmt.Println()
	//
	////Logic 2-Soal 6
	//slice8 := make([][]int, n)
	//num8 := 1
	//
	//for i := 0; i < n; i++ {
	//	slice8[i] = make([]int, n)
	//	if i%2 == 0 {
	//		for j := 0; j < n; j++ {
	//			slice8[i][j] = num8
	//			num8 += 3
	//		}
	//	} else {
	//		for j := n - 1; j >= 0; j-- {
	//			slice8[i][j] = num8
	//			num8 += 3
	//		}
	//	}
	//}
	//
	//go_print_slice.PrintSlice2D(slice8)
}
