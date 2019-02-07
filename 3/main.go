package main

import (
	"log"
)

func main() {
	log.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	totalLen := len(s)
	setMap := make(map[byte]bool)
	i, j, maxLen := 0, 0, 0
	for i < totalLen && j < totalLen {

		newC := s[j]
		oldC := s[i]
		if _, ok := setMap[newC]; ok {
			delete(setMap, oldC)
			i++
		} else {
			setMap[newC] = true
			j++
			if tempLen := len(setMap); tempLen > maxLen {
				maxLen = tempLen
			}
		}
	}
	return maxLen
}
