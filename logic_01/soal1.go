package logic1

import "fmt"

func Soal1() {
	var n = 10
	var result [10]int

	for i := 1; i <= n; i++ {
		result[i-1] = 2*i - 1
	}

	fmt.Println(result)
}
