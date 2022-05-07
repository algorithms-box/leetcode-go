package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	start_cursor, max_length := -1, 0
	// use map keys cannot be same
	charsMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := charsMap[s[i]]; ok && charsMap[s[i]] > start_cursor {
			// print the string without repeating chars for now
			// fmt.Println(charsMap)
			// let start_cursor move to char ch last one position before this loop
			start_cursor = charsMap[s[i]]
			charsMap[s[i]] = i
		} else {
			charsMap[s[i]] = i
			max_length = max(max_length, i-start_cursor)
		}
	}

	return max_length

}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	s := "dfasdfwefwfe"
	rest := lengthOfLongestSubstring(s)
	fmt.Println(rest)
}
