package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	return reverse(x) == x
}
func reverse(x int) int {
	y := math.Abs(float64(x))
	s := strconv.FormatInt(int64(y), 10)
	s1 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s1 += string(s[i])
	}

	n, e := strconv.ParseInt(s1, 10, 64)
	if e != nil {
		return 0
	}

	return int(n)
}
