package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var url string
	for {
		//fmt.Print("输入b站地址回车, 0退出: ")
		//_, err := fmt.Scanf("%s", &url)
		//if err != nil {
		//	panic(err)
		//}
		//if url == "0" {
		//	break
		//}
		url = "https://www.bilibili.com/video/BV1mQ4y1q7wC?spm_id_from=333.999.0.0"
		
		resp, bvid, p := GetJson(url)
		
		videoUrl := resp.Data.Dash.Video[2].BaseUrl
		audioUrl := resp.Data.Dash.Audio[2].BaseUrl
	
		res := GetHtml(videoUrl)
		// Create the file
		//fmt.Println(cid)
		out, err := os.Create(fmt.Sprintf("%s_%s_%s.mp4", bvid, p, "video"))
		if err != nil {
						  panic(err)
						  }
		defer out.Close()
		// Write the body to file
		_, err = io.Copy(out, strings.NewReader(res))
		fmt.Println("下载视频成功")
		
		res = GetHtml(audioUrl)
		// Create the file
		//fmt.Println(cid)
		out, err = os.Create(fmt.Sprintf("%s_%s_%s.mp4", bvid, p, "audio"))
		if err != nil {
			panic(err)
		}
		defer out.Close()
		// Write the body to file
		_, err = io.Copy(out, strings.NewReader(res))
		fmt.Println("下载音频成功")
		
	}

	
	
	
	
	
}
