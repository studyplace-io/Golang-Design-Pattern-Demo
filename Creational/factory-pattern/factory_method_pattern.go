package factory_pattern

const (
	ProductStudyBook = iota
	ProductMachineryCar
)

/*
	当有不同关连种类的对象需要创建，可以用工厂方法
*/

// IProductFactory 总工厂
type IProductFactory interface {
	// 注意：这里返回接口对象，可以看成工厂的"对象"
	CreateProduct(t ProductType) IProduct
}

// IProduct 产品接口，"所有要被创建的对象，都要实现此接口"
type IProduct interface {
	// 可以抽象出不同的方法
	GetInfo() string
	SetInfo()
}

// ProductType 区分标示
type ProductType int

type StudyFactory struct{}

func (*StudyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductStudyBook:
		return &Book{}
	}

	return nil
}

type MachineryFactory struct{}

func (*MachineryFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductMachineryCar:
		return &Car{}
	}

	return nil
}
