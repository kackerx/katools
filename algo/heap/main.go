package main

import "fmt"

type MaxHeap struct {
    data []int
}

func NewMaxHeap() *MaxHeap {
    return &MaxHeap{
        data: make([]int, 0),
    }
}

func (m *MaxHeap) size() int {
    return len(m.data)
}

func (m *MaxHeap) isEmpty() bool {
    return len(m.data) == 0
}

func (m *MaxHeap) parent(index int) int {
    return (index - 1) / 2
}

func (m *MaxHeap) left(index int) int {
    return 2*index + 1
}

func (m *MaxHeap) right(index int) int {
    return 2*index + 2
}

func (m *MaxHeap) add(e int) {
    m.data = append(m.data, e)
    m.siftUp(len(m.data) - 1)
}

func (m *MaxHeap) siftUp(k int) {
    for k > 0 && m.data[k] > m.data[m.parent(k)] {
        m.data[k], m.data[m.parent(k)] = m.data[m.parent(k)], m.data[k]
        k = m.parent(k)
    }
}

// 看最大元素
func (m *MaxHeap) findMax() int {
    return m.data[0]
}

// 取出最大元素
func (m *MaxHeap) getMax() int {
    ret := m.findMax()
    m.data[0], m.data[m.size()-1] = m.data[m.size()-1], m.data[0] // 交换堆顶和最后一个元素
    m.data = m.data[:m.size()-1]
    m.siftDown(0)
    
    return ret
}

func (m *MaxHeap) siftDown(k int) {
    for m.left(k) < m.size() { // 左孩子如果大于元素数, 说明没有了孩子
        l := m.left(k)
        r := m.right(k)
        j := l // j是左孩子索引, 如果右孩子更大, j就是右孩子索引
        if r < m.size() && m.data[r] > m.data[l] {
            j = r
        }
        // 目前保证了j是左右孩子中的最大值
        if m.data[k] >= m.data[j] {
            break
        }
        m.data[k], m.data[j] = m.data[j], m.data[k]
        k = j
    }
}

func (m *MaxHeap) sort(arr []int) {
    m.data = make([]int, 0)
    for _, v := range arr {
        m.add(v)
    }
    
    for i := len(arr) - 1; i >= 0; i-- {
        arr[i] = m.getMax()
    }
    
    fmt.Println(arr)
}

// 数组维护成堆: 从最后一个非叶子节点向前遍历, 每个进行siftDown
func (m *MaxHeap) heapify(arr []int) {
    m.data = arr
    for i := m.parent(len(m.data) - 1); i >= 0; i-- {
        m.siftDown(i)
    }
}

// 堆排序
func heapSort(arr []int) {
    // heapify: 从最后一个非叶子节点开始执行下沉操作, 把数组维护成一个堆
    for i := (len(arr) - 1) / 2; i >= 0; i-- {
        siftDown(arr, i, len(arr))
    }
    
    // 每次交换堆顶到数组末尾, 去掉末尾继续执行
    for i := len(arr) - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        siftDown(arr, 0, i) // 每次堆总数比上次少一: i
    }
    
    fmt.Println(arr)
}

func siftDown(arr []int, i int, n int) {
    for i * 2 + 1 < n {
        l := i * 2 + 1
        r := i * 2 + 2
        j := l
        if r < n && arr[j] < arr[r] {
            j = r
        }
        
        if arr[i] >= arr[j] {
            break
        }
        
        arr[i], arr[j] = arr[j], arr[i]
        i = j
    }
}



func main() {
    heapSort([]int{1, 8, 4, 6, 3, 10, 11, 19, 0, 1})
}
