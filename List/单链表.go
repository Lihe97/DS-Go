package main

import "fmt"

//定义结点
type ListNode struct {
	Val int
	Next *ListNode
}
type List struct {
	head *ListNode //头节点
}
//判空
func(list *List) IsEmpty()bool{
	if list.head == nil{
		return true
	}else{
		return false
	}
}
//求长度
func(list *List) GetLength()int{
	res := 0
	temp := list.head
	for temp!=nil{
		res ++
		temp = temp.Next
	}
	return res
}
//尾部插入
func(list *List)PushBack(a int){
	node := &ListNode{
		Val:  a,
		Next: nil,
	}
	temp := list.head
	if temp == nil{
		list.head = node
	}else{
		for temp.Next != nil{
			temp = temp.Next
		}
		temp.Next = node
	}

}

//遍历
func (list *List)Range(){
	if list.IsEmpty(){
		fmt.Println("链表为空")
	}else{
		temp := list.head
		for temp != nil{
			fmt.Println(temp.Val)
			temp = temp.Next
		}
	}
}
//头插
func (list *List) Add(value int)  {
	node := &ListNode{Val:value}
	node.Next = list.head
	list.head = node
}
//在i处插入
func(list *List)Insert(x,j int){
	if x > list.GetLength()+1{
		fmt.Println("超出最大长度,插入失败")
	}else{
	temp := list.head
	node := &ListNode{Val:j}
	for i :=1 ; i <x-1; i ++{
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = node
	}
}
//删除第i个元素
func(list *List)Remove(x int){
	if x > list.GetLength(){
		fmt.Println("超出最大长度,删除失败")
	}else{
		if x == 1{
			list.head = list.head.Next

		}else {
			temp := list.head
			for i := 1; i < x-1; i++ {
				temp = temp.Next
			}
			temp.Next = temp.Next.Next
		}
	}
}
//删除值为i的元素
func(list *List)Delete(x int){
	temp := list.head
	if temp.Val == x{
		list.head = temp.Next
	}else {
	for temp.Next != nil{
		if temp.Next.Val == x{
			temp.Next = temp.Next.Next
		}else {
			temp = temp.Next
		}
	}
	}
}
//单链表反转
func(list *List)Reverse(){
	cur := list.head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	list.head = pre

}
//func (list *List) Add(value int)  {
//	node := &ListNode{Val:value}
//	node.Next = list.head
//	list.head = node
//}


func main(){
	a := List{}
	a.Add(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)
	//a.Insert(6,5)
	//a.PushBack(4)
	//a.Delete(1)
	 a.Reverse()
	a.Range()
	//fmt.Println("aaa",b.GetLength())





}