package main

import "fmt"

// 发奖品的接口
type ISendJiangPin interface {
    send()
}

type GoodsJiang struct{}

func (g GoodsJiang) send() {
    fmt.Println("发产品奖品")
}

type CardJiang struct{}

func (c CardJiang) send() {
    fmt.Println("发卡券奖品")
}

func simpleFactory(typ int) ISendJiangPin {
    if typ == 1 {
        return GoodsJiang{}
    }
    if typ == 2 {
        return CardJiang{}
    }
    return nil
}

func main() {
    var jiangpin1 ISendJiangPin = simpleFactory(1)
    jiangpin1.send()

    var jiangpin2 ISendJiangPin = simpleFactory(2)
    jiangpin2.send()
}
