package main

import (
    "fmt"
    "github.com/kackerx/katools/requests"
    "github.com/tidwall/gjson"
    "log"
)

var (
    SEC_UID = "MS4wLjABAAAASIXegoQCyol8Vty3JdAcYoM_NegtJcUMEJET1nhkhsI"
    SEC_USER_INFO = "https://www.iesdouyin.com/web/api/v2/user/info/?sec_uid=%s"

)

func main() {
    
    resp, err := requests.Get(fmt.Sprintf(SEC_USER_INFO, SEC_UID))
    if err != nil {
        log.Fatalln(err)
    }
    
    nickname := gjson.Get(string(resp), "user_info.nickname").Str
    followerNum := gjson.Get(string(resp), "user_info.follower_count").Int()
    favoriteNum := gjson.Get(string(resp), "user_info.favoriting_count").Int()
    sign := gjson.Get(string(resp), "user_info.signature").Str
    
    fmt.Println(nickname)
    fmt.Println(followerNum)
    fmt.Println(favoriteNum)
    fmt.Println(sign)
    
    


}
