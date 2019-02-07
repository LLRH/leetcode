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
	log.Println(findMedianSortedArraysTest(num1, num2))
	log.Println(findMedianSortedArraysTest(num1, num3))

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
		if res != 10 {
			log.Fatalf("Expected res == 10 but Got %v", res)
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
		if res != 10.5 {
			log.Fatalf("Expected res == 10.5 but Got %v", res)
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

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mid1 := getMid(nums1)
	mid2 := getMid(nums2)
	log.Printf("%v(%v) %v(%v)", nums1, mid1, nums2, mid2)

	if len(nums1) == 0 {
		return mid2
	}
	if len(nums2) == 0 {
		return mid1
	}

	if mid1 == mid2 {
		return mid1
	}

	s1Len := len(nums1)
	s2Len := len(nums2)
	minLen := s1Len
	if s2Len < minLen {
		minLen = s2Len
	}
	cutLen := minLen / 2

	if s1Len == 1 {
		return processOne(nums2, nums1[0])
	} else if s2Len == 1 {
		return processOne(nums1, nums2[0])
	} else {
		if mid1 < mid2 {
			nums1 = nums1[cutLen:]
			nums2 = nums2[:s2Len-cutLen]
		} else {
			nums1 = nums1[:s1Len-cutLen]
			nums2 = nums2[cutLen:]
		}
		return findMedianSortedArrays(nums1, nums2)
	}
}

func getMid(num []int) float64 {
	lenNum := len(num)
	if lenNum == 0 {
		return 0
	}
	if lenNum%2 == 0 {
		return (float64(num[lenNum/2]) + float64(num[lenNum/2-1])) / 2
	} else {
		return float64(num[lenNum/2])
	}
}

func processOne(num []int, v int) float64 {
	lenNum := len(num)
	if lenNum == 0 {
		return float64(v)
	} else if lenNum == 1 {
		return (float64(num[0]) + float64(v)) / 2
	} else if lenNum%2 == 0 {
		leftMid := num[lenNum/2-1]
		rightMid := num[lenNum/2]
		/*log.Printf("leftMid:%v, rightMid:%v, v:%v",
			leftMid, rightMid, v)*/
		if v < leftMid {
			return float64(leftMid)
		} else if v > rightMid {
			return float64(rightMid)
		} else {
			return float64(v)
		}
	} else {

		p1, p2 := 0, 0
		L := num[lenNum/2-1]
		M := num[lenNum/2]
		R := num[lenNum/2+1]

		/*log.Printf("L:%v, M:%v, R:%v, v:%v",
			L, M, R, v)*/

		if v < L {
			p1 = L
			p2 = M
		} else if v > R {
			p1 = M
			p2 = R
		} else {
			p1 = M
			p2 = v
		}
		return (float64(p1) + float64(p2)) / 2
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
