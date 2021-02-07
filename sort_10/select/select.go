package main

import (
	"fmt"
	common_util "gogogo/sort_10/common"
)

/*windows使用git时出现：warning: LF will be replaced by CRLF
以及使windows10支持区分大写小的问题处理
解决办法（二）

找到windows项目的.git目录,修改config文件，在 [core] 配置项添加一行即可。
autocrlf = false*/

//选择排序的要点是游标为minIndex，下一个值不断的和该游标进行比较，将最小值放到minIndex处
//==>和冒泡排序的区别？
//冒泡是逐次交换num[j],num[j+1]
//==>与插入排序的区别？
//插入是保持队伍最前面的有序（常用于排队）

func sort_select(nums []int) []int {
	n := len(nums)
	var swapCount, handleCount int
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[minIndex] > nums[j] {
				//nums[minIndex] , nums[j] = nums[j] , nums[minIndex]
				common_util.SwapInt(&nums[j], &nums[minIndex])
				swapCount++
			}
			handleCount++
		}
	}
	fmt.Println("[sort_select]handleCount:", handleCount, " swapCount:", swapCount)
	return nums
}

func main() {
	n := []int{1, 2, 6, 4}
	fmt.Println(sort_select(n))
}
