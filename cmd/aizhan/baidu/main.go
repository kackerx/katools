package main

import (
    "bufio"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/pkg/errors"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "regexp"
    "strings"
    "sync"
    "time"
)

var (
    wg  sync.WaitGroup
    err error
    res *http.Response
    //client = &http.Client{}
    USER_AGENTS = []string{
        "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
        "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
        "Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
        "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
        "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
        "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
        "Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
        "Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
        "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
        "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
        "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
    }
    COOKIES    []string
    PROXY      []string
    ZHAN_PROXY []string
    akey       string
)

type Proxy struct {
    Ips     []string `json:"ips"`
    Success string   `json:"success"`
}

type ZhanProxy struct {
    Code string `json:"code"`
    Msg  string `json:"msg"`
    Data struct {
        Count     int `json:"count"`
        ProxyList []struct {
            IP   string `json:"ip"`
            Port int    `json:"port"`
        } `json:"proxy_list"`
    } `json:"data"`
}

func init() {
    h := md5.New()
    h.Write([]byte("49903389"))
    akey = hex.EncodeToString(h.Sum(nil))[8:24]
}

func main() {
    st := time.Now()
    //test()
    //time.Sleep(time.Second * 10)
    //bindIp()
    proxy, err := GetAli()
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }

    PROXY = proxy.Ips
    fmt.Println("成功获取代理ip数量: ", len(PROXY))

    GetProxy()
    fmt.Println("成功获取站大爷代理ip数量: ", len(ZHAN_PROXY))
    // 三分钟换一批ip
    go func() {
        t := time.NewTicker(time.Second * 12)
        for {
            select {
            case <-t.C:
                GetProxy()
                fmt.Println("更新ip池...", len(ZHAN_PROXY))
            }
        }
    }()

    sig := make(chan interface{}, 22)
    urlCh := make(chan string, 100)
    retCh := make(chan string, 5000)
    var count int = 0

    defer close(sig)

    go func() {
        defer close(urlCh)
        file, err := os.Open("./domain.txt")
        if err != nil {
            log.Fatalln(err)
        }

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            urlCh <- scanner.Text()
        }
    }()

    for url := range urlCh {
        sig <- nil
        wg.Add(1)
        count++
        fmt.Println(url, "抓取中...", count)
        go GetBaidu(url, retCh, sig)
        if count%50 == 0 {
            fmt.Println(count, "已爬取")
            time.Sleep(time.Second * 2)
        }
    }

    wg.Wait()
    close(retCh)

    outFile, err := os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR, 0666)
    defer outFile.Close()
    if err != nil {
        fmt.Println(err)

    }

    fmt.Println("结果数目: ", len(retCh))
    for v := range retCh {
        fmt.Println(v)
        outFile.WriteString(v + "\n")
    }

    fmt.Println("任务结束")
    fmt.Println("耗时: ", time.Since(st))
    time.Sleep(time.Second * 10)
}

