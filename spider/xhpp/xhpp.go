package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
)

func main() {
	var name = pflag.StringP("name", "n", "kackerr", "help usage for name")
	pflag.Parse()

	m3Url := "https://m3u8.34cdn.com"
	url := "https://www.b4j77.com/html/202204/45444.html"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	resByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	reStr := `(?U).*?"(.*?m3u8)`
	re := regexp.MustCompile(reStr)
	ret := re.FindStringSubmatch(string(resByte))
	m3Url = m3Url + ret[1]

	args := []string{"-i", m3Url, *name + ".mp4"}

	cmd := exec.Command("ffmpeg", args...)

	fmt.Println(cmd.String())
	//out, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer out.Close()

	//err = cmd.Start()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//go func() {
	//	scanner := bufio.NewScanner(out)
	//	for scanner.Scan() {
	//		fmt.Println(string(scanner.Bytes()))
	//		if err != nil || err == io.EOF {
	//			fmt.Println(err)
	//			break
	//		}
	//	}
	//}()
	//
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
