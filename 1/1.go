package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		j, found := seen[val]
		if found {
			return []int{i, j}
		}
		seen[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 22))
}
