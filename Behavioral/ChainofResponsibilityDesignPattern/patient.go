package ChainofResponsibilityDesignPattern

// 输入对象。 这个事例是病人。
type Patient struct {
	// 对象数据
	Name              string

	// 用来区分是否执行过这个部门
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}