func GetBaidu(uri string, retCh chan string, sig chan interface{}) (string, error) {
    defer wg.Done()
    defer func() {
        <-sig
    }()
    targetUrl := fmt.Sprintf(`https://www.baidu.com/s?ie=UTF-8&wd="%s"`, uri)
    //targetUrl := "https://www.baidu.com/s?ie=UTF-8&wd=%22" + uri + "%22"
    //targetUrl := fmt.Sprintf(`https://www.baidu.com/s?q1="%s"`, uri)
    //targetUrl := ""
    method := "GET"

    req, err := http.NewRequest(method, targetUrl, nil)
    if err != nil {
        fmt.Println(err)
        return "", nil
    }

    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5")
    req.Header.Add("Cache-Control", "no-cache")
    req.Header.Add("Connection", "keep-alive")
    //req.Header.Add("Cookie", "BIDUPSID=7697652968A15AB9AAAC12836A7316A3; PSTM=1652965889; BAIDUID=7697652968A15AB9F7541392C857762D:FG=1; delPer=0; BD_CK_SAM=1; PSINO=3; BD_UPN=123253; BA_HECTOR=048l8ga08k0g2k05201h8cgg315; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; ZFY=MBRzdwzXQPG6CRXmJsQvFQZDT5zPsqBnOAsg4m5:B9Cw:C; H_PS_PSSID=31253_36452_35910_36166_35979_36055_36235_26350_36299_22157_36447; H_PS_645EC=3e20Yf4mqMEbV8Yxu8V5fvQ2TN3Rt14OQY1XrawqdSlLeRSZntPgvDMEPlU; BDSVRTM=239; channel=baidusearch; baikeVisitId=b074a1fa-ad6a-4348-9cd5-1ac488b713e5")
    req.Header.Add("Pragma", "no-cache")
    req.Header.Add("Sec-Fetch-Dest", "document")
    req.Header.Add("Sec-Fetch-Mode", "navigate")
    req.Header.Add("Sec-Fetch-Site", "same-origin")
    req.Header.Add("Sec-Fetch-User", "?1")
    req.Header.Add("Upgrade-Insecure-Requests", "1")
    req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Microsoft Edge\";v=\"101\"")
    //req.Header.Add("sec-ch-ua-mobile", "?0")
    //req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
    //req.Header.Add("sec-ch-ua-platform", "\"windows\"")
    recount := 0
REPEAT:
    rand.Seed(time.Now().UnixNano())

    ua := USER_AGENTS[rand.Intn(len(USER_AGENTS))]
    req.Header.Add("User-Agent", ua)

    //cookie := COOKIES[rand.Intn(len(COOKIES))]
    //req.Header.Add("Cookie", cookie)
    ran := rand.Intn(100)
    var proxy *url.URL

    if ran < 5 {
        proxyUrl := "http://proxy:orderId=O21072616193919678168&sign=3014798228629a3c2ee159b7a0bd35c7&time=1627290066&pid=-1&cid=@proxy-service2.vpsnb.net:14223"
        proxy, _ = url.Parse(proxyUrl)
        //p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        //proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran >= 5 && ran <= 10 {
        p := PROXY[rand.Intn(len(PROXY))]
        proxy, _ = url.Parse("http://" + p)
        //p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        //proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran > 10 && ran < 25 {
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
        //proxy = nil
    } else {
        //p := PROXY[rand.Intn(len(PROXY))]
        //proxy, _ = url.Parse("http://" + p)
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    }
    recount++

    //if proxy != nil {
    //    fmt.Println("使用代理: ", proxy.Host, ran)
    //}

    client := &http.Client{
        Transport: &http.Transport{
            Proxy:               http.ProxyURL(proxy),
            MaxIdleConns:        200,
            MaxIdleConnsPerHost: 100,
        },
        Timeout: time.Second * 10,
    }
    //client := http.DefaultClient

    //for i := 0; i < 5; i++ {
    time.Sleep(time.Second * time.Duration(rand.Intn(3)))
    //ua := USER_AGENTS[rand.Intn(len(USER_AGENTS))]
    //req.Header.Add("User-Agent", ua)
    res, err = client.Do(req)
    if err != nil {
        //urlCh <- uri
        if recount < 3 {
            fmt.Println("request: ", "重试...", uri, ", 第几次: ", recount)
            goto REPEAT
        } else {
            return "", nil
        }
        //return "", err
        //continue
    }

    //}
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        if recount < 3 {
            fmt.Println("read: ", "重试...", uri, ", 第几次: ", recount)
            goto REPEAT
        } else {
            return "", nil
        }
        //return "", err
    }

    ret := string(body)

    //fmt.Println(ret)
    //f, err := os.OpenFile("./index.html", os.O_APPEND|os.O_WRONLY, 777)
    //if err != nil {
    //    fmt.Println(err)
    //    return "", err
    //}
    //f.WriteString(ret)

    if strings.Contains(ret, "安全验证") {
        //fmt.Println(ret)
        if recount < 3 {
            fmt.Println("触发百度验证: ", "重试...", uri, ", 第几次: ", recount)
            goto REPEAT
        } else {
            return "", nil
        }
        //urlCh <- uri
        //goto REPEAT
        //goto REPEAT
        //return "", nil
        //fmt.Println(res.Cookies())
        //time.Sleep(time.Second * 20)
        //goto REPEAT
    }

    re := regexp.MustCompile(fmt.Sprintf("(?s).*?class=\"c-container\".*?<em>(%s)</em>.*?</div></div></div></div><div></div></div>", uri))
    retBody := re.FindAllStringSubmatch(ret, -1)

    if len(retBody) == 0 {
        //fmt.Println(uri, "no: ", len(retBody))
        return "", nil
    }

    count := len(retBody)

    for _, v := range retBody {
        //fmt.Println(v[0])
        if strings.Contains(v[0], "查询网") ||
            strings.Contains(v[0], "站长工具") ||
            strings.Contains(v[0], "站长之家") ||
            strings.Contains(v[0], "时代互联") ||
            strings.Contains(v[0], "名站在线") ||
            strings.Contains(v[0], "网站排名") ||
            //strings.Contains(v[0], "企查查") ||
            strings.Contains(v[0], "爱企查") ||
            //strings.Contains(v[0], "外链查询") ||
            //strings.Contains(v[0], "www.beianw.com") ||
            strings.Contains(v[0], "爱站网") ||
            strings.Contains(v[0], "西部数码") {
            count--
        }
    }

    r := fmt.Sprintf("%s\t结果条目: %d\t过滤注册商: %d", uri, len(retBody), count)
    retCh <- r
    // 写入html
    if _, err := os.Stat("./html"); os.IsNotExist(err) {
        os.Mkdir("./html", 0777)
        os.Chmod("./html", 0777)
    }
    f, err := os.OpenFile(fmt.Sprintf("./html/%s.html", uri), os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        fmt.Println(err)
    }
    _, err = f.WriteString(ret)
    if err != nil {
        fmt.Println("文件写入错误: ", err)
    }

    return ret, nil
}

