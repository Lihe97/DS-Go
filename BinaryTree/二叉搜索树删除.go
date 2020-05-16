package main

func deleteBST(root *TreeNode,val int) *TreeNode {

	res := root

	father := &TreeNode{}
	father = nil
	temp := root
	for temp != nil && temp.Val != val{
		father = temp
		if val > temp.Val{
			temp = temp.Right
		}else{
			temp = temp.Left
		}
	}
	if temp == nil{
		return res
	}

	if temp.Left != nil && temp.Right != nil{

		min := temp.Right
		minP := temp

		for min.Left != nil{
			minP = min
			min = min.Left
		}
		temp.Val = min.Val
		father = minP
		temp = min
	}
	child := &TreeNode{}

	if temp.Left != nil{
		child = temp.Left
	}else if temp.Right != nil{
		child = temp.Right
	}else{
		child = nil
	}

	if father == nil{
		res = child
	}else if father.Left == temp{
		father.Left = child
	}else if father.Right == temp{
		father.Right = child
	}

	return res
}

func main() {
	
}
