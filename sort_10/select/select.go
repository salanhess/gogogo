package main

import "fmt"

/*windows使用git时出现：warning: LF will be replaced by CRLF
以及使windows10支持区分大写小的问题处理
解决办法（二）

找到windows项目的.git目录,修改config文件，在 [core] 配置项添加一行即可。
autocrlf = false*/

func sort_select(nums []int) []int {
	n := len(nums)
	for i, _ := range nums {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				nums[minIndex], nums[j] = nums[j], nums[minIndex]
			}
		}
	}
	return nums
}

func main() {
	n := []int{1, 2, 6, 4}
	fmt.Println(sort_select(n))
}
