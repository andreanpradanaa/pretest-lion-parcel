package main

import "fmt"

/*
	Buatlah fungsi rekursif untuk menghitung faktorial dari suatu bilangan.

	Input: 5
	Output: 120 // karena 5! = 5 x 4 x 3 x 2 x 1 = 120
*/

func countFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * countFactorial(n-1)
}

func main() {
	n := 5
	res := countFactorial(n)
	fmt.Println(res)
}
