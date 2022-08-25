package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "regexp"
    "strconv"
    "testing"
    "time"
)

func TestRand(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 50; i++ {
        fmt.Println(rand.Intn(3))
    }
    //slice := []int{0, 2, 3, 4, 5, 6, 7, 8}
    //
    //rand.Seed(time.Now().UnixNano())
    //
    //var rm int
    //for i := 0; i < 20; i++ {
    //    rm = rand.Intn(len(slice))
    //    fmt.Println(rm)
    //}
}

func TestGetBaidu(t *testing.T) {
    body := `
le>&quot;zjkrgssr.com&quot;_百度搜索</title>
<style data-for="result" type="text/css" id="css_newi_result">body{color:#333;background:#fff;padding:6px 0 0;margin:0;position:relative}
body,th,td,.p1,.p2{font-family:arial}
p,form,ol,ul,li,dl,dt,dd,h3{margin:0;padding:0;list-style:none}
input{padding-top:0;padding-bottom:0;-moz-box-sizing:border-box;-webkit-box-sizing:border-box;box-sizing:border-box}
table,img{border:0}
<em>zjkrgsr.com</em>
td{font-size:9pt;line-h<em>zjkrssr.com</em>eight:18px}
<em>zjkrgsr.com</em>
em{font-style:normal}
    `

    url := "zjkrgssr.com"

    re := regexp.MustCompile(fmt.Sprintf("(?s).*?<em>(%s)</em>.*?", url))

    retBody := re.FindAllStringSubmatch(body, -1)

    fmt.Println(len(retBody))
    for _, v := range retBody {
        fmt.Println("======")
        fmt.Println(v[1])
    }
}

func TestPerson(t *testing.T) {
    fmt.Println("kacker")
}

func TestProxy(t *testing.T) {
    cli := redis.NewClient(&redis.Options{
        Addr: "10.0.4.105:6379",
        DB:   0,
    })

    res, err := cli.Get(context.Background(), "public:iproxy:aliyun_num").Result()
    if err != nil {
        fmt.Println(err)
        return
    }

    i, _ := strconv.Atoi(res)
    ret, err := cli.Get(context.Background(), fmt.Sprintf("public:iproxy:91vps:ip_%d", rand.Intn(i))).Result()
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(ret)
}

func TestFoo(t *testing.T) {
    a := []string{"0", "1", "2"}
    a = a[:0]

    fmt.Println(len(a))
    fmt.Println(cap(a))
    fmt.Println(a)

    var b []int
    fmt.Println(len(b))
    fmt.Println(cap(b))
    fmt.Println(b)

    b = append(b, 1)

    fmt.Println(len(b))
    fmt.Println(cap(b))
    fmt.Println(b)

    c := make([]int, 3)
    copy(c, b)

    //fmt.Println(c)
    c[0] = 0
    fmt.Println(b)
    fmt.Println(c)
}

func TestGetProxy(t *testing.T) {
    //bindIp()
    got := GetProxy()
    fmt.Println(got)

}

func TestProx(t *testing.T) {
    //targetUrl := "https://www.baidu.com/s?ie=UTF-8&wd=%22" + uri + "%22"
    //targetUrl := fmt.Sprintf(`https://www.baidu.com/s?q1="%s"`, uri)
    GetProxy()
    targetUrl := "http://current.ip.16yun.cn:802/"
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
        //proxyUrl := "http://proxy:orderId=O21072616193919678168&sign=3014798228629a3c2ee159b7a0bd35c7&time=1627290066&pid=-1&cid=@proxy-service2.vpsnb.net:14223"
        //proxy, _ = url.Parse(proxyUrl)
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran >= 10 && ran <= 20 {
        //p := PROXY[rand.Intn(len(PROXY))]
        //proxy, _ = url.Parse("http://" + p)
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    } else if ran > 20 && ran < 25 {
        proxy = nil
    } else {
        p := ZHAN_PROXY[rand.Intn(len(ZHAN_PROXY))]
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + p)
    }

    //if proxy != nil {
    //    fmt.Println("使用代理: ", proxy.Host, ran)
    //}

    for _, v := range ZHAN_PROXY {
        proxy, _ = url.Parse("http://202205271640405995:49903389@" + v)
        client := http.Client{
            Transport: &http.Transport{
                Proxy: http.ProxyURL(proxy),
            },
            Timeout: time.Second * 7,
        }

        req, _ := http.NewRequest(method, targetUrl, nil)
        res, _ := client.Do(req)
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
