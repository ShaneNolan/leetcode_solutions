package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	answer := []int{0, 1}
	result := twoSum(nums, target)
	if answer[0] != result[0] || answer[1] != result[1] {
		t.Errorf("Invalid")
	}
}
