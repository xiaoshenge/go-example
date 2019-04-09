package single

func ExampleInsertToHead() {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
	// Output:
	// 0->10->9->8->7->6->5->4->3->2->1
}
