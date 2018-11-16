package main

import (
	"fmt"
	)

type Node struct {
	name string
	children []*Node
}

func (n *Node)InitChildren(name string)  {
	n.name = name
	child := &Node{}
	n.children = []*Node{child}
	n = child
	n.name = "child"
}

func InitChildren(n *Node, name string)  {
	n.name = name
	child := &Node{}
	n.children = []*Node{child}
	n = child
	n.name = "child---"
}

func main() {
	root := &Node{}
	root.InitChildren("root")
	fmt.Println(root)
	for _,v := range root.children{
		fmt.Println(v)
	}


	root1 := &Node{}
	InitChildren(root1, "root1")
	fmt.Println(root1)
	for _,v := range root1.children{
		fmt.Println(v)
	}
}
