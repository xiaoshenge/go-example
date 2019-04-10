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
	for cur.next != nil {
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

func (list *LinkedList)Revert() {
	if list.head == nil || list.head.next == nil || list.head.next.next == nil {
		return 
	}
	
	var pre *ListNode
	cur := list.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	list.head.next = pre
}

func MergeSorted(l1,l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	l := &LinkedList{head:&ListNode{}}
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	if cur1 != nil && cur2 != nil{
		if cur1.Value().(int) > cur2.Value().(int){
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil{
		cur.next = cur2
	}
	return l
}

func (list *LinkedList)DeleteBottomN(n int)  {
	if n < 0 || list.head == nil || list.head.next == nil {
		return
	}
	fast := list.head
	for i := 0; i <= n; i++ {
		fast = fast.next
	}

	if fast == nil{
		return
	}

	slow := list.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

func (list *LinkedList) FindMiddleNode() *ListNode {
	if list.head == nil || list.head.next == nil {
		return nil
	}
	if list.head.next.next == nil {
		return list.head.next
	}
	slow, fast := list.head, list.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}