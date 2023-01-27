package main

import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Behavioral/ChainofResponsibilityDesignPattern"

func main() {

	// 依序 patient(对象) 经过：  reception -> doctor -> medical -> cashier

	cashier := &ChainofResponsibilityDesignPattern.Cashier{}

	//Set next for medical department
	medical := &ChainofResponsibilityDesignPattern.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &ChainofResponsibilityDesignPattern.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &ChainofResponsibilityDesignPattern.Reception{}
	reception.SetNext(doctor)

	patient := &ChainofResponsibilityDesignPattern.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
