package main

func insertionSortList(head *ListNode) *ListNode {

	if head == nil{
		return head
	}
	res := &ListNode{
		Val:  11,
		Next: nil,
	}
	cur := res
	cur.Next = head
	head = head.Next
	cur.Next.Next = nil
	for head!= nil{
		cur = res
		for cur.Next != nil && head != nil{
			if head.Val <= cur.Next.Val{
				Next := head.Next
				head.Next = cur.Next
				cur.Next = head
				head = Next
				break
			}else{
				cur = cur.Next
			}
		}
		if cur.Next == nil{
			Next := head.Next
			cur.Next = head
			head.Next = nil
			head = Next
		}

	}

	return res.Next
}
func main() {
	
}