func GetAli() (proxy Proxy, err error) {
    resp, err := http.Get("http://101.33.117.86:9999/get")
    if err != nil {
        return
    }
    defer resp.Body.Close()

    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }

    err = json.Unmarshal(res, &proxy)
    if err != nil {
        return proxy, errors.Wrap(err, "序列化错误: ")
    }
    return
}

func bindIp() string {
    url := "http://www.zdopen.com/ShortProxy/BindIP/"

    url = fmt.Sprintf("%s?api=%s&akey=%s&i=2", url, "202205271640405995", akey)

    get, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    resp, err := ioutil.ReadAll(get.Body)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer get.Body.Close()

    fmt.Println("绑定代理使用ip...")

    return string(resp)
}

func GetProxy() string {
    url := fmt.Sprintf("http://www.zdopen.com/ShortProxy/GetIP/?api=%s&akey=%s&count=30&timespan=4&type=3", "202205271640405995", akey)

    get, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    resp, err := ioutil.ReadAll(get.Body)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer get.Body.Close()

    var proxyList ZhanProxy
    err = json.Unmarshal(resp, &proxyList)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    if len(proxyList.Data.ProxyList) > 0 {
        var list []string
        for _, v := range proxyList.Data.ProxyList {
            list = append(list, fmt.Sprintf("%s:%d", v.IP, v.Port))
        }
        ZHAN_PROXY = list
    } else {
        fmt.Println("获取代理为空: ", string(resp))
    }
    return string(resp)
}

