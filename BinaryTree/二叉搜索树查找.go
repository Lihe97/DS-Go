package main



func searchBST(root *TreeNode,val int) *TreeNode{

	for root != nil{
		if root.Val == val{
			return root
		}
		if root.Val < val{
			root = root.Right
		}else if root.Val > val{
			root = root.Left
		}
	}
	return nil
}
func main() {
	
}
