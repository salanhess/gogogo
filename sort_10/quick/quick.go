package main

import (
	"fmt"
	common_util "gogogo/sort_10/common"
)

/*PartSort()函数是进行一次快排的算法。
对于快速排序的一次排序，有很多种算法，我这里列举一种。
https://blog.csdn.net/qq_36528114/article/details/78667034
###挖坑法
*/

func QuickSort1(l []int, left int, right int) {
	if len(l) <= 1 {
		fmt.Println("l is sorted.")
		return
	}
	if left >= right {
		return
	}
	index := PartSort1(l, left, right)
	QuickSort1(l, left, index-1)
	QuickSort1(l, index+1, right)
}

func QuickSort(l []int, left int, right int, handleCount *int, swapCount *int) {
	if len(l) <= 1 {
		fmt.Println("l is sorted.")
		return
	}
	if left >= right {
		return
	}
	index := PartSort(l, left, right, handleCount, swapCount)
	QuickSort(l, left, index-1, handleCount, swapCount)
	QuickSort(l, index+1, right, handleCount, swapCount)
}

func PartSort1(l []int, left int, right int) (rightindex int) {
	key := l[right]
	for left < right {
		for left < right && l[left] <= key {
			left++
		}
		l[right] = l[left]
		for left < right && l[right] >= key {
			right--
		}
		l[left] = l[right]
	}
	l[right] = key
	return right
}

/*
###挖坑法
1.选取一个关键字(key)作为枢轴，一般取整组记录的第一个数/最后一个，
这里采用选取序列最后一个数为枢轴，也是初始的坑位。
2.设置两个变量left = 0;right = N - 1;
3.从left一直向后走，直到找到一个大于key的值，然后将该数放入坑中，坑位变成了array[left]。
4.right一直向前走，直到找到一个小于key的值，然后将该数放入坑中，坑位变成了array[right]。
5.重复3和4的步骤，直到left和right相遇，然后将key放入最后一个坑位。
*/

func PartSort(l []int, left int, right int, handleCount *int, swapCount *int) (rightindex int) {
	key := l[right]
	for left < right {
		for left < right && l[left] <= key {
			left++
			*handleCount++
		}
		l[right] = l[left]
		*swapCount++
		for left < right && l[right] >= key {
			right--
			*handleCount++
		}
		l[left] = l[right]
		*swapCount++
	}
	l[right] = key
	return right
}

func main() {
	n := common_util.Nums1
	//n = []int{4,1,7,6,9,2,8,0,3,5}
	fmt.Printf("Before sort: n is %v \n========\n", n)
	swapCount, handleCount := 0, 0
	QuickSort(n, 0, len(n)-1, &handleCount, &swapCount)
	fmt.Printf("After sort: n is %v \n handleCount:%d swapCount:%d========\n", n, handleCount, swapCount)
	n = []int{4, 1, 7, 6, 9, 2, 8, 0, 3, 5}
	QuickSort1(n, 0, len(n)-1)
	fmt.Printf("After sort: n is %v", n)
}
