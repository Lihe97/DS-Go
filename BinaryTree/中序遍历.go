package main

import (
	"fmt"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//递归
func inOrder(t *TreeNode){
	if t == nil{
		return
	}
	inOrder(t.Left)
	fmt.Println(t.Val)
	inOrder(t.Right)
}
//非递归
func inOrderr(t *TreeNode){

	stack := []*TreeNode{}


	for len(stack) != 0 || t != nil{

		for t != nil{
			stack = append(stack,t)
			t = t.Left
		}
		t = stack[len(stack) - 1]
		fmt.Println(t.Val)
		stack = stack[:len(stack) - 1]

		t = t.Right

	}
}

func main() {
	g := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	f := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	e := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	d := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	c := &TreeNode{
		Val:   3,
		Left:  f,
		Right: g,
	}
	b := &TreeNode{
		Val:   2,
		Left:  d,
		Right: e,
	}

	a := &TreeNode{
		Val:   1,
		Left:  b,
		Right: c,
	}
	inOrderr(a)
}
