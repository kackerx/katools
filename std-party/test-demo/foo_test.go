package foo_test

import (
    "fmt"
    "io/ioutil"
    "testing"
)

// 测试数据目录
func TestFoo(t *testing.T) {
    if res, err := ioutil.ReadFile("testdata/test.txt"); err != nil {
        t.Errorf("%s", err)
    } else {
        fmt.Println(string(res))
    }
}
