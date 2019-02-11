package __Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		base := 1
		for temp := x; temp >= 10; {
			base = base * 10
			temp = temp / 10
		}
		first := x / base
		last := x % 10
		if first != last {
			return false
		} else {
			x = (x - first*base) / 10
			return isPalindrome(x)
		}
	}
}
