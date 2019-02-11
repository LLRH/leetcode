package __Palindrome_Number

import (
	"testing"
	"leetcode.com/common"
)

func TestSolve(t *testing.T) {
	{
		res := isPalindrome(-5)
		common.EXPECT(res, false)
	}
}
