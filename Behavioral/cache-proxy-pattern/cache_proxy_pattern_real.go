package cache_proxy_pattern

import (
	"fmt"
	"time"
)

var _ RequestProvider = &RealRequestProvider{}

type RealRequestProvider struct {}

func NewRealRequestProvider() *RealRequestProvider {
	return &RealRequestProvider{}
}

func (r *RealRequestProvider) DoRequest(requestId string) string {

	time.Sleep(time.Second * 5)
	doRequest()
	return fmt.Sprintf("do request for id: %v", requestId)

}

func doRequest() {
	fmt.Println("do something request.")
}
