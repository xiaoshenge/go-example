package single

import (
	"fmt"
)


func ExampleInsertToHead() {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
	// Output:
	// 0->10->9->8->7->6->5->4->3->2->1
}
func ExampleInsertToTail()  {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i+1)
	}
	l.Print()
	// Output:
	// 0->1->2->3->4->5->6->7->8->9->10
}
func ExampleFindIdex()  {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i+1)
	}
	var i uint = 0
	fmt.Println(l.FindIndex(i).Value())
	// Output:
	// 1
}

func ExampleDeletNode()  {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i+1)
	}
	l.Delete(l.head.next)
	l.Print()
	// Output:
	// 0->2->3->4->5->6->7->8->9->10
}


func ExampleInsertBefore()  {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i+1)
	}
	l.InsertBefore(l.head.next, 1)
	l.Print()
	// Output:
	// 0->1->1->2->3->4->5->6->7->8->9->10
}

func ExampleRevert()  {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i+1)
	}
	l.Revert()
	l.Print()
	// Output:
	// 0->10->9->8->7->6->5->4->3->2->1
}