package main

import "fmt"
import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/TemplatePattern"

/*
	让我们来考虑一个一次性密码功能 （OTP） 的例子。
	将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。 但无论是短信还是邮件， 整个 OTP 流程都是相同的：

	1.生成随机的 n 位数字。
	2.在缓存中保存这组数字以便进行后续验证。
	3.准备内容。
	4.发送通知。

	后续引入的任何新 OTP 类型都很有可能需要进行相同的上述步骤。
	因此， 我们会有这样的一个场景， 其中某个特定操作的步骤是相同的，
	但实现方式却可能有所不同。 这正是适合考虑使用模板方法模式的情况。

 */



func main() {
	//otp := TemplatePattern.Otp{}
	//
	//smsOTP := &TemplatePattern.Sms{
	// Otp: otp,
	//}
	//
	//smsOTP.GenAndSendOTP(smsOTP,4)
	//
	//emailOTP := &TemplatePattern.Email{
	// Otp: otp,
	//}
	//emailOTP.GenAndSendOTP(emailOTP,4)
	//fmt.Scanln()

	// 建立第一个对象
	smsOTP := &TemplatePattern.Sms{}
	// 定义模版
	o := TemplatePattern.Otp{
		// 不同对象
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")

	// 建立第二个对象
	emailOTP := &TemplatePattern.Email{}
	// 定义模版
	o = TemplatePattern.Otp{
		// 不同对象
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
