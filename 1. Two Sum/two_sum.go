package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, x := range nums {
		if value, ok := hashMap[target-x]; ok {
			return []int{value, i}
		}
		hashMap[x] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
