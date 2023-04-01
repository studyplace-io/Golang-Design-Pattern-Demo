### 命令模式介绍：
命令模式主要使用场景：客户端通过调用者发送命令,命令调用的接收者执行相应操作，可以解耦执行方与调用方。

- 一句话概括：将一个请求封装成一个对象，从而是你可用不同的的请求参数对客户进行订制化；ex: 对请求排队或记录请求日志，以及支持可撤销的操作。
- 实现方法：
    1. 所有待实现的命令对象都要实现ICommand接口对象。
- 使用方法：
    1. 初始化各个命令对象，并调用AddCommand加入命令调用方。
    2. 执行。

### 示例：
```go

// ICommand 所有的命令都要实现此接口
// 执行者
type ICommand interface {
    Exec(args ...interface{}) error
}

// Invoker 调用者
type Invoker struct {
    // 重点：命令模式内部需要维护一个切片，用来存储不同命令
    cmds []ICommand
}

func NewInvoker() *Invoker {
    return &Invoker{cmds: make([]ICommand, 0)}
}

// Execute 执行所有命令，传入args 需要注意：传入的个数与顺序，避免报错
func (iv *Invoker) Execute(args ...interface{}) {
    for _, cmd := range iv.cmds {
        if err := cmd.Exec(args...); err != nil {
            // 如果当中有错误，直接失败
            break
        }
    }
}

func (iv *Invoker) AddCommand(cmds ...ICommand) {
    iv.cmds = append(iv.cmds, cmds...)
}

```
