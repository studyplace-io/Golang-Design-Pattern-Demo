package factory_pattern

const (
	ProductTechBook = iota
	ProductDailyCar
)

/*
	当有不同关连种类的对象需要创建，可以用工厂方法
 */

// 总工厂
type IProductFactory interface {
	CreateProduct(t int) IProduct
}

type IProduct interface {
	GetInfo() string
}


type ProductType int

type TechFactory struct {}

func (*TechFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductTechBook:
		return &Book{}
	}

	return nil
}

type DailyFactory struct {}

func (*DailyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductDailyCar:
		return &Car{}
	}

	return nil
}


