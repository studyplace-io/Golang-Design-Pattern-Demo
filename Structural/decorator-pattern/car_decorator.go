package decorator_pattern

import "fmt"

var _ ICar = &Ambulance{}

type Ambulance struct {
	ICar
	StaffNum int
}

func NewAmbulance(car ICar, staffNum int) *Ambulance {
	return &Ambulance{StaffNum: staffNum}
}

func (a *Ambulance) Parking() string {
	fmt.Println("这台车有救人能力")
	return fmt.Sprintf("Parking!!")
}

func (a *Ambulance) Driving() string {
	fmt.Println("这台车有救人能力")
	return fmt.Sprintf("Driving!!")
}

func (a *Ambulance) SavePeople() {
	fmt.Printf("目前正在救人，主要有%v人在救人\n", a.StaffNum)
}

type Chariot struct {
	ICar
	Weapon   string
	StaffNum int
}

func NewChariot(car ICar, weapon string, staffNum int) *Chariot {
	return &Chariot{Weapon: weapon, StaffNum: staffNum}
}

func (c *Chariot) AttackPeople() {
	fmt.Printf("目前正在攻击，主要有%v人在攻击\n", c.StaffNum)
}

func (c *Chariot) Parking() string {
	fmt.Println("这台车有攻击能力")
	return fmt.Sprintf("Driving!!")
}

func (c *Chariot) Driving() string {
	fmt.Println("这台车有攻击能力")
	return fmt.Sprintf("Driving!!")
}

var _ ICar = &Chariot{}

