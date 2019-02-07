package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func main() {
	//243
	l1_2 := NewListNode(2)
	l1_4 := NewListNode(4)
	l1_3 := NewListNode(3)

	l1_3.Next = l1_4
	l1_4.Next = l1_2

	res := addTwoNumbers(l1_3, l1_3)
	for res != nil {
		log.Println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var pre *ListNode = nil
	var res *ListNode = nil
	add := 0
	for l1 != nil || l2 != nil || add != 0 {
		num1 := 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		num2 := 0
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		num := num1 + num2 + add

		add = num / 10
		num = num % 10

		newNode := &ListNode{
			Val:  num,
			Next: nil,
		}

		if pre != nil {
			pre.Next = newNode
		} else {
			res = newNode
		}
		pre = newNode
	}

	return res
}
