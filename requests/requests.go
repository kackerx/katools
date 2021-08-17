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
