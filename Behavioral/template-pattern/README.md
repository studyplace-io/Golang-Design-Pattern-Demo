### 模版模式介绍：
模版模式主要使用场景：一个抽象模版对象定义执行的方法，再实现具体模版对象，并使用嵌套方式组合，调用时将以抽象模版的方式调用。

- 一句话概括：区分抽象模版与具体模版对象，可以相互替换。
- 实现方法：
    1. 定义具体模版需要实现的接口，定义抽象模版对象与具体模版对象，并使用组合嵌套。
- 使用方法：
    1. 初始化不同的具体模版对象。
    2. 使用抽象模版对象赋值，并调用抽象模版对象的方法。

### 示例：
```go
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

// Email 发送消息的对象
// 具体模版对象
type Email struct {
    *Otp
}

func NewEmail() *Email {
    // 初始化就把抽象模版中的具体模版注入
    email := &Email{}
    email.Otp = &Otp{email}
    return email
}

```
