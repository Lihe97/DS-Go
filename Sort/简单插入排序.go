package main

import "fmt"

func insertArray(nums []int) []int {

	for i := 1; i < len(nums) ; i ++{
		temp := i

		for j := i - 1; j >= 0 && nums[temp] < nums[j]; j --{
			nums[temp],nums[j] = nums[j],nums[temp]
			temp = j

		}
	}
	return nums
}


func main() {
	fmt.Println(insertArray([]int{5,2,3,1}))
	
}
