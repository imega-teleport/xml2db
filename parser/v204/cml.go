package v204

import (
	"encoding/xml"

	"github.com/imega-teleport/xml2db/commerceml"
)

type parser struct {
	storage commerceml.Storage
}

func NewParser204(storage commerceml.Storage) parser {
	return parser{
		storage: storage,
	}
}

func (p parser) Parse(data []byte) (err error) {
	cml := &commerceml.CommerceML{}
	err = xml.Unmarshal(data, cml)

	for _, g := range cml.Classifier.Groups {
		p.CreateGroup("", g)
	}

	for _, i := range cml.Classifier.Properties {
		p.CreateProperty(i)
	}

	for _, i := range cml.Catalog.Products {
		p.CreateProduct(i)
	}
	p.CreateProducts(cml.Catalog.Products)

	return
}

func (p parser) CreateProducts(products []commerceml.Product) (err error) {
	err = p.storage.CreateProducts(products)
	return
}

func (p parser) CreateGroup(parentId string, group commerceml.Group) (err error) {
	err = p.storage.CreateGroup(parentId, group)
	if err != nil {
		return
	}

	if len(group.Groups) == 0 {
		return
	}

	for _, g := range group.Groups {
		err = p.CreateGroup(group.Id, g)
		if err != nil {
			return
		}
	}

	return
}

func (p parser) CreateProperty(property commerceml.Property) (err error) {
	err = p.storage.CreateProperty(property)

	return
}

func (p parser) CreateProduct(product commerceml.Product) (err error) {
	for _, i := range product.Groups {
		err = p.storage.CreateProductGroup(product.Id, i)
	}

	for _, i := range product.Images {
		err = p.storage.CreateProductImage(product.Id, i)
	}

	for _, i := range product.Properties {
		err = p.storage.CreateProductProperty(product.Id, i)
	}

	for _, i := range product.Taxes {
		err = p.storage.CreateProductTax(product.Id, i)
	}
	for _, i := range product.Requisites {
		err = p.storage.CreateProductRequisite(product.Id, i)
	}

	if len(product.Manufacturer.Id) > 0 {
		err = p.storage.CreateProductContractor(product.Id, product.Manufacturer)
	}

	if len(product.OwnerBrand.Id) > 0 {
		err = p.storage.CreateProductContractor(product.Id, product.OwnerBrand)
	}

	for _, i := range product.Excises {
		err = p.storage.CreateProductExcise(product.Id, i)
	}

	for _, i := range product.Components {
		err = p.storage.CreateProductComponent(i)
	}

	return
}

func (p parser) ParseBundling(data []byte) (err error) {
	cml := &commerceml.CommerceML{}
	err = xml.Unmarshal(data, cml)

	err = p.storage.CreateBundling(cml.Bundling)
	/*if len(cml.Bundling.Owner.Id) > 0 {
		err = p.storage.CreateProductContractor(product.Id, product.Manufacturer)
	}*/
	err = p.storage.CreatePricesTypes(cml.Bundling, cml.Bundling.PriceTypes)
	err = p.CreateOffers(cml.Bundling, cml.Bundling.Offers)

	return
}

func (p parser) CreateOffers(bundling commerceml.Bundling, offers []commerceml.Offer) (err error) {
	err = p.storage.CreateOffers(bundling, offers)

	for _, i := range offers {
		p.storage.CreatePrices(i, i.Prices)
	}

	return
}
