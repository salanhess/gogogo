package main

import "fmt"

//类似排队，从队伍里面拉人，进行标杆比较。
//第一个、第二个。。。依次确定
//此处num[j]才是挑出来的值，nums[i]是游标，故nums[i] < nums[j]
func sort_insert(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
func main() {
	nums := []int{3, 1, 4}
	fmt.Println(sort_insert(nums))
}
