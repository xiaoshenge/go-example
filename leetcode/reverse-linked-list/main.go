package main

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil{
		return head
	}
	cur := head.Next
	pre := &ListNode{}
	for cur != nil{
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return cur
}