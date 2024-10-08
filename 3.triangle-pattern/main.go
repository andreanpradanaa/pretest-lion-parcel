package main

import "fmt"

/*
	Buatlah sebuah program yang mencetak pola segitiga dengan menggunakan karakter '*'.
	
	Input: 5
	Output:
	*
	**
	***
	****
	*****
*/

func printTriangle(n int) {
	for row := 0; row < n; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	n := 5
	printTriangle(n)
}