func test() {
    //targetUrl := "https://www.baidu.com/s?ie=UTF-8&wd=%22" + uri + "%22"
    //targetUrl := fmt.Sprintf(`https://www.baidu.com/s?q1="%s"`, uri)
    GetProxy()
    //targetUrl := "http://current.ip.16yun.cn:802/"
    targetUrl := "http://httpbin.org/get"
    method := "GET"

    req, err := http.NewRequest(method, targetUrl, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5")
    req.Header.Add("Cache-Control", "no-cache")
    req.Header.Add("Connection", "keep-alive")
    //req.Header.Add("Cookie", "BIDUPSID=7697652968A15AB9AAAC12836A7316A3; PSTM=1652965889; BAIDUID=7697652968A15AB9F7541392C857762D:FG=1; delPer=0; BD_CK_SAM=1; PSINO=3; BD_UPN=123253; BA_HECTOR=048l8ga08k0g2k05201h8cgg315; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; ZFY=MBRzdwzXQPG6CRXmJsQvFQZDT5zPsqBnOAsg4m5:B9Cw:C; H_PS_PSSID=31253_36452_35910_36166_35979_36055_36235_26350_36299_22157_36447; H_PS_645EC=3e20Yf4mqMEbV8Yxu8V5fvQ2TN3Rt14OQY1XrawqdSlLeRSZntPgvDMEPlU; BDSVRTM=239; channel=baidusearch; baikeVisitId=b074a1fa-ad6a-4348-9cd5-1ac488b713e5")
    req.Header.Add("Pragma", "no-cache")
    req.Header.Add("Sec-Fetch-Dest", "document")
    req.Header.Add("Sec-Fetch-Mode", "navigate")
    req.Header.Add("Sec-Fetch-Site", "same-origin")
    req.Header.Add("Sec-Fetch-User", "?1")
    req.Header.Add("Upgrade-Insecure-Requests", "1")
    req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Microsoft Edge\";v=\"101\"")
    //req.Header.Add("sec-ch-ua-mobile", "?0")
    //req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
    //req.Header.Add("sec-ch-ua-platform", "\"windows\"")
    rand.Seed(time.Now().UnixNano())

    ua := USER_AGENTS[rand.Intn(len(USER_AGENTS))]
    req.Header.Add("User-Agent", ua)

    //cookie := COOKIES[rand.Intn(len(COOKIES))]
    //req.Header.Add("Cookie", cookie)

    ran := rand.Intn(100)
    var proxy *url.URL

    if ran < 10 {
        proxyUrl := "http://proxy:orderId=O21072616193919678168&sign=3014798228629a3c2ee159b7a0bd35c7&time=1627290066&pid=-1&cid=@proxy-service2.vpsnb.net:14223"
        proxy, _ = url.Parse(proxyUrl)
        //p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        //proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran >= 10 && ran <= 20 {
        p := PROXY[rand.Intn(len(PROXY))]
        proxy, _ = url.Parse("http://" + p)
        //p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        //proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran > 20 && ran < 25 {
        //proxy = nil
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else {
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    }

    //if proxy != nil {
    //    fmt.Println("使用代理: ", proxy.Host, ran)
    //}

    fmt.Println(ZHAN_PROXY)
    for _, v := range ZHAN_PROXY {
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + v)
        client := http.Client{
            Transport: &http.Transport{
                Proxy: http.ProxyURL(proxy),
            },
            Timeout: time.Second * 7,
        }

        req, _ := http.NewRequest(method, targetUrl, nil)
        res, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
            continue
        }

        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)
        //if err != nil {
        //    fmt.Println("read: ", err)
        //    return
        //}

        ret := string(body)
        fmt.Println(ret)

    }

    //client := http.DefaultClient

    //REPEAT:

    //for i := 0; i < 5; i++ {
    //time.Sleep(time.Second * time.Duration(rand.Intn(3)))
    //ua := USER_AGENTS[rand.Intn(len(USER_AGENTS))]
    //req.Header.Add("User-Agent", ua)
    //res, err = client.Do(req)
    //if err != nil {
    //    fmt.Println("request: ", err)
    //    return
    //    //continue
    //}

    //}
}
