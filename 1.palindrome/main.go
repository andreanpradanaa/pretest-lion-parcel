package main

import "fmt"

/*
	Buatlah sebuah fungsi yang mengembalikan True jika sebuah string adalah palindrom, dan
	False jika tidak.

	Input: "radar"
	Output: True

	Input: "hello"
	Output: False
*/

func isPalindrome(word string) bool {

	// using two pointer to compare character
	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	s1 := "radar"
	res1 := isPalindrome(s1)
	fmt.Println(res1)

	s2 := "hello"
	res2 := isPalindrome(s2)
	fmt.Println(res2)
}
