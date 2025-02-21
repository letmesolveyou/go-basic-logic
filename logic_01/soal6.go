package logic1

import "fmt"

func Soal6() {
	var n = 10
	var result [10]int

	for i := 1; i <= n; i++ {
		result[i-1] = 33 - 3*i
	}

	fmt.Println(result)
}
