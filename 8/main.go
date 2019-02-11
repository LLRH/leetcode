package main

import "log"

//case1: "42" 42
//case2: "    -42" -42
//case3: "4193 with word" 4193
//case4: "words and 987" 0
//case5: "-91283472322" -2147483648
//[-2^31, 2^31-1]

func main() {
	{
		res := myAtoi("42")
		log.Printf("res:%v", res)
	}
	{
		res := myAtoi("-42")
		log.Printf("res:%v", res)
	}
	{
		res := myAtoi("4193 with word")
		log.Printf("res:%v", res)
	}
	{
		res := myAtoi("words and 987")
		log.Printf("res:%v", res)
	}
	{
		res := myAtoi("-91283472322")
		log.Printf("res:%v", res)
	}

	{
		res := myAtoi("9223372036854775808")
		log.Printf("res:%v", res)
	}
}

func myAtoi(str string) int {
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	neg := 1
	if len(str) > 0 {
		if str[0] == '+' {
			str = str[1:]
			neg = 1
		} else if str[0] == '-' {
			str = str[1:]
			neg = -1
		}
	}
	if len(str) == 0 {
		return 0
	}

	min := (-1) * (1 << 31)
	max := 1<<31 - 1

	numSet := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	res := 0
	for len(str) > 0 {
		if value, ok := numSet[str[0]]; ok {
			res = res*10 + value
			str = str[1:]

			if res*neg < min {
				return min
			}
			if res*neg > max {
				return max
			}
		} else {
			break
		}
	}

	log.Printf("res:%v, neg:%v", res, neg)

	return res * neg
}
