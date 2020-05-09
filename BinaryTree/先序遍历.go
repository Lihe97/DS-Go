package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//递归遍历
func preOrder(t *TreeNode){
	if t == nil{
		return
	}
	fmt.Println(t.Val)
	preOrder(t.Left)
	preOrder(t.Right)
}
//非递归遍历
func preOrderr(t *TreeNode){

	stack := []*TreeNode{}

	stack = append(stack,t)
	for len(stack) != 0{

		temp := stack[len(stack)-1]
		fmt.Println(temp.Val)
		stack = stack[:len(stack)-1]
		if temp.Right != nil{
			stack = append(stack,temp.Right)
		}
		if temp.Left != nil{
			stack = append(stack,temp.Left)
		}

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

	preOrderr(a)
	
}
