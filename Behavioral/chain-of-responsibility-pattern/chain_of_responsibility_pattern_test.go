package chain_of_responsibility_pattern

import "testing"

func TestChainOfResponsibilityPattern(t *testing.T) {
	// 依序 patient(对象) 经过：  reception -> doctor -> medical -> cashier

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.SetNext(doctor)

	patient := NewPatient("abc")
	//Patient visiting
	reception.Execute(patient)
}
