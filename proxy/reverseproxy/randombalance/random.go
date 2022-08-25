package randombalance

import (
    "errors"
    "math/rand"
)

type RandomBalance struct {
    curIndex int
    rss      []string
}

func (r *RandomBalance) Add(addrs ...string) error {
    if len(addrs) == 0 {
        return errors.New("addrs is empty")
    }

    for _, addr := range addrs {
        r.rss = append(r.rss, addr)
    }

    return nil
}

func (r *RandomBalance) Next() string {
    if len(r.rss) == 0 {
        return ""
    }

    r.curIndex = rand.Intn(len(r.rss))
    return r.rss[r.curIndex]
}

func (r *RandomBalance) Get() (string, error) {
    return r.Next(), nil
}
