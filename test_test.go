package ktools

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {

	var retMap sync.Map
	retMap.Store("st", "st")
	retMap.Store("s", "s")
	retMap.Store("sst", "sst")
	retMap.Range(func(key, value interface{}) bool {
		fmt.Println(value)
		return true
	})
}

