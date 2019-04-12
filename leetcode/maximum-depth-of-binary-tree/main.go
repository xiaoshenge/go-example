package main

import (
	"fmt"
	"math"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

type StackNode struct{
	Depth int
	Tree *TreeNode
}

func maxDepth(root *TreeNode) int {
	stack := make([]StackNode,0)
	if root != nil {
		stack = append(stack, StackNode{1, root})
	}

	depth := 0
	for len(stack) > 0 {
		v := stack[0]
		stack = stack[1:]
		curDepth := v.Depth
		node := v.Tree
		if node != nil {
			depth = int(math.Max(float64(depth), float64(curDepth)))
			stack = append(stack, StackNode{curDepth+1, node.Left})
			stack = append(stack, StackNode{curDepth+1, node.Right})
		}
	}
	return depth
}

func main()  {
	a := []int{1}
	// for  _,v := range a{
	// 	fmt.Println(v)
		
	// }
	// fmt.Println(a)
	for len(a) >0 {
		v := a[0]
		fmt.Println(v)
		a = a[1:]
		a = append(a, 2)
	}
}



func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(maxDepthByRecursion(root.Left)) - float64(maxDepthByRecursion(root.Right))) > 1 {
		return false
	} else {
		if (isBalanced(root.Left) && isBalanced(root.Right)){
			return true
		} else {
			return false
		}
	}
}
func maxDepthByRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return int(math.Max(float64(maxDepthByRecursion(root.Left)), float64(maxDepthByRecursion(root.Right)))) + 1
	}
}