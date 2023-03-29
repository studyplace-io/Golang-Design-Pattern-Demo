### 责任链模式介绍：
责任链模式主要使用场景：有流程化的场景时非常适合使用。

- 一句话概括：某个请求创建一个对象链，每个对象依次检查此请求，并对其进行处理，或者将它传给链中的下一个对象。
- 实现方法：主要分为两种：
    1. 适配器模式将一个类的接口，转换成被调用方期望的另一个接口。
    2. 可以在主对象中维护一个数组或链表，每走一个流程就调用Next()，走入下一流程。
- 使用方法：
    1. 创建输入的"请求"或"对象"中，放入流程字段。
    2. 实现流程接口Process。

### 示例：
```go
// Patient 输入对象。病人。
type Patient struct {
    // 对象数据
    Name          string // 名
    cost          int    // 费用
    prescriptions string // 处方
    illness       string // 病情
    
    // 用来区分是否执行过这个部门
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

// Process 流程接口，表示相关但负责不同功能的对象均实现相同方法
type Process interface {
    Execute(*Patient) // 执行该流程的方法。
    SetNext(Process)  // 交给下一个流程
}

// Reception 对象
type Reception struct {
    // 可以有自己业务逻辑的对象
    // 需要关连下一个部门
    next Process
}

// Execute 执行
func (r *Reception) Execute(p *Patient) {
    // 需要判断一下是否执行过
    if p.registrationDone {
      fmt.Println("Patient registration already done")
      r.next.Execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.SetCost(p.GetCost() + 50) // 挂号费
    p.registrationDone = true
    r.next.Execute(p) // 执行下个流程
}

// SetNext 关连到下一个
func (r *Reception) SetNext(next Process) {
    r.next = next
}
```
