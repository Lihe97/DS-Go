package main

import "fmt"

func MergeSort(nums []int) []int {
	if len(nums) < 2{
		return nums
	}

	i := len(nums) / 2

	left := MergeSort(nums[:i])
	right := MergeSort(nums[i:])

	return merge(left,right)
}
func merge(left,right []int) []int{

	m,n := 0,0
	res := []int{}

	for m < len(left) && n < len(right){
		if left[m] > right[n]{
			res = append(res, right[n])
			n++
		}else{
			res = append(res, left[m])
			m++
		}
	}
	res = append(res, right[n:]...)
	res = append(res, left[m:]...)
	return res
}

func main() {

	a := []int{3,5,7,2,3,3,2,6}
	fmt.Println(MergeSort(a))
	
}
