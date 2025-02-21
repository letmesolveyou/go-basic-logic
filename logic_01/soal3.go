package logic1

import "fmt"

func Soal3() {
	var n = 10
	var result [10]int

	for i := 1; i <= n; i++ {
		result[i-1] = 3 * i
	}

	fmt.Println(result)
}
