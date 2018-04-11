package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output:  321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range.
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
	check := func(x int) int {
		if x > math.MaxInt32 || x < -math.MaxInt32 {
			return 0
		}
		return x
	}

	if check(x) == 0 {
		return 0
	}

	y := math.Abs(float64(x))
	s := strconv.FormatInt(int64(y), 10)
	s1 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s1 += string(s[i])
	}

	n, e := strconv.ParseInt(s1, 10, 32)
	if e != nil {
		return 0
	}

	if x < 0 {
		return -int(n)
	}
	return int(n)
}
