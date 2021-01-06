package ktools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	err error
)

type (
	Token struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			BucketInfo struct {
				Endpoint string   `json:"endpoint"`
				Region   string   `json:"region"`
				Bucket   string   `json:"bucket"`
				Keys     []string `json:"keys"`
			}
			StsToken struct {
				AccessKeyId     string `json:"accessKeyId"`
				AccessKeySecret string `json:"accessKeySecret"`
				SecurityToken   string `json:"securityToken"`
			}
		}
	}
)

func getToken() (tokenStr string) {
	var (
		//uri string
		request *http.Request
		header  map[string]string
		resp    *http.Response
		token   []byte
	)
	tokenUri := url.URL{
		Scheme:   "https",
		Host:     "upload-ugc.bz.mgtv.com",
		Path:     "upload/image/getStsToken",
		RawQuery: "uuid=2c9ce736a2694b5b8dd4cf309b677a2f&ticket=BRAEK1IU3DU717N47ANG&biz=1&num=1&callback=jQuery182005832266842939937_1591011852911&_support=10000000&_=1591011884333",
	}
	if request, err = http.NewRequest(http.MethodGet, tokenUri.String(), nil); err != nil {
		fmt.Println(err)
		return
	}
	header = map[string]string{
		"authority":       "upload-ugc.bz.mgtv.com",
		"pragma":          "no-cache",
		"cache-control":   "no-cache",
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
		"accept":          "*/*",
		"sec-fetch-site":  "same-site",
		"sec-fetch-mode":  "no-cors",
		"sec-fetch-dest":  "script",
		"referer":         "https://www.mgtv.com/b/338408/8231766.html?fpa=se&lastp=so_result",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,la;q=0.7",
		"cookie":          "__STKUUID=56977c5b-7bb9-4dce-9bfb-4504d3b03b35; mba_deviceid=d5de0866-2126-36b1-d6bc-d81dc685a555; MQGUID=1263467303526383616; __MQGUID=1263467303526383616; pc_v6=v6; _source_=C; PLANB_FREQUENCY=XsaHekEI7nBda-UG; __random_seed=0.09710746912465718; __gads=ID=5a426adf65fcb414:T=1590069115:S=ALNI_MbjtIzVPfp7KjnKslTuXdr6aWe8yw; PM_CHKID=a344546caaf6da69; sessionid=1591011742282; mba_sessionid=f8cd99a8-42ef-7ee8-6e6f-ed46a85ab06c; beta_timer=1591011743950; id=63530547; rnd=rnd; seqid=braek1sr1q1gmk1e4egg; uuid=2c9ce736a2694b5b8dd4cf309b677a2f; vipStatus=3; wei=1a119c627f24dd8ab8bad24e573cb744; wei2=5c14DHntEgaWKDdjHc6c62QLi8HGYlktLJSLQ8abWHVGt91Nqo13SLrvaW%2B9H07trkaV3NRRBy30j1CX0NQ7kJl2JJm0ZS4ck0sVgT9EWndzBVNl4tRlOkfSh0wosk3tIXn%2BjEqTxCfBLH%2BGDPe%2Fw2Slb12JjW3xLubeJaXTE6J6bD6f9WakDdLUIG6xbkT0OqpbMqPOC4Lrhv3R0td%2B0HHSKzo; HDCN=BRAEK1IU3DU717N47ANG-248244420; mba_last_action_time=1591011860129; lastActionTime=1591011881735",
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	if resp, err = http.DefaultClient.Do(request); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if token, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return
	}
	return string(token)
}

func UploadImg(rawImg string) (ret string, err error) {
	var (
		tokenStr        string
		token           Token
		jsonStr         string
		accessKeyId     string
		accessKeySecret string
		securityToken   string
		keys            string
		resp            *http.Response
		res             []byte
		authJson        map[string]string
		payload         []byte
		request         *http.Request
	)
	tokenStr = getToken()
	reg := regexp.MustCompile(`.*?\((.*?)\)`)
	jsonStr = reg.FindStringSubmatch(tokenStr)[1]
	if err = json.Unmarshal([]byte(jsonStr), &token); err != nil {
		fmt.Println(err)
		return
	}

	accessKeySecret = token.Data.StsToken.AccessKeySecret
	accessKeyId = token.Data.StsToken.AccessKeyId
	securityToken = token.Data.StsToken.SecurityToken
	keys = token.Data.BucketInfo.Keys[0]

	// 拿到token发送post请求给node执行
	if resp, err = http.PostForm("http://localhost:3000/getsign", url.Values{
		"accessKeyId":     {accessKeyId},
		"accessKeySecret": {accessKeySecret},
		"securityToken":   {securityToken},
		"keys":            {keys},
	}); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if res, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(res, &authJson)
	auth := authJson["auth"]
	date := authJson["date"]

	ret = fmt.Sprintf("https://mgtv-bbqn.oss-cn-beijing.aliyuncs.com/%s.jpeg", keys)
	if resp, err = http.Get(rawImg); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	payload, err = ioutil.ReadAll(resp.Body)
	if request, err = http.NewRequest(http.MethodPut, ret, strings.NewReader(string(payload))); err != nil {
		fmt.Println(err)
		return
	}
	header := map[string]string{
		"connection":           "keep-alive",
		"pragma":               "no-cache",
		"cache-control":        "no-cache",
		"x-oss-user-agent":     "aliyun-sdk-js/5.2.0 Chrome 81.0.4044.122 on OS X 10.15.5 64-bit",
		"authorization":        auth,
		"x-oss-date":           date,
		"x-oss-security-token": securityToken,
		"user-agent":           "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36",
		"content-type":         "image/jpeg",
		"accept":               "*/*",
		"origin":               "https://www.mgtv.com",
		"sec-fetch-site":       "cross-site",
		"sec-fetch-mode":       "cors",
		"sec-fetch-dest":       "empty",
		"referer":              "https://www.mgtv.com/b/316458/3998946.html",
		"accept-language":      "zh-CN,zh;q=0.9,en;q=0.8,la;q=0.7",
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	if resp, err = http.DefaultClient.Do(request); err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	return
}

//func main() {
//	var ret string
//	if ret, err = uploadImg("https://entgo.io/assets/gopher-schema-as-code.png"); err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//}
