package main

import "fmt"

/*冒泡需要比较的值为nums[i],需要大值逐渐冒泡到最底层 保证最小的在最前面
num[j],num[j+1]进行比较,注意边界
[r1]1 2,2 3,3 4,4 5
[r2]1 2,2 3,3 4
[r3]1 2,2 3
[r4]1 2
可优化*/

//优化一仅仅适用于连片有序而整体无序的数据(例如：1， 2，3 ，4 ，7，6，5)。
//但是对于前面大部分是无序而后边小半部分有序的数据(1，2，5，7，4，3，6，8，9，10)排序效率也不可观
//我们可以继续优化。既我们可以记下最后一次交换的位置，后边没有交换，必然是有序的，
//然后下一次排序从第一个比较到上次记录的位置结束即可。
func sort_bubble3(nums []int) []int {
	n := len(nums)
	var swapCount, handleCount, swapIndex int
	k := n - 1
	swapIndex = 0
	for _ = range nums {
		swapFlag := false
		for j := 0; j < k; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapFlag = true
				swapIndex = j
				swapCount++
			}
			handleCount++
		}
		if !swapFlag {
			break
		}
		k = swapIndex
	}
	fmt.Println("[bubble3]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}

//优化1：使用flag标记本轮是否有冒泡，没有的话就退出
func sort_bubble2(nums []int) []int {
	n := len(nums)
	var swapCount, handleCount int
	for i, _ := range nums {
		swapFlag := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapCount++
				swapFlag = true
			}
			handleCount++
		}
		if !swapFlag {
			break
		}
	}
	fmt.Println("[bubble2]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}

func sort_bubble(nums []int) []int {
	n := len(nums)
	var swapCount, handleCount int
	for i, _ := range nums {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapCount++
			}
			handleCount++
		}
	}
	fmt.Println("[bubble]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}

func main() {
	n := []int{1, 2, 6, 4, 7, 8, 9}
	fmt.Println(sort_bubble(n))
	n = []int{1, 2, 6, 4, 7, 8, 9}
	fmt.Println(sort_bubble2(n))
	n = []int{1, 2, 6, 4, 7, 8, 9}
	fmt.Println(sort_bubble3(n))

	fmt.Println("====================")
	n = []int{1, 2, 3, 4, 7, 8, 9}
	fmt.Println(sort_bubble(n))
	n = []int{1, 2, 3, 4, 7, 8, 9}
	fmt.Println(sort_bubble2(n))
	n = []int{1, 2, 3, 4, 7, 8, 9}
	fmt.Println(sort_bubble3(n))
}
