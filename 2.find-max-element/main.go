package main

import "fmt"

/*
	Buatlah sebuah fungsi yang mengembalikan nilai maksimum dari sebuah array bilangan
	bulat tanpa menggunakan fungsi bawaan seperti `max()`.

	Input: [3, 5, 1, 9, 2]
	Output: 9
*/

func findMaxElement(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	nums := []int{3, 5, 1, 9, 2}
	res := findMaxElement(nums)
	fmt.Println(res)
}
