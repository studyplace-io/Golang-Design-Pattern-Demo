package ChainofResponsibilityDesignPattern

// 部门接口，表示相关但负责不同功能的对象均实现相同方法
type Department interface {
	// 两个方法
	Execute(*Patient)
	SetNext(Department)
}