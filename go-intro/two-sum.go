package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	counts := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		remainder := target - nums[i]
		if v, ok := counts[remainder]; ok {
			result[0] = i
			result[1] = v
			return result
		}

		counts[nums[i]] = i
	}

	return make([]int, 0)
}

func main() {
	nums := []int{2,7,11,15}
	target := 9

	result := twoSum(nums, target)

	fmt.Print("Answer: ", result)
}