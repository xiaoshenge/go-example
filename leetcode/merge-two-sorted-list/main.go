package main

import (
	"fmt"
)

func main()  {
	a1 := [...]int{2,4,7}
	a2 := [...]int{3,5}

	l1 := &ListNode{1,nil}
	l2 := &ListNode{1,nil}
	cur := l1
	for _,i := range a1{
		tmp := &ListNode{}
		tmp.Val = i
		cur.Next = tmp
		cur = cur.Next
	}

	cur = l2
	for _,i := range a2{
		tmp := &ListNode{}
		tmp.Val = i
		cur.Next = tmp
		cur = cur.Next
	}
	
	

	l3 := MergeTwoLists(l1,l2)
	for l3 != nil{
		fmt.Println(l3)
		l3 = l3.Next
	}
}

type ListNode struct{
	Val int
	Next *ListNode
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil{
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}


	list := &ListNode{}
	if l1.Val <= l2.Val {
		list = l1
	} else {
		list = l2
	}

	cur := list
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if cur != l1 {
				cur.Next = l1
				cur = cur.Next
			}
			l1 = l1.Next
		} else {
			if cur != l2 {
				cur.Next = l2
				cur = cur.Next
			}
			l2 = l2.Next
		}
	}
	if l1 != nil{
		cur.Next = l1
	}
	if l2 != nil{
		cur.Next = l2
	}
	return list
}