package main

import "fmt"

/*冒泡需要比较的值为nums[i],需要大值逐渐冒泡到最底层 保证最小的在最前面
[r1]1 2,2 3,3 4,4 5
[r2]1 2,2 3,3 4
[r3]1 2,2 3
[r4]1 2
可优化*/
func sort_bubble(nums []int) []int {
	n := len(nums)
	for i, _ := range nums {
		for j := 0; j < n-1; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	n := []int{1, 2, 6, 4}
	fmt.Println(sort_bubble(n))
}
