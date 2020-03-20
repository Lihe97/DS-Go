package main

import "fmt"

type Stack struct {
	values    []int
	top int
}
//构造栈
func Constructor(maxSize int) Stack {
	if maxSize > 0{
		values := make([]int,maxSize)
		return Stack{
			values: values,
			top:    -1,
		}
	}else{
		panic("数组")
	}
}
//入栈
func (this *Stack) Push(x int)  {
	this.top++
	this.values[this.top] = x
}
//出栈
func (this *Stack) Pop() int {
	res := this.values[this.top]
	this.top --
	return res
}
//判空
func (this *Stack) isEmpty()bool{

	return this.top == -1
}
//获取栈顶元素
func (this *Stack) getTop() int{
	return this.values[this.top]
}
//遍历栈
func(this *Stack)Range(){
	if this.isEmpty(){
		fmt.Println("栈空")
		return
	}
	for i := this.top; i>= 0 ; i--{
		fmt.Println("this.values[%d]:=%d",i,this.values[i])
	}
}




