package requests

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func Get(url string) (ret string, err error) {
    var (
        request *http.Request
        resp    *http.Response
        rest    []byte
    )
    if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
        fmt.Println(err)
        return "", err
    }
    request.Header["user-agent"] = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32"}

    if resp, err = http.DefaultClient.Do(request); err != nil {
        fmt.Println(err)
        return "", err
    }
    defer resp.Body.Close()

    if rest, err = ioutil.ReadAll(resp.Body); err != nil {
        fmt.Println(err)
        return
    }
    ret = string(rest)
    return
}
