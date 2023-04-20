package template_pattern

import (
	"fmt"
	"testing"
)

func TestTemplatePattern1(t *testing.T) {

	// 建立第一个具体对象
	smsOTP := NewSms()

	fmt.Println("")

	// 建立第二个具体对象
	emailOTP := NewEmail()

	// 初始化抽象模版
	o := NewOtp(emailOTP) // 不同具体模版对象
	o.GenAndSendOTP(4)

	fmt.Println("")

	o = NewOtp(smsOTP)
	o.GenAndSendOTP(4)
}
