package logic1

import "fmt"

func Soal4() {
	var n = 10
	var result [10]int

	for i := 1; i <= n; i++ {
		result[i-1] = 21 - 2*i
	}

	fmt.Println(result)
}
