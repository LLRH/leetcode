package main

import "log"

func main() {
	nums := []int{
		3, 2, 4,
	}
	res := twoSum(nums, 6)
	log.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if index2, ok := m[value]; ok {
			if 2*value == target {
				return []int{
					index, index2,
				}
			}
		}
		m[value] = index
	}
	for value, index := range m {
		otherValue := target - value
		log.Printf("index:%v value:%v otherValue:%v m:%v", index, value, otherValue, m)
		if index2, ok := m[otherValue]; ok {
			if index2 != index {
				return []int{
					index, index2,
				}
			}
		}
	}
	return nil
}
