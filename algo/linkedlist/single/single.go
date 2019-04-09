package single
import (
	"fmt"
)
/*
在某个节点后面插入节点
在某个节点前面插入节点
在链表头部插入节点
链表尾部插入节点
通过索引查找节点
删除传入的节点
打印链表
*/

type ListNode struct{
	next *ListNode
	value interface{}
}

type LinkedList struct{
	head *ListNode
	len uint
}
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}
func (node *ListNode) Value() interface{}{
	return node.value
}
func (node *ListNode) Next() *ListNode {
	return node.next
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0),0}
}

func (list *LinkedList)InsertAfter(node *ListNode, v interface{}) bool {
	if node == nil {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = node.next
	node.next = newNode
	list.len++
	return true
}
func (list *LinkedList)InsertBefore(node *ListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	cur := list.head.next
	pre := list.head
	for cur != nil{
		if cur == node{
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = cur
	pre.next=newNode
	list.len++
	return true
}
func (list *LinkedList)InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.head,v)
}
func (list *LinkedList)InsertToTail(v interface{}) bool {
	cur := list.head
	if cur != nil {
		cur = cur.next
	}
	return list.InsertAfter(cur, v)
}
func (list *LinkedList)FindIndex(index uint) *ListNode{
	if index >= list.len{
		return nil
	}
	cur := list.head.next
	var i uint
	for ; i <index;i++{
		cur = cur.next		
	}
	return cur
}
func (list *LinkedList)Delete(node *ListNode) bool{
	if node == nil {
		return false
	}

	cur := list.head.next
	pre := list.head
	for cur != nil{
		if cur == node{
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil{
		return false
	}
	pre.next = cur.next
	node = nil
	list.len--
	return true
}
func (list *LinkedList)Print(){
	cur := list.head
	format := ""
	for cur != nil{
		format += fmt.Sprintf("%+v", cur.Value())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}