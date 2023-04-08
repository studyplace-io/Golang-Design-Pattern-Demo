package template_pattern

// IOtp 具体模版需要实现的接口
type IOtp interface {
	// 各种方法
	// 生成随机的 n 位数字
	genRandomOTP(int) string
	// 存到缓存中
	saveOTPCache(string)
	// 取得要发送消息的内容
	getMessage(string) string
	// 发送消息通知方法
	sendNotification(string) error
}

// Otp 抽象模版对象
type Otp struct {
	// 内部有IOtp接口对象
	IOtp IOtp
}

func NewOtp(IOtp IOtp) *Otp {
	return &Otp{IOtp: IOtp}
}

// GenAndSendOTP 所有流程在这里做完 ，对外只有暴露这个方法。
func (o *Otp) GenAndSendOTP(otpLength int) error {
	// 流程
	otp := o.IOtp.genRandomOTP(otpLength)
	o.IOtp.saveOTPCache(otp)
	message := o.IOtp.getMessage(otp)
	err := o.IOtp.sendNotification(message)

	// 错误处理
	if err != nil {
		return err
	}
	return nil
}
