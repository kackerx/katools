package main

import (
    "fmt"
)

// ISendJiangPin 发奖品的接口
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

type JiangFactory interface {
    CreateJiang() ISendJiangPin
}

type GoodsJiangFactory struct{}

func (g GoodsJiangFactory) CreateJiang() ISendJiangPin {
    return GoodsJiang{}
}

type CardJiangFactory struct{}

func (c CardJiangFactory) CreateJiang() ISendJiangPin {
    return CardJiang{}
}

func factory(jiangFactory JiangFactory) ISendJiangPin {
    return jiangFactory.CreateJiang()
}

func main() {

    m := make([]int, 5)
    fmt.Println(len(m))
    fmt.Println(cap(m))

    var jiangpin1 ISendJiangPin = factory(GoodsJiangFactory{})
    jiangpin1.send()

    var jiangpin2 ISendJiangPin = factory(CardJiangFactory{})
    jiangpin2.send()

}
