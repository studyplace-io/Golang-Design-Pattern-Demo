package main

import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Behavioral/chain-of-responsibility-pattern"

func main() {

	// 依序 patient(对象) 经过：  reception -> doctor -> medical -> cashier

	cashier := &chain_of_responsibility_pattern.Cashier{}

	//Set next for medical department
	medical := &chain_of_responsibility_pattern.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &chain_of_responsibility_pattern.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &chain_of_responsibility_pattern.Reception{}
	reception.SetNext(doctor)

	patient := &chain_of_responsibility_pattern.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
