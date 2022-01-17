package array

import "fmt"

func GetSparseArray() [][3]int {
    a := [11][11]int{}
    //1黑, 2白
    a[1][2] = 1
    a[2][3] = 2
    
    // sum
    sum := 0
    for _, v := range a {
        for _, data := range v {
            if data != 0 {
                sum++
            }
        }
    }
    
    //创建稀疏数组
    //rol := sum+1
    sparseArray := make([][3]int, sum+1)
    //fmt.Println(sparseArray)
    sparseArray[0][0] = 11
    sparseArray[0][1] = 11
    sparseArray[0][2] = sum
    
    var count int //用于记录第几个非零数据
    for i, v := range a {
        for j, data := range v {
            if data != 0 {
                count++
                sparseArray[count][0] = i
                sparseArray[count][1] = j
                sparseArray[count][2] = data
            }
        }
    }
    
    for _, v := range sparseArray {
        for _, data := range v {
            fmt.Print(data, "\t")
        }
        fmt.Println()
    }
    return sparseArray
}

func Get2Array(sparseArray [][3]int) (arr [11][11]int) {
    for _, v := range sparseArray[1:] {
        arr[v[0]][v[1]] = v[2]
    }
    
    for _, v := range arr {
        for _, data := range v {
            fmt.Print(data, "\t")
        }
        fmt.Println()
    }
    return
}



