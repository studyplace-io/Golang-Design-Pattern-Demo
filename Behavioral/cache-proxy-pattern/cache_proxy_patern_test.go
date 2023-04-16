package cache_proxy_pattern

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheProxyPattern(test *testing.T) {
	r := NewRealRequestProvider()
	startTime := time.Now()

	cp := NewCacheRequestProvider(r)
	res := cp.DoRequest("1")
	fmt.Println("request res: ", res)
	fmt.Println("spend time: ", time.Since(startTime))

	startTime = time.Now()
	res1 := cp.DoRequest("1")
	fmt.Println("request res: ", res1)
	fmt.Println("spend time: ", time.Since(startTime))
}
