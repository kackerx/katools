package main

import (
    "log"
    "math/rand"
    "sync"
    "time"
)

func main() {
    c := sync.NewCond(&sync.Mutex{})
    var ready int

    for i := 0; i < 10; i++ {
        go func(i int) {
            time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

            // 加锁更改等待条件
            c.L.Lock()
            ready++
            c.L.Unlock()

            log.Printf("运动员#%d 已准备就绪\n", i)

            // 广播唤醒所有的等待者
            c.Broadcast()
        }(i)
    }

    c.L.Lock()
    for ready != 10 {
        c.Wait()
        log.Println("裁判员被唤醒一次")
    }
    c.L.Unlock()

    //所有的运动员是否就绪
    log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}

//一，条件变量更改的时候需要是原子操作或者互斥锁保护；二，Wait 的操作需要加锁保护；三，Wait 唤醒之后仍然需要检查条件。
//
//使用 Cond 最常见的2个错误
//Wait 调用没有加锁
//Wait 调用一定要加锁。如果不加锁，Wait 内部执行 Unlock 操作的时候就会解锁一个未加锁的 Locker，直接报 panic 错误。
//
//为什么这个加锁操作不放在 Wait 内部执行呢？因为和 Wait 同时执行的还有对条件变量的判断，条件变量的判断是需要加锁的（个人理解）。
//
//Wait 只调用一次
//细心的读者会发现，Wait 调用的时候，外面嵌套了 for 循环，这个 for 循环不能丢。因为 Wait 收到唤醒通知，并不能确定这一组 goroutine 是否都完成了任务，有可能只是其中一个完成了，因此还需要再检查一次条件是否满足。
//
//这也是和 WaitGroup 不同的地方。如果这个条件就是全组 goroutine 都完成任务，检查一次确实就够了。
