package main

import (
	"fmt"
)

func sortArray(nums []int) []int {

	return quickSort(nums,0,len(nums)-1)

}
func quickSort(nums []int,start,end int)[]int{
	if start < end{
		partitionkey := partition(nums,start,end)
		quickSort(nums,start,partitionkey-1)
		quickSort(nums,partitionkey+1,end)
	}
	return nums
}
func partition(nums []int,start,end int)int{
	pivot := nums[start]
	for start < end{
		for start < end && nums[end] >= pivot{
			end --
		}
		nums[start],nums[end] = nums[end],nums[start]

		for start < end && nums[start] <= pivot{
			start++
		}
		nums[start],nums[end] = nums[end],nums[start]

	}
	return start

}
func main() {
	fmt.Println("?????????")
	a:= []int{45,23,23,67,123,4,98,17,76,54}
	//a:= []int{123,98,67,76,54}
	fmt.Println(sortArray(a))
	
}
