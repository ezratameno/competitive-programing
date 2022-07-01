package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	roteated, pos := IsRoteated(nums)
	if roteated {
		return nums[pos]
	}
	return nums[0]
}

func IsRoteated(nums []int) (bool, int) {
	for i := range nums[:len(nums)-1] {
		if nums[i] > nums[i+1] {
			return true, i + 1
		}
	}
	return false, 0
}
