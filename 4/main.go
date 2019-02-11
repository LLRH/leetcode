package main

import (
	"log"
	"sort"
)

func main() {

	num1 := []int{
		1, 3, 5, 7,
	}
	num2 := []int{
		2, 4, 6, 8,
	}
	num3 := []int{
		1, 9, 11, 13, 100,
	}

	{
		res := processOne(num1, 4)
		if res != 4 {
			log.Fatalf("Expected res == 4 but Got %v", res)
		}
	}
	{
		res := processOne(num1, 0)
		if res != 3 {
			log.Fatalf("Expected res == 3 but Got %v", res)
		}
	}
	{
		res := processOne(num1, 44)
		if res != 5 {
			log.Fatalf("Expected res == 5 but Got %v", res)
		}
	}

	{
		res := processOne(num3, 0)
		if res != 9 {
			log.Fatalf("Expected res == 9.Palindrome Number but Got %v", res)
		}
	}
	{
		res := processOne(num3, 11)
		if res != 11 {
			log.Fatalf("Expected res == 11 but Got %v", res)
		}
	}
	{
		res := processOne(num3, 10)
		if res != 10 {
			log.Fatalf("Expected res == 10 but Got %v", res)
		}
	}
	{
		res := processOne(nil, 0)
		if res != 0 {
			log.Fatalf("Expected res == 0 but Got %v", res)
		}
	}
	{
		log.Println(findMedianSortedArrays(num1, num2))
		log.Println(findMedianSortedArraysTest(num1, num2))
	}
	{
		A := []int{
			1, 2,
		}
		B := []int{
			-1, 3,
		}
		log.Println(findMedianSortedArrays(A, B))
		log.Println(findMedianSortedArraysTest(A, B))
	}
	{
		A := []int{
			1,
		}
		B := []int{
			2, 3, 4,
		}
		log.Println(findMedianSortedArrays(A, B))
		log.Println(findMedianSortedArraysTest(A, B))
	}

}

func findMedianSortedArraysTest(nums1 []int, nums2 []int) float64 {

	num := make([]int, 0, 0)
	num = append(num, nums1...)
	num = append(num, nums2...)
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	lenNum := len(num)

	if lenNum == 0 {
		return -1
	}

	if lenNum%2 == 1 {
		return float64(num[lenNum/2])
	} else {
		return (float64(num[lenNum/2]) + float64(num[lenNum/2-1])) / 2
	}

	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenAll := len(nums1) + len(nums2)
	if lenAll%2 != 0 {
		return findMedianSortedArraysLeft(nums1, nums2)
	} else {
		l := findMedianSortedArraysLeft(nums1, nums2)
		r := findMedianSortedArraysRight(nums1, nums2)
		return (l + r) / 2.0
	}

}

func findMedianSortedArraysLeft(nums1 []int, nums2 []int) float64 {

	mid1 := getMid(nums1)
	mid2 := getMid(nums2)
	log.Printf("Left: %v(%v) %v(%v)", nums1, mid1, nums2, mid2)

	if len(nums1) == 0 {
		return float64(mid2)
	}
	if len(nums2) == 0 {
		return float64(mid1)
	}

	if mid1 == mid2 {
		return float64(mid1)
	}

	s1Len := len(nums1)
	s2Len := len(nums2)
	minLen := s1Len
	if s2Len < minLen {
		minLen = s2Len
	}
	cutLen := minLen / 2

	if s1Len == 1 {
		return float64(processOne(nums2, nums1[0]))
	} else if s2Len == 1 {
		return float64(processOne(nums1, nums2[0]))
	} else {
		if mid1 < mid2 {
			nums1 = nums1[cutLen:]
			nums2 = nums2[:s2Len-cutLen]
		} else {
			nums1 = nums1[:s1Len-cutLen]
			nums2 = nums2[cutLen:]
		}
		return findMedianSortedArraysLeft(nums1, nums2)
	}
}

func getMid(num []int) int {
	lenNum := len(num)
	if lenNum == 0 {
		return 0
	}
	if lenNum%2 == 0 {
		return num[lenNum/2-1]
	} else {
		return num[lenNum/2]
	}
}

//up mid
func processOne(num []int, v int) int {
	lenNum := len(num)
	if lenNum == 0 {
		return v
	} else if lenNum == 1 {
		if v < num[0] {
			return v
		} else {
			return num[0]
		}
	} else {
		leftIndex := lenNum/2 - 1
		rightIndex := lenNum / 2
		if v < num[leftIndex] {
			return num[leftIndex]
		} else if v > num[rightIndex] {
			return num[rightIndex]
		} else {
			return v
		}
	}
}

func findMedianSortedArraysRight(nums1 []int, nums2 []int) float64 {

	mid1 := getMidRight(nums1)
	mid2 := getMidRight(nums2)
	log.Printf("Right: %v(%v) %v(%v)", nums1, mid1, nums2, mid2)

	if len(nums1) == 0 {
		return float64(mid2)
	}
	if len(nums2) == 0 {
		return float64(mid1)
	}

	if mid1 == mid2 {
		return float64(mid1)
	}

	s1Len := len(nums1)
	s2Len := len(nums2)
	minLen := s1Len
	if s2Len < minLen {
		minLen = s2Len
	}
	cutLen := minLen / 2

	if s1Len == 1 {
		return float64(processOneRight(nums2, nums1[0]))
	} else if s2Len == 1 {
		return float64(processOneRight(nums1, nums2[0]))
	} else {
		if mid1 < mid2 {
			nums1 = nums1[cutLen:]
			nums2 = nums2[:s2Len-cutLen]
		} else {
			nums1 = nums1[:s1Len-cutLen]
			nums2 = nums2[cutLen:]
		}
		return findMedianSortedArraysRight(nums1, nums2)
	}
}

func getMidRight(num []int) int {
	lenNum := len(num)
	if lenNum == 0 {
		return 0
	}
	return num[lenNum/2]
}

//up mid
func processOneRight(num []int, v int) int {
	lenNum := len(num)
	if lenNum == 0 {
		return v
	} else if lenNum == 1 {
		if v > num[0] {
			return v
		} else {
			return num[0]
		}
	} else if lenNum%2 == 0 {
		leftIndex := lenNum/2 - 1
		rightIndex := lenNum / 2
		if v < num[leftIndex] {
			return num[leftIndex]
		} else if v > num[rightIndex] {
			return num[rightIndex]
		} else {
			return v
		}
	} else {
		leftIndex := lenNum / 2
		rightIndex := lenNum/2 + 1
		if v < num[leftIndex] {
			return num[leftIndex]
		} else if v > num[rightIndex] {
			return num[rightIndex]
		} else {
			return v
		}
	}
}
