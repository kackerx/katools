package main

import "fmt"

func sort(arr []int, l, r int) {
    if l >= r {
        return
    }
    
    mid := int((l + r) / 2)
    sort(arr, l, mid)
    sort(arr, mid + 1, r)
    
    merge(arr, l, mid, r)
}

func merge(arr []int, l int, mid int, r int) {
    temp := make([]int, len(arr))
    i := l
    j := mid+1
    k := 0
    
    
    for i <= mid && j <= r {
        if arr[i] < arr[j] {
            temp[k] = arr[i]
            i++
            k++
        } else {
            temp[k] = arr[j]
            j++
            k++
        }
    }
    
    for i <= mid {
        temp[k] = arr[i]
        i++
        k++
    }
    
    for j <= r {
        temp[k] = arr[j]
        j++
        k++
    }
    
    k = 0
    tempLeft := l
    for tempLeft <= r {
        arr[tempLeft] = temp[k]
        k++
        tempLeft++
    }
    
}

func main() {
    a := []int{1, 3, 5, 7, 2, 4, 6, 8}
    sort(a, 0, 7)
    fmt.Println(a)
}
