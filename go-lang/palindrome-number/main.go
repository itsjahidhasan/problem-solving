package main

import "log"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r := 0
	for x > r {
		r = r*10 + x%10
		x /= 10
	}

	return r == x
}

func main() {
	log.Println(isPalindrome(232))
}
