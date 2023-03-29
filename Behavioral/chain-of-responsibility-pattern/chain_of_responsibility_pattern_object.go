package chain_of_responsibility_pattern

// Patient 输入对象。病人。
type Patient struct {
	// 对象数据
	Name          string // 名
	cost          int    // 费用
	prescriptions string // 处方
	illness       string // 病情

	// 用来区分是否执行过此流程
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func NewPatient(name string) *Patient {
	return &Patient{Name: name}
}

func (p *Patient) SetCost(cost int) {
	p.cost = cost
	return
}

func (p *Patient) GetCost() int {
	return p.cost
}

func (p *Patient) SetPrescriptions(prescriptions string) {
	p.prescriptions = prescriptions
	return
}

func (p *Patient) GetPrescriptions() string {
	return p.prescriptions
}

func (p *Patient) SetIllness(illness string) {
	p.illness = illness
	return
}

func (p *Patient) GetIllness() string {
	return p.illness
}
