package main

import "fmt"

func binFind(arr []int, target int) int {
    l := 0
    r := len(arr) - 1
    for l <= r {
        mid := int((l + r) / 2)
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1
}

func recursionBinFind(arr []int, l, r, target int) int {
    if l > r {
        return -1
    }
    
    mid := int((l + r) / 2)
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        return recursionBinFind(arr, mid+1, r, target)
    } else {
        return recursionBinFind(arr, l, mid-1, target)
    }
}

// 查找大于target的最小值
func binFindMin(arr []int, target int) int {
    
    // 在[l, r]寻找
    l := 0
    r := len(arr) // 范围囊括最后一个元素
    for l < r {
        mid := int((l+r) / 2)
        if arr[mid] > target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    
    return l
}




func main() {
    a := []int{1, 2, 3, 4, 5}
    ret := recursionBinFind(a, 0, 4, 5)
    ret2 := binFind(a, 5)
    fmt.Println(ret)
    fmt.Println(ret2)
    ret3 := binFindMin(a, 3)
    fmt.Println(ret3)
}
