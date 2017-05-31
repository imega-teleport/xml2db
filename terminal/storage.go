package terminal

import (
	"fmt"

	"github.com/imega-teleport/xml2db/commerceml"
)

type storage struct{}

func NewStorage() *storage {
	return &storage{}
}

func (s storage) CreateGroup(parentID string, group commerceml.Group) (err error) {
	//fmt.Println(parentID, group)
	return
}

func (s storage) CreateProperty(property commerceml.Property) (err error) {
	//fmt.Println(property)

	return
}

func (s storage) CreateProduct(product commerceml.Product) (err error) {
	//fmt.Println(product)

	//prod := product.NewProduct(p)
	//err = prod.CreateImage()

	return
}

func (s storage) CreateProductGroup(parentID string, group commerceml.Group) (err error) {
	return
}

func (s storage) CreateProductImage(parentID string, image commerceml.Image) (err error) {
	//fmt.Println(image.String())
	return
}

func (s storage) CreateProductProperty(parentID string, property commerceml.IdValue) (err error) {
	fmt.Println(property)
	return
}

func (s storage) CreateProductTax(parentID string, tax commerceml.Tax) (err error) {
	return
}

func (s storage) CreateProductRequisite(parentID string, requisite commerceml.Requisite) (err error) {
	return
}

func (s storage) CreateProductContractor(parentID string, contractor commerceml.Contractor) (err error) {
	return
}

func (s storage) CreateProductExcise(parentID string, excise commerceml.Excise) (err error) {
	return
}

func (s storage) CreateProductComponent(parentID string, component commerceml.Component) (err error) {
	return
}
