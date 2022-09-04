package TemplatePattern

// IOtp接口
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

// 对象
type Otp struct {
	// 接口对象
	IOtp IOtp
}

// 所有流程在这里做完 ，对外只有暴露这个方法。
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



// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  return nil
// }

