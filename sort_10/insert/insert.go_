package main

import "fmt"

func sort_insert(nums []int) []int {
	var swapCount, handleCount int
	n := len(nums)
	for assumeSortedScope := 0; assumeSortedScope < n; assumeSortedScope++ {
		for j := 0; j < assumeSortedScope; j++ {
			if nums[assumeSortedScope] < nums[j] {
				nums[assumeSortedScope], nums[j] = nums[j], nums[assumeSortedScope]
				swapCount++
			}
			handleCount++
		}
	}
	fmt.Println("[sort_insert]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}

//类似排队，从队伍里面拉人，范围依次扩大(j<i),进行标杆比较,保持nums[j]最小
//第一个、第二个。。。依次确定
//此处num[j]才是挑出来的值，nums[i]是游标，故nums[i] < nums[j]
func sort_insert_p(nums []int) []int {
	var swapCount, handleCount int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				swapCount++
			}
			handleCount++
		}
	}
	fmt.Println("[sort_insert]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}
func main() {
	nums := []int{3, 1, 4}
	fmt.Println(sort_insert(nums))
}
