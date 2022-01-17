package pkg

import (
    "io/ioutil"
    "net/http"
)

func Get(url string) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header["User-Agent"] = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44"}
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        resp.Body.Close()
        return nil, err
    }
    
    ret, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    return ret, err
}




