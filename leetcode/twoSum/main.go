package main

import "fmt"

func twoSum(nums []int, target int) []int {
	leftPointer, rightPointer := 0, 1
	result := []int{nums[leftPointer], nums[rightPointer]}
	if len(nums) == 2 {
		return result
	}
	for ; leftPointer < len(nums); leftPointer++ {
		for ; rightPointer < len(nums); rightPointer++ {
			if nums[leftPointer]+nums[rightPointer] == target {
				result = []int{leftPointer, rightPointer}
			}
		}
		rightPointer = leftPointer + 2
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 22))
}
