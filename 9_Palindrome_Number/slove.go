package __Palindrome_Number

func isPalindrome(x int) bool {
	y, z := 0, x
	for x > 0 {
		y = 10*y + x%10
		x = x / 10
	}
	if y == z {
		return true
	} else {
		return false
	}
}
