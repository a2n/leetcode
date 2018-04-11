package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func main() {
	s := []int{2, 7, 11, 15}
	t := 9
	fmt.Println(twoSum(s, t))
}

func twoSum(nums []int, target int) []int {
	for {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				if i == j {
					continue
				}
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
		return nil
	}
}
