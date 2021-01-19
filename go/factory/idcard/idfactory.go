package idcard

import (
	"github.com/hitsuji-haneta/design-pattern/go/factory/framework"
)

type idFactory struct {
	*framework.AbstractFactory
	owners []string
}

func NewFactory() *idFactory {
	idf := &idFactory{
		&framework.AbstractFactory{},
		[]string{},
	}
	idf.Factory = idf
	return idf
}

func (idf *idFactory) CreateProduct(owner string) framework.Product {
	return &idCard{owner}
}

func (idf *idFactory) RegisterProduct(product framework.Product) {
	owner := product.(*idCard).GetOwner()
	idf.owners = append(idf.owners, owner)
}

func (idf *idFactory) GetOwners() []string {
	return idf.owners
}
