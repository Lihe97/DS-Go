package main

import "fmt"

func postOrder(root *TreeNode){
	if root == nil{
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println(root.Val)
}

func postOrderr(root *TreeNode){

	stack1 := []*TreeNode{}
	stack2 := []*TreeNode{}

	stack1 = append(stack1,root)

	for len(stack1) != 0{
		temp := stack1[len(stack1) - 1]
		stack1 = stack1[:len(stack1) - 1]
		stack2 = append(stack2,temp)
		if temp.Left != nil{
			stack1 = append(stack1,temp.Left)
		}
		if temp.Right != nil{
			stack1 = append(stack1,temp.Right)
		}
	}
	for len(stack2) != 0{
		temp := stack2[len(stack2) - 1]
		fmt.Println(temp.Val)
		stack2 = stack2[:len(stack2) - 1]
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
	postOrderr(a)
}
