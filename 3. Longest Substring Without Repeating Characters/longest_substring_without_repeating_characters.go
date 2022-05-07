package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	tmpChar := []byte(s)
	fmt.Println(tmpChar)
	max_length := 0

	for left_cursor := 0; left_cursor < len(s); left_cursor++ {
		tmpMap := make(map[byte]int)
		for right_cursor := 0; right_cursor < len(s); right_cursor++ {
			if _, ok := tmpMap[s[right_cursor]]; ok || right_cursor == len(s)-1 {
				break
			}
			tmpMap[s[right_cursor]] = right_cursor
		}
		fmt.Printf("length %d, map is %+v \n", len(tmpMap), tmpMap)
		max_length = max(max_length, len(tmpMap))
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
	s := "asfdgdsfgsdffdgergsdfsadfwefdsacv"
	rest := lengthOfLongestSubstring(s)
	fmt.Println(rest)
}
