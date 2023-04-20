package template_pattern

import "fmt"

// Email 发送消息的对象
// 具体模版对象
type Email struct {
	*Otp
}

func NewEmail() *Email {
	// 初始化就把抽象模版中的具体模版注入
	email := &Email{}
	//email.Otp = &Otp{email}
	return email
}

// 接口实现的方法

// 方法
func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

// Sms 发送消息的对象
// 具体模版对象
type Sms struct {
	*Otp
}

func NewSms() *Sms {
	// 初始化就把抽象模版中的具体模版注入
	sms := &Sms{}
	sms.Otp = &Otp{sms}
	return sms
}

// 接口实现的方法

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
