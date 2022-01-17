package main

import (
	"fmt"
)

func Select(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println(arr)
}

func Insert(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j-1 >= 0 && arr[j] < arr[j-1]; j-- { // j从下标i开始, 比较j和j-1, j-1 >= 0, 小的话就交换
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	fmt.Println(arr)
}

func Merge(arr []int, left, mid, right int) {
	i := left
	j := mid + 1
	t := 0
	temp := make([]int, len(arr))
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[t] = arr[i]
			i++
			t++
		} else {
			temp[t] = arr[j]
			j++
			t++
		}
	}

	for i <= mid {
		temp[t] = arr[i]
		i++
		t++
	}
	for j <= right {
		temp[t] = arr[j]
		j++
		t++
	}
	t = 0
	tempLeft := left
	for tempLeft <= right {
		arr[tempLeft] = temp[t]
		t++
		tempLeft++
	}
}

// 注意递归的宏观语义: 对一个数组, 排序l到r之间的数据
func MergeSort(arr []int, l, r int) {
	if l >= r { // l == r 要么空要么一个元素, 直接返回
		return // 递归结束条件
	}

	mid := int((l + r) / 2)
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	Merge(arr, l, mid, r)
}

func find(arr []int, k int) int {
	return selectk(arr, 0, len(arr)-1, len(arr)-k)
}

func find2(arr []int, k int) int {
	return selectk(arr, 0, len(arr)-1, k-1)
}

func selectk(arr []int, l, r, k int) int {
	p := partition2(arr, l, r)

	if p == k {
		return arr[p]
	} else if p < k {
		return selectk(arr, p+1, r, k)
	} else {
		return selectk(arr, l, p-1, k)
	}
}

func quickSork(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition2(arr, l, r)
	quickSork(arr, l, p-1)
	quickSork(arr, p+1, r)
}

func partition(arr []int, l int, r int) int {
	target := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < target {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition2(arr []int, l int, r int) int {
	target := arr[l]
	i := l + 1
	j := r

	for {
		for i <= j && arr[i] < target { // 左指针扫描到比target大的
			i++
		}

		for j >= i && arr[j] > target { // 右指针扫描比target小的
			j--
		}

		if i >= j {
			break
		} // 相等扫描完毕

		arr[i], arr[j] = arr[j], arr[i] // 找到左右两边需要交换的位置
		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition3(arr []int, l int, r int) {
	// arr[l+1], lt], arr[lt+1, i-1], arr[gt, r]
	if l >= r {
		return
	}
	lt := l
	i := l + 1
	gt := r + 1
	for i < gt {
		if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > arr[l] {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}

	arr[l], arr[lt] = arr[lt], arr[l]

	partition3(arr, l, lt-1)
	partition3(arr, gt, r)
}

func main() {
	a := []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(a)
	fmt.Println(find2(a, 4))
	fmt.Println(a[:find2(a, 4)])
}
