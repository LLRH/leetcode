package __Palindrome_Number

import (
	"testing"
	"leetcode.com/common"
)

func TestSolve(t *testing.T) {
	common.EXPECT(isPalindrome(-5), false)
	common.EXPECT(isPalindrome(121), true)
	common.EXPECT(isPalindrome(122), false)
	common.EXPECT(isPalindrome(10), false)
	common.EXPECT(isPalindrome(1000021), false)
}
