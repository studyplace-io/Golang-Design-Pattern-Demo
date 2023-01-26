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
	client := NewHttpClient(
		WithTimeout(100),
		WithMaxIdle(100),
		WithErrCallback(callback),
	)
	fmt.Println(client)
}
