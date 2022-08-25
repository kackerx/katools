package main

import (
    "bufio"
    "fmt"
    "github.com/kackerx/katools/requests"
    "github.com/pkg/errors"
    "github.com/tidwall/gjson"
    "golang.org/x/sync/errgroup"
    "os"
    "strconv"
    "sync"
    "sync/atomic"
    "time"
)

const LIMIT = 100

var (
    err       error
    eg        errgroup.Group
    wg        sync.WaitGroup
    workerNum int32 = 0
    workCh          = make(chan string)

    FANLINK_API = "https://apistore.aizhan.com/site/linkinfos/9858573c9f1da3d398676623b207604a?domain=%s&type=home&page=1"
    INDEX_API   = "https://apistore.aizhan.com/baidu/sladdtasks/9858573c9f1da3d398676623b207604a?domains=%s"
    SlIdApi     = "https://apistore.aizhan.com/baidu/sladdtasks/9858573c9f1da3d398676623b207604a?domains=%s"
    GetSlApi    = "https://apistore.aizhan.com/baidu/sltasksdata/9858573c9f1da3d398676623b207604a?taskids=%s"
    workDone    = make(chan string, 100)
)

func handle(url string) error {
    //fmt.Println("request url", url)
    defer wg.Done()
    target := fmt.Sprintf(FANLINK_API, url)
    for {
        resp, err := requests.Get(target)
        if err != nil {
            continue
            return err
        }
        totalNum := gjson.Get(resp, "data.total_num").Int()
        if totalNum > 0 {
            workDone <- url + ": " + strconv.Itoa(int(totalNum))
        }
        atomic.AddInt32(&workerNum, -1)
        return nil
    }

    //return nil
}

func GetTaskId(out chan string, url string) error {
    defer wg.Done()

    //fmt.Println("get url: ", url)

    resp, err := requests.Get(url)
    if err != nil {
        fmt.Println(err)
        return errors.Wrap(err, "get taskid error")
    }

    ret := gjson.Get(resp, "data.success").Array()
    var taskIds string
    for _, v := range ret {
        taskId := v.Get("taskid")
        taskIds += taskId.String() + ","
    }

    for {
        resp, err = requests.Get(fmt.Sprintf(GetSlApi, taskIds[:len(taskIds)-1]))
        if err != nil {
            fmt.Println(err)
            return errors.Wrap(err, "get sl error")
        }

        if gjson.Get(resp, "code").Int() == 200000 {
            break
        } else {
            fmt.Println("获取失败: ", resp)
            time.Sleep(time.Second * 2)
        }
    }
    var list map[string]gjson.Result
    list = gjson.Get(resp, "data.list").Map()
    for _, v := range list {
        //fmt.Println(v)
        if v.Get("baidusl").Int() > 0 {
            //fmt.Println("send out")
            out <- v.Get("domain").String() + ": " + v.Get("baidusl").String()
        }
    }
    return nil
}

func main() {

    fmt.Println("开始检测")
    f, err := os.Open("./domain.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    ch := make(chan string, 50)
    resCh := make(chan string, 50)

    go func() {
        scanner := bufio.NewScanner(f)
        flag := 0
        var targetUrl string
        for scanner.Scan() {
            url := scanner.Text()
            targetUrl += url + "|"
            flag++
            if flag%50 == 0 {
                ch <- targetUrl
                targetUrl = ""
            }
        }
        if targetUrl != "" {
            ch <- targetUrl
        }
        close(ch)
    }()

    for v := range ch {
        query := v[:len(v)-1]
        wg.Add(1)
        go GetTaskId(resCh, fmt.Sprintf(SlIdApi, query))
        //if err != nil {
        //    fmt.Printf("%+v\n", err)
        //}
    }

    go func() {
        wg.Wait()
        fmt.Println("wg done")
        close(resCh)
    }()

    f, err = os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    for v := range resCh {
        f.WriteString(v + "\n")
    }

    //    ctx, cancel := context.WithCancel(context.Background())
    //
    //    go func() {
    //        scanner := bufio.NewScanner(f)
    //        for scanner.Scan() {
    //            url := scanner.Text()
    //            //fmt.Println(url)
    //            workCh <- strings.Trim(url, "\n")
    //        }
    //
    //        // 读取完毕全部任务
    //        cancel()
    //    }()
    //
    //    for {
    //        select {
    //        case url := <-workCh:
    //            //fmt.Println("task num start", workerNum)
    //            if workerNum > LIMIT {
    //                fmt.Println("多线程等待中...")
    //                time.Sleep(time.Second)
    //                continue
    //            }
    //            workerNum++
    //            wg.Add(1)
    //            go handle(url)
    //
    //        case <-ctx.Done():
    //            //fmt.Println("完成任务")
    //            goto Finish
    //        }
    //    }
    //Finish:
    //    wg.Wait()
    //    fmt.Println("end: ", len(workDone))
    //    close(workDone)
    //    //fmt.Println(workerNum)
    //    f, err = os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR, 0666)
    //    if err != nil {
    //        fmt.Println(err)
    //        return
    //    }
    //
    //    for v := range workDone {
    //        f.WriteString(v + "\n")
    //    }
    fmt.Println("任务结束")
}
