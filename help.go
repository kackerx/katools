package katools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Img struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

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
		RawQuery: "uuid=18d6436f26ce4d7ba0d258c439a736a5&ticket=E29530A8ABC7390A92E8A481B8767814&biz=1&num=1&callback=jQuery18205619504208532247_1631539923435&_support=10000000&_=1631539954417",
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
		"cookie":          "_source_=C; __STKUUID=cd318ade-b019-4154-8107-ba041fec0a78; PLANB_FREQUENCY=YJZ9HSCYw0r8lQWz_21050819; MQGUID=1390998916581019648; __MQGUID=1390998916581019648; mba_deviceid=7784f841-fee4-6fe6-d607-431af05bd9a8; __random_seed=0.9222008335131051; PM_CHKID=f4f3ce5cc212f7d8; mg_uuid=ba18b8b7-9b22-4e2c-95b9-bae9537dc53b; NUC_STATE=1629192769.xolvoAWiekrbca-B8z86q4u9MuA; loginAccount=18838974677; uuid=389edfe9aedc4e48afc06487f8f90861; vipStatus=3; rnd=rnd; id=77573232; wei=7a2e64709ca753b3338e667d546153cb; wei2=7c83gWi6PMZ1jieQaKmcVeR%2FVuzCxxyEgJBKwOVmgDbpao%2B7NWoT%2BnFkHG%2FTLlmWm1DnskX53e1nbdPtzGL1JWZCj30QvUmouZpzxV9TETFDVF7XRQT6Ij3k10vHDb4zGtyNSejTq4LEEvnpJFi78zAA%2Fxn4w%2FpkW374DKBhF4C229%2Bd0f2lVo1GJreMDAtRts%2BFHDMfacd5; HDCN=7D1BE0DD226339228743A98FBCC69AC4-701210564; sessionid=1629192797312; mba_sessionid=45aec5c8-d304-a20e-aef2-d6434230aee1; mba_last_action_time=1629192802446; beta_timer=1629192803404; lastActionTime=1629192819464",
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
	if resp, err = http.Get(strings.Replace(rawImg, "\n", "", -1)); err != nil {
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

func UploadImgTx(rawImg string) (ret string, err error) {
	targetUrl := "https://om.qq.com/image/orginalupload"
	resp, err := http.Get(rawImg)
	if err != nil {
		return "http://www.hulupa.com/zuoz/img/load.gif", nil
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	fileWriter, err := w.CreateFormFile("Filedata", rawImg)
	if err != nil {
		panic(err)
	}

	io.Copy(fileWriter, bytes.NewReader(payload))
	contentType := w.FormDataContentType()
	w.Close()

	resp, err = http.Post(targetUrl, contentType, buf)
	if err != nil {
		panic(err)
	}

	retImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var m Img
	err = json.Unmarshal(retImg, &m)
	if err != nil {
		return "http://www.hulupa.com/zuoz/img/load.gif", nil
	}
	
	return m.Data.Url, nil
}

//func main() {
//	var ret string
//	if _, err = UploadImgTx("https://s0.lgstatic.com/i/image/M00/5B/9C/Ciqc1F9_0LeARfzsAAE16bQXukg851.png"); err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//}
