package main

func twoSum(nums []int, target int) []int {
	values := make(map[int]int)
	for index, value := range nums {
		if number, exists := values[target-value]; exists {
			return []int{number, index}
		}
		values[value] = index
	}
	return nums
}
