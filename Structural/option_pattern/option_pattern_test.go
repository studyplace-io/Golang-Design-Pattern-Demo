package option_pattern

import (
	"fmt"
	"testing"
)

func TestOptionPattern(t *testing.T) {
	callback := func(err error) {
		if err != nil {
			fmt.Errorf("error: %s", err)
		}
	}

	// 如此，调用方初始化时，就能自由配置对象
	client := NewHttpClient(
		WithTimeout(100),
		WithMaxIdle(100),
		WithErrCallback(callback),
	)
	fmt.Println(client)
}
