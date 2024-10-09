package main

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	numMap "= make(map[iont]int)
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complemet]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	result := twoSum(nums, target)_
	fmt.Println("Input: nums = %v, target = %d\n", nums, target)
	fmt.Println("Output: %v\n, result)
}

func TestTwoSum(t *testing.T){
	testCases := []struct {
		nums []int
		target int
		want []int
	}{
		{[]int{2,7,11,15}, 9, []int{0,1}}},
		{[]int{3,2,4}, 6, []int{1,2}}},
		{[]int{2,7,4,15}, 11, []int{1,2}}},
	}

	for _, tc:= range testCases {
		got := twoSum(tc.nums, tc.target)
		if !equalSlices(got, tc.want){
			t.Errorf("twoSum(%v,%d) = %v, want %v", tc.nums, tc.target, got, tc.want)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i:= range a {
		if a [i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2,7,11,15}
	target := 9
	for i:=0; i< b.N; i++ {
		twoSum(nums, target)
	}
}
