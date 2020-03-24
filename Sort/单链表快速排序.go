package main


func swap(a,b *ListNode){
	a.Val , b.Val = b.Val,a.Val
}
func quicksort(head,end *ListNode){

	if head == end || head.Next == nil{
		return
	}
	slow , fast := head, head.Next
	pivot := head.Val
	for fast != nil{
		if fast.Val <= pivot{
			slow = slow.Next
			swap(slow,fast)
		}
		fast = fast.Next
	}
	swap(head,slow)
	quicksort(head,slow)
	quicksort(slow.Next,end)
}
func main() {
	g := & ListNode{
		Val:  27,
		Next: nil,
	}
	f := & ListNode{
		Val:  13,
		Next: g,
	}
	e := & ListNode{
		Val:  76,
		Next: f,
	}
	d := & ListNode{
		Val:  97,
		Next: e,
	}
	c := & ListNode{
		Val:  65,
		Next: d,
	}
	b := & ListNode{
		Val:  38,
		Next: c,
	}
	a := & ListNode{
		Val:  49,
		Next: b,
	}
	quicksort(a,g)

	
}
