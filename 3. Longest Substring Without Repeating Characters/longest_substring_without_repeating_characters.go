package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max_length := 0
	for left_cursor := 0; left_cursor < len(s); left_cursor++ {
		tmpMap := make(map[byte]int)
		for right_cursor := 0; right_cursor < len(s); right_cursor++ {
			if right_cursor >= left_cursor {
				// Note, if the substring without repeating chars is end with tail,
				// then we need take care of that
				if _, ok := tmpMap[s[right_cursor]]; ok || right_cursor == len(s) {
					break
				}
				// in fact, here we can store anything in tmpMap, we don't care the value
				tmpMap[s[right_cursor]] = right_cursor
			}
		}
		// fmt.Printf("length %d, map is %+v \n", len(tmpMap), tmpMap)
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
