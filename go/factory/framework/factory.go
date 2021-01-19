package framework

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type AbstractFactory struct {
	Factory Factory
}

func (af *AbstractFactory) Create(owner string) Product {
	p := af.Factory.CreateProduct(owner)
	af.Factory.RegisterProduct(p)
	return p
}
