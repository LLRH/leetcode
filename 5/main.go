package main

import (
	"log"
)

func main() {
	res := longestPalindrome("zcabcb")
	log.Printf("res=%v", res)

	res = longestPalindrome("a")
	log.Printf("res=%v", res)

	res = longestPalindrome("aacdefcaa")
	log.Printf("res=%v", res)
}

func longestPalindrome(s string) string {
	sR := ""
	for _, v := range s {
		sR = string(v) + sR
	}
	l := len(s)
	mm := make([][]int, l, l)
	for idx := 0; idx < l; idx ++ {
		mm[idx] = make([]int, l, l)
	}

	maxIdx := 0
	maxLen := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == 0 || j == 0 {
				if s[i] == sR[j] {
					mm[i][j] = 1
					if mm[i][j] > maxLen {
						maxLen = mm[i][j]
						maxIdx = i
					}
				} else {
					mm[i][j] = 0
				}
			} else {
				if s[i] == sR[j] {
					mm[i][j] = mm[i-1][j-1] + 1
					if mm[i][j] > maxLen {
						if isPalindrome(getString(s, mm[i][j], i)) {
							maxLen = mm[i][j]
							maxIdx = i
						}
					}
				} else {
					mm[i][j] = 0
				}
			}
		}
	}

	//for i := 0; i < l; i++ {
	//	for j := 0; j < l; j++ {
	//		fmt.Printf("%v", mm[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	//log.Printf("maxIdx:%v maxLen:%v", maxIdx, maxLen)

	return getString(s, maxLen, maxIdx)
}

func getString(s string, maxLen int, maxIdx int) string {
	res := ""
	for i := 0; i < maxLen; i++ {
		res = res + string(s[maxIdx-i])
	}
	return res
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
